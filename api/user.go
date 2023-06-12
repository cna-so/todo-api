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
	var userDto dto.UserSignUpRequestDto
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
	if mUser, err := us.s.SignUpUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"message": fmt.Sprintf("user with email : `%s` successfully created ", mUser.Email),
		})
	}
}

func (us *UserApi) SignInUser(ctx *gin.Context) {
	var userDto dto.UserSignInRequest
	err := ctx.ShouldBindJSON(&userDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "email and password are require",
		})
		return
	}
	if user, err := us.s.SignInUser(userDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"user": dto.UserSignInResponseDto{
				ID:        user.ID,
				Email:     user.Email,
				Password:  user.Password,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
		})
	}
}
