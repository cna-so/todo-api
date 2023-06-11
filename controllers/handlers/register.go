package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cna-so/todo-api/healpers"
	"github.com/cna-so/todo-api/models"
)

func CreateUserHandler(ctx *gin.Context) {
	var user models.UserModels
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "email and password are require",
		})
		return
	}
	user.Password, _ = healpers.GenerateHashPassword(user.Password)
	status, err := user.CreateUserAccount()
	if err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
}
