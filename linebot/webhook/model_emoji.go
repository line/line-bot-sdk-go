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

// Emoji
// Emoji

type Emoji struct {

	/**
	 * Index position for a character in text, with the first character being at position 0. (Required)
	 */
	Index int32 `json:"index"`

	/**
	 * The length of the LINE emoji string. For LINE emoji (hello), 7 is the length. (Required)
	 */
	Length int32 `json:"length"`

	/**
	 * Product ID for a LINE emoji set. (Required)
	 */
	ProductId string `json:"productId"`

	/**
	 * ID for a LINE emoji inside a set. (Required)
	 */
	EmojiId string `json:"emojiId"`
}
