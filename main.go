package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/santiagoBotina/fiber_api_golang/db"
	"github.com/santiagoBotina/fiber_api_golang/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	db.Initialize()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "OK",
		})
	})

	routes.UseMoviesRoutes(app, db.DB)

	app.Listen(":3000")
}
