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

// CreateRichMenuAliasRequest
// CreateRichMenuAliasRequest
// https://developers.line.biz/en/reference/messaging-api/#create-rich-menu-alias
type CreateRichMenuAliasRequest struct {

	/**
	 * Rich menu alias ID, which can be any ID, unique for each channel. (Required)
	 */
	RichMenuAliasId string `json:"richMenuAliasId"`

	/**
	 * The rich menu ID to be associated with the rich menu alias. (Required)
	 */
	RichMenuId string `json:"richMenuId"`
}
