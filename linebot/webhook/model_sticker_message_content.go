/**
 * Webhook Type Definition
 * Webhook event definition of the LINE Messaging API
 *
 * The version of the OpenAPI document: 1.0.0
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
package webhook

import (
	"encoding/json"
)

// StickerMessageContent
// StickerMessageContent
// https://developers.line.biz/en/reference/messaging-api/#wh-sticker
type StickerMessageContent struct {
	MessageContent

	/**
	 * Message ID (Required)
	 */
	Id string `json:"id"`

	/**
	 * Package ID (Required)
	 */
	PackageId string `json:"packageId"`

	/**
	 * Sticker ID (Required)
	 */
	StickerId string `json:"stickerId"`

	/**
	 * Get StickerResourceType
	 */
	StickerResourceType StickerMessageContentSTICKER_RESOURCE_TYPE `json:"stickerResourceType"`

	/**
	 * Array of up to 15 keywords describing the sticker. If a sticker has 16 or more keywords, a random selection of 15 keywords will be returned. The keyword selection is random for each event, so different keywords may be returned for the same sticker.
	 */
	Keywords []string `json:"keywords"`

	/**
	 * Any text entered by the user. This property is only included for message stickers. Max character limit: 100
	 */
	Text string `json:"text,omitempty"`

	/**
	 * Quote token to quote this message.  (Required)
	 */
	QuoteToken string `json:"quoteToken"`

	/**
	 * Message ID of a quoted message. Only included when the received message quotes a past message.
	 */
	QuotedMessageId string `json:"quotedMessageId,omitempty"`
}

// MarshalJSON customizes the JSON serialization of the StickerMessageContent struct.
func (r *StickerMessageContent) MarshalJSON() ([]byte, error) {

	type Alias StickerMessageContent
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "sticker",
	})
}

// StickerMessageContentSTICKER_RESOURCE_TYPE type

type StickerMessageContentSTICKER_RESOURCE_TYPE string

// StickerMessageContentSTICKER_RESOURCE_TYPE constants
const (
	StickerMessageContentSTICKER_RESOURCE_TYPE_STATIC StickerMessageContentSTICKER_RESOURCE_TYPE = "STATIC"

	StickerMessageContentSTICKER_RESOURCE_TYPE_ANIMATION StickerMessageContentSTICKER_RESOURCE_TYPE = "ANIMATION"

	StickerMessageContentSTICKER_RESOURCE_TYPE_SOUND StickerMessageContentSTICKER_RESOURCE_TYPE = "SOUND"

	StickerMessageContentSTICKER_RESOURCE_TYPE_ANIMATION_SOUND StickerMessageContentSTICKER_RESOURCE_TYPE = "ANIMATION_SOUND"

	StickerMessageContentSTICKER_RESOURCE_TYPE_POPUP StickerMessageContentSTICKER_RESOURCE_TYPE = "POPUP"

	StickerMessageContentSTICKER_RESOURCE_TYPE_POPUP_SOUND StickerMessageContentSTICKER_RESOURCE_TYPE = "POPUP_SOUND"

	StickerMessageContentSTICKER_RESOURCE_TYPE_CUSTOM StickerMessageContentSTICKER_RESOURCE_TYPE = "CUSTOM"

	StickerMessageContentSTICKER_RESOURCE_TYPE_MESSAGE StickerMessageContentSTICKER_RESOURCE_TYPE = "MESSAGE"

	StickerMessageContentSTICKER_RESOURCE_TYPE_NAME_TEXT StickerMessageContentSTICKER_RESOURCE_TYPE = "NAME_TEXT"

	StickerMessageContentSTICKER_RESOURCE_TYPE_PER_STICKER_TEXT StickerMessageContentSTICKER_RESOURCE_TYPE = "PER_STICKER_TEXT"
)
