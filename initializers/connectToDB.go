package initializers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/cna-so/todo-api/healpers"
)

var (
	Db    *sql.DB
	dbCfg = healpers.DatabaseConfig{
		User:     "cna",
		Database: "todo_db",
		SSLMode:  "disable",
		Port:     "5432",
		Password: "1234",
	}
)

func ConnectToDatabase() *sql.DB {
	var err error
	Db, err = sql.Open("postgres", dbCfg.DbValues())
	if err != nil {
		log.Fatal("error while connecting to postgresSQL ", err)
	}
	err = Db.Ping()
	if err != nil {
		log.Fatal("error while pinging postgresSQL ", err)
	}
	fmt.Printf("successfully connect to database , connections details = %s \n", dbCfg.DbValues())
	return Db
}
