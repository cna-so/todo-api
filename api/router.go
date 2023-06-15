package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/cna-so/todo-api/controllers/authorization/jwt"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()
	// middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// setup api
	userApi := initUserApi(db)
	tagApi := initTagApi(db)
	todoApi := initTodoApi(db)
	// test server
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	// groups
	auth := router.Group("/api/v1/auth")
	tag := router.Group("/api/v1/tag/")
	todo := router.Group("/api/v1/todo")
	// authorization (login - register )
	auth.POST("/signup", userApi.CreateUserAccount)
	auth.POST("/login", userApi.SignInUser)
	auth.GET("/check", jwt.RequireLogin, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "login")
	})
	// Tag
	tag.POST("/create", jwt.RequireLogin, tagApi.CreateTag)
	tag.GET("/all", jwt.RequireLogin, tagApi.GetAllTags)
	tag.GET("/search/:name", jwt.RequireLogin, tagApi.GetTagByName)
	tag.DELETE("/delete/:id", jwt.RequireLogin, tagApi.DeleteTagById)
	// todo
	todo.GET("/all", jwt.RequireLogin, todoApi.GetAllTodo)
	todo.POST("/create", jwt.RequireLogin, todoApi.CreateTodo)
	todo.POST("/change/status", jwt.RequireLogin, todoApi.ChangeTodoStatus)
	return router
}
