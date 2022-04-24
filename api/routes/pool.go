package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pruales/defi-yield-oracle/api/handlers"
)

func PoolRouter(app fiber.Router) {
	app.Get("/pools", handlers.GetPools())
}
