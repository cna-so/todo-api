package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/cna-so/todo-api/helpers"
)

var (
	Db    *gorm.DB
	dbCfg = helpers.DatabaseConfig{
		User:     "cna",
		Database: "todo_db",
		SSLMode:  "disable",
		Port:     "5432",
		Password: "1234",
		Host:     "localhost",
	}
)

func ConnectToDatabase() *gorm.DB {
	var err error
	Db, err = gorm.Open(postgres.Open(dbCfg.DbValues()), &gorm.Config{})
	if err != nil {
		log.Fatal("error while connecting to postgresSQL ", err)
	}
	fmt.Printf("successfully connect to database , connections details = %s \n", dbCfg.DbValues())
	return Db
}
