package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gustafer/go-games/cmd/api/configs"
	"github.com/gustafer/go-games/cmd/api/database"
	"github.com/gustafer/go-games/cmd/api/router"
)

func main() {
	if err := database.AutoMigrate(); err != nil {
		panic(err)
	}

	app := fiber.New()

	router.SetupRoutes(app)

	port := configs.GetPort()
	log.Fatal(app.Listen(port)) // Default port is :4040, change it in ./configs/.env
}
