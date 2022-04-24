package fetchers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type ScrapeFrax struct {
		Info struct {
			Version       string `json:"version"`
			StatusCode    int    `json:"statusCode"`
			StatusMessage string `json:"statusMessage"`
			Headers       struct {
				Date                       string `json:"date"`
				ContentType                string `json:"content-type"`
				AccessControlAllowOrigin   string `json:"access-control-allow-origin"`
				CacheControl               string `json:"cache-control"`
				AccessControlExposeHeaders string `json:"access-control-expose-headers"`
				FraxSrc                    string `json:"frax-src"`
				ExpectCt                   string `json:"expect-ct"`
				Vary                       string `json:"vary"`
				Server                     string `json:"server"`
				CfRay                      string `json:"cf-ray"`
				ContentEncoding            string `json:"content-encoding"`
			} `json:"headers"`
		} `json:"info"`
		Body string `json:"body"`
	}
type FraxPool struct {
	Identifier      string   `json:"identifier"`
	Chain           string   `json:"chain"`
	Platform        string   `json:"platform"`
	Logo            string   `json:"logo"`
	Pair            string   `json:"pair"`
	PairLink        string   `json:"pairLink"`
	PoolTokens      []string `json:"pool_tokens"`
	PoolRewards     []string `json:"pool_rewards"`
	LiquidityLocked float64  `json:"liquidity_locked"`
	Apy             float64  `json:"apy"`
	ApyMax          float64  `json:"apy_max"`
	IsDeprecated    bool     `json:"is_deprecated"`
}

type PageResponse []FraxPool

type FraxResponse struct {
	Pools []FraxPool `json:"pools"`
}

const fraxPoolsURL = "https://api.frax.finance/pools"


func GetFraxPools() (*FraxResponse, error) {
	url := os.Getenv("SCRAPE_HOST")
	if url == "" {
		return nil, fmt.Errorf("oracleScraper env var not set")
	}
	u := launcher.MustResolveURL(url)
	page := rod.New().ControlURL(u).MustConnect().MustPage(fraxPoolsURL)
	element := page.MustWaitLoad().MustSearch("pre")
	pageData := new(PageResponse)
	err := json.Unmarshal([]byte(element.MustText()), pageData)
	if err != nil {
		return nil, err
	}

	fraxPools := new(FraxResponse)
	fraxPools.Pools = *pageData

	return fraxPools, nil
}
