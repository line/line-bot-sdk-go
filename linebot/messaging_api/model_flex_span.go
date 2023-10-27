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

// FlexSpan
// FlexSpan

// Deprecated: Use OpenAPI based classes instead.
type FlexSpan struct {
	FlexComponent

	/**
	 * Get Text
	 */
	Text string `json:"text,omitempty"`

	/**
	 * Get Size
	 */
	Size string `json:"size,omitempty"`

	/**
	 * Get Color
	 */
	Color string `json:"color,omitempty"`

	/**
	 * Get Weight
	 */
	Weight FlexSpanWEIGHT `json:"weight,omitempty"`

	/**
	 * Get Style
	 */
	Style FlexSpanSTYLE `json:"style,omitempty"`

	/**
	 * Get Decoration
	 */
	Decoration FlexSpanDECORATION `json:"decoration,omitempty"`
}

// MarshalJSON customizes the JSON serialization of the FlexSpan struct.
func (r *FlexSpan) MarshalJSON() ([]byte, error) {

	type Alias FlexSpan
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "span",
	})
}

// FlexSpanWEIGHT type

type FlexSpanWEIGHT string

// FlexSpanWEIGHT constants
const (
	FlexSpanWEIGHT_REGULAR FlexSpanWEIGHT = "regular"

	FlexSpanWEIGHT_BOLD FlexSpanWEIGHT = "bold"
)

// FlexSpanSTYLE type

type FlexSpanSTYLE string

// FlexSpanSTYLE constants
const (
	FlexSpanSTYLE_NORMAL FlexSpanSTYLE = "normal"

	FlexSpanSTYLE_ITALIC FlexSpanSTYLE = "italic"
)

// FlexSpanDECORATION type

type FlexSpanDECORATION string

// FlexSpanDECORATION constants
const (
	FlexSpanDECORATION_NONE FlexSpanDECORATION = "none"

	FlexSpanDECORATION_UNDERLINE FlexSpanDECORATION = "underline"

	FlexSpanDECORATION_LINE_THROUGH FlexSpanDECORATION = "line-through"
)
