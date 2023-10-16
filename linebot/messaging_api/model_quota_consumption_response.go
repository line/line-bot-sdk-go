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

// QuotaConsumptionResponse
// QuotaConsumptionResponse
// https://developers.line.biz/en/reference/messaging-api/#get-consumption
type QuotaConsumptionResponse struct {

	/**
	 * The number of sent messages in the current month (Required)
	 */
	TotalUsage int64 `json:"totalUsage"`
}

func NewQuotaConsumptionResponse(

	TotalUsage int64,

) *QuotaConsumptionResponse {
	e := &QuotaConsumptionResponse{}

	e.TotalUsage = TotalUsage

	return e
}
