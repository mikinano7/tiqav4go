package tiqav4go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	SearchApiUrl = "http://api.tiqav.com/search.json"
)

// Tiqav API client that wrapped http.Client struct.
type TiqavClient struct {
	client *http.Client
}

// Return new Tiqav API client that includes new http.Client.
func NewTiqavClient() *TiqavClient {
	return &TiqavClient{&http.Client{}}
}

// Request Tiqav Search API and then return response.
// Close connection automatically when request is finished.
func (tc TiqavClient) Search(query string) ([]Tiqav, error) {
	url := fmt.Sprintf("%s?q=%s", SearchApiUrl, query)
	response, err := tc.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var results []Tiqav
	if err = json.Unmarshal(body, &results); err != nil {
		return nil, err
	}

	return results, nil
}
