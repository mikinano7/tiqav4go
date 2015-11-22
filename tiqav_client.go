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
//
// TODO:
//   - error handling
//   - specify the number of response
func (tc TiqavClient) Search(query string) ([]Tiqav, error) {
	url := fmt.Sprintf("%s?q=%s", SearchApiUrl, query)
	if response, err := tc.client.Get(url); err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if body, err := ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	var results []Tiqav
	if err = json.Unmarshal(body, &results); err != nil {
		return nil, err
	}

	return results, nil
}
