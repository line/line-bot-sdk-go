/**
 * LINE Messaging API(Insight)
 * This document describes LINE Messaging API(Insight).
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
package insight

// GetNumberOfFollowersResponse
// Get number of followers
// https://developers.line.biz/en/reference/messaging-api/#get-number-of-followers
type GetNumberOfFollowersResponse struct {

	/**
	 * Calculation status.
	 */
	Status GetNumberOfFollowersResponseSTATUS `json:"status,omitempty"`

	/**
	 * The number of times, as of the specified date, that a user added this LINE Official Account as a friend for the first time. The number doesn&#39;t decrease even if a user later blocks the account or when they delete their LINE account.
	 */
	Followers int64 `json:"followers"`

	/**
	 * The number of users, as of the specified date, that the LINE Official Account can reach through targeted messages based on gender, age, and/or region. This number only includes users who are active on LINE or LINE services and whose demographics have a high level of certainty.
	 */
	TargetedReaches int64 `json:"targetedReaches"`

	/**
	 * The number of users blocking the account as of the specified date. The number decreases when a user unblocks the account.
	 */
	Blocks int64 `json:"blocks"`
}

// GetNumberOfFollowersResponseSTATUS type
/* Calculation status. */
type GetNumberOfFollowersResponseSTATUS string

// GetNumberOfFollowersResponseSTATUS constants
const (
	GetNumberOfFollowersResponseSTATUS_READY GetNumberOfFollowersResponseSTATUS = "ready"

	GetNumberOfFollowersResponseSTATUS_UNREADY GetNumberOfFollowersResponseSTATUS = "unready"

	GetNumberOfFollowersResponseSTATUS_OUT_OF_SERVICE GetNumberOfFollowersResponseSTATUS = "out_of_service"
)
