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
	"encoding/json"
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/context"
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

func (call *IssueAccessTokenCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		GrantType    string `json:"grant_type"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{
		GrantType:    "client_credentials",
		ClientID:     call.channelID,
		ClientSecret: call.channelSecret,
	})
}

// Do method
func (call *IssueAccessTokenCall) Do() (*AccessTokenResponse, error) {
	vs := url.Values{}
	vs.Set("grant_type", "client_credentials")
	vs.Set("client_id", call.channelID)
	vs.Set("client_secret", call.channelSecret)
	body := strings.NewReader(vs.Encode())

	res, err := call.c.postform(call.ctx, APIEndpointIssueAccessToken, body)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToAccessTokenResponse(res)
}
