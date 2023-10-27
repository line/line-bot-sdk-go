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

// GetNumberOfMessageDeliveriesResponse
// Get number of message deliveries
// https://developers.line.biz/en/reference/messaging-api/#get-number-of-delivery-messages
// Deprecated: Use OpenAPI based classes instead.
type GetNumberOfMessageDeliveriesResponse struct {

	/**
	 * Status of the counting process.
	 */
	Status GetNumberOfMessageDeliveriesResponseSTATUS `json:"status,omitempty"`

	/**
	 * Number of messages sent to all of this LINE Official Account&#39;s friends (broadcast messages).
	 */
	Broadcast int64 `json:"broadcast"`

	/**
	 * Number of messages sent to some of this LINE Official Account&#39;s friends, based on specific attributes (targeted messages).
	 */
	Targeting int64 `json:"targeting"`

	/**
	 * Number of auto-response messages sent.
	 */
	AutoResponse int64 `json:"autoResponse"`

	/**
	 * Number of greeting messages sent.
	 */
	WelcomeResponse int64 `json:"welcomeResponse"`

	/**
	 * Number of messages sent from LINE Official Account Manager [Chat screen](https://www.linebiz.com/jp/manual/OfficialAccountManager/chats/) (only available in Japanese).
	 */
	Chat int64 `json:"chat"`

	/**
	 * Number of broadcast messages sent with the `Send broadcast message` Messaging API operation.
	 */
	ApiBroadcast int64 `json:"apiBroadcast"`

	/**
	 * Number of push messages sent with the `Send push message` Messaging API operation.
	 */
	ApiPush int64 `json:"apiPush"`

	/**
	 * Number of multicast messages sent with the `Send multicast message` Messaging API operation.
	 */
	ApiMulticast int64 `json:"apiMulticast"`

	/**
	 * Number of narrowcast messages sent with the `Send narrowcast message` Messaging API operation.
	 */
	ApiNarrowcast int64 `json:"apiNarrowcast"`

	/**
	 * Number of replies sent with the `Send reply message` Messaging API operation.
	 */
	ApiReply int64 `json:"apiReply"`
}

// GetNumberOfMessageDeliveriesResponseSTATUS type
/* Status of the counting process. */
type GetNumberOfMessageDeliveriesResponseSTATUS string

// GetNumberOfMessageDeliveriesResponseSTATUS constants
const (
	GetNumberOfMessageDeliveriesResponseSTATUS_READY GetNumberOfMessageDeliveriesResponseSTATUS = "ready"

	GetNumberOfMessageDeliveriesResponseSTATUS_UNREADY GetNumberOfMessageDeliveriesResponseSTATUS = "unready"

	GetNumberOfMessageDeliveriesResponseSTATUS_OUT_OF_SERVICE GetNumberOfMessageDeliveriesResponseSTATUS = "out_of_service"
)
