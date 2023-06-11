package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cna-so/todo-api/controllers/routes"
	"github.com/cna-so/todo-api/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabase()
}

func main() {
	server := http.Server{
		Addr:    os.Getenv("port"),
		Handler: routes.Routes(),
	}
	fmt.Printf("start server on port %s", os.Getenv("port"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
