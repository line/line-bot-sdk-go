/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py
package messaging_api

import (
	"encoding/json"
	"fmt"
)

// ImagemapMessage
// ImagemapMessage
// https://developers.line.biz/en/reference/messaging-api/#imagemap-message
type ImagemapMessage struct {
	Message

	/**
	 * Get QuickReply
	 */
	QuickReply *QuickReply `json:"quickReply,omitempty"`

	/**
	 * Get Sender
	 */
	Sender *Sender `json:"sender,omitempty"`

	/**
	 * Get BaseUrl
	 */
	BaseUrl string `json:"baseUrl"`

	/**
	 * Get AltText
	 */
	AltText string `json:"altText"`

	/**
	 * Get BaseSize
	 */
	BaseSize *ImagemapBaseSize `json:"baseSize"`

	/**
	 * Get Actions
	 */
	Actions []ImagemapActionInterface `json:"actions"`

	/**
	 * Get Video
	 */
	Video *ImagemapVideo `json:"video,omitempty"`
}

func NewImagemapMessage(

	BaseUrl string,

	AltText string,

	BaseSize *ImagemapBaseSize,

	Actions []ImagemapActionInterface,

) *ImagemapMessage {
	e := &ImagemapMessage{}

	e.Type = "imagemap"

	e.BaseUrl = BaseUrl

	e.AltText = AltText

	e.BaseSize = BaseSize

	e.Actions = Actions

	return e
}

func (cr *ImagemapMessage) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["type"], &cr.Type)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["quickReply"], &cr.QuickReply)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["sender"], &cr.Sender)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["baseUrl"], &cr.BaseUrl)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["altText"], &cr.AltText)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["baseSize"], &cr.BaseSize)
	if err != nil {
		return err
	}

	var rawactions []json.RawMessage
	err = json.Unmarshal(raw["actions"], &rawactions)
	if err != nil {
		return err
	}

	for _, data := range rawactions {
		e, err := UnmarshalImagemapAction(data)
		if err != nil {
			return fmt.Errorf("JSON parse error in UnmarshalImagemapAction: %w, body: %s", err, string(data))
		}
		cr.Actions = append(cr.Actions, e)
	}

	err = json.Unmarshal(raw["video"], &cr.Video)
	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the ImagemapMessage struct.
func (r *ImagemapMessage) MarshalJSON() ([]byte, error) {

	newActions := make([]ImagemapActionInterface, len(r.Actions))
	for i, v := range r.Actions {
		newActions[i] = setDiscriminatorPropertyImagemapAction(v)
	}

	type Alias ImagemapMessage
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`

		Actions []ImagemapActionInterface `json:"actions"`
	}{
		Alias: (*Alias)(r),

		Type: "imagemap",

		Actions: newActions,
	})
}
