package initializers

import (
	"database/sql"
	"fmt"
	"github.com/cna-so/todo-api/controllers/handlers"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB
var dbCfg = handlers.DatabaseConfig{
	User:     "cna",
	Database: "todo_db",
	SSLMode:  "disable",
	Port:     "5432",
	Password: "1234",
}

func ConnectToDatabase() *sql.DB {
	var err error
	DB, err = sql.Open("postgres", dbCfg.DbValues())
	if err != nil {
		log.Fatal("error while connecting to postgresSQL ", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("error while pinging postgresSQL ", err)
	}
	fmt.Printf("successfully connect to database , connections details = %s \n", dbCfg.DbValues())
	return DB
}
