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
)

// FlexVideo
// FlexVideo

type FlexVideo struct {
	FlexComponent

	/**
	 * Get Url
	 */
	Url string `json:"url"`

	/**
	 * Get PreviewUrl
	 */
	PreviewUrl string `json:"previewUrl"`

	/**
	 * Get AltContent
	 */
	AltContent FlexComponentInterface `json:"altContent"`

	/**
	 * Get AspectRatio
	 */
	AspectRatio string `json:"aspectRatio,omitempty"`

	/**
	 * Get Action
	 */
	Action ActionInterface `json:"action,omitempty"`
}

func NewFlexVideo(

	Url string,

	PreviewUrl string,

	AltContent FlexComponentInterface,

) *FlexVideo {
	e := &FlexVideo{}

	e.Type = "video"

	e.Url = Url

	e.PreviewUrl = PreviewUrl

	e.AltContent = AltContent

	return e
}

func (cr *FlexVideo) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["type"], &cr.Type)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["url"], &cr.Url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["previewUrl"], &cr.PreviewUrl)
	if err != nil {
		return err
	}

	if rawaltContent, ok := raw["altContent"]; ok && rawaltContent != nil {
		AltContent, err := UnmarshalFlexComponent(rawaltContent)
		if err != nil {
			return err
		}
		cr.AltContent = AltContent
	}

	err = json.Unmarshal(raw["aspectRatio"], &cr.AspectRatio)
	if err != nil {
		return err
	}

	if rawaction, ok := raw["action"]; ok && rawaction != nil {
		Action, err := UnmarshalAction(rawaction)
		if err != nil {
			return err
		}
		cr.Action = Action
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the FlexVideo struct.
func (r *FlexVideo) MarshalJSON() ([]byte, error) {

	r.AltContent = setDiscriminatorPropertyFlexComponent(r.AltContent)

	r.Action = setDiscriminatorPropertyAction(r.Action)

	type Alias FlexVideo
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "video",
	})
}