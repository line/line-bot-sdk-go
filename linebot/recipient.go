// Copyright 2020 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package linebot

import "encoding/json"

// Recipient interface
type Recipient interface {
	Recipient()
}

// AudienceObject type is created to be used with specific recipient objects
// Deprecated: Use OpenAPI based classes instead.
type AudienceObject struct {
	Type    string `json:"type"`
	GroupID int    `json:"audienceGroupId"`
}

// NewAudienceObject function
// Deprecated: Use OpenAPI based classes instead.
func NewAudienceObject(groupID int) *AudienceObject {
	return &AudienceObject{
		Type:    "audience",
		GroupID: groupID,
	}
}

// Recipient implements Recipient interface
func (*AudienceObject) Recipient() {}

// RedeliveryObject type is created to be used with specific recipient objects
// Deprecated: Use OpenAPI based classes instead.
type RedeliveryObject struct {
	Type      string `json:"type"`
	RequestID string `json:"requestId"`
}

// NewRedeliveryObject function
// Deprecated: Use OpenAPI based classes instead.
func NewRedeliveryObject(requestID string) *RedeliveryObject {
	return &RedeliveryObject{
		Type:      "redelivery",
		RequestID: requestID,
	}
}

// Recipient implements Recipient interface
func (*RedeliveryObject) Recipient() {}

// RecipientOperator struct
// Deprecated: Use OpenAPI based classes instead.
type RecipientOperator struct {
	ConditionAnd []Recipient `json:"and,omitempty"`
	ConditionOr  []Recipient `json:"or,omitempty"`
	ConditionNot Recipient   `json:"not,omitempty"`
}

// RecipientOperatorAnd method
// Deprecated: Use OpenAPI based classes instead.
func RecipientOperatorAnd(conditions ...Recipient) *RecipientOperator {
	return &RecipientOperator{
		ConditionAnd: conditions,
	}
}

// RecipientOperatorOr method
// Deprecated: Use OpenAPI based classes instead.
func RecipientOperatorOr(conditions ...Recipient) *RecipientOperator {
	return &RecipientOperator{
		ConditionOr: conditions,
	}
}

// RecipientOperatorNot method
// Deprecated: Use OpenAPI based classes instead.
func RecipientOperatorNot(condition Recipient) *RecipientOperator {
	return &RecipientOperator{
		ConditionNot: condition,
	}
}

// MarshalJSON method of Operator
func (o *RecipientOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type         string      `json:"type"`
		ConditionAnd []Recipient `json:"and,omitempty"`
		ConditionOr  []Recipient `json:"or,omitempty"`
		ConditionNot Recipient   `json:"not,omitempty"`
	}{
		Type:         "operator",
		ConditionAnd: o.ConditionAnd,
		ConditionOr:  o.ConditionOr,
		ConditionNot: o.ConditionNot,
	})
}

// Recipient implements Recipient interface
func (*RecipientOperator) Recipient() {}
