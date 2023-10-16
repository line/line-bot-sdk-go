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

// NumberOfMessagesResponse
// NumberOfMessagesResponse

type NumberOfMessagesResponse struct {

	/**
	 * Aggregation process status. One of:  `ready`: The number of messages can be obtained. `unready`: We haven&#39;t finished calculating the number of sent messages for the specified in date. For example, this property is returned when the delivery date or a future date is specified. Calculation usually takes about a day. `unavailable_for_privacy`: The total number of messages on the specified day is less than 20. `out_of_service`: The specified date is earlier than the date on which we first started calculating sent messages (March 31, 2018).  (Required)
	 */
	Status NumberOfMessagesResponseSTATUS `json:"status"`

	/**
	 * The number of messages delivered using the phone number on the date specified in `date`. The response has this property only when the value of `status` is `ready`.
	 */
	Success int64 `json:"success"`
}

func NewNumberOfMessagesResponse(

	Status NumberOfMessagesResponseSTATUS,

) *NumberOfMessagesResponse {
	e := &NumberOfMessagesResponse{}

	e.Status = Status

	return e
}

// NumberOfMessagesResponseSTATUS type
/* Aggregation process status. One of:  `ready`: The number of messages can be obtained. `unready`: We haven't finished calculating the number of sent messages for the specified in date. For example, this property is returned when the delivery date or a future date is specified. Calculation usually takes about a day. `unavailable_for_privacy`: The total number of messages on the specified day is less than 20. `out_of_service`: The specified date is earlier than the date on which we first started calculating sent messages (March 31, 2018).  */
type NumberOfMessagesResponseSTATUS string

// NumberOfMessagesResponseSTATUS constants
const (
	NumberOfMessagesResponseSTATUS_READY NumberOfMessagesResponseSTATUS = "ready"

	NumberOfMessagesResponseSTATUS_UNREADY NumberOfMessagesResponseSTATUS = "unready"

	NumberOfMessagesResponseSTATUS_UNAVAILABLE_FOR_PRIVACY NumberOfMessagesResponseSTATUS = "unavailable_for_privacy"

	NumberOfMessagesResponseSTATUS_OUT_OF_SERVICE NumberOfMessagesResponseSTATUS = "out_of_service"
)
