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

type MessageContentInterface interface {
	GetType() string
}

func (e MessageContent) GetType() string {
	return e.Type
}

// Deprecated: Use OpenAPI based classes instead.
type UnknownMessageContent struct {
	MessageContentInterface
	Type string
	Raw  map[string]json.RawMessage
}

func (e UnknownMessageContent) GetType() string {
	return e.Type
}

// Deprecated: Use OpenAPI based classes instead.
func setDiscriminatorPropertyMessageContent(r MessageContentInterface) MessageContentInterface {
	switch v := r.(type) {
	case *AudioMessageContent:
		if v.Type == "" {
			v.Type = "audio"
		}
		return v
	case AudioMessageContent:
		if v.Type == "" {
			v.Type = "audio"
		}
		return v
	case *FileMessageContent:
		if v.Type == "" {
			v.Type = "file"
		}
		return v
	case FileMessageContent:
		if v.Type == "" {
			v.Type = "file"
		}
		return v
	case *ImageMessageContent:
		if v.Type == "" {
			v.Type = "image"
		}
		return v
	case ImageMessageContent:
		if v.Type == "" {
			v.Type = "image"
		}
		return v
	case *LocationMessageContent:
		if v.Type == "" {
			v.Type = "location"
		}
		return v
	case LocationMessageContent:
		if v.Type == "" {
			v.Type = "location"
		}
		return v
	case *StickerMessageContent:
		if v.Type == "" {
			v.Type = "sticker"
		}
		return v
	case StickerMessageContent:
		if v.Type == "" {
			v.Type = "sticker"
		}
		return v
	case *TextMessageContent:
		if v.Type == "" {
			v.Type = "text"
		}
		return v
	case TextMessageContent:
		if v.Type == "" {
			v.Type = "text"
		}
		return v
	case *VideoMessageContent:
		if v.Type == "" {
			v.Type = "video"
		}
		return v
	case VideoMessageContent:
		if v.Type == "" {
			v.Type = "video"
		}
		return v

	default:
		return v
	}
}

// MessageContent

// https://developers.line.biz/en/reference/messaging-api/#message-event

// Deprecated: Use OpenAPI based classes instead.
type MessageContent struct {
	// Type

	Type string `json:"type,omitempty"`
	// Message ID

	Id string `json:"id"`
}

// Deprecated: Use OpenAPI based classes instead.
func UnmarshalMessageContent(data []byte) (MessageContentInterface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalMessageContent: %w", err)
	}

	var discriminator string
	err = json.Unmarshal(raw["type"], &discriminator)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalMessageContent: Cannot read type: %w", err)
	}

	switch discriminator {
	case "audio":
		var audio AudioMessageContent
		if err := json.Unmarshal(data, &audio); err != nil {
			return nil, fmt.Errorf("UnmarshalMessageContent: Cannot read audio: %w", err)
		}
		return audio, nil
	case "file":
		var file FileMessageContent
		if err := json.Unmarshal(data, &file); err != nil {
			return nil, fmt.Errorf("UnmarshalMessageContent: Cannot read file: %w", err)
		}
		return file, nil
	case "image":
		var image ImageMessageContent
		if err := json.Unmarshal(data, &image); err != nil {
			return nil, fmt.Errorf("UnmarshalMessageContent: Cannot read image: %w", err)
		}
		return image, nil
	case "location":
		var location LocationMessageContent
		if err := json.Unmarshal(data, &location); err != nil {
			return nil, fmt.Errorf("UnmarshalMessageContent: Cannot read location: %w", err)
		}
		return location, nil
	case "sticker":
		var sticker StickerMessageContent
		if err := json.Unmarshal(data, &sticker); err != nil {
			return nil, fmt.Errorf("UnmarshalMessageContent: Cannot read sticker: %w", err)
		}
		return sticker, nil
	case "text":
		var text TextMessageContent
		if err := json.Unmarshal(data, &text); err != nil {
			return nil, fmt.Errorf("UnmarshalMessageContent: Cannot read text: %w", err)
		}
		return text, nil
	case "video":
		var video VideoMessageContent
		if err := json.Unmarshal(data, &video); err != nil {
			return nil, fmt.Errorf("UnmarshalMessageContent: Cannot read video: %w", err)
		}
		return video, nil

	default:
		log.Println("MessageContent fallback: ", discriminator)
		var unknown UnknownMessageContent
		unknown.Type = discriminator
		unknown.Raw = raw
		return unknown, nil
	}
}
