package main

import (
	"kw/database"
	"kw/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	err := database.DatabaseConnection()
	if err != nil {
		log.Fatal("Unable to Connext")
	}

	app := fiber.New()
	app.Use(cors.New())
	router.SetupRouter(app)

	log.Fatal(app.Listen(":3000"))

}
