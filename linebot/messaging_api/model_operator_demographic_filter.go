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

import (
	"encoding/json"
	"fmt"
)

// OperatorDemographicFilter
// OperatorDemographicFilter

type OperatorDemographicFilter struct {
	DemographicFilter

	/**
	 * Get And
	 */
	And []DemographicFilterInterface `json:"and,omitempty"`

	/**
	 * Get Or
	 */
	Or []DemographicFilterInterface `json:"or,omitempty"`

	/**
	 * Get Not
	 */
	Not DemographicFilterInterface `json:"not,omitempty"`
}

func (cr *OperatorDemographicFilter) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw["type"], &cr.Type)
	if err != nil {
		return err
	}

	var rawand []json.RawMessage
	err = json.Unmarshal(raw["and"], &rawand)
	if err != nil {
		return err
	}

	for _, data := range rawand {
		e, err := UnmarshalDemographicFilter(data)
		if err != nil {
			return fmt.Errorf("JSON parse error in UnmarshalDemographicFilter: %w, body: %s", err, string(data))
		}
		cr.And = append(cr.And, e)
	}

	var rawor []json.RawMessage
	err = json.Unmarshal(raw["or"], &rawor)
	if err != nil {
		return err
	}

	for _, data := range rawor {
		e, err := UnmarshalDemographicFilter(data)
		if err != nil {
			return fmt.Errorf("JSON parse error in UnmarshalDemographicFilter: %w, body: %s", err, string(data))
		}
		cr.Or = append(cr.Or, e)
	}

	if rawnot, ok := raw["not"]; ok && rawnot != nil {
		Not, err := UnmarshalDemographicFilter(rawnot)
		if err != nil {
			return err
		}
		cr.Not = Not
	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the OperatorDemographicFilter struct.
func (r *OperatorDemographicFilter) MarshalJSON() ([]byte, error) {

	newAnd := make([]DemographicFilterInterface, len(r.And))
	for i, v := range r.And {
		newAnd[i] = setDiscriminatorPropertyDemographicFilter(v)
	}

	newOr := make([]DemographicFilterInterface, len(r.Or))
	for i, v := range r.Or {
		newOr[i] = setDiscriminatorPropertyDemographicFilter(v)
	}

	r.Not = setDiscriminatorPropertyDemographicFilter(r.Not)

	type Alias OperatorDemographicFilter
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type,omitempty"`

		And []DemographicFilterInterface `json:"and,omitempty"`

		Or []DemographicFilterInterface `json:"or,omitempty"`
	}{
		Alias: (*Alias)(r),

		Type: "operator",

		And: newAnd,

		Or: newOr,
	})
}
