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
package manage_audience

// CreateClickBasedAudienceGroupResponse
// Create audience for click-based retargeting
// https://developers.line.biz/en/reference/messaging-api/#create-click-audience-group
type CreateClickBasedAudienceGroupResponse struct {

	/**
	 * The audience ID.
	 */
	AudienceGroupId int64 `json:"audienceGroupId"`

	/**
	 * Get Type
	 */
	Type AudienceGroupType `json:"type,omitempty"`

	/**
	 * The audience&#39;s name.
	 */
	Description string `json:"description,omitempty"`

	/**
	 * When the audience was created (in UNIX time).
	 */
	Created int64 `json:"created"`

	/**
	 * The request ID that was specified when the audience was created.
	 */
	RequestId string `json:"requestId,omitempty"`

	/**
	 * The URL that was specified when the audience was created.
	 */
	ClickUrl string `json:"clickUrl,omitempty"`

	/**
	 * How the audience was created. `MESSAGING_API`: An audience created with Messaging API.
	 */
	CreateRoute CreateClickBasedAudienceGroupResponseCREATE_ROUTE `json:"createRoute,omitempty"`

	/**
	 * Audience&#39;s update permission. Audiences linked to the same channel will be READ_WRITE.  - `READ`: Can use only. - `READ_WRITE`: Can use and update.
	 */
	Permission CreateClickBasedAudienceGroupResponsePERMISSION `json:"permission,omitempty"`

	/**
	 * Time of audience expiration. Only returned for specific audiences.
	 */
	ExpireTimestamp int64 `json:"expireTimestamp"`

	/**
	 * The value indicating the type of account to be sent, as specified when creating the audience for uploading user IDs. One of:  true: Accounts are specified with IFAs. false (default): Accounts are specified with user IDs.
	 */
	IsIfaAudience bool `json:"isIfaAudience"`
}

// CreateClickBasedAudienceGroupResponseCREATE_ROUTE type
/* How the audience was created. `MESSAGING_API`: An audience created with Messaging API.  */
type CreateClickBasedAudienceGroupResponseCREATE_ROUTE string

// CreateClickBasedAudienceGroupResponseCREATE_ROUTE constants
const (
	CreateClickBasedAudienceGroupResponseCREATE_ROUTE_MESSAGING_API CreateClickBasedAudienceGroupResponseCREATE_ROUTE = "MESSAGING_API"
)

// CreateClickBasedAudienceGroupResponsePERMISSION type
/* Audience's update permission. Audiences linked to the same channel will be READ_WRITE.  - `READ`: Can use only. - `READ_WRITE`: Can use and update.  */
type CreateClickBasedAudienceGroupResponsePERMISSION string

// CreateClickBasedAudienceGroupResponsePERMISSION constants
const (
	CreateClickBasedAudienceGroupResponsePERMISSION_READ CreateClickBasedAudienceGroupResponsePERMISSION = "READ"

	CreateClickBasedAudienceGroupResponsePERMISSION_READ_WRITE CreateClickBasedAudienceGroupResponsePERMISSION = "READ_WRITE"
)
