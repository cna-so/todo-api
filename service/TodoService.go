package service

import (
	dto "github.com/cna-so/todo-api/models/DTO"
	"github.com/cna-so/todo-api/models/db"
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
			resTodos = append(resTodos, dto.TodoResponseDto{
				ID: todo.ID, Title: todo.Title, Status: todo.Status,
				CreateAt: todo.CreateAt, ExpireDate: todo.ExpireDate,
				Tags: todo.Tags,
			})
		}
	}
	return resTodos, nil
}

func (ts *TodoService) CreateTodo(todo dto.TodoCreateRequestDto) (dto.TodoResponseDto, error) {
	todoModel := db.Todo{
		UserID:      todo.UserID,
		Title:       todo.Title,
		Tags:        todo.Tags,
		Description: todo.Description,
		ExpireDate:  todo.ExpireDate,
		Status:      todo.Status,
	}
	if createTodo, err := ts.repository.CreateTodo(todoModel); err != nil {
		return dto.TodoResponseDto{}, err
	} else {
		var resTodos dto.TodoResponseDto
		resTodos = dto.TodoResponseDto{
			ID: createTodo.ID, Title: createTodo.Title, Status: createTodo.Status,
			CreateAt: createTodo.CreateAt, ExpireDate: createTodo.ExpireDate,
			Tags: createTodo.Tags,
		}
		return resTodos, err
	}
}

func (ts *TodoService) ChangeTodoStatus(todoStatus dto.TodoChangeStatusRequestDto) (dto.TodoResponseDto, error) {
	todo, err := ts.repository.ChangeTodoStatus(todoStatus)
	if err != nil {
		return dto.TodoResponseDto{}, err
	}
	return dto.TodoResponseDto{
		ID: todo.ID, Title: todo.Title, Status: todo.Status,
		CreateAt: todo.CreateAt, ExpireDate: todo.ExpireDate,
		Tags: todo.Tags,
	}, nil
}
