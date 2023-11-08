package repositories

import (
	"github.com/doganarif/todo/internal/models"
	"gorm.io/gorm"
)

type ITodoRepository interface {
	Create(*models.Todo) error
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (t *TodoRepository) Create(todo *models.Todo) error {
	if err := t.db.Create(todo).Error; err != nil {
		return err
	}
	return nil

}
