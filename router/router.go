package router

import (
	"github.com/glup3/jetti-ui/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup routes
func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.GetIndex)
	app.Get("/hi", handler.GetHi)
}
