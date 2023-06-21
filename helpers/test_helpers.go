package helpers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/cna-so/todo-api/models/db"
)

var (
	TestDb    *gorm.DB
	dbTestCfg = DatabaseConfig{
		User:     "cna",
		Database: "todo_db_test",
		SSLMode:  "disable",
		Port:     "5432",
		Password: "1234",
		Host:     "localhost",
	}
)

func SetupDatabase() *gorm.DB {
	TestDb, _ = gorm.Open(postgres.Open(dbTestCfg.DbValues()), &gorm.Config{})
	_ = TestDb.AutoMigrate(&db.Users{}, &db.Tag{}, &db.Todo{})
	return TestDb
}

func DropTables() {
	_ = TestDb.Migrator().DropTable(&db.Users{}, &db.Tag{}, &db.Todo{})
}
