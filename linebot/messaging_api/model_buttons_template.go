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

// ButtonsTemplate
// ButtonsTemplate

type ButtonsTemplate struct {
	Template

	/**
	 * Get ThumbnailImageUrl
	 */
	ThumbnailImageUrl string `json:"thumbnailImageUrl,omitempty"`

	/**
	 * Get ImageAspectRatio
	 */
	ImageAspectRatio string `json:"imageAspectRatio,omitempty"`

	/**
	 * Get ImageSize
	 */
	ImageSize string `json:"imageSize,omitempty"`

	/**
	 * Get ImageBackgroundColor
	 */
	ImageBackgroundColor string `json:"imageBackgroundColor,omitempty"`

	/**
	 * Get Title
	 */
	Title string `json:"title,omitempty"`

	/**
	 * Get Text
	 */
	Text string `json:"text"`

	/**
	 * Get DefaultAction
	 */
	DefaultAction ActionInterface `json:"defaultAction,omitempty"`

	/**
	 * Get Actions
	 */
	Actions []ActionInterface `json:"actions"`
}

func NewButtonsTemplate(

	Text string,

	Actions []ActionInterface,

) *ButtonsTemplate {
	e := &ButtonsTemplate{}

	e.Type = "buttons"

	e.Text = Text

	e.Actions = Actions

	return e
}

func (cr *ButtonsTemplate) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["type"], &cr.Type)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["thumbnailImageUrl"], &cr.ThumbnailImageUrl)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["imageAspectRatio"], &cr.ImageAspectRatio)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["imageSize"], &cr.ImageSize)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["imageBackgroundColor"], &cr.ImageBackgroundColor)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["title"], &cr.Title)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["text"], &cr.Text)
	if err != nil {
		return err
	}

	if rawdefaultAction, ok := raw["defaultAction"]; ok && rawdefaultAction != nil {
		DefaultAction, err := UnmarshalAction(rawdefaultAction)
		if err != nil {
			return err
		}
		cr.DefaultAction = DefaultAction
	}

	var rawactions []json.RawMessage
	err = json.Unmarshal(raw["actions"], &rawactions)
	if err != nil {
		return err
	}

	for _, data := range rawactions {
		e, err := UnmarshalAction(data)
		if err != nil {
			return fmt.Errorf("JSON parse error in UnmarshalAction: %w, body: %s", err, string(data))
		}
		cr.Actions = append(cr.Actions, e)
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the ButtonsTemplate struct.
func (r *ButtonsTemplate) MarshalJSON() ([]byte, error) {

	r.DefaultAction = setDiscriminatorPropertyAction(r.DefaultAction)

	newActions := make([]ActionInterface, len(r.Actions))
	for i, v := range r.Actions {
		newActions[i] = setDiscriminatorPropertyAction(v)
	}

	type Alias ButtonsTemplate
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`

		Actions []ActionInterface `json:"actions"`
	}{
		Alias: (*Alias)(r),

		Type: "buttons",

		Actions: newActions,
	})
}
