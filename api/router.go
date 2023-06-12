package api

import (
	"github.com/cna-so/todo-api/controllers/jwt"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()
	// middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// setup api
	userApi := initUserApi(db)
	// test server
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// authorization (login - register )
	auth := router.Group("/api/v1/auth")
	auth.POST("/signup", userApi.CreateUserAccount)
	auth.POST("/login", userApi.SignInUser)
	auth.GET("/check", jwt.RequireLogin, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "login")
	})
	return router
}
