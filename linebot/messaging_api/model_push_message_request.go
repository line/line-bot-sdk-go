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

import (
	"encoding/json"
	"fmt"
)

// PushMessageRequest
// PushMessageRequest
// https://developers.line.biz/en/reference/messaging-api/#send-push-message
type PushMessageRequest struct {

	/**
	 * ID of the receiver. (Required)
	 */
	To string `json:"to"`

	/**
	 * List of Message objects. (Required)
	 */
	Messages []MessageInterface `json:"messages"`

	/**
	 * `true`: The user doesn’t receive a push notification when a message is sent. `false`: The user receives a push notification when the message is sent (unless they have disabled push notifications in LINE and/or their device). The default value is false.
	 */
	NotificationDisabled bool `json:"notificationDisabled"`

	/**
	 * List of aggregation unit name. Case-sensitive. This functions can only be used by corporate users who have submitted the required applications.
	 */
	CustomAggregationUnits []string `json:"customAggregationUnits"`
}

func (cr *PushMessageRequest) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return fmt.Errorf("JSON parse error in map: %w", err)
	}

	if raw["to"] != nil {

		err = json.Unmarshal(raw["to"], &cr.To)
		if err != nil {
			return fmt.Errorf("JSON parse error in string(To): %w", err)
		}

	}

	if raw["messages"] != nil {

		var rawmessages []json.RawMessage
		err = json.Unmarshal(raw["messages"], &rawmessages)
		if err != nil {
			return fmt.Errorf("JSON parse error in messages(array): %w", err)
		}

		for _, data := range rawmessages {
			e, err := UnmarshalMessage(data)
			if err != nil {
				return fmt.Errorf("JSON parse error in Message(discriminator array): %w", err)
			}
			cr.Messages = append(cr.Messages, e)
		}

	}

	if raw["notificationDisabled"] != nil {

		err = json.Unmarshal(raw["notificationDisabled"], &cr.NotificationDisabled)
		if err != nil {
			return fmt.Errorf("JSON parse error in bool(NotificationDisabled): %w", err)
		}

	}

	if raw["customAggregationUnits"] != nil {

		err = json.Unmarshal(raw["customAggregationUnits"], &cr.CustomAggregationUnits)
		if err != nil {
			return fmt.Errorf("JSON parse error in array(CustomAggregationUnits): %w", err)
		}

	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the PushMessageRequest struct.
func (r *PushMessageRequest) MarshalJSON() ([]byte, error) {

	newMessages := make([]MessageInterface, len(r.Messages))
	for i, v := range r.Messages {
		newMessages[i] = setDiscriminatorPropertyMessage(v)
	}

	type Alias PushMessageRequest
	return json.Marshal(&struct {
		*Alias

		Messages []MessageInterface `json:"messages"`
	}{
		Alias: (*Alias)(r),

		Messages: newMessages,
	})
}
