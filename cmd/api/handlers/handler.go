package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gustafer/go-games/cmd/api/models"
)

func CreateGame(c fiber.Ctx) error {
	game := new(models.Game)

	if err := c.Bind().Body(game); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "game is format (title, description)."})
	}

	gameId, err := models.InsertGame(game)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not create game."})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "game created with ease.", "gameId": gameId})
}

func GetGame(c fiber.Ctx) error {
	id := c.Params("id")
	games, err := models.GameById(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("could not find the game with id: %v.", id)})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"games": games})
}

func DeleteGame(c fiber.Ctx) error {
	id := c.Params("id")
	rowsAffected, err := models.DeleteGame(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "something happened."})
	}

	if rowsAffected < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("no game with id: %v were found", id)})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "game deleted with ease.", "rowsAffected": rowsAffected})
}
func UpdateGame(c fiber.Ctx) error {
	game := new(models.Game)

	if err := c.Bind().Body(game); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "something happened."})
	}

	id := c.Params("id")

	updatedGameId, err := models.PatchGame(id, game)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("game with id: %v not found", id)})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "game updated with ease.", "updatedGameId": updatedGameId})
}

func GetGames(c fiber.Ctx) error {
	games, err := models.QueryAllGames()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "no games found"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"games": games})
}
