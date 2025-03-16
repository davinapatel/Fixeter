// The package controller is concerned with all operations for handler functions.
// It  provides handler functions for managing the Issue model.

package controller

import (
	"log"
	"os"
	"path/filepath"

	"github.com/davinapatel/Fixeter/database"
	"github.com/davinapatel/Fixeter/model"
	"github.com/gofiber/fiber/v2"
)

// IssueList handles the HTTP Request to retrieve a list of all Issue records from the DB.
// Returns a JSON response containing the records or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON
func IssueList(c *fiber.Ctx) error {

	// Setting initial JSON Response
	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Issue List",
	}

	// A slice of Issues, where result of DB query will be stored
	var records []model.Issue

	err := database.GetRecords(&records) // Retrieves records from Issue table and loads into records

	if err != nil {
		log.Println("Failed to get Issue Records")
	} else {
		log.Println("Successfully Retrieved Issue Records")
	}

	// Include retrieved issue records in HTTP Response
	context["issue_records"] = records

	// Set status code and return JSON
	c.Status(200)
	return c.JSON(context)

}

// IssueListStatus handles the HTTP Request to retrieve a list of all Issue Statuses and the number of issues with
// each status from the DB.
// Returns a JSON response containing the records or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON
func IssueListStatus(c *fiber.Ctx) error {
	// Set initial statusText and message
	context := fiber.Map{

		"statusText": "Ok",
		"message":    "Issue Status Detail",
	}

	// Define struct to store results from DB Query in
	var results []struct {
		Status string
		Count  int64
	}

	err := database.GetIssueStatusCount(&results) // Retrieve count of issues per status from DB and load into results struct
	if err != nil {
		log.Println("Error in getting Issue List Status")
	} else {
		log.Println("Successfully Retrieve Issue List Status")
	}
	// Update HTTP Response with retrieved records in
	context["issueStatus_records"] = results

	// Set status code and return JSON
	c.Status(200)
	return c.JSON(context)
}

// IssueDetail handles the HTTP Request to retrieve a single record from the Issue table based on the id passed
// as a parameter with the request
// Returns a JSON response containing the specific record or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON
func IssueDetail(c *fiber.Ctx) error {

	// Sets intial statusText and message to be included in HTTP Response
	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Issue Detail",
	}

	id := c.Params("id") // Get id from route

	// record is an instance of the struct Issue, where results of DB Query will be stored
	var record model.Issue

	err := database.GetRecordByID(&record, id) // Retrieves Issue Record from querying on id, and result in loaded into record var

	if err != nil {
		log.Println("Failed to get Issue Record of ID", id)
	} else {
		log.Println("Successfully Retrieved Issue Record of ID", id)
	}

	// Checks record exists
	if record.ID == 0 {
		log.Println("Record of Issue ID", id, "not found.")
		context["statusText"] = "Bad Request" // Set failed statusText and messages for HTTP Response
		context["message"] = "Record of Issue ID " + id + " not found."
		c.Status(400)
		return c.JSON(context)
	}

	// Include retrieved record in HTTP Response and set status code
	context["record"] = record
	c.Status(200)
	return c.JSON(context)

}

