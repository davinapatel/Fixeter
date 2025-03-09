package tests

import (
	"bytes"
	"encoding/json"
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

func setUpResourceTable() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("testDb"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
	db.AutoMigrate(&model.Resource{})

	database.DBConn = db
	return db
}

func tearDownResourceTable(db *gorm.DB) {
	db.Exec("DROP TABLE resources")
}

func TestResourceListSuccess(t *testing.T) {

	// Setup test db
	db := setUpResourceTable()

	// Will be called at the end of test to clear up db
	defer tearDownResourceTable(db)

	//Create test records
	db.Create(&model.Resource{
		Type:       "TestType1",
		Department: "TestDepartment1",
	})

	db.Create(&model.Resource{
		Type:       "TestType2",
		Department: "TestDepartment2",
	})

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Get("/resource", controller.ResourceList)

	// Send mock http request
	request := httptest.NewRequest("GET", "/resource", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert no errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	// Parse body from response to check data returned
	body, _ := io.ReadAll(response.Body)
	var responseBody struct {
		ResourceRecords []model.Resource `json:"resource_records"`
		Message         string           `json:"message"`
		StatusText      string           `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		log.Println(err)
	}

	expectedRecords := []model.Resource{
		{
			ID:         1,
			Type:       "TestType1",
			Department: "TestDepartment1",
		},
		{
			ID:         2,
			Type:       "TestType2",
			Department: "TestDepartment2",
		},
	}

	assert.Equal(t, expectedRecords, responseBody.ResourceRecords)

}

func TestResourceListNoRecords(t *testing.T) {

	// Setup test db
	db := setUpResourceTable()

	// Will be called at the end of test to clear up db
	defer tearDownResourceTable(db)

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Get("/resource", controller.ResourceList)

	// Send mock http request
	request := httptest.NewRequest("GET", "/resource", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert no errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	// Parse body from response to check data returned
	body, _ := io.ReadAll(response.Body)
	var responseBody struct {
		ResourceRecords []model.Resource `json:"resource_records"`
		Message         string           `json:"message"`
		StatusText      string           `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		log.Println(err)
	}

	expectedRecords := []model.Resource{}

	assert.Equal(t, expectedRecords, responseBody.ResourceRecords)

}

func TestResourceDetailSuccess(t *testing.T) {

	// Setup test db
	db := setUpResourceTable()

	// Will be called at the end of test to clear up db
	defer tearDownResourceTable(db)

	//Create test records
	db.Create(&model.Resource{
		Type:       "TestType1",
		Department: "TestDepartment1",
	})

	db.Create(&model.Resource{
		Type:       "TestType2",
		Department: "TestDepartment2",
	})

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Get("/resource/:id", controller.ResourceDetail)

	// Send mock http request
	request := httptest.NewRequest("GET", "/resource/1", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert no errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	// Parse body from response to check data returned
	body, _ := io.ReadAll(response.Body)
	var responseBody struct {
		ResourceRecords model.Resource `json:"record"`
		Message         string         `json:"message"`
		StatusText      string         `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		log.Println(err)
	}

	expectedRecords := model.Resource{
		ID:         1,
		Type:       "TestType1",
		Department: "TestDepartment1",
	}

	assert.Equal(t, expectedRecords, responseBody.ResourceRecords)

}

func TestResourceDetailRecordNotFound(t *testing.T) {

	// Setup test db
	db := setUpResourceTable()

	// Will be called at the end of test to clear up db
	defer tearDownResourceTable(db)

	//Create test records
	db.Create(&model.Resource{
		Type:       "TestType1",
		Department: "TestDepartment1",
	})

	db.Create(&model.Resource{
		Type:       "TestType2",
		Department: "TestDepartment2",
	})

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Get("/resource/:id", controller.ResourceDetail)

	// Send mock http request
	request := httptest.NewRequest("GET", "/resource/999", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	if err != nil {
		log.Println("Error with running mock app:", err)
	}

	assert.Equal(t, fiber.ErrBadRequest.Code, response.StatusCode)

	// Parse body from response to check data returned
	body, _ := io.ReadAll(response.Body)
	var responseBody struct {
		ResourceRecords model.Resource `json:"record"`
		Message         string         `json:"message"`
		StatusText      string         `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Bad Request", responseBody.StatusText)
	assert.Equal(t, "Record of Resource ID 999 not found.", responseBody.Message)

}

func TestResourceCreateSuccess(t *testing.T) {

	// Setup test db
	db := setUpResourceTable()

	// Will be called at the end of test to clear up db
	defer tearDownResourceTable(db)

	// Record to post
	recordToCreate := model.Resource{
		Type:       "TestType",
		Department: "TestDepartment",
	}

	// Create new App
	app := fiber.New()

	//Register route to test
	app.Post("/resource", controller.ResourceCreate)

	body, _ := json.Marshal(recordToCreate)

	req := httptest.NewRequest("POST", "/resource", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	response, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, response.StatusCode)

	// Parse body from response to check data returned
	body, _ = io.ReadAll(response.Body)
	var responseBody struct {
		ResourceRecords model.Resource `json:"record"`
		Message         string         `json:"message"`
		StatusText      string         `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "New Resource Record saved successfully.", responseBody.StatusText)

	// Check if the resource was actually saved to the database

	var savedResource model.Resource

	expectedRecord := model.Resource{
		ID:         1,
		Type:       "TestType",
		Department: "TestDepartment",
	}
	db.Find(&savedResource)
	assert.Equal(t, expectedRecord.ID, savedResource.ID)
	assert.Equal(t, expectedRecord.Type, savedResource.Type)
	assert.Equal(t, expectedRecord.Department, savedResource.Department)
}
func TestResourceUpdateSuccess(t *testing.T) {

	// Setup test db
	db := setUpResourceTable()

	// Will be called at the end of test to clear up db
	defer tearDownResourceTable(db)

	//Create test records
	db.Create(&model.Resource{
		Type:       "TestType1",
		Department: "TestDepartment1",
	})

	// Record to post
	recordToUpdate := model.Resource{
		Type:       "TestTypeUpdated",
		Department: "TestDepartment",
	}

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Put("/resource/:id", controller.ResourceUpdate)

	// Send mock http request
	body, _ := json.Marshal(recordToUpdate)

	request := httptest.NewRequest("PUT", "/resource/1", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert no errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	body, _ = io.ReadAll(response.Body)
	var responseBody struct {
		ResourceRecords model.Resource `json:"data"`
		Message         string         `json:"message"`
		StatusText      string         `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "Record updated successfully", responseBody.Message)
	assert.Equal(t, "Ok", responseBody.StatusText)
	assert.Equal(t, recordToUpdate.Type, responseBody.ResourceRecords.Type)
}

func TestResourceUpdateFail(t *testing.T) {

	// Setup test db
	db := setUpResourceTable()

	// Will be called at the end of test to clear up db
	defer tearDownResourceTable(db)

	//Create test records
	db.Create(&model.Resource{
		Type:       "TestType1",
		Department: "TestDepartment1",
	})

	// Record to post
	recordToUpdate := model.Resource{
		Type:       "TestTypeUpdated",
		Department: "TestDepartment",
	}

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Put("/resource/:id", controller.ResourceUpdate)

	// Send mock http request
	body, _ := json.Marshal(recordToUpdate)

	request := httptest.NewRequest("PUT", "/resource/999", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.ErrBadRequest.Code, response.StatusCode)

	body, _ = io.ReadAll(response.Body)
	var responseBody struct {
		Data       model.Resource `json:"data"`
		Message    string         `json:"message"`
		StatusText string         `json:"statusText"`
	}
	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "Record of ID 999 not found.", responseBody.Message)
	assert.Equal(t, "Ok.", responseBody.StatusText)
	//assert.Equal(t, "TestType1", responseBody.Data.Type)
}

func TestResourceDeleteSuccess(t *testing.T) {

	// Setup test db
	db := setUpResourceTable()

	// Will be called at the end of test to clear up db
	defer tearDownResourceTable(db)

	//Create test records
	db.Create(&model.Resource{
		Type:       "TestType1",
		Department: "TestDepartment1",
	})

	db.Create(&model.Resource{
		Type:       "TestType2",
		Department: "TestDepartment2",
	})

	var countOfRecordsAtStart int64
	result := db.Model(&model.Resource{}).Count(&countOfRecordsAtStart)
	if result.Error != nil {
		log.Println("Error in retrieving number of records.")
	}

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Delete("/resource/:id", controller.ResourceDelete)

	// Send mock http request
	request := httptest.NewRequest("DELETE", "/resource/1", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert no errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	// Parse body from response to check data returned
	body, _ := io.ReadAll(response.Body)
	var responseBody struct {
		Message    string `json:"message"`
		StatusText string `json:"statusText"`
	}

	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "Record deleted successfully.", responseBody.Message)
	assert.Equal(t, "Ok.", responseBody.StatusText)

	var countOfRecordsAtEnd int64
	result = db.Model(&model.Resource{}).Count(&countOfRecordsAtEnd)
	assert.NoError(t, result.Error)

	assert.Equal(t, countOfRecordsAtStart-1, countOfRecordsAtEnd)

}

func TestResourceDeleteFail(t *testing.T) {

	// Setup test db
	db := setUpResourceTable()

	// Will be called at the end of test to clear up db
	defer tearDownResourceTable(db)

	//Create test records
	db.Create(&model.Resource{
		Type:       "TestType1",
		Department: "TestDepartment1",
	})

	db.Create(&model.Resource{
		Type:       "TestType2",
		Department: "TestDepartment2",
	})

	var countOfRecordsAtStart int64
	result := db.Model(&model.Resource{}).Count(&countOfRecordsAtStart)
	if result.Error != nil {
		log.Println("Error in retrieving number of records.")
	}

	// Create new App
	app := fiber.New()

	// Register route to test
	app.Delete("/resource/:id", controller.ResourceDelete)

	// Send mock http request
	request := httptest.NewRequest("DELETE", "/resource/999", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)

	// Assert no errors and status codes are equal
	assert.NoError(t, err)
	assert.Equal(t, fiber.ErrBadRequest.Code, response.StatusCode)

	// Parse body from response to check data returned
	body, _ := io.ReadAll(response.Body)
	var responseBody struct {
		Message    string `json:"message"`
		StatusText string `json:"statusText"`
	}

	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "Record of Resource ID 999 not found.", responseBody.Message)
	assert.Equal(t, "Ok.", responseBody.StatusText)

	var countOfRecordsAtEnd int64
	result = db.Model(&model.Resource{}).Count(&countOfRecordsAtEnd)
	assert.NoError(t, result.Error)

	assert.Equal(t, countOfRecordsAtStart, countOfRecordsAtEnd)

}
