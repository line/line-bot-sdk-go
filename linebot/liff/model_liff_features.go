/**
 * LIFF server API
 * LIFF Server API.
 *
 * The version of the OpenAPI document: 1.0
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
package liff

// LiffFeatures
// LiffFeatures

// Deprecated: Use OpenAPI based classes instead.
type LiffFeatures struct {

	/**
	 * `true` if the LIFF app supports Bluetooth® Low Energy for LINE Things. `false` otherwise.
	 */
	Ble bool `json:"ble"`

	/**
	 * `true` to use the 2D code reader in the LIFF app. false otherwise. The default value is `false`.
	 */
	QrCode bool `json:"qrCode"`
}
