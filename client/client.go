package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)


var (
	once sync.Once
	client *http.Client
)


func initClient(){
	once.Do(
		func() {
			client = &http.Client{
					Timeout: time.Second * 30,
			}
		})
}

func Do(req *http.Request, response interface{}) error{
	if client == nil {
		initClient()
	}

	if req == nil {
		return errors.New("nil request")
	}

	url := req.URL.RequestURI()

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failure on %s: %w", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return makeHTTPClientError(url, resp)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HTTP Read error on response for %s: %w", url, err)
	}

	log.Printf("HTTP response for %s: %s", url, string(b))

	err = json.Unmarshal(b, response)
	if err != nil {
		return fmt.Errorf("JSON decode failed on %s:\n%s\nerror: %w", url, string(b), err)
	}

	return nil
}

