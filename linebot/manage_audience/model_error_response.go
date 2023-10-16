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
package manage_audience

// ErrorResponse
// ErrorResponse
// https://developers.line.biz/en/reference/messaging-api/#error-responses
type ErrorResponse struct {

	/**
	 * Message containing information about the error. (Required)
	 */
	Message string `json:"message"`

	/**
	 * An array of error details. If the array is empty, this property will not be included in the response.
	 */
	Details []ErrorDetail `json:"details,omitempty"`
}

func NewErrorResponse(

	Message string,

) *ErrorResponse {
	e := &ErrorResponse{}

	e.Message = Message

	return e
}
