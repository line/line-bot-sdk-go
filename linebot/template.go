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

// TemplateType type
type TemplateType string

// TemplateType constants
const (
	TemplateTypeButtons       TemplateType = "buttons"
	TemplateTypeConfirm       TemplateType = "confirm"
	TemplateTypeCarousel      TemplateType = "carousel"
	TemplateTypeImageCarousel TemplateType = "image_carousel"
)

// TemplateActionType type
type TemplateActionType string

// TemplateActionType constants
const (
	TemplateActionTypeURI            TemplateActionType = "uri"
	TemplateActionTypeMessage        TemplateActionType = "message"
	TemplateActionTypePostback       TemplateActionType = "postback"
	TemplateActionTypeDatetimePicker TemplateActionType = "datetimepicker"
)

// Template interface
type Template interface {
	json.Marshaler
	template()
}

// ButtonsTemplate type
type ButtonsTemplate struct {
	ThumbnailImageURL string
	Title             string
	Text              string
	Actions           []TemplateAction
}

// MarshalJSON method of ButtonsTemplate
func (t *ButtonsTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type              TemplateType     `json:"type"`
		ThumbnailImageURL string           `json:"thumbnailImageUrl,omitempty"`
		Title             string           `json:"title,omitempty"`
		Text              string           `json:"text"`
		Actions           []TemplateAction `json:"actions"`
	}{
		Type:              TemplateTypeButtons,
		ThumbnailImageURL: t.ThumbnailImageURL,
		Title:             t.Title,
		Text:              t.Text,
		Actions:           t.Actions,
	})
}

// ConfirmTemplate type
type ConfirmTemplate struct {
	Text    string
	Actions []TemplateAction
}

// MarshalJSON method of ConfirmTemplate
func (t *ConfirmTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    TemplateType     `json:"type"`
		Text    string           `json:"text"`
		Actions []TemplateAction `json:"actions"`
	}{
		Type:    TemplateTypeConfirm,
		Text:    t.Text,
		Actions: t.Actions,
	})
}

// CarouselTemplate type
type CarouselTemplate struct {
	Columns []*CarouselColumn
}

// CarouselColumn type
type CarouselColumn struct {
	ThumbnailImageURL string           `json:"thumbnailImageUrl,omitempty"`
	Title             string           `json:"title,omitempty"`
	Text              string           `json:"text"`
	Actions           []TemplateAction `json:"actions"`
}

// MarshalJSON method of CarouselTemplate
func (t *CarouselTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    TemplateType      `json:"type"`
		Columns []*CarouselColumn `json:"columns"`
	}{
		Type:    TemplateTypeCarousel,
		Columns: t.Columns,
	})
}

// ImageCarouselTemplate type
type ImageCarouselTemplate struct {
	Columns []*ImageCarouselColumn
}

// ImageCarouselColumn type
type ImageCarouselColumn struct {
	ImageURL string         `json:"imageUrl"`
	Action   TemplateAction `json:"action"`
}

// MarshalJSON method of ImageCarouselTemplate
func (t *ImageCarouselTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    TemplateType           `json:"type"`
		Columns []*ImageCarouselColumn `json:"columns"`
	}{
		Type:    TemplateTypeImageCarousel,
		Columns: t.Columns,
	})
}

// implements Template interface
func (*ConfirmTemplate) template()       {}
func (*ButtonsTemplate) template()       {}
func (*CarouselTemplate) template()      {}
func (*ImageCarouselTemplate) template() {}

// NewConfirmTemplate function
func NewConfirmTemplate(text string, left, right TemplateAction) *ConfirmTemplate {
	return &ConfirmTemplate{
		Text:    text,
		Actions: []TemplateAction{left, right},
	}
}

// NewButtonsTemplate function
// `thumbnailImageURL` and `title` are optional. they can be empty.
func NewButtonsTemplate(thumbnailImageURL, title, text string, actions ...TemplateAction) *ButtonsTemplate {
	return &ButtonsTemplate{
		ThumbnailImageURL: thumbnailImageURL,
		Title:             title,
		Text:              text,
		Actions:           actions,
	}
}

// NewCarouselTemplate function
func NewCarouselTemplate(columns ...*CarouselColumn) *CarouselTemplate {
	return &CarouselTemplate{
		Columns: columns,
	}
}

