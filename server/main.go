package main

import (
	"log"

	"github.com/davinapatel/Fixeter/database"
	"github.com/davinapatel/Fixeter/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading .env file.")
	}
	database.ConnectDB()

}

func main() {

	sqlDB, err := database.DBConn.DB()

	if err != nil {
		panic("Error in SQL Connection.")
	}

	defer sqlDB.Close()

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":8000")

}
