// The package controller is concerned with all operations for handler functions.
// It  provides handler functions for managing the IssueHistory model.

package controller

import (
	"log"

	"github.com/davinapatel/Fixeter/database"
	"github.com/davinapatel/Fixeter/model"
	"github.com/gofiber/fiber/v2"
)

// IssueHistoryList handles the HTTP Request to retrieve a list of all IssueHistory records from the DB.
// Returns a JSON response containing the records or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON

func IssueHistoryList(c *fiber.Ctx) error {

	// Set intial statusText and message for response
	context := fiber.Map{
		"statusText": "Ok",
		"message":    "IssueHistory List",
	}

	var records []model.IssueHistory

	// Retrieves records from IssueHistory table and loads into var records
	err := database.GetRecords(&records)
	if err != nil {
		log.Println("Failed to get IssueHistory Records")
	} else {
		log.Println("Successfully Retrieved IssueHistory Records")
	}

	// Include retrieved records in HTTP Response
	context["issueHistory_records"] = records

	// Set status code and return JSON
	c.Status(200)
	return c.JSON(context)
}

// IssueHistoryUpdate handles updating the count by one for a Category in the IssueHistory table when a new Issue is created.
// This function will be called when the handler for IssueCreate is executing and it updates the count of the Category
// for which the new Issue was created for.
//
// Parameters:
// - category: A string of the category to increment the count for

func IssueHistoryUpdate(category string) {

	var record model.IssueHistory

	err := database.GetRecordByCategory(&record, category) // Retrieves record which has the specified category in params
	if err != nil {
		log.Println("Failed to get IssueHistory Record of Category", category)
		log.Println(err)
	} else {
		log.Println("Successfully Retrieved IssueHistory Record of Category", category)
	}

	// Checks if record exists
	if record.ID == 0 {
		log.Println("Record with a Category of", category, "not found.")
		log.Println(err)
	}

	count := record.Count + 1 // Increment count by one
	record.Count = count      // Update the records' count field to new count

	err = database.SaveRecord(&record)

	if err != nil {
		log.Println("Error in updating Issue Count for Category:", category)
	} else {
		log.Println("Successfully updated Issue Count for Category:", category)
	}
}
