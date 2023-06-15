package repository

import (
	"fmt"
	dto "github.com/cna-so/todo-api/models/DTO"
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

func (tr *TodoRepository) CreateTodo(todo dto.TodoCreateRequestDto) {
	test := db.Todo{
		UserID: "5d883b8e-0f85-409e-875e-cc1a31d648d2",
		Title:  "test",
		Tags: []*db.Tag{
			{BaseModel: db.BaseModel{ID: "081b5fe0-4e3f-4852-a7e3-151a7ce082ad"}},
			{BaseModel: db.BaseModel{ID: "5c344cab-ad36-4ba9-a070-5da5a615b1ab"}},
		},
	}
	tr.db.Create(&test)
	fmt.Println(todo)
}
