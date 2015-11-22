package tiqav4go

import "fmt"

const (
	BaseUrl = "http://img.tiqav.com"
)

// Tiqav is a Entity of Tiqav Search and Image API response.
type Tiqav struct {
	Id        string `json:"id"`
	Ext       string `json:"ext"`
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	SourceUrl string `json:"source_url"`
}

// Return image entity.
func (class Tiqav) Image() Image {
	return Image{
		Original: fmt.Sprintf("%s/%s.%s", BaseUrl, class.Id, class.Ext),
		Thumbnail: fmt.Sprintf("%s/%s.th.jpg", BaseUrl, class.Id),
		Glitch: fmt.Sprintf("%s/%s.glitch", BaseUrl, class.Id),
	}
}

// Image is a Entity of image url container.
type Image struct {
	Original string
	Thumbnail string
	Glitch string
}
