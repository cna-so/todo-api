package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cna-so/todo-api/controllers/handlers"
)

func Routes() *gin.Engine {
	router := gin.New()
	// middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// test server
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// authorization (login - register )
	auth := router.Group("/api/v1/auth")
	auth.POST("/create", handlers.CreateUserHandler)
	return router
}
