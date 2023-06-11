package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cna-so/todo-api/models"
)

func CreateUserHandler(ctx *gin.Context) {
	var user models.UserModels
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "email and password are require",
		})
		return
	}
	status, err := user.CreateUserAccount()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "email and password are require",
		})
	}
	fmt.Println(status)
}
