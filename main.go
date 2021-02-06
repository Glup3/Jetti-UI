package main

import (
	"github.com/glup3/jetti-ui/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main() {
	engine := handlebars.New("./view", ".hbs")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.SetupRoutes(app)

	app.Listen(":3000")
}
