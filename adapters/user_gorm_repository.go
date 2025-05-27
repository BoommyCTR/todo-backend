package adapters

import (
	"os"
	"time"

	"github.com/boomctr/todo-backend-go/entities"
	"github.com/boomctr/todo-backend-go/usecases"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) usecases.UserRepository {
	return &GormUserRepository{
		db: db,
	}
}

// create a new user
func (g *GormUserRepository) AddUser(user *entities.Users) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashPassword)
	// Insert the user into the database
	result := g.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// login user
func (g *GormUserRepository) CheckUser(user *entities.Users) (string, error) {
	loginUser := new(entities.Users)

	result := g.db.First(&loginUser, "email = ?", user.Email)

	if result.Error != nil {
		return "", result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(user.Password))

	if err != nil {
		return "", err
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"user_id": loginUser.ID,
		"admin":   false,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return t, nil

}

func (g *GormUserRepository) WhoAmI(id uint64) (string, error) {
	user := new(entities.Users)

	result := g.db.First(&user, id)

	if result.Error != nil {
		return "", result.Error
	}

	return user.Name, nil
}
