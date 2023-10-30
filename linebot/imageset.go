// Copyright 2021 LINE Corporation
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

// ImageSet type
// Deprecated: Use OpenAPI based classes instead.
type ImageSet struct {
	ID    string `json:"id"`
	Index int    `json:"index"`
	Total int    `json:"total"`
}

// NewImageSet function
// Deprecated: Use OpenAPI based classes instead.
func NewImageSet(ID string, index, total int) *ImageSet {
	return &ImageSet{
		ID:    ID,
		Index: index,
		Total: total,
	}
}
