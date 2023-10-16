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
package manage_audience

// CreateClickBasedAudienceGroupRequest
// Create audience for click-based retargeting
// https://developers.line.biz/en/reference/messaging-api/#create-click-audience-group
type CreateClickBasedAudienceGroupRequest struct {

	/**
	 * The audience&#39;s name. This is case-insensitive, meaning AUDIENCE and audience are considered identical. Max character limit: 120
	 */
	Description string `json:"description,omitempty"`

	/**
	 * The request ID of a broadcast or narrowcast message sent in the past 60 days. Each Messaging API request has a request ID.
	 */
	RequestId string `json:"requestId,omitempty"`

	/**
	 * The URL clicked by the user. If empty, users who clicked any URL in the message are added to the list of recipients. Max character limit: 2,000
	 */
	ClickUrl string `json:"clickUrl,omitempty"`
}

func NewCreateClickBasedAudienceGroupRequest() *CreateClickBasedAudienceGroupRequest {
	e := &CreateClickBasedAudienceGroupRequest{}

	return e
}
