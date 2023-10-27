/**
 * LIFF server API
 * LIFF Server API.
 *
 * The version of the OpenAPI document: 1.0
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
package liff

// LiffApp
// LiffApp

// Deprecated: Use OpenAPI based classes instead.
type LiffApp struct {

	/**
	 * LIFF app ID
	 */
	LiffId string `json:"liffId,omitempty"`

	/**
	 * Get View
	 */
	View *LiffView `json:"view,omitempty"`

	/**
	 * Name of the LIFF app
	 */
	Description string `json:"description,omitempty"`

	/**
	 * Get Features
	 */
	Features *LiffFeatures `json:"features,omitempty"`

	/**
	 * How additional information in LIFF URLs is handled. concat is returned.
	 */
	PermanentLinkPattern string `json:"permanentLinkPattern,omitempty"`

	/**
	 * Get Scope
	 */
	Scope []LiffScope `json:"scope,omitempty"`

	/**
	 * Get BotPrompt
	 */
	BotPrompt LiffBotPrompt `json:"botPrompt,omitempty"`
}
