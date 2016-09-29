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

// ImagemapActionType type
type ImagemapActionType string

// ImagemapActionType constants
const (
	ImagemapActionTypeURI     ImagemapActionType = "uri"
	ImagemapActionTypeMessage ImagemapActionType = "message"
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

// MarshalJSON method of URIImagemapAction
func (a *URIImagemapAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    ImagemapActionType `json:"type"`
		LinkURL string             `json:"linkUri"`
		Area    ImagemapArea       `json:"area"`
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
		Type ImagemapActionType `json:"type"`
		Text string             `json:"text"`
		Area ImagemapArea       `json:"area"`
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