// IssueCreate handles the HTTP Request to create a new Issue record
// Returns a JSON response containing the record created  or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON
func IssueCreate(c *fiber.Ctx) error {

	// Set initial statusText and message for HTTP Response
	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Create an Issue",
	}

	// Create an instance of Issue struct which data from request can be loaded into
	record := new(model.Issue)

	// Parses the body from the POST Request into record var above
	if err := c.BodyParser(record); err != nil {
		log.Println(err)
		log.Println("Error in parsing request.")
		context["statusText"] = "Bad Request"
		context["message"] = "Parsing Request Failed."
		c.Status(400)
	}

	//File upload
	file, err := c.FormFile("file")

	// Define the upload directory for static uploads
	uploadDir := "./static/uploads/"

	// Dynamically create the folder if it doesn't exist
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Println("Error creating upload directory:", err)
		return c.Status(500).SendString("Failed to create upload directory")
	}

	// Build the file path
	filePath := filepath.Join(uploadDir, file.Filename)
	if err != nil {
		log.Println("Error in file upload.", err)
	}

	// Check that file isn't empty
	if file.Size > 0 {
		if err := c.SaveFile(file, filePath); err != nil {
			log.Println("Error in file uploading...", err)
		}

		//Set image path to the struct
		record.Image = filePath
	}

	// Save data in the DB
	err = database.SaveRecord(&record)

	if err != nil {
		log.Println("Error in saving data for new Issue.")
		context["statusText"] = "Bad Request"
		context["message"] = "Saving New Issue Failed."
		c.Status(400)
	} else {
		log.Println("Successfully saved new Issue")
	}

	//Update issue history count
	category := record.Category
	IssueHistoryUpdate(category)

	// Inclyude success statusText and new issue record that has been created in HTTP Response
	context["statusText"] = "New Issue Record saved successfully."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

// IssueUpdate handles the HTTP Request to update a single record from the Issue table based on the id passed
// as a parameter with the request
// Returns a JSON response containing the specific record that has been updated or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON
func IssueUpdate(c *fiber.Ctx) error {

	// Set initial statusText and message for HTTP Response
	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Update Issue",
	}

	// http:localhost:8000/issue/2 - retrieve the id at the end of the route
	id := c.Params("id")

	var record model.Issue

	// Finds the record in the DB that matches that given ID
	// And populates into the Issue struct record

	//database.DBConn.First(&record, id)
	err := database.GetRecordByID(&record, id)

	if err != nil {
		log.Println("Failed to get Issue Record of ID", id)
	} else {
		log.Println("Successfully Retrieved Issue Record of ID", id)
	}

	// If issue ID does not exist in DB
	if record.ID == 0 {
		log.Println("Record of ID", id, "not found.")
		context["statusText"] = "Bad Request"
		context["message"] = "Record of ID " + id + " not found."
		c.Status(400)
		return c.JSON(context)
	}

	// Parse the data of the updated record from HTTP Request into record which is a struct of Issue
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
	}

	err = database.SaveRecord(&record) // Save uodated record to DB

	if err != nil {
		log.Println("Error in updating Issue of ID:", id)
	} else {
		log.Println("Successfully updated Issue of ID:", id)
	}

	// Update HTTP Response message with success message
	// Update HTTP Response data with updated record data
	context["message"] = "Record updated successfully"
	context["data"] = record
	c.Status(200)
	return c.JSON(context)

}

// IssueDelete handles the HTTP Request to delete a single record from the Issue table based on the id passed
// as a parameter with the request
// Returns a JSON response containing the specific record or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON
func IssueDelete(c *fiber.Ctx) error {

	// Set initial statusText and message for HTTP Response
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"message":    "",
	}

	id := c.Params("id") // Get id from route

	var record model.Issue

	err := database.GetRecordByID(&record, id) // Get record from Issue table using id and load into record

	if err != nil {
		log.Println("Failed to get Issue Record of ID", id)
	} else {
		log.Println("Successfully Retrieved Issue Record of ID", id)
	}

	// Check if record exists
	if record.ID == 0 {
		log.Println("Record of ID", id, "not found.")
		context["message"] = "Record of ID " + id + " not found."
		return c.JSON(context)
	}

	err = database.DeleteRecord(&record) // Delete record from db

	if err != nil {
		log.Println("Failed to delete Issue Record of ID:", id)
		context["message"] = "Failure to delete Issue from Database."
		return c.JSON(context)
	} else {
		log.Println("Successfully deleted Issue Record of ID:", id)
	}

	// Set success message and statusText
	context["message"] = "Record deleted successfully."
	context["statusText"] = "Ok."
	c.Status(200)
	return c.JSON(context)
}
