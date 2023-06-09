package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/cna-so/todo-api/initializers"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	fmt.Printf("start server on port %s", os.Getenv("port"))
	app.Listen(os.Getenv("port"))
}
