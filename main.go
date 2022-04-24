package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/pruales/defi-yield-oracle/api/routes"
)

func main () {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			c.Status(code).SendString(err.Error())

			return nil
		},
	})
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	api := app.Group("/api")
	routes.TaskRouter(api)
	routes.PoolRouter(api)

	log.Println("Listening on port", port)
	log.Fatal(app.Listen(":" + port))
}
