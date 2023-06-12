package api

import (
	"github.com/cna-so/todo-api/repository"
	"github.com/cna-so/todo-api/service"
	"gorm.io/gorm"
)

func initUserApi(db *gorm.DB) UserApi {
	userRepository := repository.UserRepositoryProvider(db)
	userService := service.UserServiceProvider(userRepository)
	userApi := UserApi{
		s: userService,
	}
	return userApi
}
