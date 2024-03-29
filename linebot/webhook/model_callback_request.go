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
	"fmt"
)

// CallbackRequest
// The request body contains a JSON object with the user ID of a bot that should receive webhook events and an array of webhook event objects.
// https://developers.line.biz/en/reference/messaging-api/#request-body
type CallbackRequest struct {

	/**
	 * User ID of a bot that should receive webhook events. The user ID value is a string that matches the regular expression, `U[0-9a-f]{32}`.  (Required)
	 */
	Destination string `json:"destination"`

	/**
	 * Array of webhook event objects. The LINE Platform may send an empty array that doesn&#39;t include a webhook event object to confirm communication.  (Required)
	 */
	Events []EventInterface `json:"events"`
}

func (cr *CallbackRequest) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return fmt.Errorf("JSON parse error in map: %w", err)
	}

	if raw["destination"] != nil {

		err = json.Unmarshal(raw["destination"], &cr.Destination)
		if err != nil {
			return fmt.Errorf("JSON parse error in string(Destination): %w", err)
		}

	}

	if raw["events"] != nil {

		var rawevents []json.RawMessage
		err = json.Unmarshal(raw["events"], &rawevents)
		if err != nil {
			return fmt.Errorf("JSON parse error in events(array): %w", err)
		}

		for _, data := range rawevents {
			e, err := UnmarshalEvent(data)
			if err != nil {
				return fmt.Errorf("JSON parse error in Event(discriminator array): %w", err)
			}
			cr.Events = append(cr.Events, e)
		}

	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the CallbackRequest struct.
func (r *CallbackRequest) MarshalJSON() ([]byte, error) {

	newEvents := make([]EventInterface, len(r.Events))
	for i, v := range r.Events {
		newEvents[i] = setDiscriminatorPropertyEvent(v)
	}

	type Alias CallbackRequest
	return json.Marshal(&struct {
		*Alias

		Events []EventInterface `json:"events"`
	}{
		Alias: (*Alias)(r),

		Events: newEvents,
	})
}
