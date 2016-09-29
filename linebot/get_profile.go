// Copyright 2016 LINE Corporation
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

import (
	"fmt"

	"golang.org/x/net/context"
)

// GetProfile method
func (client *Client) GetProfile(userID string) *GetProfileCall {
	return &GetProfileCall{
		c:      client,
		userID: userID,
	}
}

// GetProfileCall type
type GetProfileCall struct {
	c   *Client
	ctx context.Context

	userID string
}

// WithContext method
func (call *GetProfileCall) WithContext(ctx context.Context) *GetProfileCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetProfileCall) Do() (*UserProfileResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointGetProfile, call.userID)
	res, err := call.c.get(call.ctx, endpoint)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToUserProfileResponse(res)
}