// NewCarouselColumn function
// `thumbnailImageURL` and `title` are optional. they can be empty.
func NewCarouselColumn(thumbnailImageURL, title, text string, actions ...TemplateAction) *CarouselColumn {
	return &CarouselColumn{
		ThumbnailImageURL: thumbnailImageURL,
		Title:             title,
		Text:              text,
		Actions:           actions,
	}
}

// NewImageCarouselTemplate function
func NewImageCarouselTemplate(columns ...*ImageCarouselColumn) *ImageCarouselTemplate {
	return &ImageCarouselTemplate{
		Columns: columns,
	}
}

// NewImageCarouselColumn function
func NewImageCarouselColumn(imageURL string, action TemplateAction) *ImageCarouselColumn {
	return &ImageCarouselColumn{
		ImageURL: imageURL,
		Action:   action,
	}
}

// TemplateAction interface
type TemplateAction interface {
	json.Marshaler
	templateAction()
}

// URITemplateAction type
type URITemplateAction struct {
	Label string
	URI   string
}

// MarshalJSON method of URITemplateAction
func (a *URITemplateAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type  TemplateActionType `json:"type"`
		Label string             `json:"label"`
		URI   string             `json:"uri"`
	}{
		Type:  TemplateActionTypeURI,
		Label: a.Label,
		URI:   a.URI,
	})
}

// MessageTemplateAction type
type MessageTemplateAction struct {
	Label string
	Text  string
}

// MarshalJSON method of MessageTemplateAction
func (a *MessageTemplateAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type  TemplateActionType `json:"type"`
		Label string             `json:"label"`
		Text  string             `json:"text"`
	}{
		Type:  TemplateActionTypeMessage,
		Label: a.Label,
		Text:  a.Text,
	})
}

// PostbackTemplateAction type
type PostbackTemplateAction struct {
	Label string
	Data  string
	Text  string
}

// MarshalJSON method of PostbackTemplateAction
func (a *PostbackTemplateAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type  TemplateActionType `json:"type"`
		Label string             `json:"label"`
		Data  string             `json:"data"`
		Text  string             `json:"text,omitempty"`
	}{
		Type:  TemplateActionTypePostback,
		Label: a.Label,
		Data:  a.Data,
		Text:  a.Text,
	})
}

// DatetimePickerTemplateAction type
type DatetimePickerTemplateAction struct {
	Label   string
	Data    string
	Mode    string
	Initial string
	Max     string
	Min     string
}

// MarshalJSON method of DatetimePickerTemplateAction
func (a *DatetimePickerTemplateAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    TemplateActionType `json:"type"`
		Label   string             `json:"label"`
		Data    string             `json:"data"`
		Mode    string             `json:"mode"`
		Initial string             `json:"initial,omitempty"`
		Max     string             `json:"max,omitempty"`
		Min     string             `json:"min,omitempty"`
	}{
		Type:    TemplateActionTypeDatetimePicker,
		Label:   a.Label,
		Data:    a.Data,
		Mode:    a.Mode,
		Initial: a.Initial,
		Max:     a.Max,
		Min:     a.Min,
	})
}

// implements TemplateAction interface
func (*URITemplateAction) templateAction()            {}
func (*MessageTemplateAction) templateAction()        {}
func (*PostbackTemplateAction) templateAction()       {}
func (*DatetimePickerTemplateAction) templateAction() {}

// NewURITemplateAction function
func NewURITemplateAction(label, uri string) *URITemplateAction {
	return &URITemplateAction{
		Label: label,
		URI:   uri,
	}
}

// NewMessageTemplateAction function
func NewMessageTemplateAction(label, text string) *MessageTemplateAction {
	return &MessageTemplateAction{
		Label: label,
		Text:  text,
	}
}

// NewPostbackTemplateAction function
func NewPostbackTemplateAction(label, data, text string) *PostbackTemplateAction {
	return &PostbackTemplateAction{
		Label: label,
		Data:  data,
		Text:  text,
	}
}

// NewDatetimePickerTemplateAction function
func NewDatetimePickerTemplateAction(label, data, mode, initial, max, min string) *DatetimePickerTemplateAction {
	return &DatetimePickerTemplateAction{
		Label:   label,
		Data:    data,
		Mode:    mode,
		Initial: initial,
		Max:     max,
		Min:     min,
	}
}
