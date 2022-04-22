package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main () {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Println("Listening on port 3000")
	log.Fatal(app.Listen(":3000"))
}
