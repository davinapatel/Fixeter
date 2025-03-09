package main

import (
	"log"

	"github.com/davinapatel/Fixeter/database"
	"github.com/davinapatel/Fixeter/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading .env file.")
	}
	database.ConnectDB()

}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	sqlDB, err := database.DBConn.DB()

	if err != nil {
		panic("Error in SQL Connection.")
	}

	defer sqlDB.Close()

	app := fiber.New()

	app.Static("/static", "./static")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.SetupRoutes(app)

	app.Listen(":8000")

}
