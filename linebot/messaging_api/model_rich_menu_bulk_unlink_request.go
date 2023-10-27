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

// RichMenuBulkUnlinkRequest
// RichMenuBulkUnlinkRequest
// https://developers.line.biz/en/reference/messaging-api/#unlink-rich-menu-from-users
// Deprecated: Use OpenAPI based classes instead.
type RichMenuBulkUnlinkRequest struct {

	/**
	 * Array of user IDs. Found in the `source` object of webhook event objects. Do not use the LINE ID used in LINE. (Required)
	 */
	UserIds []string `json:"userIds"`
}
