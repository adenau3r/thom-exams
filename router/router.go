package router

import (
	"log"

	"github.com/adenau3r/thom-exams/handlers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
)

func Run() {
	engine := html.New("../views", ".html")

	router := fiber.New(fiber.Config{
		Immutable: true,
		Views:     engine,
	})

	router.Get("/", handlers.HandleRoot)
	router.Get("/search", handlers.HandleSearch)

	log.Fatal(router.Listen(":3000"))
}
