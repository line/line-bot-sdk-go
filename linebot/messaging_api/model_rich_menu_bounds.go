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

// RichMenuBounds
// Rich menu bounds
// https://developers.line.biz/en/reference/messaging-api/#bounds-object
type RichMenuBounds struct {

	/**
	 * Horizontal position relative to the top-left corner of the area.
	 * minimum: 0
	 * maximum: 2147483647
	 */
	X int64 `json:"x"`

	/**
	 * Vertical position relative to the top-left corner of the area.
	 * minimum: 0
	 * maximum: 2147483647
	 */
	Y int64 `json:"y"`

	/**
	 * Width of the area.
	 * minimum: 1
	 * maximum: 2147483647
	 */
	Width int64 `json:"width"`

	/**
	 * Height of the area.
	 * minimum: 1
	 * maximum: 2147483647
	 */
	Height int64 `json:"height"`
}
