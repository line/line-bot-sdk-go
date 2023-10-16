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

// AppTypeTile
// AppTypeTile

type AppTypeTile struct {

	/**
	 * users&#39; OS
	 */
	AppType AppTypeTileAPP_TYPE `json:"appType,omitempty"`

	/**
	 * Percentage
	 */
	Percentage float64 `json:"percentage"`
}

func NewAppTypeTile() *AppTypeTile {
	e := &AppTypeTile{}

	return e
}

// AppTypeTileAPP_TYPE type
/* users' OS */
type AppTypeTileAPP_TYPE string

// AppTypeTileAPP_TYPE constants
const (
	AppTypeTileAPP_TYPE_IOS AppTypeTileAPP_TYPE = "ios"

	AppTypeTileAPP_TYPE_ANDROID AppTypeTileAPP_TYPE = "android"

	AppTypeTileAPP_TYPE_OTHERS AppTypeTileAPP_TYPE = "others"
)
