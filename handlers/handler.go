package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func HandleRoot(c fiber.Ctx) error {
	return c.SendString("Welcome to /")
}

func HandleSearch(c fiber.Ctx) error {
	query := c.Query("query")

	if query != "" {
		return c.SendString("Found query")
	}
	return c.Render("index", fiber.Map{
		"Title": "Hello, world!",
	})
}
