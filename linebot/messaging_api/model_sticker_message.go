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

// StickerMessage
// StickerMessage
// https://developers.line.biz/en/reference/messaging-api/#sticker-message
type StickerMessage struct {
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
	 * Get PackageId
	 */
	PackageId string `json:"packageId"`

	/**
	 * Get StickerId
	 */
	StickerId string `json:"stickerId"`

	/**
	 * Quote token of the message you want to quote.
	 */
	QuoteToken string `json:"quoteToken,omitempty"`
}

// MarshalJSON customizes the JSON serialization of the StickerMessage struct.
func (r *StickerMessage) MarshalJSON() ([]byte, error) {

	type Alias StickerMessage
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "sticker",
	})
}
