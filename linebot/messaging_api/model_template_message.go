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
)

// TemplateMessage
// TemplateMessage
// https://developers.line.biz/en/reference/messaging-api/#template-messages
type TemplateMessage struct {
	Message

	/**
	 * Get QuickReply
	 */
	QuickReply *QuickReply `json:"quickReply,omitempty"`

	/**
	 * Get Sender
	 */
	Sender *Sender `json:"sender,omitempty"`

	/**
	 * Get AltText
	 */
	AltText string `json:"altText"`

	/**
	 * Get Template
	 */
	Template TemplateInterface `json:"template"`
}

func NewTemplateMessage(

	AltText string,

	Template TemplateInterface,

) *TemplateMessage {
	e := &TemplateMessage{}

	e.Type = "template"

	e.AltText = AltText

	e.Template = Template

	return e
}

func (cr *TemplateMessage) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["type"], &cr.Type)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["quickReply"], &cr.QuickReply)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["sender"], &cr.Sender)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["altText"], &cr.AltText)
	if err != nil {
		return err
	}

	if rawtemplate, ok := raw["template"]; ok && rawtemplate != nil {
		Template, err := UnmarshalTemplate(rawtemplate)
		if err != nil {
			return err
		}
		cr.Template = Template
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the TemplateMessage struct.
func (r *TemplateMessage) MarshalJSON() ([]byte, error) {

	r.Template = setDiscriminatorPropertyTemplate(r.Template)

	type Alias TemplateMessage
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "template",
	})
}
