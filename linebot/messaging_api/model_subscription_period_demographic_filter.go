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

// SubscriptionPeriodDemographicFilter
// SubscriptionPeriodDemographicFilter

type SubscriptionPeriodDemographicFilter struct {
	DemographicFilter

	/**
	 * Get Gte
	 */
	Gte SubscriptionPeriodDemographic `json:"gte,omitempty"`

	/**
	 * Get Lt
	 */
	Lt SubscriptionPeriodDemographic `json:"lt,omitempty"`
}

// MarshalJSON customizes the JSON serialization of the SubscriptionPeriodDemographicFilter struct.
func (r *SubscriptionPeriodDemographicFilter) MarshalJSON() ([]byte, error) {

	type Alias SubscriptionPeriodDemographicFilter
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type,omitempty"`
	}{
		Alias: (*Alias)(r),

		Type: "subscriptionPeriod",
	})
}
