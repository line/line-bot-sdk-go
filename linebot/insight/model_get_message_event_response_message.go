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

// GetMessageEventResponseMessage
// GetMessageEventResponseMessage

// Deprecated: Use OpenAPI based classes instead.
type GetMessageEventResponseMessage struct {

	/**
	 * Bubble&#39;s serial number.
	 */
	Seq int32 `json:"seq"`

	/**
	 * Number of times the bubble was displayed.
	 */
	Impression int64 `json:"impression"`

	/**
	 * Number of times audio or video in the bubble started playing.
	 */
	MediaPlayed int64 `json:"mediaPlayed"`

	/**
	 * Number of times audio or video in the bubble started playing and was played 25% of the total time.
	 */
	MediaPlayed25Percent int64 `json:"mediaPlayed25Percent"`

	/**
	 * Number of times audio or video in the bubble started playing and was played 50% of the total time.
	 */
	MediaPlayed50Percent int64 `json:"mediaPlayed50Percent"`

	/**
	 * Number of times audio or video in the bubble started playing and was played 75% of the total time.
	 */
	MediaPlayed75Percent int64 `json:"mediaPlayed75Percent"`

	/**
	 * Number of times audio or video in the bubble started playing and was played 100% of the total time.
	 */
	MediaPlayed100Percent int64 `json:"mediaPlayed100Percent"`

	/**
	 * Number of users that started playing audio or video in the bubble.
	 */
	UniqueMediaPlayed int64 `json:"uniqueMediaPlayed"`

	/**
	 * Number of users that started playing audio or video in the bubble and played 25% of the total time.
	 */
	UniqueMediaPlayed25Percent int64 `json:"uniqueMediaPlayed25Percent"`

	/**
	 * Number of users that started playing audio or video in the bubble and played 50% of the total time.
	 */
	UniqueMediaPlayed50Percent int64 `json:"uniqueMediaPlayed50Percent"`

	/**
	 * Number of users that started playing audio or video in the bubble and played 75% of the total time.
	 */
	UniqueMediaPlayed75Percent int64 `json:"uniqueMediaPlayed75Percent"`

	/**
	 * Number of users that started playing audio or video in the bubble and played 100% of the total time.
	 */
	UniqueMediaPlayed100Percent int64 `json:"uniqueMediaPlayed100Percent"`
}
