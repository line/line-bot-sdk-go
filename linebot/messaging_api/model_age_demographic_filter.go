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

// AgeDemographicFilter
// AgeDemographicFilter

type AgeDemographicFilter struct {
	DemographicFilter

	/**
	 * Get Gte
	 */
	Gte AgeDemographic `json:"gte,omitempty"`

	/**
	 * Get Lt
	 */
	Lt AgeDemographic `json:"lt,omitempty"`
}

func NewAgeDemographicFilter() *AgeDemographicFilter {
	e := &AgeDemographicFilter{}

	e.Type = "age"

	return e
}

// MarshalJSON customizes the JSON serialization of the AgeDemographicFilter struct.
func (r *AgeDemographicFilter) MarshalJSON() ([]byte, error) {

	type Alias AgeDemographicFilter
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type,omitempty"`
	}{
		Alias: (*Alias)(r),

		Type: "age",
	})
}