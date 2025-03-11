// The package controller is concerned with all operations for handler functions.
// It  provides handler functions for managing the Resource model.
package controller

import (
	"log"

	"github.com/davinapatel/Fixeter/database"
	"github.com/davinapatel/Fixeter/model"
	"github.com/gofiber/fiber/v2"
)

// ResourceList handles the HTTP Request to retrieve a list of all Resource records from the DB.
// Returns a JSON response containing the records or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON
func ResourceList(c *fiber.Ctx) error {

	// Set intial statusText and message for response
	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Resource List",
	}

	var records []model.Resource

	// Retrieves records from Resource table and loads into var records
	err := database.GetRecords(&records)
	if err != nil {
		log.Println("Failed to get Resource Records")
	} else {
		log.Println("Successfully Retrieved Resource Records")
	}

	// Include retrieved records in HTTP Response
	context["resource_records"] = records

	// Set status code and return JSON
	c.Status(200)
	return c.JSON(context)
}

// ResourceDetail handles the HTTP Request to retrieve a single record from the Resource table based on the id passed
// as a parameter with the request
// Returns a JSON response containing the specific record or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON
func ResourceDetail(c *fiber.Ctx) error {

	// Set inital statusText and message for HTTP Response
	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Resource Detail",
	}

	id := c.Params("id") // Retrieve id of resource to retrieve from route

	var record model.Resource

	err := database.GetRecordByID(&record, id) // Retrieve record from DB using id and load into var record aboce

	if err != nil {
		log.Println("Failed to get Resource Record of ID", id)
	} else {
		log.Println("Successfully Retrieved Resource Record of ID", id)
	}

	// Checks if the record exists, if not sets error statusText and message
	if record.ID == 0 {
		log.Println("Record of ID", id, "not found.")
		context["statusText"] = "Bad Request"
		context["message"] = "Record of Resource ID " + id + " not found."
		c.Status(400)
		return c.JSON(context)
	}

	// Send retrieved record as part of HTTP Response
	context["record"] = record

	// Set status code and return JSON
	c.Status(200)
	return c.JSON(context)

}

// ResourceCreate handles the HTTP Request to create a new resource record
// Returns a JSON response containing the record created  or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON
func ResourceCreate(c *fiber.Ctx) error {

	// Set initial statusText and message
	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Create a Resource",
	}

	record := new(model.Resource)

	// Parses the body from the POST Request into the record struct above of the type Resource
	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = "Bad Request"
		context["message"] = "Parsing Request Failed."
		c.Status(400)
	}

	// Save data in the DB
	err := database.SaveRecord(&record)

	if err != nil {
		log.Println("Error in saving data for a Resource.")
		context["statusText"] = "Bad Request"
		context["message"] = "Saving New Resource Failed."
		c.Status(400)
	}

	// Set success statusText and include record that was just saved in response
	context["statusText"] = "New Resource Record saved successfully."
	context["data"] = record

	// Set status code and return JSON
	c.Status(201)
	return c.JSON(context)
}

// ResourceUpdate handles the HTTP Request to update a single record from the Resource table based on the id passed
// as a parameter with the request
// Returns a JSON response containing the specific record that has been updated or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON
func ResourceUpdate(c *fiber.Ctx) error {

	// Set initial statusText and message to be included in HTTP Response
	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Update Resource",
	}

	// http:localhost:8000/issue/2 - retrieves the id at the end of the route
	id := c.Params("id")

	var record model.Resource

	// Finds the record in the DB that matches that given ID
	// And populates into the Resource struct record

	err := database.GetRecordByID(&record, id)

	if err != nil {
		log.Println("Failed to get Resource Record of ID", id)
	} else {
		log.Println("Successfully Retrieved Resource Record of ID", id)
	}

	// If issue ID does not exist in DB
	if record.ID == 0 {
		log.Println("Record of ID", id, "not found.")
		context["statusText"] = "Bad Request"
		context["message"] = "Record of ID " + id + " not found."
		c.Status(400)
		return c.JSON(context)
	}

	// Parses the record to be updated (with updated data) from request into record struct
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
	}

	err = database.SaveRecord(&record) // Saves updated record to DB

	if err != nil {
		log.Println("Error in updating Resource of ID:", id)
	} else {
		log.Println("Successfully updated Resource of ID:", id)
	}

	// Updates message and data with updated record and returns JSON
	context["message"] = "Record updated successfully"
	context["data"] = record
	c.Status(200)
	return c.JSON(context)

}

// ResourceDelete handles the HTTP Request to delete a single record from the Resource table based on the id passed
// as a parameter with the request
// Returns a JSON response containing the specific record or an appropriate error message.
//
// Parameters:
// - c: A pointer to fiber.Ctx, a struct which contains information about the HTTP Request and Response
//
// Returns:
// - error: Returns an appropriate error when returning JSON
func ResourceDelete(c *fiber.Ctx) error {

	// Set intial status. statusText and message to be included in HTTP Response
	c.Status(400)
	context := fiber.Map{
		"statusText": "Ok.",
		"message":    "Delete Resource.",
	}

	id := c.Params("id") // Retrieves if of record to delete from route

	var record model.Resource

	err := database.GetRecordByID(&record, id) // Retrieves record to be deleted from id and loads data into var record

	if err != nil {
		log.Println("Failed to get Resource Record of ID", id)
	} else {
		log.Println("Successfully Retrieved Resource Record of ID", id)
	}

	// Checks if record exists
	if record.ID == 0 {
		log.Println("Record of Resource ID", id, "not found.")
		context["message"] = "Record of Resource ID " + id + " not found."
		return c.JSON(context)
	}

	err = database.DeleteRecord(&record) // Deletes record from Resource table

	if err != nil {
		log.Println("Failed to delete Issue Record of ID:", id)
		context["message"] = "Failure to delete Resource from Database."
		return c.JSON(context)
	} else {
		log.Println("Successfully deleted Issue Record of ID:", id)
	}

	// Include success message and statusText in HTTP Response
	// Return JSON
	context["message"] = "Record deleted successfully."
	context["statusText"] = "Ok."
	c.Status(200)
	return c.JSON(context)
}
