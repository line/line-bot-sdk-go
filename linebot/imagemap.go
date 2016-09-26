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
	imagemapAction()
}

// URIImagemapAction type
type URIImagemapAction struct {
	LinkURL string
	Area    ImagemapArea
}

// MarshalJSON method of ImagemapURIAction
func (a *URIImagemapAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    string       `json:"type"`
		LinkURL string       `json:"linkUri"`
		Area    ImagemapArea `json:"area"`
	}{
		Type:    ImagemapActionTypeURI,
		LinkURL: a.LinkURL,
		Area:    a.Area,
	})
}

// MessageImagemapAction type
type MessageImagemapAction struct {
	Text string
	Area ImagemapArea
}

// MarshalJSON method of MessageImagemapAction
func (a *MessageImagemapAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string       `json:"type"`
		Text string       `json:"text"`
		Area ImagemapArea `json:"area"`
	}{
		Type: ImagemapActionTypeMessage,
		Text: a.Text,
		Area: a.Area,
	})
}

// implements ImagemapAction interface
func (a *URIImagemapAction) imagemapAction()     {}
func (a *MessageImagemapAction) imagemapAction() {}

// NewURIImagemapAction function
func NewURIImagemapAction(linkURL string, area ImagemapArea) *URIImagemapAction {
	return &URIImagemapAction{
		LinkURL: linkURL,
		Area:    area,
	}
}

// NewMessageImagemapAction function
func NewMessageImagemapAction(text string, area ImagemapArea) *MessageImagemapAction {
	return &MessageImagemapAction{
		Text: text,
		Area: area,
	}
}
