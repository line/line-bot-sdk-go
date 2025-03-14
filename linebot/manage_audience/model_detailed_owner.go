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

// DetailedOwner
// Owner of this audience group.

type DetailedOwner struct {

	/**
	 * Service name where the audience group has been created.
	 */
	ServiceType string `json:"serviceType,omitempty"`

	/**
	 * Owner ID in the service.
	 */
	Id string `json:"id,omitempty"`

	/**
	 * Owner account name.
	 */
	Name string `json:"name,omitempty"`
}
