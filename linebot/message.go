// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package linebot

import (
	"encoding/json"
)

// MessageType type
type MessageType string

// MessageType constants
const (
	MessageTypeText     MessageType = "text"
	MessageTypeImage    MessageType = "image"
	MessageTypeVideo    MessageType = "video"
	MessageTypeAudio    MessageType = "audio"
	MessageTypeLocation MessageType = "location"
	MessageTypeSticker  MessageType = "sticker"
	MessageTypeTemplate MessageType = "template"
	MessageTypeImagemap MessageType = "imagemap"
)

// Message inteface
type Message interface {
	json.Marshaler
	message()
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

// VideoMessage type
type VideoMessage struct {
	ID                 string
	OriginalContentURL string
	PreviewImageURL    string
}

// MarshalJSON method of VideoMessage
func (m *VideoMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type               MessageType `json:"type"`
		OriginalContentURL string      `json:"originalContentUrl"`
		PreviewImageURL    string      `json:"previewImageUrl"`
	}{
		Type:               MessageTypeVideo,
		OriginalContentURL: m.OriginalContentURL,
		PreviewImageURL:    m.PreviewImageURL,
	})
}

// AudioMessage type
type AudioMessage struct {
	ID                 string
	OriginalContentURL string
	Duration           int
}

// MarshalJSON method of AudioMessage
func (m *AudioMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type               MessageType `json:"type"`
		OriginalContentURL string      `json:"originalContentUrl"`
		Duration           int         `json:"duration"`
	}{
		Type:               MessageTypeAudio,
		OriginalContentURL: m.OriginalContentURL,
		Duration:           m.Duration,
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

// TemplateMessage type
type TemplateMessage struct {
	AltText  string
	Template Template
}

// MarshalJSON method of TemplateMessage
func (m *TemplateMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type     MessageType `json:"type"`
		AltText  string      `json:"altText"`
		Template Template    `json:"template"`
	}{
		Type:     MessageTypeTemplate,
		AltText:  m.AltText,
		Template: m.Template,
	})
}

// ImagemapMessage type
type ImagemapMessage struct {
	BaseURL  string
	AltText  string
	BaseSize ImagemapBaseSize
	Actions  []ImagemapAction
}

// MarshalJSON method of ImagemapMessage
func (m *ImagemapMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type     MessageType      `json:"type"`
		BaseURL  string           `json:"baseUrl"`
		AltText  string           `json:"altText"`
		BaseSize ImagemapBaseSize `json:"baseSize"`
		Actions  []ImagemapAction `json:"actions"`
	}{
		Type:     MessageTypeImagemap,
		BaseURL:  m.BaseURL,
		AltText:  m.AltText,
		BaseSize: m.BaseSize,
		Actions:  m.Actions,
	})
}

// implements Message interface
func (*TextMessage) message()     {}
func (*ImageMessage) message()    {}
func (*VideoMessage) message()    {}
func (*AudioMessage) message()    {}
func (*LocationMessage) message() {}
func (*StickerMessage) message()  {}
func (*TemplateMessage) message() {}
func (*ImagemapMessage) message() {}

// NewTextMessage function
func NewTextMessage(content string) *TextMessage {
	return &TextMessage{
		Text: content,
	}
}

// NewImageMessage function
func NewImageMessage(originalContentURL, previewImageURL string) *ImageMessage {
	return &ImageMessage{
		OriginalContentURL: originalContentURL,
		PreviewImageURL:    previewImageURL,
	}
}

// NewVideoMessage function
func NewVideoMessage(originalContentURL, previewImageURL string) *VideoMessage {
	return &VideoMessage{
		OriginalContentURL: originalContentURL,
		PreviewImageURL:    previewImageURL,
	}
}

// NewAudioMessage function
func NewAudioMessage(originalContentURL string, duration int) *AudioMessage {
	return &AudioMessage{
		OriginalContentURL: originalContentURL,
		Duration:           duration,
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

// NewTemplateMessage function
func NewTemplateMessage(altText string, template Template) *TemplateMessage {
	return &TemplateMessage{
		AltText:  altText,
		Template: template,
	}
}

// NewImagemapMessage function
func NewImagemapMessage(baseURL, altText string, baseSize ImagemapBaseSize, actions ...ImagemapAction) *ImagemapMessage {
	return &ImagemapMessage{
		BaseURL:  baseURL,
		AltText:  altText,
		BaseSize: baseSize,
		Actions:  actions,
	}
}
