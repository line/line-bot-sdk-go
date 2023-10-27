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

// RichMenuBatchLinkOperation
// Replace the rich menu with the rich menu specified in the `to` property for all users linked to the rich menu specified in the `from` property.

// Deprecated: Use OpenAPI based classes instead.
type RichMenuBatchLinkOperation struct {
	RichMenuBatchOperation

	/**
	 * Get From
	 */
	From string `json:"from"`

	/**
	 * Get To
	 */
	To string `json:"to"`
}

// MarshalJSON customizes the JSON serialization of the RichMenuBatchLinkOperation struct.
func (r *RichMenuBatchLinkOperation) MarshalJSON() ([]byte, error) {

	type Alias RichMenuBatchLinkOperation
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "link",
	})
}
