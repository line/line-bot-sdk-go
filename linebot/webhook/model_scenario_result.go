/**
 * Webhook Type Definition
 * Webhook event definition of the LINE Messaging API
 *
 * The version of the OpenAPI document: 1.0.0
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
package webhook

// ScenarioResult
// ScenarioResult
// https://developers.line.biz/en/reference/messaging-api/#scenario-result-event
type ScenarioResult struct {

	/**
	 * Scenario ID executed
	 */
	ScenarioId string `json:"scenarioId,omitempty"`

	/**
	 * Revision number of the scenario set containing the executed scenario
	 */
	Revision int32 `json:"revision"`

	/**
	 * Timestamp for when execution of scenario action started (milliseconds, LINE app time) (Required)
	 */
	StartTime int64 `json:"startTime"`

	/**
	 * Timestamp for when execution of scenario was completed (milliseconds, LINE app time) (Required)
	 */
	EndTime int64 `json:"endTime"`

	/**
	 * Scenario execution completion status (Required)
	 */
	ResultCode string `json:"resultCode"`

	/**
	 * Execution result of individual operations specified in action. Only included when things.result.resultCode is success.
	 */
	ActionResults []ActionResult `json:"actionResults,omitempty"`

	/**
	 * Data contained in notification.
	 */
	BleNotificationPayload string `json:"bleNotificationPayload,omitempty"`

	/**
	 * Error reason.
	 */
	ErrorReason string `json:"errorReason,omitempty"`
}
