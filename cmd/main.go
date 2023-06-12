package main

import (
	"fmt"
	"github.com/cna-so/todo-api/api"
	"log"
	"net/http"
	"os"

	"github.com/cna-so/todo-api/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabase()
	initializers.InitMigrations()
}

func main() {
	server := http.Server{
		Addr:    os.Getenv("port"),
		Handler: api.Routes(),
	}
	fmt.Printf("start server on port %s", os.Getenv("port"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
