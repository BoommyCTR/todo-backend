package usecases

import (
	"errors"

	"github.com/boomctr/todo-backend-go/entities"
)

type UserUseCase interface {
	CreateUser(user *entities.Users) error
	VerifyUser(user *entities.Users) (string, error)
	WhoAmIUser(id uint64) (string, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserUseCase {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(user *entities.Users) error {
	if user.Name == "" || user.Password == "" || user.Email == "" {
		return errors.New("name, password, and email are required")
	}

	return u.repo.AddUser(user)
}

func (u *UserService) VerifyUser(user *entities.Users) (string, error) {
	return u.repo.CheckUser(user)
}

func (u *UserService) WhoAmIUser(id uint64) (string, error) {
	return u.repo.WhoAmI(id)
}
