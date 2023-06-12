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
	dsn := "host=localhost user=cna password=1234 dbname=todo_db port=5432 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error while connecting to postgresSQL ", err)
	}
	fmt.Printf("successfully connect to database , connections details = %s \n", dbCfg.DbValues())
	return Db
}
