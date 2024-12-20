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

// Limit
// Limit of the Narrowcast
// https://developers.line.biz/en/reference/messaging-api/#send-narrowcast-message
type Limit struct {

	/**
	 * The maximum number of narrowcast messages to send. Use this parameter to limit the number of narrowcast messages sent. The recipients will be chosen at random.
	 * minimum: 1
	 */
	Max int32 `json:"max,omitempty"`

	/**
	 * If true, the message will be sent within the maximum number of deliverable messages. The default value is `false`.  Targets will be selected at random.
	 */
	UpToRemainingQuota bool `json:"upToRemainingQuota"`
}
