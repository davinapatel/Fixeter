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

func GetRecords[T any](records *[]T) error {
	if DBConn == nil {
		log.Println("Database Connection is not initialized.")
	}
	return DBConn.Find(records).Error
}

func GetRecordByID[T any](record *T, id string) error {
	if DBConn == nil {
		log.Println("Database Connection is not initialized.")
	}
	return DBConn.Find(record, id).Error
}

func GetIssueStatusCount(results *[]struct {
	Status string
	Count  int64
}) error {
	if DBConn == nil {
		log.Println("Database Connection is not initialized.")
		return DBConn.Error
	}
	return DBConn.Model(&model.Issue{}).Select("status, COUNT(*) as count").Group("status").Scan(results).Error
}

func SaveRecord[T any](record *T) error {
	if DBConn == nil {
		log.Println("Database Connection is not initialized.")
		return DBConn.Error
	}
	return DBConn.Save(record).Error
}

func GetRecordByCategory(record *model.IssueHistory, category string) error {
	if DBConn == nil {
		log.Println("Database Connection is not initialized.")
		return DBConn.Error
	}
	return DBConn.First(record, "category = ?", category).Error
}

func DeleteRecord[T any](record *T) error {
	if DBConn == nil {
		log.Println("Database Connection is not initialized.")
		return DBConn.Error
	}
	return DBConn.Delete(record).Error
}
