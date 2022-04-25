package fetchers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/pruales/defi-yield-oracle/client"
)

const (
	LCD_URL = "https://lcd.terra.dev"
	ContractStorePathForAnchorEarn = "/terra/wasm/v1beta1/contracts/terra1tmnqgvg567ypvsvk6rwsga3srp7e3lg6u0elp8/store?query_msg=eyJlcG9jaF9zdGF0ZSI6e319"
	BlocksPerYear = 4_656_810
	Identifier = "anchor_earn"
	PoolTokens = "UST"
	PoolRewards = "UST"
	Chain = "terra"
	Link = "https://app.anchorprotocol.com/earn"
)

type EpochStateResponse struct {
	QueryResult struct {
		DepositRate        string `json:"deposit_rate"`
		PrevAterraSupply   string `json:"prev_aterra_supply"`
		PrevExchangeRate   string `json:"prev_exchange_rate"`
		PrevInterestBuffer string `json:"prev_interest_buffer"`
		LastExecutedHeight int    `json:"last_executed_height"`
	} `json:"query_result"`
}

type AnchorEarnResponse struct {
	Identifier string `json:"identifier"`
	Apy float64 `json:"apy"`
	PoolTokens []string `json:"pool_tokens"`
	PoolRewards []string `json:"pool_rewards"`
	Chain string `json:"chain"`
	Link string `json:"link"`
}

func GetAnchorYield() (*AnchorEarnResponse, error) {
	url := fmt.Sprintf("%s%s", LCD_URL, ContractStorePathForAnchorEarn)
	log.Println("Fetching Anchor Earn from: ", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	epochStateResponse := new(EpochStateResponse)
	client.Do(req, epochStateResponse)

	log.Println("Anchor Earn: ", epochStateResponse)

	depositRateFloat, err := strconv.ParseFloat(epochStateResponse.QueryResult.DepositRate, 64)
	if err != nil {
		return nil, err
	}
	apy := depositRateFloat * BlocksPerYear

	return generateAnchorEarnResponse(apy), nil
}

func generateAnchorEarnResponse(apy float64) *AnchorEarnResponse {
	return &AnchorEarnResponse{
		Identifier: Identifier,
		Apy: apy,
		PoolTokens: []string{PoolTokens},
		PoolRewards: []string{PoolRewards},
		Chain: Chain,
		Link: Link,
	}
}
