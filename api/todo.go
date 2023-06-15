package api

import (
	"github.com/cna-so/todo-api/controllers/authorization/jwt"
	dto "github.com/cna-so/todo-api/models/DTO"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"github.com/cna-so/todo-api/service"
)

type TodoApi struct {
	ts *service.TodoService
}

func (ta *TodoApi) GetAllTodo(ctx *gin.Context) {
	token := strings.Split(ctx.GetHeader("Authorization"), " ")
	tokenValues, _ := jwt.GetValueFromJWT(token[1], []string{"user_id"})
	todos, err := ta.ts.GetAllTodos(tokenValues["user_id"])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func (ta *TodoApi) CreateTodo(ctx *gin.Context) {
	var todo dto.TodoCreateRequestDto
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	token := strings.Split(ctx.GetHeader("Authorization"), " ")
	tokenValues, _ := jwt.GetValueFromJWT(token[1], []string{"user_id"})
	todo.UserID = tokenValues["user_id"]
	ta.ts.CreateTodo(todo)
}
