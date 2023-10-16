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

// LocationMessageContent
// LocationMessageContent

type LocationMessageContent struct {
	MessageContent

	/**
	 * Message ID (Required)
	 */
	Id string `json:"id"`

	/**
	 * Title
	 */
	Title string `json:"title,omitempty"`

	/**
	 * Address
	 */
	Address string `json:"address,omitempty"`

	/**
	 * Latitude (Required)
	 */
	Latitude float64 `json:"latitude"`

	/**
	 * Longitude (Required)
	 */
	Longitude float64 `json:"longitude"`
}

func NewLocationMessageContent(

	Id string,

	Latitude float64,

	Longitude float64,

) *LocationMessageContent {
	e := &LocationMessageContent{}

	e.Type = "location"

	e.Id = Id

	e.Latitude = Latitude

	e.Longitude = Longitude

	return e
}

// MarshalJSON customizes the JSON serialization of the LocationMessageContent struct.
func (r *LocationMessageContent) MarshalJSON() ([]byte, error) {

	type Alias LocationMessageContent
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type,omitempty"`
	}{
		Alias: (*Alias)(r),

		Type: "location",
	})
}
