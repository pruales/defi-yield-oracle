package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pruales/defi-yield-oracle/data"
	"github.com/pruales/defi-yield-oracle/fetchers"
	"github.com/pruales/defi-yield-oracle/fetchers/cache"
)

//TODO: make this a an array of tasks and handle them in parallel using map of task keys and handlers
type TaskHandlerBody struct {
	Task string `json:"task"`
}

func TaskHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := new(TaskHandlerBody)
		if err := c.BodyParser(t); err != nil {
			return err
		}

		log.Println("TaskHandler triggered with task: ", t.Task)

		switch t.Task {
			case data.FRAX_POOLS:
				response, err := fetchers.GetFraxPools()
				if err != nil {
					log.Println("Error getting FRAX pools: ", err)
					return err
				}

				err = cache.Put(data.FRAX_POOLS, response)
				if err != nil {
					log.Println("Error putting FRAX pools into cache: ", err)
					return err
				}
				return c.SendStatus(fiber.StatusOK)
			default:
				return c.Status(fiber.ErrBadRequest.Code).SendString("Unsupported task")
		}
	}
}

