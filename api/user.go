package api

import (
	"fmt"
	"github.com/cna-so/todo-api/models/DTO"
	"github.com/cna-so/todo-api/models/db"
	"github.com/cna-so/todo-api/service"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cna-so/todo-api/helpers"
)

type UserApi struct {
	s service.UserService
}

func (us *UserApi) CreateUserAccount(ctx *gin.Context) {
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
	mUser, err := us.s.SignUpUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("user with email : `%s` successfully created ", mUser.Email),
	})
}

func (us *UserApi) SignInUser(ctx *gin.Context) {

}
