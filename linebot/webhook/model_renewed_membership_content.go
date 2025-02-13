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

// RenewedMembershipContent
// RenewedMembershipContent

type RenewedMembershipContent struct {
	MembershipContent

	/**
	 * The ID of the membership that the user renewed. This is defined for each membership. (Required)
	 */
	MembershipId int32 `json:"membershipId"`
}

// MarshalJSON customizes the JSON serialization of the RenewedMembershipContent struct.
func (r *RenewedMembershipContent) MarshalJSON() ([]byte, error) {

	type Alias RenewedMembershipContent
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "renewed",
	})
}
