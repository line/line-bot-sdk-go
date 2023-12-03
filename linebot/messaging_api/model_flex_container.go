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

type FlexContainerInterface interface {
	GetType() string
}

func (e FlexContainer) GetType() string {
	return e.Type
}

type UnknownFlexContainer struct {
	FlexContainerInterface
	Type string
	Raw  map[string]json.RawMessage
}

func (e UnknownFlexContainer) GetType() string {
	return e.Type
}

func setDiscriminatorPropertyFlexContainer(r FlexContainerInterface) FlexContainerInterface {
	switch v := r.(type) {
	case *FlexBubble:
		if v.Type == "" {
			v.Type = "bubble"
		}
		return v
	case FlexBubble:
		if v.Type == "" {
			v.Type = "bubble"
		}
		return v
	case *FlexCarousel:
		if v.Type == "" {
			v.Type = "carousel"
		}
		return v
	case FlexCarousel:
		if v.Type == "" {
			v.Type = "carousel"
		}
		return v

	default:
		return v
	}
}

// FlexContainer

type FlexContainer struct {
	Type string `json:"type"`
}

func UnmarshalFlexContainer(data []byte) (FlexContainerInterface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalFlexContainer: %w", err)
	}

	var discriminator string
	err = json.Unmarshal(raw["type"], &discriminator)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalFlexContainer: Cannot read type: %w", err)
	}

	switch discriminator {
	case "bubble":
		var bubble FlexBubble
		if err := json.Unmarshal(data, &bubble); err != nil {
			return nil, fmt.Errorf("UnmarshalFlexContainer: Cannot read bubble: %w", err)
		}
		return bubble, nil
	case "carousel":
		var carousel FlexCarousel
		if err := json.Unmarshal(data, &carousel); err != nil {
			return nil, fmt.Errorf("UnmarshalFlexContainer: Cannot read carousel: %w", err)
		}
		return carousel, nil

	default:
		var unknown UnknownFlexContainer
		unknown.Type = discriminator
		unknown.Raw = raw
		return unknown, nil
	}
}
