package main

import (
	"beerscraper/handlers"
	"beerscraper/repos"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "BeerScraper",
		ServerHeader: "Fiber",
	})

	beerRepo := repos.NewBeerRepo()

	server := app.Group("/api")

	handlers.NewBeerSearchHandler(server.Group("/beer"), beerRepo)

	app.Listen(":8080")
}
