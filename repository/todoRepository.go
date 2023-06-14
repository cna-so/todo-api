package repository

import (
	"gorm.io/gorm"

	"github.com/cna-so/todo-api/models/db"
)

type TodoRepository struct {
	db *gorm.DB
}

func TodoRepositoryProvider(db *gorm.DB) TodoRepository {
	return TodoRepository{
		db: db,
	}
}

func (tr *TodoRepository) FindAllTodos(userId string) (*[]db.Todo, error) {
	var todos []db.Todo
	if err := tr.db.Model(&db.Todo{}).Where(db.Todo{UserID: userId}).Find(&todos).Error; err != nil {
		return nil, err
	} else {
		return &todos, nil
	}
}
