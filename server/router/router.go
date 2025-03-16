package router

import (
	"github.com/davinapatel/Fixeter/controller"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all the API routes for the GoFiber Application
// Each route is associated with a corresponding controller function which handles the request

// Parameters:
// - app: A pointer to the GoFiber Appplication instance where the routes are registered
func SetupRoutes(app *fiber.App) {
	// Setup Issue Routes
	app.Get("/issue", controller.IssueList)
	app.Get("/issue/:id", controller.IssueDetail)
	app.Get("/issueStatus", controller.IssueListStatus)
	app.Post("/issue", controller.IssueCreate)
	app.Put("/issue/:id", controller.IssueUpdate)
	app.Delete("/issue/:id", controller.IssueDelete)

	// Setup Resource Routes
	app.Get("/resource", controller.ResourceList)
	app.Get("/resource/:id", controller.ResourceDetail)
	app.Post("/resource", controller.ResourceCreate)
	app.Put("/resource/:id", controller.ResourceUpdate)
	app.Delete("/resource/:id", controller.ResourceDelete)

	// Setup IssueHistory Routes
	app.Get("/issuehistory", controller.IssueHistoryList)
}
