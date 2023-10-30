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

// AgeTile
// AgeTile

type AgeTile struct {

	/**
	 * users&#39; age
	 */
	Age AgeTileAGE `json:"age,omitempty"`

	/**
	 * Percentage
	 */
	Percentage float64 `json:"percentage"`
}

// AgeTileAGE type
/* users' age */
type AgeTileAGE string

// AgeTileAGE constants
const (
	AgeTileAGE_FROM0TO14 AgeTileAGE = "from0to14"

	AgeTileAGE_FROM15TO19 AgeTileAGE = "from15to19"

	AgeTileAGE_FROM20TO24 AgeTileAGE = "from20to24"

	AgeTileAGE_FROM25TO29 AgeTileAGE = "from25to29"

	AgeTileAGE_FROM30TO34 AgeTileAGE = "from30to34"

	AgeTileAGE_FROM35TO39 AgeTileAGE = "from35to39"

	AgeTileAGE_FROM40TO44 AgeTileAGE = "from40to44"

	AgeTileAGE_FROM45TO49 AgeTileAGE = "from45to49"

	AgeTileAGE_FROM50 AgeTileAGE = "from50"

	AgeTileAGE_UNKNOWN AgeTileAGE = "unknown"
)
