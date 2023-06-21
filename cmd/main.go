package main

import (
	"fmt"
	"github.com/cna-so/todo-api/api"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/cna-so/todo-api/initializers"
)

var db *gorm.DB

func init() {
	initializers.LoadEnv()
	db = initializers.ConnectToDatabase()
	initializers.InitMigrations(db)
}

func main() {
	server := http.Server{
		Addr:    os.Getenv("port"),
		Handler: api.SetupRouter(db),
	}
	fmt.Printf("start server on port %s\n", os.Getenv("port"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
