package service

import (
	dto "github.com/cna-so/todo-api/models/DTO"
	"github.com/cna-so/todo-api/repository"
)

type TodoService struct {
	repository repository.TodoRepository
}

func TodoServiceProvider(repository repository.TodoRepository) TodoService {
	return TodoService{
		repository: repository,
	}
}

func (ts *TodoService) GetAllTodos(userId string) ([]dto.TodoResponseDto, error) {
	var resTodos []dto.TodoResponseDto
	if dbTodos, err := ts.repository.FindAllTodos(userId); err != nil {
		return nil, err
	} else {
		for _, todo := range *dbTodos {
			resTodos = append(resTodos, dto.TodoResponseDto{ID: todo.ID, Title: todo.Title, Status: todo.Status, CreateAt: todo.CreateAt, ExpireDate: todo.ExpireDate})
		}
	}
	return resTodos, nil
}
