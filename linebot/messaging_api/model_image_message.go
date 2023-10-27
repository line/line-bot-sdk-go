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

// ImageMessage
// ImageMessage
// https://developers.line.biz/en/reference/messaging-api/#image-message
type ImageMessage struct {
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
	 * Get OriginalContentUrl
	 */
	OriginalContentUrl string `json:"originalContentUrl"`

	/**
	 * Get PreviewImageUrl
	 */
	PreviewImageUrl string `json:"previewImageUrl"`
}

func NewImageMessage(

	OriginalContentUrl string,

	PreviewImageUrl string,

) *ImageMessage {
	e := &ImageMessage{}

	e.Type = "image"

	e.OriginalContentUrl = OriginalContentUrl

	e.PreviewImageUrl = PreviewImageUrl

	return e
}

// MarshalJSON customizes the JSON serialization of the ImageMessage struct.
func (r *ImageMessage) MarshalJSON() ([]byte, error) {

	type Alias ImageMessage
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "image",
	})
}