package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	auth.POST("/create", CreateUserHandler)
	return router
}
