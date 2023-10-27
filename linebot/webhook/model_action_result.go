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

// ActionResult
// ActionResult

type ActionResult struct {

	/**
	 * Get Type
	 */
	Type ActionResultTYPE `json:"type"`

	/**
	 * Base64-encoded binary data
	 */
	Data string `json:"data,omitempty"`
}

func NewActionResult(

	Type ActionResultTYPE,

) *ActionResult {
	e := &ActionResult{}

	e.Type = Type

	return e
}

// ActionResultTYPE type

type ActionResultTYPE string

// ActionResultTYPE constants
const (
	ActionResultTYPE_VOID ActionResultTYPE = "void"

	ActionResultTYPE_BINARY ActionResultTYPE = "binary"
)