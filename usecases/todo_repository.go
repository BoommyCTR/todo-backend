package usecases

import (
	"github.com/boomctr/todo-backend-go/entities"
)

type TodoRepository interface {
	Save(todo *entities.Todos) error
	GetID(id uint64) (*entities.Todos, error)
	GetAll(id uint64) (*[]entities.Todos, error)
	Update(todo *entities.Todos) error
	Delete(id uint64) error
}
