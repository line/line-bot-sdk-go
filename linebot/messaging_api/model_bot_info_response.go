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

// BotInfoResponse
// BotInfoResponse
// https://developers.line.biz/en/reference/messaging-api/#get-bot-info
type BotInfoResponse struct {

	/**
	 * Bot&#39;s user ID (Required)
	 */
	UserId string `json:"userId"`

	/**
	 * Bot&#39;s basic ID (Required)
	 */
	BasicId string `json:"basicId"`

	/**
	 * Bot&#39;s premium ID. Not included in the response if the premium ID isn&#39;t set.
	 */
	PremiumId string `json:"premiumId,omitempty"`

	/**
	 * Bot&#39;s display name (Required)
	 */
	DisplayName string `json:"displayName"`

	/**
	 * Profile image URL. `https` image URL. Not included in the response if the bot doesn&#39;t have a profile image.
	 */
	PictureUrl string `json:"pictureUrl,omitempty"`

	/**
	 * Chat settings set in the LINE Official Account Manager. One of:  `chat`: Chat is set to \&quot;On\&quot;. `bot`: Chat is set to \&quot;Off\&quot;.  (Required)
	 */
	ChatMode BotInfoResponseCHAT_MODE `json:"chatMode"`

	/**
	 * Automatic read setting for messages. If the chat is set to \&quot;Off\&quot;, auto is returned. If the chat is set to \&quot;On\&quot;, manual is returned.  `auto`: Auto read setting is enabled. `manual`: Auto read setting is disabled.   (Required)
	 */
	MarkAsReadMode BotInfoResponseMARK_AS_READ_MODE `json:"markAsReadMode"`
}

func NewBotInfoResponse(

	UserId string,

	BasicId string,

	DisplayName string,

	ChatMode BotInfoResponseCHAT_MODE,

	MarkAsReadMode BotInfoResponseMARK_AS_READ_MODE,

) *BotInfoResponse {
	e := &BotInfoResponse{}

	e.UserId = UserId

	e.BasicId = BasicId

	e.DisplayName = DisplayName

	e.ChatMode = ChatMode

	e.MarkAsReadMode = MarkAsReadMode

	return e
}

// BotInfoResponseCHAT_MODE type
/* Chat settings set in the LINE Official Account Manager. One of:  `chat`: Chat is set to \"On\". `bot`: Chat is set to \"Off\".  */
type BotInfoResponseCHAT_MODE string

// BotInfoResponseCHAT_MODE constants
const (
	BotInfoResponseCHAT_MODE_CHAT BotInfoResponseCHAT_MODE = "chat"

	BotInfoResponseCHAT_MODE_BOT BotInfoResponseCHAT_MODE = "bot"
)

// BotInfoResponseMARK_AS_READ_MODE type
/* Automatic read setting for messages. If the chat is set to \"Off\", auto is returned. If the chat is set to \"On\", manual is returned.  `auto`: Auto read setting is enabled. `manual`: Auto read setting is disabled.   */
type BotInfoResponseMARK_AS_READ_MODE string

// BotInfoResponseMARK_AS_READ_MODE constants
const (
	BotInfoResponseMARK_AS_READ_MODE_AUTO BotInfoResponseMARK_AS_READ_MODE = "auto"

	BotInfoResponseMARK_AS_READ_MODE_MANUAL BotInfoResponseMARK_AS_READ_MODE = "manual"
)
