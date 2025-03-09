package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/davinapatel/Fixeter/controller"
	"github.com/davinapatel/Fixeter/database"
	"github.com/davinapatel/Fixeter/model"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type IssueHistoryRecord struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Count    int    `json:"count"`
}

type IssueHistoryResponse struct {
	IssueHistoryRecords []IssueHistoryRecord `json:"issueHistory_records"`
	Message             string               `json:"message"`
	StatusText          string               `json:"statusText"`
}

func setUpIssueHistoryTable() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("testDb"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
	db.AutoMigrate(&model.IssueHistory{})

	database.DBConn = db
	return db
}

func tearDownIssueHistoryTable(db *gorm.DB) {
	db.Exec("DROP TABLE issue_histories")
}

func TestIssueHistoryList(t *testing.T) {

	// Setup test db
	db := setUpIssueHistoryTable()

	// Will be called at the end of test to clear up db
	defer tearDownIssueHistoryTable(db)

	//Create test records
	db.Create(&model.IssueHistory{
		Category: "TestCategory1",
		Count:    100,
	})

	db.Create(&model.IssueHistory{
		Category: "TestCategory2",
		Count:    200,
	})

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Get("/issuehistory", controller.IssueHistoryList)

	// Send mock http request
	request := httptest.NewRequest("GET", "/issuehistory", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert no errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	// Parse body from response to check data returned
	body, _ := io.ReadAll(response.Body)
	var responseBody struct {
		IssueHistoryRecords []IssueHistoryRecord `json:"issueHistory_records"`
		Message             string               `json:"message"`
		StatusText          string               `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		log.Println(err)
	}
	fmt.Println()

	expectedRecords := []IssueHistoryRecord{
		{
			ID:       1,
			Category: "TestCategory1",
			Count:    100,
		},
		{
			ID:       2,
			Category: "TestCategory2",
			Count:    200,
		},
	}

	assert.Equal(t, expectedRecords, responseBody.IssueHistoryRecords)

}

func TestIssueHistoryUpdateSuccess(t *testing.T) {
	// Setup test db
	db := setUpIssueHistoryTable()

	// Will be called at the end of test to clear up db
	defer tearDownIssueHistoryTable(db)

	//Create test records
	category := "TestCategory1"
	count := 100
	db.Create(&model.IssueHistory{
		Category: category,
		Count:    count,
	})

	controller.IssueHistoryUpdate("TestCategory1")

	var updatedRecord model.IssueHistory
	db.First(&updatedRecord, "category=?", category)

	assert.Equal(t, count+1, updatedRecord.Count)

}

func TestIssueHistoryUpdateNoRecordFound(t *testing.T) {
	// Setup test db
	db := setUpIssueHistoryTable()

	// Will be called at the end of test to clear up db
	defer tearDownIssueHistoryTable(db)

	//Create test records
	category := "TestCategory1"
	count := 100
	db.Create(&model.IssueHistory{
		Category: category,
		Count:    count,
	})

	controller.IssueHistoryUpdate("TestCategory5")

	var updatedRecord model.IssueHistory
	err := db.First(&updatedRecord, "category=?", category).Error

	assert.Equal(t, count, updatedRecord.Count)
	assert.Error(t, gorm.ErrRecordNotFound, err)

}

func TestIssueHistoryListNoRecords(t *testing.T) {

	// Setup test db
	db := setUpIssueHistoryTable()

	// Will be called at the end of test to clear up db
	defer tearDownIssueHistoryTable(db)

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Get("/issuehistory", controller.IssueHistoryList)

	// Send mock http request
	request := httptest.NewRequest("GET", "/issuehistory", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert no errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	// Parse body from response to check data returned
	body, _ := io.ReadAll(response.Body)
	var responseBody struct {
		IssueHistoryRecords []IssueHistoryRecord `json:"issueHistory_records"`
		Message             string               `json:"message"`
		StatusText          string               `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		log.Println(err)
	}
	fmt.Println()

	expectedRecords := []IssueHistoryRecord{}

	assert.Equal(t, expectedRecords, responseBody.IssueHistoryRecords)

}
