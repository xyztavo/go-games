package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gustafer/go-games/cmd/api/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Simple CRUD Fiber API about games using POSTGRESQL", "health": "ok"})
	})
	app.Post("/game", handlers.CreateGame)       // Create a game
	app.Get("/game/:id", handlers.GetGame)       // Get a game by ID
	app.Get("/games", handlers.GetGames)         // Lists All Games
	app.Put("/game/:id", handlers.UpdateGame)    // Update a game by ID
	app.Delete("/game/:id", handlers.DeleteGame) // Delete a game by ID
}
