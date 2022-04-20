package main

import (
	"SitesBackend/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		BodyLimit:     1024 * 1024 * 1024,
	})
	app.Use(cors.New())

	route.InitRoutes(app)

	app.Listen(":8091")
}
