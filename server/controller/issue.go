package controller

import (
	"log"
	"os"
	"path/filepath"

	"github.com/davinapatel/Fixeter/database"
	"github.com/davinapatel/Fixeter/model"
	"github.com/gofiber/fiber/v2"
)

func IssueList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Issue List",
	}

	db := database.DBConn

	var records []model.Issue

	db.Find(&records)

	context["issue_records"] = records

	c.Status(200)
	return c.JSON(context)

}

func IssueDetail(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Issue Detail",
	}

	id := c.Params("id")

	var record model.Issue

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record of ID", id, "not found.")
		context["statusText"] = "Bad Request"
		context["message"] = "Record of ID " + id + " not found."
		c.Status(400)
		return c.JSON(context)
	}

	context["record"] = record
	c.Status(200)
	return c.JSON(context)

}

func IssueCreate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Create an Issue",
	}

	record := new(model.Issue)

	// Parses the body from the POST Request
	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = "Bad Request"
		context["message"] = "Parsing Request Failed."
		c.Status(400)
	}

	//File upload
	file, err := c.FormFile("file")

	// Define the upload directory
	uploadDir := "./static/uploads"

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

	if file.Size > 0 {
		if err := c.SaveFile(file, filePath); err != nil {
			log.Println("Error in file uploading...", err)
		}

		//Set image path to the struct
		record.Image = filePath
	}

	// Save data in the DB
	result := database.DBConn.Create(record)

	if result.Error != nil {
		log.Println("Error in saving data for N Issue.")
		context["statusText"] = "Bad Request"
		context["message"] = "Saving New Issue Failed."
		c.Status(400)
	}

	context["statusText"] = "New Issue Record saved successfully."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

func IssueUpdate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Update Issue",
	}

	// http:localhost:8000/issue/2
	id := c.Params("id")

	var record model.Issue

	// Finds the record in the DB that matches that given ID
	// And populates into the Issue struct record

	database.DBConn.First(&record, id)

	// If issue ID does not exist in DB
	if record.ID == 0 {
		log.Println("Record of ID", id, "not found.")
		context["statusText"] = "Bad Request"
		context["message"] = "Record of ID " + id + " not found."
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
	}

	result := database.DBConn.Save(record)

	if result.Error != nil {
		log.Println("Error in saving data.")
	}

	context["message"] = "Record updated successfully"
	context["data"] = record
	c.Status(200)
	return c.JSON(context)

}

func IssueDelete(c *fiber.Ctx) error {

	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"message":    "",
	}

	id := c.Params("id")

	var record model.Issue

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record of ID", id, "not found.")
		context["message"] = "Record of ID " + id + " not found."
		return c.JSON(context)
	}

	result := database.DBConn.Delete(record)

	if result.Error != nil {
		context["message"] = "Failure to delete Issue from Database."
		return c.JSON(context)

	}

	context["message"] = "Record deleted successfully."
	context["statusText"] = "Ok."
	c.Status(200)
	return c.JSON(context)
}
