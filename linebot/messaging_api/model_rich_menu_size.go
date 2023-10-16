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

// RichMenuSize
// Rich menu size

type RichMenuSize struct {

	/**
	 * width
	 * minimum: 1
	 * maximum: 2147483647
	 */
	Width int64 `json:"width"`

	/**
	 * height
	 * minimum: 1
	 * maximum: 2147483647
	 */
	Height int64 `json:"height"`
}

func NewRichMenuSize() *RichMenuSize {
	e := &RichMenuSize{}

	return e
}
