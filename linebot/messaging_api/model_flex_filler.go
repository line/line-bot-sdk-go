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

// FlexFiller
// FlexFiller

type FlexFiller struct {
	FlexComponent

	/**
	 * Get Flex
	 */
	Flex int32 `json:"flex"`
}

func NewFlexFiller() *FlexFiller {
	e := &FlexFiller{}

	e.Type = "filler"

	return e
}

// MarshalJSON customizes the JSON serialization of the FlexFiller struct.
func (r *FlexFiller) MarshalJSON() ([]byte, error) {

	type Alias FlexFiller
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "filler",
	})
}
