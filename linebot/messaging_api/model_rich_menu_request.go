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

// RichMenuRequest
// RichMenuRequest

type RichMenuRequest struct {

	/**
	 * Get Size
	 */
	Size *RichMenuSize `json:"size,omitempty"`

	/**
	 * `true` to display the rich menu by default. Otherwise, `false`.
	 */
	Selected bool `json:"selected"`

	/**
	 * Name of the rich menu. This value can be used to help manage your rich menus and is not displayed to users.
	 */
	Name string `json:"name,omitempty"`

	/**
	 * Text displayed in the chat bar
	 */
	ChatBarText string `json:"chatBarText,omitempty"`

	/**
	 * Array of area objects which define the coordinates and size of tappable areas
	 */
	Areas []RichMenuArea `json:"areas,omitempty"`
}

func NewRichMenuRequest() *RichMenuRequest {
	e := &RichMenuRequest{}

	return e
}
