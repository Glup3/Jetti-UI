package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetIndex returns index
func GetIndex(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Render("index", fiber.Map{
			"Title": "Jetti UI",
		}, "layout/main")
	}
}
