package usecases

import (
	"github.com/boomctr/todo-backend-go/entities"
)

type TodoUseCase interface {
	CreateTodo(todo *entities.Todos) error
	GetTodo(id uint64) (*entities.Todos, error)
	GetAllTodos(id uint64) (*[]entities.Todos, error)
	UpdateTodo(todo *entities.Todos) error
	DeleteTodo(id uint64) error
}

type TodoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) TodoUseCase {
	return &TodoService{
		repo: repo,
	}
}

func (t *TodoService) GetTodo(id uint64) (*entities.Todos, error) {
	return t.repo.GetID(id)
}

func (t *TodoService) CreateTodo(todo *entities.Todos) error {
	return t.repo.Save(todo)
}

func (t *TodoService) GetAllTodos(id uint64) (*[]entities.Todos, error) {
	return t.repo.GetAll(id)
}

func (t *TodoService) UpdateTodo(todo *entities.Todos) error {
	return t.repo.Update(todo)
}

func (t *TodoService) DeleteTodo(id uint64) error {
	return t.repo.Delete(id)
}
