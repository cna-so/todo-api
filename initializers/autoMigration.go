package initializers

import (
	"github.com/cna-so/todo-api/models/db"
	"gorm.io/gorm"
	"log"
)

func InitMigrations(Db *gorm.DB) {
	err := Db.AutoMigrate(&db.Users{}, &db.Tag{}, &db.Todo{})
	if err != nil {
		log.Fatal("error while migrations", err)
	}
}
