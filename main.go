package main

import (
	"github.com/glup3/jetti-ui/config"
	"github.com/glup3/jetti-ui/database"
	"github.com/glup3/jetti-ui/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main() {
	engine := handlebars.New("./view", ".hbs")
	port := config.Get("APP_PORT")

	db := database.Connect()

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.SetupRoutes(app, db)

	app.Listen(port)
}
