package cache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/pruales/defi-yield-oracle/client"
)

const CACHE_URL = "https://global-square-mule-32511.upstash.io"

type PutBody struct {
	Value interface{} `json:"value"`
	UpdatedAt string `json:"updatedAt"`
}

type PutResponse struct {
	Result *string `json:"result"`
	Error *string `json:"error"`
}

func Put(key string, value interface{}) error {
	url := fmt.Sprintf("%s/%s/%s", CACHE_URL, "set", key)

	body, err := json.Marshal(PutBody{Value: value, UpdatedAt: time.Now().UTC().Format(time.RFC3339)})

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", getToken()))

	response := new(PutResponse)
	err = client.Do(req, response)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return fmt.Errorf("%s", *response.Error)
	}
	return nil
}

func Get(key string, target interface{}) (error) {
	url := fmt.Sprintf("%s/%s/%s", CACHE_URL, "get", key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", getToken()))

	err = client.Do(req, target)
	if err != nil {
		return err
	}
	return nil
}


func getToken() string {
	token := os.Getenv("UPSTASH_TOKEN")
	if token == "" {
		panic("UPSTASH_TOKEN is not set")
	}
	return token
}

