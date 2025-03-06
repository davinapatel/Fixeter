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

	db := database.DBConn

	var records []model.IssueHistory

	db.Find(&records)

	context["issueHistory_records"] = records

	c.Status(200)
	return c.JSON(context)
}

func IssueHistoryUpdate(category string) {

	var record model.IssueHistory

	database.DBConn.First(&record, "category = ?", category)

	if record.ID == 0 {
		log.Println("Record with a Category of", category, "not found.")
	}

	count := record.Count + 1
	record.Count = count

	result := database.DBConn.Save(record)

	if result.Error != nil {
		log.Println("Error in updating Issue Count for Category:", category)
	}
}
