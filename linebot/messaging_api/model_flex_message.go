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

// FlexMessage
// FlexMessage
// https://developers.line.biz/en/reference/messaging-api/#flex-message
// Deprecated: Use OpenAPI based classes instead.
type FlexMessage struct {
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
	 * Get AltText
	 */
	AltText string `json:"altText"`

	/**
	 * Get Contents
	 */
	Contents FlexContainerInterface `json:"contents"`
}

func (cr *FlexMessage) UnmarshalJSON(data []byte) error {
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

	err = json.Unmarshal(raw["altText"], &cr.AltText)
	if err != nil {
		return err
	}

	if rawcontents, ok := raw["contents"]; ok && rawcontents != nil {
		Contents, err := UnmarshalFlexContainer(rawcontents)
		if err != nil {
			return err
		}
		cr.Contents = Contents
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the FlexMessage struct.
func (r *FlexMessage) MarshalJSON() ([]byte, error) {

	r.Contents = setDiscriminatorPropertyFlexContainer(r.Contents)

	type Alias FlexMessage
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "flex",
	})
}
