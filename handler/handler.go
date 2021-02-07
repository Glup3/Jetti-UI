package handler

import "github.com/gofiber/fiber/v2"

// GetIndex returns index
func GetIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello",
	}, "layout/main")
}

// GetHi returns Hi
func GetHi(c *fiber.Ctx) error {
	return c.Render("hi", fiber.Map{
		"Title": "Jetti UI",
		"Name":  "Glup",
	}, "layout/main")
}
