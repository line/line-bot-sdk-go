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
	"context"
	"net/url"
	"strings"
)

// IssueAccessToken method
func (client *Client) IssueAccessToken(channelID, channelSecret string) *IssueAccessTokenCall {
	return &IssueAccessTokenCall{
		c:             client,
		channelID:     channelID,
		channelSecret: channelSecret,
	}
}

// IssueAccessTokenCall type
// Deprecated: Use OpenAPI based classes instead.
type IssueAccessTokenCall struct {
	c   *Client
	ctx context.Context

	channelID     string
	channelSecret string
}

// WithContext method
func (call *IssueAccessTokenCall) WithContext(ctx context.Context) *IssueAccessTokenCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *IssueAccessTokenCall) Do() (*AccessTokenResponse, error) {
	vs := url.Values{}
	vs.Set("grant_type", "client_credentials")
	vs.Set("client_id", call.channelID)
	vs.Set("client_secret", call.channelSecret)
	body := strings.NewReader(vs.Encode())

	res, err := call.c.postForm(call.ctx, APIEndpointIssueAccessToken, body)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToAccessTokenResponse(res)
}

// RevokeAccessToken method
func (client *Client) RevokeAccessToken(accessToken string) *RevokeAccessTokenCall {
	return &RevokeAccessTokenCall{
		c:           client,
		accessToken: accessToken,
	}
}

// RevokeAccessTokenCall type
// Deprecated: Use OpenAPI based classes instead.
type RevokeAccessTokenCall struct {
	c   *Client
	ctx context.Context

	accessToken string
}

// WithContext method
func (call *RevokeAccessTokenCall) WithContext(ctx context.Context) *RevokeAccessTokenCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *RevokeAccessTokenCall) Do() (*BasicResponse, error) {
	vs := url.Values{}
	vs.Set("access_token", call.accessToken)
	body := strings.NewReader(vs.Encode())

	res, err := call.c.postForm(call.ctx, APIEndpointRevokeAccessToken, body)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// VerifyAccessToken method
func (client *Client) VerifyAccessToken(accessToken string) *VerifyAccessTokenCall {
	return &VerifyAccessTokenCall{
		c:           client,
		accessToken: accessToken,
	}
}

// VerifyAccessTokenCall type
// Deprecated: Use OpenAPI based classes instead.
type VerifyAccessTokenCall struct {
	c   *Client
	ctx context.Context

	accessToken string
}

// WithContext method
func (call *VerifyAccessTokenCall) WithContext(ctx context.Context) *VerifyAccessTokenCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *VerifyAccessTokenCall) Do() (*VerifiedAccessTokenResponse, error) {
	vs := url.Values{}
	vs.Set("access_token", call.accessToken)
	body := strings.NewReader(vs.Encode())

	res, err := call.c.postForm(call.ctx, APIEndpointVerifyAccessToken, body)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToVerifiedAccessTokenResponse(res)
}
