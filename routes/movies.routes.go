package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Movie struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	Year     int                `json:"year"`
	Duration int                `json:"duration"`
	Director string             `json:"director"`
}

func UseMoviesRoutes(router fiber.Router, dbInstance *mongo.Database) {
	db := dbInstance.Collection("movies")
	router.Get("/movies", func(c *fiber.Ctx) error {
		cursor, err := db.Find(c.Context(), bson.M{})

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"success": false,
				"message": "Error fetching data",
				"error":   err.Error(),
			})
		}

		var movies []Movie = make([]Movie, 0)

		if err := cursor.All(c.Context(), &movies); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"success": false,
				"message": "Error decoding data",
				"error":   err.Error(),
			})

		}

		return c.JSON(fiber.Map{
			"success": true,
			"data":    movies,
		})
	})
}
