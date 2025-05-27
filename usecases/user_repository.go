package usecases

import "github.com/boomctr/todo-backend-go/entities"

type UserRepository interface {
	AddUser(user *entities.Users) error
	CheckUser(user *entities.Users) (string, error)
	WhoAmI(id uint64) (string, error)
}
