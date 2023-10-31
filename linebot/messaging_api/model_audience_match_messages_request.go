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

// AudienceMatchMessagesRequest
// AudienceMatchMessagesRequest
// https://developers.line.biz/en/reference/partner-docs/#phone-audience-match
type AudienceMatchMessagesRequest struct {

	/**
	 * Destination of the message (A value obtained by hashing the telephone number, which is another value normalized to E.164 format, with SHA256). (Required)
	 */
	Messages []MessageInterface `json:"messages"`

	/**
	 * Message to send. (Required)
	 */
	To []string `json:"to"`

	/**
	 * `true`: The user doesn’t receive a push notification when a message is sent. `false`: The user receives a push notification when the message is sent (unless they have disabled push notifications in LINE and/or their device). The default value is false.
	 */
	NotificationDisabled bool `json:"notificationDisabled"`
}

func (cr *AudienceMatchMessagesRequest) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return fmt.Errorf("JSON parse error in map: %w", err)
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

	if raw["to"] != nil {

		err = json.Unmarshal(raw["to"], &cr.To)
		if err != nil {
			return fmt.Errorf("JSON parse error in array(To): %w", err)
		}

	}

	if raw["notificationDisabled"] != nil {

		err = json.Unmarshal(raw["notificationDisabled"], &cr.NotificationDisabled)
		if err != nil {
			return fmt.Errorf("JSON parse error in bool(NotificationDisabled): %w", err)
		}

	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the AudienceMatchMessagesRequest struct.
func (r *AudienceMatchMessagesRequest) MarshalJSON() ([]byte, error) {

	newMessages := make([]MessageInterface, len(r.Messages))
	for i, v := range r.Messages {
		newMessages[i] = setDiscriminatorPropertyMessage(v)
	}

	type Alias AudienceMatchMessagesRequest
	return json.Marshal(&struct {
		*Alias

		Messages []MessageInterface `json:"messages"`
	}{
		Alias: (*Alias)(r),

		Messages: newMessages,
	})
}
