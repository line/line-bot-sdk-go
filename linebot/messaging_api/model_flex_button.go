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

// FlexButton
// FlexButton

type FlexButton struct {
	FlexComponent

	/**
	 * Get Flex
	 */
	Flex int32 `json:"flex"`

	/**
	 * Get Color
	 */
	Color string `json:"color,omitempty"`

	/**
	 * Get Style
	 */
	Style FlexButtonSTYLE `json:"style,omitempty"`

	/**
	 * Get Action
	 */
	Action ActionInterface `json:"action"`

	/**
	 * Get Gravity
	 */
	Gravity FlexButtonGRAVITY `json:"gravity,omitempty"`

	/**
	 * Get Margin
	 */
	Margin string `json:"margin,omitempty"`

	/**
	 * Get Position
	 */
	Position FlexButtonPOSITION `json:"position,omitempty"`

	/**
	 * Get OffsetTop
	 */
	OffsetTop string `json:"offsetTop,omitempty"`

	/**
	 * Get OffsetBottom
	 */
	OffsetBottom string `json:"offsetBottom,omitempty"`

	/**
	 * Get OffsetStart
	 */
	OffsetStart string `json:"offsetStart,omitempty"`

	/**
	 * Get OffsetEnd
	 */
	OffsetEnd string `json:"offsetEnd,omitempty"`

	/**
	 * Get Height
	 */
	Height FlexButtonHEIGHT `json:"height,omitempty"`

	/**
	 * Get AdjustMode
	 */
	AdjustMode FlexButtonADJUST_MODE `json:"adjustMode,omitempty"`

	/**
	 * Get Scaling
	 */
	Scaling bool `json:"scaling"`
}

func NewFlexButton(

	Action ActionInterface,

) *FlexButton {
	e := &FlexButton{}

	e.Type = "button"

	e.Action = Action

	return e
}

func (cr *FlexButton) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["type"], &cr.Type)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["flex"], &cr.Flex)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["color"], &cr.Color)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["style"], &cr.Style)
	if err != nil {
		return err
	}

	if rawaction, ok := raw["action"]; ok && rawaction != nil {
		Action, err := UnmarshalAction(rawaction)
		if err != nil {
			return err
		}
		cr.Action = Action
	}

	err = json.Unmarshal(raw["gravity"], &cr.Gravity)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["margin"], &cr.Margin)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["position"], &cr.Position)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["offsetTop"], &cr.OffsetTop)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["offsetBottom"], &cr.OffsetBottom)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["offsetStart"], &cr.OffsetStart)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["offsetEnd"], &cr.OffsetEnd)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["height"], &cr.Height)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["adjustMode"], &cr.AdjustMode)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["scaling"], &cr.Scaling)
	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the FlexButton struct.
func (r *FlexButton) MarshalJSON() ([]byte, error) {

	r.Action = setDiscriminatorPropertyAction(r.Action)

	type Alias FlexButton
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "button",
	})
}

// FlexButtonSTYLE type

type FlexButtonSTYLE string

// FlexButtonSTYLE constants
const (
	FlexButtonSTYLE_PRIMARY FlexButtonSTYLE = "primary"

	FlexButtonSTYLE_SECONDARY FlexButtonSTYLE = "secondary"

	FlexButtonSTYLE_LINK FlexButtonSTYLE = "link"
)

// FlexButtonGRAVITY type

type FlexButtonGRAVITY string

// FlexButtonGRAVITY constants
const (
	FlexButtonGRAVITY_TOP FlexButtonGRAVITY = "top"

	FlexButtonGRAVITY_BOTTOM FlexButtonGRAVITY = "bottom"

	FlexButtonGRAVITY_CENTER FlexButtonGRAVITY = "center"
)

// FlexButtonPOSITION type

type FlexButtonPOSITION string

// FlexButtonPOSITION constants
const (
	FlexButtonPOSITION_RELATIVE FlexButtonPOSITION = "relative"

	FlexButtonPOSITION_ABSOLUTE FlexButtonPOSITION = "absolute"
)

// FlexButtonHEIGHT type

type FlexButtonHEIGHT string

// FlexButtonHEIGHT constants
const (
	FlexButtonHEIGHT_MD FlexButtonHEIGHT = "md"

	FlexButtonHEIGHT_SM FlexButtonHEIGHT = "sm"
)

// FlexButtonADJUST_MODE type

type FlexButtonADJUST_MODE string

// FlexButtonADJUST_MODE constants
const (
	FlexButtonADJUST_MODE_SHRINK_TO_FIT FlexButtonADJUST_MODE = "shrink-to-fit"
)
