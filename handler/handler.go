package handler

import "github.com/gofiber/fiber/v2"

// GetIndex returns index
func GetIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello World",
	}, "layout/main")
}

// GetHi returns Hi
func GetHi(c *fiber.Ctx) error {
	return c.Render("hi", fiber.Map{
		"Name": "Glup",
	}, "layout/main")
}
