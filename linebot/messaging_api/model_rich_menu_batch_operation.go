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
	"log"
)

type RichMenuBatchOperationInterface interface {
	GetType() string
}

func (e RichMenuBatchOperation) GetType() string {
	return e.Type
}

type UnknownRichMenuBatchOperation struct {
	RichMenuBatchOperationInterface
	Type string
	Raw  map[string]json.RawMessage
}

func (e UnknownRichMenuBatchOperation) GetType() string {
	return e.Type
}

func setDiscriminatorPropertyRichMenuBatchOperation(r RichMenuBatchOperationInterface) RichMenuBatchOperationInterface {
	switch v := r.(type) {
	case *RichMenuBatchLinkOperation:
		if v.Type == "" {
			v.Type = "link"
		}
		return v
	case RichMenuBatchLinkOperation:
		if v.Type == "" {
			v.Type = "link"
		}
		return v
	case *RichMenuBatchUnlinkOperation:
		if v.Type == "" {
			v.Type = "unlink"
		}
		return v
	case RichMenuBatchUnlinkOperation:
		if v.Type == "" {
			v.Type = "unlink"
		}
		return v
	case *RichMenuBatchUnlinkAllOperation:
		if v.Type == "" {
			v.Type = "unlinkAll"
		}
		return v
	case RichMenuBatchUnlinkAllOperation:
		if v.Type == "" {
			v.Type = "unlinkAll"
		}
		return v

	default:
		return v
	}
}

// RichMenuBatchOperation
// Rich menu operation object represents the batch operation to the rich menu linked to the user.

// https://developers.line.biz/en/reference/messaging-api/#batch-control-rich-menus-of-users-operations

type RichMenuBatchOperation struct {
	// The type of operation to the rich menu linked to the user. One of link, unlink, or unlinkAll.

	Type string `json:"type"`
}

func UnmarshalRichMenuBatchOperation(data []byte) (RichMenuBatchOperationInterface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalRichMenuBatchOperation: %w", err)
	}

	var discriminator string
	err = json.Unmarshal(raw["type"], &discriminator)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalRichMenuBatchOperation: Cannot read type: %w", err)
	}

	switch discriminator {
	case "link":
		var link RichMenuBatchLinkOperation
		if err := json.Unmarshal(data, &link); err != nil {
			return nil, fmt.Errorf("UnmarshalRichMenuBatchOperation: Cannot read link: %w", err)
		}
		return link, nil
	case "unlink":
		var unlink RichMenuBatchUnlinkOperation
		if err := json.Unmarshal(data, &unlink); err != nil {
			return nil, fmt.Errorf("UnmarshalRichMenuBatchOperation: Cannot read unlink: %w", err)
		}
		return unlink, nil
	case "unlinkAll":
		var unlinkAll RichMenuBatchUnlinkAllOperation
		if err := json.Unmarshal(data, &unlinkAll); err != nil {
			return nil, fmt.Errorf("UnmarshalRichMenuBatchOperation: Cannot read unlinkAll: %w", err)
		}
		return unlinkAll, nil

	default:
		log.Println("RichMenuBatchOperation fallback: ", discriminator)
		var unknown UnknownRichMenuBatchOperation
		unknown.Type = discriminator
		unknown.Raw = raw
		return unknown, nil
	}
}