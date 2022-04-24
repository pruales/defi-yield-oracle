package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pruales/defi-yield-oracle/api/handlers"
)

func TaskRouter(app fiber.Router) {
	app.Post("/task", handlers.TaskHandler())
}
