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

// GetStatisticsPerUnitResponseOverview
// Statistics related to messages.
// https://developers.line.biz/en/reference/messaging-api/#get-statistics-per-unit-response
type GetStatisticsPerUnitResponseOverview struct {

	/**
	 * Number of users who opened the message, meaning they displayed at least 1 bubble.
	 */
	UniqueImpression int64 `json:"uniqueImpression"`

	/**
	 * Number of users who opened any URL in the message.
	 */
	UniqueClick int64 `json:"uniqueClick"`

	/**
	 * Number of users who started playing any video or audio in the message.
	 */
	UniqueMediaPlayed int64 `json:"uniqueMediaPlayed"`

	/**
	 * Number of users who played the entirety of any video or audio in the message.
	 */
	UniqueMediaPlayed100Percent int64 `json:"uniqueMediaPlayed100Percent"`
}

func NewGetStatisticsPerUnitResponseOverview() *GetStatisticsPerUnitResponseOverview {
	e := &GetStatisticsPerUnitResponseOverview{}

	return e
}
