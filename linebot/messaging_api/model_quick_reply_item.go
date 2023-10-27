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

// QuickReplyItem
// QuickReplyItem
// https://developers.line.biz/en/reference/messaging-api/#items-object
type QuickReplyItem struct {

	/**
	 * URL of the icon that is displayed at the beginning of the button
	 */
	ImageUrl string `json:"imageUrl,omitempty"`

	/**
	 * Get Action
	 */
	Action ActionInterface `json:"action,omitempty"`

	/**
	 * `action`
	 */
	Type string `json:"type,omitempty"`
}

func (cr *QuickReplyItem) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["imageUrl"], &cr.ImageUrl)
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

	err = json.Unmarshal(raw["type"], &cr.Type)
	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the QuickReplyItem struct.
func (r *QuickReplyItem) MarshalJSON() ([]byte, error) {

	r.Action = setDiscriminatorPropertyAction(r.Action)

	type Alias QuickReplyItem
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	})
}
