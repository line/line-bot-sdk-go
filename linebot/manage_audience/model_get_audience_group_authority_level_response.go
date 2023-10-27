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

// GetAudienceGroupAuthorityLevelResponse
// Get the authority level of the audience
// https://developers.line.biz/en/reference/messaging-api/#get-authority-level
// Deprecated: Use OpenAPI based classes instead.
type GetAudienceGroupAuthorityLevelResponse struct {

	/**
	 * Get AuthorityLevel
	 */
	AuthorityLevel AudienceGroupAuthorityLevel `json:"authorityLevel,omitempty"`
}
