package router

import (
	"github.com/davinapatel/Fixeter/controller"
	"github.com/gofiber/fiber/v2"
)

// Setup routes
func SetupRoutes(app *fiber.App) {
	app.Get("/issue", controller.IssueList)
	app.Get("/issue/:id", controller.IssueDetail)
	app.Get("/issueStatus", controller.IssueListStatus)
	app.Post("/issue", controller.IssueCreate)
	app.Put("/issue/:id", controller.IssueUpdate)
	app.Delete("/issue/:id", controller.IssueDelete)

	app.Get("/resource", controller.ResourceList)
	app.Get("/resource/:id", controller.ResourceDetail)
	app.Post("/resource", controller.ResourceCreate)
	app.Put("/resource/:id", controller.ResourceUpdate)
	app.Delete("/resource/:id", controller.ResourceDelete)

	app.Get("/issuehistory", controller.IssueHistoryList)
}
