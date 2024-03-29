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

import (
	"encoding/json"
)

// AttachedModuleContent
// AttachedModuleContent

type AttachedModuleContent struct {
	ModuleContent

	/**
	 * User ID of the bot on the attached LINE Official Account (Required)
	 */
	BotId string `json:"botId"`

	/**
	 * An array of strings indicating the scope permitted by the admin of the LINE Official Account. (Required)
	 */
	Scopes []string `json:"scopes"`
}

// MarshalJSON customizes the JSON serialization of the AttachedModuleContent struct.
func (r *AttachedModuleContent) MarshalJSON() ([]byte, error) {

	type Alias AttachedModuleContent
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "attached",
	})
}
