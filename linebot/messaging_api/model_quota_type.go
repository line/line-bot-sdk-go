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
 * One of the following values to indicate whether a target limit is set or not.
 */

// QuotaType type
type QuotaType string

// QuotaType constants
const (
	QuotaType_NONE QuotaType = "none"

	QuotaType_LIMITED QuotaType = "limited"
)
