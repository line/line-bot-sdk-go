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

// FlexBubbleStyles
// FlexBubbleStyles

// Deprecated: Use OpenAPI based classes instead.
type FlexBubbleStyles struct {

	/**
	 * Get Header
	 */
	Header *FlexBlockStyle `json:"header,omitempty"`

	/**
	 * Get Hero
	 */
	Hero *FlexBlockStyle `json:"hero,omitempty"`

	/**
	 * Get Body
	 */
	Body *FlexBlockStyle `json:"body,omitempty"`

	/**
	 * Get Footer
	 */
	Footer *FlexBlockStyle `json:"footer,omitempty"`
}
