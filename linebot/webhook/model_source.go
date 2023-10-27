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

type SourceInterface interface {
	GetType() string
}

func (e Source) GetType() string {
	return e.Type
}

// Deprecated: Use OpenAPI based classes instead.
type UnknownSource struct {
	SourceInterface
	Type string
	Raw  map[string]json.RawMessage
}

func (e UnknownSource) GetType() string {
	return e.Type
}

// Deprecated: Use OpenAPI based classes instead.
func setDiscriminatorPropertySource(r SourceInterface) SourceInterface {
	switch v := r.(type) {
	case *GroupSource:
		if v.Type == "" {
			v.Type = "group"
		}
		return v
	case GroupSource:
		if v.Type == "" {
			v.Type = "group"
		}
		return v
	case *RoomSource:
		if v.Type == "" {
			v.Type = "room"
		}
		return v
	case RoomSource:
		if v.Type == "" {
			v.Type = "room"
		}
		return v
	case *UserSource:
		if v.Type == "" {
			v.Type = "user"
		}
		return v
	case UserSource:
		if v.Type == "" {
			v.Type = "user"
		}
		return v

	default:
		return v
	}
}

// Source
// the source of the event.

// https://developers.line.biz/en/reference/messaging-api/#source-user

// Deprecated: Use OpenAPI based classes instead.
type Source struct {
	// source type

	Type string `json:"type,omitempty"`
}

// Deprecated: Use OpenAPI based classes instead.
func UnmarshalSource(data []byte) (SourceInterface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalSource: %w", err)
	}

	var discriminator string
	err = json.Unmarshal(raw["type"], &discriminator)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalSource: Cannot read type: %w", err)
	}

	switch discriminator {
	case "group":
		var group GroupSource
		if err := json.Unmarshal(data, &group); err != nil {
			return nil, fmt.Errorf("UnmarshalSource: Cannot read group: %w", err)
		}
		return group, nil
	case "room":
		var room RoomSource
		if err := json.Unmarshal(data, &room); err != nil {
			return nil, fmt.Errorf("UnmarshalSource: Cannot read room: %w", err)
		}
		return room, nil
	case "user":
		var user UserSource
		if err := json.Unmarshal(data, &user); err != nil {
			return nil, fmt.Errorf("UnmarshalSource: Cannot read user: %w", err)
		}
		return user, nil

	default:
		log.Println("Source fallback: ", discriminator)
		var unknown UnknownSource
		unknown.Type = discriminator
		unknown.Raw = raw
		return unknown, nil
	}
}
