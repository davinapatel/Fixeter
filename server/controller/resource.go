package controller

import (
	"log"

	"github.com/davinapatel/Fixeter/database"
	"github.com/davinapatel/Fixeter/model"
	"github.com/gofiber/fiber/v2"
)

func ResourceList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Resource List",
	}

	var records []model.Resource

	err := database.GetRecords(&records)
	if err != nil {
		log.Println("Failed to get Resource Records")
	} else {
		log.Println("Successfully Retrieved Resource Records")
	}

	context["resource_records"] = records

	c.Status(200)
	return c.JSON(context)
}

func ResourceDetail(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Resource Detail",
	}

	id := c.Params("id")

	var record model.Resource

	err := database.GetRecordByID(&record, id)

	if err != nil {
		log.Println("Failed to get Resource Record of ID", id)
	} else {
		log.Println("Successfully Retrieved Resource Record of ID", id)
	}

	if record.ID == 0 {
		log.Println("Record of ID", id, "not found.")
		context["statusText"] = "Bad Request"
		context["message"] = "Record of Resource ID " + id + " not found."
		c.Status(400)
		return c.JSON(context)
	}

	context["record"] = record
	c.Status(200)
	return c.JSON(context)

}

func ResourceCreate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Create a Resource",
	}

	record := new(model.Resource)

	// Parses the body from the POST Request
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

	context["statusText"] = "New Resource Record saved successfully."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

func ResourceUpdate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"message":    "Update Resource",
	}

	// http:localhost:8000/issue/2
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

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
	}

	err = database.SaveRecord(&record)

	if err != nil {
		log.Println("Error in updating Resource of ID:", id)
	} else {
		log.Println("Successfully updated Resource of ID:", id)
	}

	context["message"] = "Record updated successfully"
	context["data"] = record
	c.Status(200)
	return c.JSON(context)

}

func ResourceDelete(c *fiber.Ctx) error {

	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"message":    "",
	}

	id := c.Params("id")

	var record model.Resource

	err := database.GetRecordByID(&record, id)

	if err != nil {
		log.Println("Failed to get Resource Record of ID", id)
	} else {
		log.Println("Successfully Retrieved Resource Record of ID", id)
	}

	if record.ID == 0 {
		log.Println("Record of Resource ID", id, "not found.")
		context["message"] = "Record of Resource ID " + id + " not found."
		return c.JSON(context)
	}

	err = database.DeleteRecord(&record)

	if err != nil {
		log.Println("Failed to delete Issue Record of ID:", id)
		context["message"] = "Failure to delete Resource from Database."
		return c.JSON(context)
	} else {
		log.Println("Successfully deleted Issue Record of ID:", id)
	}

	context["message"] = "Record deleted successfully."
	context["statusText"] = "Ok."
	c.Status(200)
	return c.JSON(context)
}
