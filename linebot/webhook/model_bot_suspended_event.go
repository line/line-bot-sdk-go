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

// BotSuspendedEvent
// This event indicates that the LINE Official Account has been suspended (Suspend). Sent to the webhook URL server of the module channel.

type BotSuspendedEvent struct {
	Event

	/**
	 * Get Source
	 */
	Source SourceInterface `json:"source,omitempty"`

	/**
	 * Time of the event in milliseconds. (Required)
	 */
	Timestamp int64 `json:"timestamp"`

	/**
	 * Get Mode
	 */
	Mode EventMode `json:"mode"`

	/**
	 * Webhook Event ID. An ID that uniquely identifies a webhook event. This is a string in ULID format. (Required)
	 */
	WebhookEventId string `json:"webhookEventId"`

	/**
	 * Get DeliveryContext
	 */
	DeliveryContext *DeliveryContext `json:"deliveryContext"`
}

func (cr *BotSuspendedEvent) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["type"], &cr.Type)
	if err != nil {
		return err
	}

	if rawsource, ok := raw["source"]; ok && rawsource != nil {
		Source, err := UnmarshalSource(rawsource)
		if err != nil {
			return err
		}
		cr.Source = Source
	}

	err = json.Unmarshal(raw["timestamp"], &cr.Timestamp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["mode"], &cr.Mode)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["webhookEventId"], &cr.WebhookEventId)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["deliveryContext"], &cr.DeliveryContext)
	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the BotSuspendedEvent struct.
func (r *BotSuspendedEvent) MarshalJSON() ([]byte, error) {

	r.Source = setDiscriminatorPropertySource(r.Source)

	type Alias BotSuspendedEvent
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "botSuspended",
	})
}
