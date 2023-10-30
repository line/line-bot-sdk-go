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

// GetMessageEventResponse
// Statistics about how users interact with narrowcast messages or broadcast messages sent from your LINE Official Account.
// https://developers.line.biz/en/reference/messaging-api/#get-insight-message-event-response
type GetMessageEventResponse struct {

	/**
	 * Get Overview
	 */
	Overview *GetMessageEventResponseOverview `json:"overview,omitempty"`

	/**
	 * Array of information about individual message bubbles.
	 */
	Messages []GetMessageEventResponseMessage `json:"messages,omitempty"`

	/**
	 * Array of information about opened URLs in the message.
	 */
	Clicks []GetMessageEventResponseClick `json:"clicks,omitempty"`
}
