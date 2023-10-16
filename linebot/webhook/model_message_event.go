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

// MessageEvent
// Webhook event object which contains the sent message.

type MessageEvent struct {
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

	/**
	 * Get ReplyToken
	 */
	ReplyToken string `json:"replyToken,omitempty"`

	/**
	 * Get Message
	 */
	Message MessageContentInterface `json:"message,omitempty"`
}

func NewMessageEvent(

	Timestamp int64,

	Mode EventMode,

	WebhookEventId string,

	DeliveryContext *DeliveryContext,

) *MessageEvent {
	e := &MessageEvent{}

	e.Type = "message"

	e.Timestamp = Timestamp

	e.Mode = Mode

	e.WebhookEventId = WebhookEventId

	e.DeliveryContext = DeliveryContext

	return e
}

func (cr *MessageEvent) UnmarshalJSON(data []byte) error {
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

	err = json.Unmarshal(raw["replyToken"], &cr.ReplyToken)
	if err != nil {
		return err
	}

	if rawmessage, ok := raw["message"]; ok && rawmessage != nil {
		Message, err := UnmarshalMessageContent(rawmessage)
		if err != nil {
			return err
		}
		cr.Message = Message
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the MessageEvent struct.
func (r *MessageEvent) MarshalJSON() ([]byte, error) {

	r.Source = setDiscriminatorPropertySource(r.Source)

	r.Message = setDiscriminatorPropertyMessageContent(r.Message)

	type Alias MessageEvent
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "message",
	})
}
