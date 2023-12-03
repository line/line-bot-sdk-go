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

// FlexBoxLinearGradient
// FlexBoxLinearGradient

type FlexBoxLinearGradient struct {
	FlexBoxBackground

	/**
	 * Get Angle
	 */
	Angle string `json:"angle,omitempty"`

	/**
	 * Get StartColor
	 */
	StartColor string `json:"startColor,omitempty"`

	/**
	 * Get EndColor
	 */
	EndColor string `json:"endColor,omitempty"`

	/**
	 * Get CenterColor
	 */
	CenterColor string `json:"centerColor,omitempty"`

	/**
	 * Get CenterPosition
	 */
	CenterPosition string `json:"centerPosition,omitempty"`
}

// MarshalJSON customizes the JSON serialization of the FlexBoxLinearGradient struct.
func (r *FlexBoxLinearGradient) MarshalJSON() ([]byte, error) {

	type Alias FlexBoxLinearGradient
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "linearGradient",
	})
}
