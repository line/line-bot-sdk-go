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

// AllMentionee
// Mentioned target is entire group

type AllMentionee struct {
	Mentionee

	/**
	 * Index position of the user mention for a character in text, with the first character being at position 0. (Required)
	 */
	Index int32 `json:"index"`

	/**
	 * The length of the text of the mentioned user. For a mention @example, 8 is the length. (Required)
	 */
	Length int32 `json:"length"`
}

// MarshalJSON customizes the JSON serialization of the AllMentionee struct.
func (r *AllMentionee) MarshalJSON() ([]byte, error) {

	type Alias AllMentionee
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "all",
	})
}
