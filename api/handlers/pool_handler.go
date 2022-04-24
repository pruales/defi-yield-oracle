package handlers

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pruales/defi-yield-oracle/fetchers"
	"github.com/pruales/defi-yield-oracle/fetchers/cache"
)

type FraxCacheItem struct {
	Result *string `json:"result"`
	Error *string `json:"error"`
}

type FraxPoolsCacheItem struct {
	Pools []fetchers.FraxPool `json:"pools"`
}
type FraxCacheResponse struct {
	Value FraxPoolsCacheItem `json:"value"`
	UpdatedAt string `json:"updatedAt"`
}


func GetPools() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fraxRespose := new(FraxCacheItem)
		err := cache.Get(FRAX_POOLS, fraxRespose)
		if err != nil {
			log.Println("Error retrieving FRAX pools from cache: ", err)
			return err
		}
		if fraxRespose.Error != nil {
			log.Println("Error retrieving FRAX pools from cache: ", *fraxRespose.Error)
			return c.Status(fiber.ErrBadRequest.Code).SendString(*fraxRespose.Error)
		}
		clean := strings.ReplaceAll(*fraxRespose.Result, "\\", "")
		log.Println(clean)
		fraxPools := new(FraxCacheResponse)
		err = json.Unmarshal([]byte(clean), fraxPools)
		if err != nil {
			log.Println("Error unmarshalling FRAX pools from cache: ", err)
			return err
		}
		return c.JSON(fraxPools)
	}
}
