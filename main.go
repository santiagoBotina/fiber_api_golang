package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santiagoBotina/fiber_api_golang/routes"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "OK",
		})
	})

	routes.UseMoviesRoutes(app)

	app.Listen(":3000")
}
