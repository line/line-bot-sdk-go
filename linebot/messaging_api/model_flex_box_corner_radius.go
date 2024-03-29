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

/*
 * Radius at the time of rounding the corners of the box. This is only for `cornerRadius` in FlexBox. A value of none means that corners are not rounded; the other values are listed in order of increasing radius.
 */

// FlexBoxCornerRadius type
type FlexBoxCornerRadius string

// FlexBoxCornerRadius constants
const (
	FlexBoxCornerRadius_NONE FlexBoxCornerRadius = "none"

	FlexBoxCornerRadius_XS FlexBoxCornerRadius = "xs"

	FlexBoxCornerRadius_SM FlexBoxCornerRadius = "sm"

	FlexBoxCornerRadius_MD FlexBoxCornerRadius = "md"

	FlexBoxCornerRadius_LG FlexBoxCornerRadius = "lg"

	FlexBoxCornerRadius_XL FlexBoxCornerRadius = "xl"

	FlexBoxCornerRadius_XXL FlexBoxCornerRadius = "xxl"
)
