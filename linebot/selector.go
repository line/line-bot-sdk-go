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

// Selector interface, to be used in conjunction with logical operators
type Selector interface {
	Selector()
}

// OperatorType type
type OperatorType string

// OperatorType constants
const (
	OperatorTypeAnd OperatorType = "and"
	OperatorTypeOr  OperatorType = "or"
	OperatorTypeNot OperatorType = "not"
)

// Operator struct
type Operator struct {
	ConditionAnd []Selector `json:"and,omitempty"`
	ConditionOr  []Selector `json:"or,omitempty"`
	ConditionNot Selector   `json:"not,omitempty"`
}

// OpAnd method
func OpAnd(conditions ...Selector) *Operator {
	return &Operator{
		ConditionAnd: conditions,
	}
}

// OpOr method
func OpOr(conditions ...Selector) *Operator {
	return &Operator{
		ConditionOr: conditions,
	}
}

// OpNot method
func OpNot(condition Selector) *Operator {
	return &Operator{
		ConditionNot: condition,
	}
}

// MarshalJSON method of Operator
func (o *Operator) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type         string     `json:"type"`
		ConditionAnd []Selector `json:"and,omitempty"`
		ConditionOr  []Selector `json:"or,omitempty"`
		ConditionNot Selector   `json:"not,omitempty"`
	}{
		Type:         "operator",
		ConditionAnd: o.ConditionAnd,
		ConditionOr:  o.ConditionOr,
		ConditionNot: o.ConditionNot,
	})
}

// Selector implements Selector interface
func (*Operator) Selector() {}
