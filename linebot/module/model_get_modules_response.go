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
package module

// GetModulesResponse
// List of bots to which the module is attached
// https://developers.line.biz/en/reference/partner-docs/#get-multiple-bot-info-api
// Deprecated: Use OpenAPI based classes instead.
type GetModulesResponse struct {

	/**
	 * Array of Bot list Item objects representing basic information about the bot. (Required)
	 */
	Bots []ModuleBot `json:"bots"`

	/**
	 * Continuation token. Used to get the next array of basic bot information. This property is only returned if there are more unreturned results.
	 */
	Next string `json:"next,omitempty"`
}
