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

// Recipient interface
type Recipient interface {
	Selector
	Recipient()
}

// AudienceObject type is created to be used with specific recipient objects
type AudienceObject struct {
	Type    string `json:"type"`
	GroupID int    `json:"audienceGroupId"`
}

// NewAudienceObject function
func NewAudienceObject(groupID int) *AudienceObject {
	return &AudienceObject{
		Type:    "audience",
		GroupID: groupID,
	}
}

// Selector implements Selector interface
func (*AudienceObject) Selector() {}

// Recipient implements Recipient interface
func (*AudienceObject) Recipient() {}
