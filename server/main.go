package main

import (
	"log"

	"github.com/davinapatel/Fixeter/database"
	"github.com/davinapatel/Fixeter/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// Init is called at the very start of GoFiber instance running, it runs before the 'main' function
// This function loads in environment variables and calls the method to connect to
// the PostgreSQL Database
func init() {

	// Load env vars
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading .env file.")
	}

	database.ConnectDB() // Connect to PostgreSQL DB

}

// The main function is called when you run "go main.go" and starts the application
func main() {

	// Configure logger with flags and short file names
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Connects to PostgreSQL database
	sqlDB, err := database.DBConn.DB()

	if err != nil {
		panic("Error in SQL Connection.")
	}

	// Closes DB at the end of the main function executing
	defer sqlDB.Close()

	// Initlialise new GoFiber instance
	app := fiber.New()

	// Used to serve static files
	app.Static("/static", "./static")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Calls SetupRoutes function to set up API Routes
	router.SetupRoutes(app)

	// GoFiber application to listen to port 8000
	app.Listen(":8000")

}
