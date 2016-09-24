package linebot

import (
	"encoding/json"
)

// ImagemapActionType constants
const (
	ImagemapActionTypeURI     = "uri"
	ImagemapActionTypeMessage = "message"
)

// ImagemapBaseSize type
type ImagemapBaseSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// ImagemapArea type
type ImagemapArea struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

// ImagemapAction type
type ImagemapAction interface {
	json.Marshaler
}

// ImagemapURIAction type
type ImagemapURIAction struct {
	LinkURL string
	Area    *ImagemapArea
}

// MarshalJSON method of ImagemapURIAction
func (a *ImagemapURIAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    string        `json:"type"`
		LinkURL string        `json:"linkUri"`
		Area    *ImagemapArea `json:"area"`
	}{
		Type:    ImagemapActionTypeURI,
		LinkURL: a.LinkURL,
		Area:    a.Area,
	})
}

// ImagemapMessageAction type
type ImagemapMessageAction struct {
	Text string
	Area *ImagemapArea
}

// MarshalJSON method of ImagemapMessageAction
func (a *ImagemapMessageAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string        `json:"type"`
		Text string        `json:"text"`
		Area *ImagemapArea `json:"area"`
	}{
		Type: ImagemapActionTypeMessage,
		Text: a.Text,
		Area: a.Area,
	})
}

// NewImagemapURIAction function
func NewImagemapURIAction(linkURL string, area *ImagemapArea) *ImagemapURIAction {
	return &ImagemapURIAction{
		LinkURL: linkURL,
		Area:    area,
	}
}

// NewImagemapMessageAction function
func NewImagemapMessageAction(text string, area *ImagemapArea) *ImagemapMessageAction {
	return &ImagemapMessageAction{
		Text: text,
		Area: area,
	}
}
