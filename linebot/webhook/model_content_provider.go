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

// ContentProvider
// Provider of the media file.

// Deprecated: Use OpenAPI based classes instead.
type ContentProvider struct {

	/**
	 * Provider of the image file. (Required)
	 */
	Type ContentProviderTYPE `json:"type"`

	/**
	 * URL of the image file. Only included when contentProvider.type is external.
	 */
	OriginalContentUrl string `json:"originalContentUrl,omitempty"`

	/**
	 * URL of the preview image. Only included when contentProvider.type is external.
	 */
	PreviewImageUrl string `json:"previewImageUrl,omitempty"`
}

// ContentProviderTYPE type
/* Provider of the image file. */
type ContentProviderTYPE string

// ContentProviderTYPE constants
const (
	ContentProviderTYPE_LINE ContentProviderTYPE = "line"

	ContentProviderTYPE_EXTERNAL ContentProviderTYPE = "external"
)
