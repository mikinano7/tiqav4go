package tiqav4go

import "fmt"

const (
	BaseUrl = "http://img.tiqav.com"
)

// Entity of Tiqav Search and Image API response.
type Tiqav struct {
	Id        string `json:"id"`
	Ext       string `json:"ext"`
	SourceUrl string `json:"source_url"`
}

// Return image URL.
func (class Tiqav) Url() string {
	return fmt.Sprintf("%s/%s.%s", BaseUrl, class.Id, class.Ext)
}
