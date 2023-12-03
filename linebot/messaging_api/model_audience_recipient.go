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

// AudienceRecipient
// AudienceRecipient

type AudienceRecipient struct {
	Recipient

	/**
	 * Get AudienceGroupId
	 */
	AudienceGroupId int64 `json:"audienceGroupId"`
}

// MarshalJSON customizes the JSON serialization of the AudienceRecipient struct.
func (r *AudienceRecipient) MarshalJSON() ([]byte, error) {

	type Alias AudienceRecipient
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type,omitempty"`
	}{
		Alias: (*Alias)(r),

		Type: "audience",
	})
}
