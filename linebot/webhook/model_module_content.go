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

type ModuleContentInterface interface {
	GetType() string
}

func (e ModuleContent) GetType() string {
	return e.Type
}

type UnknownModuleContent struct {
	ModuleContentInterface
	Type string
	Raw  map[string]json.RawMessage
}

func (e UnknownModuleContent) GetType() string {
	return e.Type
}

func setDiscriminatorPropertyModuleContent(r ModuleContentInterface) ModuleContentInterface {
	switch v := r.(type) {
	case *AttachedModuleContent:
		if v.Type == "" {
			v.Type = "attached"
		}
		return v
	case AttachedModuleContent:
		if v.Type == "" {
			v.Type = "attached"
		}
		return v
	case *DetachedModuleContent:
		if v.Type == "" {
			v.Type = "detached"
		}
		return v
	case DetachedModuleContent:
		if v.Type == "" {
			v.Type = "detached"
		}
		return v

	default:
		return v
	}
}

// ModuleContent

type ModuleContent struct {
	// Type

	Type string `json:"type"`
}

func UnmarshalModuleContent(data []byte) (ModuleContentInterface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalModuleContent: %w", err)
	}

	var discriminator string
	err = json.Unmarshal(raw["type"], &discriminator)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalModuleContent: Cannot read type: %w", err)
	}

	switch discriminator {
	case "attached":
		var attached AttachedModuleContent
		if err := json.Unmarshal(data, &attached); err != nil {
			return nil, fmt.Errorf("UnmarshalModuleContent: Cannot read attached: %w", err)
		}
		return attached, nil
	case "detached":
		var detached DetachedModuleContent
		if err := json.Unmarshal(data, &detached); err != nil {
			return nil, fmt.Errorf("UnmarshalModuleContent: Cannot read detached: %w", err)
		}
		return detached, nil

	default:
		var unknown UnknownModuleContent
		unknown.Type = discriminator
		unknown.Raw = raw
		return unknown, nil
	}
}
