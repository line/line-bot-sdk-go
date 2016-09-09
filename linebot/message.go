package linebot

import (
	"encoding/json"
)

// MessageType type
type MessageType string

// MessageType constants
const (
	MessageTypeText     = "text"
	MessageTypeImage    = "image"
	MessageTypeVideo    = "video"
	MessageTypeAudio    = "audio"
	MessageTypeLocation = "location"
	MessageTypeSticker  = "sticker"
)

// Message inteface
type Message interface {
	json.Marshaler
}

// TextMessage type
type TextMessage struct {
	ID   string
	Text string
}

// MarshalJSON method of TextMessage
func (m *TextMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type MessageType `json:"type"`
		Text string      `json:"text"`
	}{
		Type: MessageTypeText,
		Text: m.Text,
	})
}

// ImageMessage type
type ImageMessage struct {
	ID                 string
	OriginalContentURL string
	PreviewImageURL    string
}

// MarshalJSON method of ImageMessage
func (m *ImageMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type               MessageType `json:"type"`
		OriginalContentURL string      `json:"originalContentUrl"`
		PreviewImageURL    string      `json:"previewImageUrl"`
	}{
		Type:               MessageTypeImage,
		OriginalContentURL: m.OriginalContentURL,
		PreviewImageURL:    m.PreviewImageURL,
	})
}

// LocationMessage type
type LocationMessage struct {
	ID        string
	Title     string
	Address   string
	Latitude  float64
	Longitude float64
}

// MarshalJSON method of LocationMessage
func (m *LocationMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      MessageType `json:"type"`
		Title     string      `json:"title"`
		Address   string      `json:"address"`
		Latitude  float64     `json:"latitude"`
		Longitude float64     `json:"longitude"`
	}{
		Type:      MessageTypeLocation,
		Title:     m.Title,
		Address:   m.Address,
		Latitude:  m.Latitude,
		Longitude: m.Longitude,
	})
}

// StickerMessage type
type StickerMessage struct {
	ID        string
	PackageID string
	StickerID string
}

// MarshalJSON method of StickerMessage
func (m *StickerMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      MessageType `json:"type"`
		PackageID string      `json:"packageId"`
		StickerID string      `json:"stickerId"`
	}{
		Type:      MessageTypeSticker,
		PackageID: m.PackageID,
		StickerID: m.StickerID,
	})
}

// NewTextMessage function
func NewTextMessage(content string) *TextMessage {
	return &TextMessage{
		Text: content,
	}
}

// NewLocationMessage function
func NewLocationMessage(title, address string, latitude, longitude float64) *LocationMessage {
	return &LocationMessage{
		Title:     title,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewStickerMessage function
func NewStickerMessage(packageID, stickerID string) *StickerMessage {
	return &StickerMessage{
		PackageID: packageID,
		StickerID: stickerID,
	}
}
