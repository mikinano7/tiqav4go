package tiqav4go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"io"
)

const (
	searchApiUrl = "http://api.tiqav.com/search.json"
	searchNewestApiUrl = "http://api.tiqav.com/search/newest.json"
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
	url := fmt.Sprintf("%s?q=%s", searchApiUrl, query)
	response, err := tc.client.Get(url)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	} else {
		return parse(response.Body)
	}
}

func (tc TiqavClient) SearchNewest() ([]Tiqav, error) {
	response, err := tc.client.Get(searchNewestApiUrl)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	} else {
		return parse(response.Body)
	}
}

// parse will deserialize JSON of response.
func parse(body io.Reader) ([]Tiqav, error) {
	res, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var results []Tiqav
	if err = json.Unmarshal(res, &results); err != nil {
		return nil, err
	}

	return results, nil
}
