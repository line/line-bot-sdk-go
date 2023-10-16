/**
 * LINE Messaging API(Insight)
 * This document describes LINE Messaging API(Insight).
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
package insight

// AreaTile
// AreaTile

type AreaTile struct {

	/**
	 * users&#39; country and region
	 */
	Area string `json:"area,omitempty"`

	/**
	 * Percentage
	 */
	Percentage float64 `json:"percentage"`
}

func NewAreaTile() *AreaTile {
	e := &AreaTile{}

	return e
}
