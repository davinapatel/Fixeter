// The package database is concerned with all operations concering the PostgreSQL DB
// This includes connecting to the database and performing queries

package database

import (
	"log"
	"os"

	"github.com/davinapatel/Fixeter/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// ConnectDB function establishes a connection to PostgreSQL DB by using DB Credentials from environment vars
// Assigns the DB Connection to global variable DBConn
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

	// Creates tables in DB Using models
	db.AutoMigrate(new(model.Issue))
	db.AutoMigrate(new(model.Resource))
	db.AutoMigrate(new(model.IssueHistory))

	DBConn = db

}

// GetRecords retrieves all records from the database for the specified model type T and stores them in provided slice
// using GORM's Find method to retrieve all records

// Type Parameters:
// -T: Represents any model type that corresponds to a DB Table
//
// Parameter:
// - records: This is a pointer to slice of type T where the retrieved records will be stored
//
// Returns:
// - error: Returns error if DB connection not intialised or query fails

func GetRecords[T any](records *[]T) error {
	if DBConn == nil {
		log.Println("Database Connection is not initialized.")
	}
	return DBConn.Find(records).Error
}

// GetRecordByID retrieves a records from the database for the specified id for the specified model type T and stores record
// in a pointer to the struct of model type T. Uses GORM Find method to retrieve record by passing id into it.

// Type Parameters:
// -T: Represents any model type that corresponds to a DB Table
//
// Parameter:
// - record: This is a pointer to the struct of type T where the retrieved recordswill be stored
//
// Returns:
// - error: Returns error if DB connection not intialised or query fails
func GetRecordByID[T any](record *T, id string) error {
	if DBConn == nil {
		log.Println("Database Connection is not initialized.")
	}
	return DBConn.Find(record, id).Error
}

// GetIssueStatusCount retrieves a count of issues grouped by the field "Status" from the database.
// This method is only to be used on the Issue model.
//
// Parameter:
//   - results: This is a pointer to a slice of structs, where each struct will store the Status i.e "Logged", "Closed" and a
//     Count of the number of issues with that status. This is where the results of the DB query will be stored.
//
// Returns:
// - error: Returns error if DB connection not intialised or query fails
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

// SaveRecord saves a record to the DB for the specified model type T.
// The data of the record to be saved is held in a pointer to a struct of the model type T.

// Type Parameters:
// -T: Represents any model type that corresponds to a DB Table
//
// Parameter:
// - record: This is a pointer to the struct of type T where the record to be saved is held
//
// Returns:
// - error: Returns error if DB connection not intialised or query fails
func SaveRecord[T any](record *T) error {
	if DBConn == nil {
		log.Println("Database Connection is not initialized.")
		return DBConn.Error
	}
	return DBConn.Save(record).Error
}

// GetRecordByCategory retrieves a records from the database for the specified id for the specified model type T and stores record
// in a pointer to the struct of model type T. Uses GORM Find method to retrieve record by passing id into it.
//
// Parameter:
// - record: This is a pointer to the struct of the model IssueHistory
// - category: This is the category that we would like to filter the records on i.e Graffiti
//
// Returns:
// - error: Returns error if DB connection not intialised or query fails
func GetRecordByCategory(record *model.IssueHistory, category string) error {
	if DBConn == nil {
		log.Println("Database Connection is not initialized.")
		return DBConn.Error
	}
	return DBConn.First(record, "category = ?", category).Error
}

// DeleteRecord deletes a record from the DB for the specified model type T.
//
// Type Parameters:
// -T: Represents any model type that corresponds to a DB Table
//
// Parameter:
// - record: This is a pointer to the struct of type T where the data for the record to be deleted is held
//
// Returns:
// - error: Returns error if DB connection not intialised or query fails
func DeleteRecord[T any](record *T) error {
	if DBConn == nil {
		log.Println("Database Connection is not initialized.")
		return DBConn.Error
	}
	return DBConn.Delete(record).Error
}
