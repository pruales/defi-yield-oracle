package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pruales/defi-yield-oracle/fetchers"
	"github.com/pruales/defi-yield-oracle/fetchers/cache"
)

const (
	FRAX_POOLS = "FRAX_POOLS"
)

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
			case FRAX_POOLS:
				response, err := fetchers.GetFraxPools()
				if err != nil {
					log.Println("Error getting FRAX pools: ", err)
					return err
				}

				err = cache.Put(FRAX_POOLS, response)
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

