package api

import (
	"github.com/cna-so/todo-api/controllers/authorization/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"github.com/cna-so/todo-api/service"
)

type TodoApi struct {
	ts *service.TodoService
}

func (ta TodoApi) GetAllTodo(ctx *gin.Context) {
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
