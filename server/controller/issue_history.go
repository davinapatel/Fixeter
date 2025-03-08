package controller

import (
	"log"

	"github.com/davinapatel/Fixeter/database"
	"github.com/davinapatel/Fixeter/model"
	"github.com/gofiber/fiber/v2"
)

func IssueHistoryList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"message":    "IssueHistory List",
	}

	var records []model.IssueHistory

	err := database.GetRecords(&records)
	if err != nil {
		log.Println("Failed to get IssueHistory Records")
	} else {
		log.Println("Successfully Retrieved IssueHistory Records")
	}

	context["issueHistory_records"] = records

	c.Status(200)
	return c.JSON(context)
}

func IssueHistoryUpdate(category string) {

	var record model.IssueHistory

	err := database.GetRecordByCategory(&record, category)
	if err != nil {
		log.Println("Failed to get IssueHistory Record of Category", category)
	} else {
		log.Println("Successfully Retrieved IssueHistory Record of Category", category)
	}

	if record.ID == 0 {
		log.Println("Record with a Category of", category, "not found.")
	}

	count := record.Count + 1
	record.Count = count

	err = database.SaveRecord(&record)

	if err != nil {
		log.Println("Error in updating Issue Count for Category:", category)
	} else {
		log.Println("Successfully updated Issue Count for Category:", category)
	}
}
