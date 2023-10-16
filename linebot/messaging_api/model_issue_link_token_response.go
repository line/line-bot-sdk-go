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

// IssueLinkTokenResponse
// IssueLinkTokenResponse
// https://developers.line.biz/en/reference/messaging-api/#issue-link-token
type IssueLinkTokenResponse struct {

	/**
	 * Link token. Link tokens are valid for 10 minutes and can only be used once.   (Required)
	 */
	LinkToken string `json:"linkToken"`
}

func NewIssueLinkTokenResponse(

	LinkToken string,

) *IssueLinkTokenResponse {
	e := &IssueLinkTokenResponse{}

	e.LinkToken = LinkToken

	return e
}
