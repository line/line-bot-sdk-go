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
 * The current status. One of:  `ongoing`: Rich menu batch control is in progress. `succeeded`: Rich menu batch control is complete. `failed`: Rich menu batch control failed.           This means that the rich menu for one or more users couldn&#39;t be controlled.            There may also be users whose operations have been successfully completed.
 */

// RichMenuBatchProgressPhase type
type RichMenuBatchProgressPhase string

// RichMenuBatchProgressPhase constants
const (
	RichMenuBatchProgressPhase_ONGOING RichMenuBatchProgressPhase = "ongoing"

	RichMenuBatchProgressPhase_SUCCEEDED RichMenuBatchProgressPhase = "succeeded"

	RichMenuBatchProgressPhase_FAILED RichMenuBatchProgressPhase = "failed"
)
