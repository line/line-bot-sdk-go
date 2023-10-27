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

// GetAudienceDataResponse
// Get audience data
// https://developers.line.biz/en/reference/messaging-api/#get-audience-group
// Deprecated: Use OpenAPI based classes instead.
type GetAudienceDataResponse struct {

	/**
	 * Get AudienceGroup
	 */
	AudienceGroup *AudienceGroup `json:"audienceGroup,omitempty"`

	/**
	 * An array of jobs. This array is used to keep track of each attempt to add new user IDs or IFAs to an audience for uploading user IDs. Empty array is returned for any other type of audience. Max: 50
	 */
	Jobs []AudienceGroupJob `json:"jobs,omitempty"`
}
