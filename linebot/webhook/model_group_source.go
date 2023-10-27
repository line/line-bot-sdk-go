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

// GroupSource
// GroupSource

// Deprecated: Use OpenAPI based classes instead.
type GroupSource struct {
	Source

	/**
	 * Group ID of the source group chat (Required)
	 */
	GroupId string `json:"groupId"`

	/**
	 * ID of the source user. Only included in message events. Only users of LINE for iOS and LINE for Android are included in userId.
	 */
	UserId string `json:"userId,omitempty"`
}

// MarshalJSON customizes the JSON serialization of the GroupSource struct.
func (r *GroupSource) MarshalJSON() ([]byte, error) {

	type Alias GroupSource
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "group",
	})
}
