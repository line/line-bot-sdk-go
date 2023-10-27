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

// FlexBox
// FlexBox

type FlexBox struct {
	FlexComponent

	/**
	 * Get Layout
	 */
	Layout FlexBoxLAYOUT `json:"layout"`

	/**
	 * Get Flex
	 */
	Flex int32 `json:"flex"`

	/**
	 * Get Contents
	 */
	Contents []FlexComponentInterface `json:"contents"`

	/**
	 * Get Spacing
	 */
	Spacing string `json:"spacing,omitempty"`

	/**
	 * Get Margin
	 */
	Margin string `json:"margin,omitempty"`

	/**
	 * Get Position
	 */
	Position FlexBoxPOSITION `json:"position,omitempty"`

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
	 * Get BackgroundColor
	 */
	BackgroundColor string `json:"backgroundColor,omitempty"`

	/**
	 * Get BorderColor
	 */
	BorderColor string `json:"borderColor,omitempty"`

	/**
	 * Get BorderWidth
	 */
	BorderWidth string `json:"borderWidth,omitempty"`

	/**
	 * Get CornerRadius
	 */
	CornerRadius string `json:"cornerRadius,omitempty"`

	/**
	 * Get Width
	 */
	Width string `json:"width,omitempty"`

	/**
	 * Get MaxWidth
	 */
	MaxWidth string `json:"maxWidth,omitempty"`

	/**
	 * Get Height
	 */
	Height string `json:"height,omitempty"`

	/**
	 * Get MaxHeight
	 */
	MaxHeight string `json:"maxHeight,omitempty"`

	/**
	 * Get PaddingAll
	 */
	PaddingAll string `json:"paddingAll,omitempty"`

	/**
	 * Get PaddingTop
	 */
	PaddingTop string `json:"paddingTop,omitempty"`

	/**
	 * Get PaddingBottom
	 */
	PaddingBottom string `json:"paddingBottom,omitempty"`

	/**
	 * Get PaddingStart
	 */
	PaddingStart string `json:"paddingStart,omitempty"`

	/**
	 * Get PaddingEnd
	 */
	PaddingEnd string `json:"paddingEnd,omitempty"`

	/**
	 * Get Action
	 */
	Action ActionInterface `json:"action,omitempty"`

	/**
	 * Get JustifyContent
	 */
	JustifyContent FlexBoxJUSTIFY_CONTENT `json:"justifyContent,omitempty"`

	/**
	 * Get AlignItems
	 */
	AlignItems FlexBoxALIGN_ITEMS `json:"alignItems,omitempty"`

	/**
	 * Get Background
	 */
	Background FlexBoxBackgroundInterface `json:"background,omitempty"`
}

func NewFlexBox(

	Layout FlexBoxLAYOUT,

	Contents []FlexComponentInterface,

) *FlexBox {
	e := &FlexBox{}

	e.Type = "box"

	e.Layout = Layout

	e.Contents = Contents

	return e
}

func (cr *FlexBox) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["type"], &cr.Type)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["layout"], &cr.Layout)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["flex"], &cr.Flex)
	if err != nil {
		return err
	}

	var rawcontents []json.RawMessage
	err = json.Unmarshal(raw["contents"], &rawcontents)
	if err != nil {
		return err
	}

	for _, data := range rawcontents {
		e, err := UnmarshalFlexComponent(data)
		if err != nil {
			return fmt.Errorf("JSON parse error in UnmarshalFlexComponent: %w, body: %s", err, string(data))
		}
		cr.Contents = append(cr.Contents, e)
	}

	err = json.Unmarshal(raw["spacing"], &cr.Spacing)
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

	err = json.Unmarshal(raw["backgroundColor"], &cr.BackgroundColor)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["borderColor"], &cr.BorderColor)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["borderWidth"], &cr.BorderWidth)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["cornerRadius"], &cr.CornerRadius)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["width"], &cr.Width)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["maxWidth"], &cr.MaxWidth)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["height"], &cr.Height)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["maxHeight"], &cr.MaxHeight)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["paddingAll"], &cr.PaddingAll)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["paddingTop"], &cr.PaddingTop)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["paddingBottom"], &cr.PaddingBottom)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["paddingStart"], &cr.PaddingStart)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["paddingEnd"], &cr.PaddingEnd)
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

	err = json.Unmarshal(raw["justifyContent"], &cr.JustifyContent)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["alignItems"], &cr.AlignItems)
	if err != nil {
		return err
	}

	if rawbackground, ok := raw["background"]; ok && rawbackground != nil {
		Background, err := UnmarshalFlexBoxBackground(rawbackground)
		if err != nil {
			return err
		}
		cr.Background = Background
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the FlexBox struct.
func (r *FlexBox) MarshalJSON() ([]byte, error) {

	newContents := make([]FlexComponentInterface, len(r.Contents))
	for i, v := range r.Contents {
		newContents[i] = setDiscriminatorPropertyFlexComponent(v)
	}

	r.Action = setDiscriminatorPropertyAction(r.Action)

	r.Background = setDiscriminatorPropertyFlexBoxBackground(r.Background)

	type Alias FlexBox
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`

		Contents []FlexComponentInterface `json:"contents"`
	}{
		Alias: (*Alias)(r),

		Type: "box",

		Contents: newContents,
	})
}

// FlexBoxLAYOUT type

type FlexBoxLAYOUT string

// FlexBoxLAYOUT constants
const (
	FlexBoxLAYOUT_HORIZONTAL FlexBoxLAYOUT = "horizontal"

	FlexBoxLAYOUT_VERTICAL FlexBoxLAYOUT = "vertical"

	FlexBoxLAYOUT_BASELINE FlexBoxLAYOUT = "baseline"
)

// FlexBoxPOSITION type

type FlexBoxPOSITION string

// FlexBoxPOSITION constants
const (
	FlexBoxPOSITION_RELATIVE FlexBoxPOSITION = "relative"

	FlexBoxPOSITION_ABSOLUTE FlexBoxPOSITION = "absolute"
)

// FlexBoxJUSTIFY_CONTENT type

type FlexBoxJUSTIFY_CONTENT string

// FlexBoxJUSTIFY_CONTENT constants
const (
	FlexBoxJUSTIFY_CONTENT_CENTER FlexBoxJUSTIFY_CONTENT = "center"

	FlexBoxJUSTIFY_CONTENT_FLEX_START FlexBoxJUSTIFY_CONTENT = "flex-start"

	FlexBoxJUSTIFY_CONTENT_FLEX_END FlexBoxJUSTIFY_CONTENT = "flex-end"

	FlexBoxJUSTIFY_CONTENT_SPACE_BETWEEN FlexBoxJUSTIFY_CONTENT = "space-between"

	FlexBoxJUSTIFY_CONTENT_SPACE_AROUND FlexBoxJUSTIFY_CONTENT = "space-around"

	FlexBoxJUSTIFY_CONTENT_SPACE_EVENLY FlexBoxJUSTIFY_CONTENT = "space-evenly"
)

// FlexBoxALIGN_ITEMS type

type FlexBoxALIGN_ITEMS string

// FlexBoxALIGN_ITEMS constants
const (
	FlexBoxALIGN_ITEMS_CENTER FlexBoxALIGN_ITEMS = "center"

	FlexBoxALIGN_ITEMS_FLEX_START FlexBoxALIGN_ITEMS = "flex-start"

	FlexBoxALIGN_ITEMS_FLEX_END FlexBoxALIGN_ITEMS = "flex-end"
)