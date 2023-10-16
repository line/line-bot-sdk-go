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

// VideoPlayCompleteEvent
// Event for when a user finishes viewing a video at least once with the specified trackingId sent by the LINE Official Account.

type VideoPlayCompleteEvent struct {
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
	 * Reply token used to send reply message to this event (Required)
	 */
	ReplyToken string `json:"replyToken"`

	/**
	 * Get VideoPlayComplete
	 */
	VideoPlayComplete *VideoPlayComplete `json:"videoPlayComplete"`
}

func NewVideoPlayCompleteEvent(

	Timestamp int64,

	Mode EventMode,

	WebhookEventId string,

	DeliveryContext *DeliveryContext,

	ReplyToken string,

	VideoPlayComplete *VideoPlayComplete,

) *VideoPlayCompleteEvent {
	e := &VideoPlayCompleteEvent{}

	e.Type = "videoPlayComplete"

	e.Timestamp = Timestamp

	e.Mode = Mode

	e.WebhookEventId = WebhookEventId

	e.DeliveryContext = DeliveryContext

	e.ReplyToken = ReplyToken

	e.VideoPlayComplete = VideoPlayComplete

	return e
}

func (cr *VideoPlayCompleteEvent) UnmarshalJSON(data []byte) error {
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

	err = json.Unmarshal(raw["videoPlayComplete"], &cr.VideoPlayComplete)
	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the VideoPlayCompleteEvent struct.
func (r *VideoPlayCompleteEvent) MarshalJSON() ([]byte, error) {

	r.Source = setDiscriminatorPropertySource(r.Source)

	type Alias VideoPlayCompleteEvent
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "videoPlayComplete",
	})
}
