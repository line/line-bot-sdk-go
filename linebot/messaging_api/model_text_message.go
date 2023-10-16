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

// TextMessage
// TextMessage
// https://developers.line.biz/en/reference/messaging-api/#text-message
type TextMessage struct {
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
	 * Get Text
	 */
	Text string `json:"text"`

	/**
	 * Get Emojis
	 */
	Emojis []Emoji `json:"emojis,omitempty"`

	/**
	 * Quote token of the message you want to quote.
	 */
	QuoteToken string `json:"quoteToken,omitempty"`
}

func NewTextMessage(

	Text string,

) *TextMessage {
	e := &TextMessage{}

	e.Type = "text"

	e.Text = Text

	return e
}

// MarshalJSON customizes the JSON serialization of the TextMessage struct.
func (r *TextMessage) MarshalJSON() ([]byte, error) {

	type Alias TextMessage
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "text",
	})
}
