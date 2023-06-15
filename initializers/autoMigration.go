package initializers

import (
	"github.com/cna-so/todo-api/models/db"
	"log"
)

func InitMigrations() {
	err := Db.AutoMigrate(&db.Users{}, &db.Tag{}, &db.Todo{})
	if err != nil {
		log.Fatal("error while migrations", err)
	}
}
