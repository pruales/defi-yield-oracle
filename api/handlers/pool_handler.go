package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pruales/defi-yield-oracle/fetchers"
)

type PoolsResponse struct {
	Frax fetchers.FraxCacheResponse `json:"frax"`
	Anchor fetchers.AnchorEarnResponse `json:"anchor"`
}

//parallelize later using goroutines.
//functions should us a common interface for waitGroup handling
//https://golangbyexample.com/template-method-design-pattern-golang/
//https://www.digitalocean.com/community/tutorials/how-to-run-multiple-functions-concurrently-in-go
func GetPools() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fraxPools, err := fetchers.GetFraxPoolsFromCache()
		if err != nil {
			log.Println("Error getting FRAX pools: ", err)
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}
		poolsResponse := new(PoolsResponse)
		poolsResponse.Frax = *fraxPools

		anchorEarn, err := fetchers.GetAnchorYield()
		if err != nil {
			log.Println("Error getting Anchor Earn: ", err)
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}
		poolsResponse.Anchor = *anchorEarn
		return c.JSON(poolsResponse)
	}
}
