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
	"log"
)

type MentioneeInterface interface {
	GetType() string
}

func (e Mentionee) GetType() string {
	return e.Type
}

type UnknownMentionee struct {
	MentioneeInterface
	Type string
	Raw  map[string]json.RawMessage
}

func (e UnknownMentionee) GetType() string {
	return e.Type
}

func setDiscriminatorPropertyMentionee(r MentioneeInterface) MentioneeInterface {
	switch v := r.(type) {
	case *AllMentionee:
		if v.Type == "" {
			v.Type = "all"
		}
		return v
	case AllMentionee:
		if v.Type == "" {
			v.Type = "all"
		}
		return v
	case *UserMentionee:
		if v.Type == "" {
			v.Type = "user"
		}
		return v
	case UserMentionee:
		if v.Type == "" {
			v.Type = "user"
		}
		return v

	default:
		return v
	}
}

// Mentionee

// https://developers.line.biz/en/reference/messaging-api/#wh-text

type Mentionee struct {
	// Mentioned target.

	Type string `json:"type,omitempty"`
	// Index position of the user mention for a character in text, with the first character being at position 0.

	Index int32 `json:"index"`
	// The length of the text of the mentioned user. For a mention @example, 8 is the length.

	Length int32 `json:"length"`
}

func UnmarshalMentionee(data []byte) (MentioneeInterface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalMentionee: %w", err)
	}

	var discriminator string
	err = json.Unmarshal(raw["type"], &discriminator)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalMentionee: Cannot read type: %w", err)
	}

	switch discriminator {
	case "all":
		var all AllMentionee
		if err := json.Unmarshal(data, &all); err != nil {
			return nil, fmt.Errorf("UnmarshalMentionee: Cannot read all: %w", err)
		}
		return all, nil
	case "user":
		var user UserMentionee
		if err := json.Unmarshal(data, &user); err != nil {
			return nil, fmt.Errorf("UnmarshalMentionee: Cannot read user: %w", err)
		}
		return user, nil

	default:
		log.Println("Mentionee fallback: ", discriminator)
		var unknown UnknownMentionee
		unknown.Type = discriminator
		unknown.Raw = raw
		return unknown, nil
	}
}
