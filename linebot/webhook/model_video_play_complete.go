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

// VideoPlayComplete
// VideoPlayComplete

// Deprecated: Use OpenAPI based classes instead.
type VideoPlayComplete struct {

	/**
	 * ID used to identify a video. Returns the same value as the trackingId assigned to the video message. (Required)
	 */
	TrackingId string `json:"trackingId"`
}
