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

// UpdateAudienceGroupDescriptionRequest
// Rename an audience
// https://developers.line.biz/en/reference/messaging-api/#set-description-audience-group
type UpdateAudienceGroupDescriptionRequest struct {

	/**
	 * The audience&#39;s name. This is case-insensitive, meaning AUDIENCE and audience are considered identical. Max character limit: 120
	 */
	Description string `json:"description,omitempty"`
}

func NewUpdateAudienceGroupDescriptionRequest() *UpdateAudienceGroupDescriptionRequest {
	e := &UpdateAudienceGroupDescriptionRequest{}

	return e
}
