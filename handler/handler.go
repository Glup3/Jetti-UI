package handler

import (
	"github.com/glup3/jetti-ui/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetIndex returns index
func GetIndex(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var matches []model.Match

		db.
			Preload("Team1.Player1").
			Preload("Team1.Player2").
			Preload("Team1.Player3").
			Preload("Team1.Player4").
			Preload("Team1.Player5").
			Preload("Team2.Player1").
			Preload("Team2.Player2").
			Preload("Team2.Player3").
			Preload("Team2.Player4").
			Preload("Team2.Player5").
			Order("ID desc").
			Find(&matches)

		return c.Render("index", fiber.Map{
			"Title":   "Jetti UI",
			"matches": matches,
		}, "layout/main")
	}
}
