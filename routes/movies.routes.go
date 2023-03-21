package routes

import "github.com/gofiber/fiber/v2"

type Movie struct {
	Title string `json:"title"`
	Id    int    `json:"id"`
}

func UseMoviesRoutes(router fiber.Router) {

	movies := []Movie{
		{
			Title: "Buscando a Nemo",
			Id:    1,
		},
		{
			Title: "Madagascar",
			Id:    2,
		},
	}
	router.Get("/movies", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"movies": movies,
		})
	})
}
