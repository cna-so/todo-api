package api

import (
	"fmt"
	"github.com/cna-so/todo-api/initializers"
	dto "github.com/cna-so/todo-api/models/DTO"
	"github.com/cna-so/todo-api/models/db"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cna-so/todo-api/helpers"
)

func CreateUserAccount(ctx *gin.Context) {
	var userDto dto.UserRequestDto
	err := ctx.ShouldBindJSON(&userDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "email and password are require",
		})
		return
	}
	userDto.Password, _ = helpers.GenerateHashPassword(userDto.Password)
	user := db.Users{
		Email:     userDto.Email,
		Password:  userDto.Password,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Role:      "U",
	}
	if err := initializers.Db.Create(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("user with email : `%s` successfully created ", user.Email),
	})
}
