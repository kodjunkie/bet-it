package main

import (
	"elivate9ja-go/configs"
	"elivate9ja-go/routes"
	"elivate9ja-go/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load env values
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading env values")
	}

	fiberConfig := configs.FiberConfig()

	// create fiber app
	app := fiber.New(fiberConfig)

	// routes
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "good"})
	})
	routes.GetAPIRoutes(app)

	utils.StartServer(app)
}
