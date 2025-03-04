package database

import (
	"log"
	"os"

	"github.com/davinapatel/Fixeter/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectDB() {

	// Access DB Credentials from environment variables

	host := os.Getenv("db_host")
	user := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbname := os.Getenv("db_name")
	port := os.Getenv("db_port")

	// Opens a DB Connection to PostgreSQL DB
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database Connection Failed.")
	}

	log.Println("Database Connection Successful.")

	db.AutoMigrate(new(model.Issue))
	db.AutoMigrate(new(model.Resource))
	db.AutoMigrate(new(model.IssueHistory))

	DBConn = db

}
