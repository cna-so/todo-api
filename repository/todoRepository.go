package repository

import (
	"errors"
	"fmt"
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
	if err := tr.db.Preload("Tags").Where(db.Todo{UserID: userId}).Find(&todos).Error; err != nil {
		return nil, err
	} else {
		return &todos, nil
	}
}

func (tr *TodoRepository) CheckTagAccessById(tags []*db.Tag) bool {
	var counts int64
	for _, tag := range tags {
		var count int64
		tr.db.Model(&db.Tag{}).Where(db.Tag{BaseModel: db.BaseModel{ID: tag.ID}}).Where(db.Tag{UserID: tag.UserID}).Count(&count)
		counts += count
	}
	fmt.Println(len(tags), counts)
	return int64(len(tags)) == counts
}

func (tr *TodoRepository) CreateTodo(todo db.Todo) (*db.Todo, error) {
	isHaveAccess := tr.CheckTagAccessById(todo.Tags)
	fmt.Println(isHaveAccess)
	if isHaveAccess != true {
		return nil, errors.New("tags id doesn't exist")
	}
	if err := tr.db.Create(&todo).Error; err != nil {
		return nil, err
	} else {
		return &todo, nil
	}
}
