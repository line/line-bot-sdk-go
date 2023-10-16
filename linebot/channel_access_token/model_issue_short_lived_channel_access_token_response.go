/**
 * Channel Access Token API
 * This document describes Channel Access Token API.
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
package channel_access_token

// IssueShortLivedChannelAccessTokenResponse
// Issued short-lived channel access token
// https://developers.line.biz/en/reference/messaging-api/#issue-shortlived-channel-access-token
type IssueShortLivedChannelAccessTokenResponse struct {

	/**
	 * A short-lived channel access token. Valid for 30 days. Note: Channel access tokens cannot be refreshed.  (Required)
	 */
	AccessToken string `json:"access_token"`

	/**
	 * Time until channel access token expires in seconds from time the token is issued. (Required)
	 */
	ExpiresIn int32 `json:"expires_in"`

	/**
	 * Token type. The value is always `Bearer`. (Required)
	 */
	TokenType string `json:"token_type"`
}

func NewIssueShortLivedChannelAccessTokenResponse(

	AccessToken string,

	ExpiresIn int32,

	TokenType string,

) *IssueShortLivedChannelAccessTokenResponse {
	e := &IssueShortLivedChannelAccessTokenResponse{}

	e.AccessToken = AccessToken

	e.ExpiresIn = ExpiresIn

	e.TokenType = TokenType

	return e
}
