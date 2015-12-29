package tiqav4go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"io"
	"time"
)

const (
	searchApiUrl = "http://api.tiqav.com/search.json"
	searchNewestApiUrl = "http://api.tiqav.com/search/newest.json"
	searchRandomApiUrl = "http://api.tiqav.com/search/random.json"
	getImagesApiUrl = "http://api.tiqav.com/images/"
	postImagesApiUrl = "http://api.tiqav.com/images.json"
)

/*
 TiqavClient is a Tiqav API client that wrapped http.Client struct.
 @see http://dev.tiqav.com/
*/
type TiqavClient struct {
	client *http.Client
}

/*
 Return new Tiqav API client that includes new http.Client.
 Default Timeout set 10 seconds.
*/
func NewTiqavClient() *TiqavClient {
	return &TiqavClient{
		&http.Client{Timeout: time.Duration(10) * time.Second},
	}
}

/*
 Request to Tiqav Search API and then return response.
 Close connection automatically when request is finished.
*/
func (tc TiqavClient) Search(query string) ([]Tiqav, error) {
	url := fmt.Sprintf("%s?q=%s", searchApiUrl, query)
	response, err := tc.client.Get(url)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	} else {
		return parseArray(response.Body)
	}
}

func (tc TiqavClient) SearchNewest() ([]Tiqav, error) {
	response, err := tc.client.Get(searchNewestApiUrl)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	} else {
		return parseArray(response.Body)
	}
}

func (tc TiqavClient) SearchRandom() ([]Tiqav, error) {
	response, err := tc.client.Get(searchRandomApiUrl)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	} else {
		return parseArray(response.Body)
	}
}

func (tc TiqavClient) GetImages(id string) (*Tiqav, error) {
	response, err := tc.client.Get(getImagesApiUrl + id + ".json")
	defer response.Body.Close()

	if err != nil {
		return nil, err
	} else {
		return parse(response.Body)
	}
}

func parseArray(body io.Reader) ([]Tiqav, error) {
	res, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var results []Tiqav
	if err = json.Unmarshal(res, &results); err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func parse(body io.Reader) (*Tiqav, error) {
	res, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	result := new(Tiqav)
	if err = json.Unmarshal(res, result); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
