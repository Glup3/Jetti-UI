package router

import (
	"github.com/glup3/jetti-ui/handler"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes setup routes
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/", handler.GetIndex(db))
}
