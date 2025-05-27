package adapters

import (
	"github.com/boomctr/todo-backend-go/entities"
	"github.com/boomctr/todo-backend-go/usecases"
	"gorm.io/gorm"
)

type GormTodoRepository struct {
	db *gorm.DB
}

func NewGormTodoRepository(db *gorm.DB) usecases.TodoRepository {
	return &GormTodoRepository{
		db: db,
	}
}

func (g *GormTodoRepository) Save(todo *entities.Todos) error {
	result := g.db.Create(&todo)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (g *GormTodoRepository) GetID(id uint64) (*entities.Todos, error) {
	Ts := new(entities.Todos)

	result := g.db.Find(&Ts, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return Ts, nil
}

func (g *GormTodoRepository) GetAll(id uint64) (*[]entities.Todos, error) {
	Ts := new([]entities.Todos)

	result := g.db.Where("users_id = ?", id).Order("id asc").Find(&Ts)

	if result.Error != nil {
		return nil, result.Error
	}

	return Ts, nil
}

func (g *GormTodoRepository) Update(todo *entities.Todos) error {
	result := g.db.Save(&todo)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (g *GormTodoRepository) Delete(id uint64) error {
	T := new(entities.Todos)
	// Delete the Todo from the database
	result := g.db.Delete(&T, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
