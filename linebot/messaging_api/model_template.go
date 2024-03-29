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

type TemplateInterface interface {
	GetType() string
}

func (e Template) GetType() string {
	return e.Type
}

type UnknownTemplate struct {
	TemplateInterface
	Type string
	Raw  map[string]json.RawMessage
}

func (e UnknownTemplate) GetType() string {
	return e.Type
}

func setDiscriminatorPropertyTemplate(r TemplateInterface) TemplateInterface {
	switch v := r.(type) {
	case *ButtonsTemplate:
		if v.Type == "" {
			v.Type = "buttons"
		}
		return v
	case ButtonsTemplate:
		if v.Type == "" {
			v.Type = "buttons"
		}
		return v
	case *CarouselTemplate:
		if v.Type == "" {
			v.Type = "carousel"
		}
		return v
	case CarouselTemplate:
		if v.Type == "" {
			v.Type = "carousel"
		}
		return v
	case *ConfirmTemplate:
		if v.Type == "" {
			v.Type = "confirm"
		}
		return v
	case ConfirmTemplate:
		if v.Type == "" {
			v.Type = "confirm"
		}
		return v
	case *ImageCarouselTemplate:
		if v.Type == "" {
			v.Type = "image_carousel"
		}
		return v
	case ImageCarouselTemplate:
		if v.Type == "" {
			v.Type = "image_carousel"
		}
		return v

	default:
		return v
	}
}

// Template

type Template struct {
	Type string `json:"type"`
}

func UnmarshalTemplate(data []byte) (TemplateInterface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalTemplate: %w", err)
	}

	var discriminator string
	err = json.Unmarshal(raw["type"], &discriminator)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalTemplate: Cannot read type: %w", err)
	}

	switch discriminator {
	case "buttons":
		var buttons ButtonsTemplate
		if err := json.Unmarshal(data, &buttons); err != nil {
			return nil, fmt.Errorf("UnmarshalTemplate: Cannot read buttons: %w", err)
		}
		return buttons, nil
	case "carousel":
		var carousel CarouselTemplate
		if err := json.Unmarshal(data, &carousel); err != nil {
			return nil, fmt.Errorf("UnmarshalTemplate: Cannot read carousel: %w", err)
		}
		return carousel, nil
	case "confirm":
		var confirm ConfirmTemplate
		if err := json.Unmarshal(data, &confirm); err != nil {
			return nil, fmt.Errorf("UnmarshalTemplate: Cannot read confirm: %w", err)
		}
		return confirm, nil
	case "image_carousel":
		var image_carousel ImageCarouselTemplate
		if err := json.Unmarshal(data, &image_carousel); err != nil {
			return nil, fmt.Errorf("UnmarshalTemplate: Cannot read image_carousel: %w", err)
		}
		return image_carousel, nil

	default:
		var unknown UnknownTemplate
		unknown.Type = discriminator
		unknown.Raw = raw
		return unknown, nil
	}
}
