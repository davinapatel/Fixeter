package tests

import (
	"encoding/json"
	"io"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/davinapatel/Fixeter/controller"
	"github.com/davinapatel/Fixeter/database"
	"github.com/davinapatel/Fixeter/model"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setUpIssueTable() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("testDb"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
	db.AutoMigrate(&model.Issue{})

	database.DBConn = db
	return db
}

func tearDownIssueTable(db *gorm.DB) {
	db.Exec("DROP TABLE issues")
}

func TestIssueListSuccess(t *testing.T) {

	// Setup test db
	db := setUpIssueTable()

	// Will be called at the end of test to clear up db
	defer tearDownIssueTable(db)

	//Create test records
	db.Create(&model.Issue{
		Category:    "TestCategory1",
		Title:       "TestTitle1",
		Date:        time.Now(),
		Description: "TestDescription1",
		Latitude:    1.234,
		Longitude:   -1.345,
		Address:     "Test Address1",
		Image:       "/test/path1",
		Status:      "Logged",
		ResourceID:  1,
		Comments:    "TestComment1",
	})
	db.Create(&model.Issue{
		Category:    "TestCategory2",
		Title:       "TestTitle2",
		Date:        time.Now(),
		Description: "TestDescription2",
		Latitude:    1.2345,
		Longitude:   -1.3456,
		Address:     "Test Address2",
		Image:       "/test/path2",
		Status:      "Logged",
		ResourceID:  1,
		Comments:    "TestComment2",
	})

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Get("/issue", controller.IssueList)

	// Send mock http request
	request := httptest.NewRequest("GET", "/issue", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert no errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	//Parse body from response to check data returned
	body, _ := io.ReadAll(response.Body)
	var responseBody struct {
		IssueRecords []model.Issue `json:"issue_records"`
		Message      string        `json:"message"`
		StatusText   string        `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		log.Println(err)
	}

	var count int64
	result := db.Model(&model.Issue{}).Count(&count)
	if result.Error != nil {
		log.Println("Error in retrieving number of records.")
	}

	assert.Equal(t, int64(2), count)

}

func TestIssueListNoRecords(t *testing.T) {

	// Setup test db
	db := setUpIssueTable()

	// Will be called at the end of test to clear up db
	defer tearDownIssueTable(db)

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Get("/issue", controller.IssueList)

	// Send mock http request
	request := httptest.NewRequest("GET", "/issue", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert no errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	// Parse body from response to check data returned
	body, _ := io.ReadAll(response.Body)
	var responseBody struct {
		IssueRecords []model.Issue `json:"issue_records"`
		Message      string        `json:"message"`
		StatusText   string        `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		log.Println(err)
	}

	expectedRecords := []model.Issue{}
	var count int64
	result := db.Model(&model.Issue{}).Count(&count)
	if result.Error != nil {
		log.Println("Error in retrieving number of records.")
	}

	assert.Equal(t, expectedRecords, responseBody.IssueRecords)
	assert.Equal(t, int64(0), count)

}

func TestIssueDetailSuccess(t *testing.T) {

	// Setup test db
	db := setUpIssueTable()

	// Will be called at the end of test to clear up db
	defer tearDownIssueTable(db)

	//Create test records
	db.Create(&model.Issue{
		Category:    "TestCategory1",
		Title:       "TestTitle1",
		Date:        time.Now(),
		Description: "TestDescription1",
		Latitude:    1.234,
		Longitude:   -1.345,
		Address:     "Test Address1",
		Image:       "/test/path1",
		Status:      "Logged",
		ResourceID:  1,
		Comments:    "TestComment1",
	})

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Get("/issue/:id", controller.IssueDetail)

	// Send mock http request
	request := httptest.NewRequest("GET", "/issue/1", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert no errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	// Parse body from response to check data returned
	body, _ := io.ReadAll(response.Body)
	var responseBody struct {
		IssueRecords model.Issue `json:"record"`
		Message      string      `json:"message"`
		StatusText   string      `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		log.Println(err)
	}

	log.Println(responseBody.IssueRecords)

	assert.Equal(t, uint(1), responseBody.IssueRecords.ID)

}
