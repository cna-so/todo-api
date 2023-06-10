package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	_ = router.Group("/api/v1/auth")
	return router
}
