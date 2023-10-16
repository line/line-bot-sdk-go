/**
 * Webhook Type Definition
 * Webhook event definition of the LINE Messaging API
 *
 * The version of the OpenAPI document: 1.0.0
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
package webhook

// ChatControl
// ChatControl

type ChatControl struct {

	/**
	 * Get ExpireAt
	 */
	ExpireAt int64 `json:"expireAt"`
}

func NewChatControl(

	ExpireAt int64,

) *ChatControl {
	e := &ChatControl{}

	e.ExpireAt = ExpireAt

	return e
}
