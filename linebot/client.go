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
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"
)

// APIEndpoint constants
const (
	APIEndpointBase = "https://api.line.me"

	APIEndpointPushMessage           = "/v2/bot/message/push"
	APIEndpointReplyMessage          = "/v2/bot/message/reply"
	APIEndpointMulticast             = "/v2/bot/message/multicast"
	APIEndpointGetMessageContent     = "/v2/bot/message/%s/content"
	APIEndpointLeaveGroup            = "/v2/bot/group/%s/leave"
	APIEndpointLeaveRoom             = "/v2/bot/room/%s/leave"
	APIEndpointGetProfile            = "/v2/bot/profile/%s"
	APIEndpointGetGroupMemberProfile = "/v2/bot/group/%s/member/%s"
	APIEndpointGetRoomMemberProfile  = "/v2/bot/room/%s/member/%s"
	APIEndpointGetGroupMemberIDs     = "/v2/bot/group/%s/members/ids"
	APIEndpointGetRoomMemberIDs      = "/v2/bot/room/%s/members/ids"
	APIEndpointCreateRichMenu        = "/v2/bot/richmenu"
	APIEndpointGetRichMenu           = "/v2/bot/richmenu/%s"
	APIEndpointListRichMenu          = "/v2/bot/richmenu/list"
	APIEndpointDeleteRichMenu        = "/v2/bot/richmenu/%s"
	APIEndpointGetUserRichMenu       = "/v2/bot/user/%s/richmenu"
	APIEndpointLinkUserRichMenu      = "/v2/bot/user/%s/richmenu/%s"
	APIEndpointUnlinkUserRichMenu    = "/v2/bot/user/%s/richmenu"
	APIEndpointDownloadRichMenuImage = "/v2/bot/richmenu/%s/content" // Download: GET / Upload: POST
	APIEndpointUploadRichMenuImage   = "/v2/bot/richmenu/%s/content" // Download: GET / Upload: POST

	APIEndpointGetAllLIFFApps = "/liff/v1/apps"
	APIEndpointAddLIFFApp     = "/liff/v1/apps"
	APIEndpointUpdateLIFFApp  = "/liff/v1/apps/%s/view"
	APIEndpointDeleteLIFFApp  = "/liff/v1/apps/%s"

	APIEndpointLinkToken = "/v2/bot/user/%s/linkToken"
)

// Client type
type Client struct {
	channelSecret string
	channelToken  string
	endpointBase  *url.URL     // default APIEndpointBase
	httpClient    *http.Client // default http.DefaultClient
}

// ClientOption type
type ClientOption func(*Client) error

// New returns a new bot client instance.
func New(channelSecret, channelToken string, options ...ClientOption) (*Client, error) {
	if channelSecret == "" {
		return nil, errors.New("missing channel secret")
	}
	if channelToken == "" {
		return nil, errors.New("missing channel access token")
	}
	c := &Client{
		channelSecret: channelSecret,
		channelToken:  channelToken,
		httpClient:    http.DefaultClient,
	}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	if c.endpointBase == nil {
		u, err := url.ParseRequestURI(APIEndpointBase)
		if err != nil {
			return nil, err
		}
		c.endpointBase = u
	}
	return c, nil
}

// WithHTTPClient function
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = c
		return nil
	}
}

// WithEndpointBase function
func WithEndpointBase(endpointBase string) ClientOption {
	return func(client *Client) error {
		u, err := url.ParseRequestURI(endpointBase)
		if err != nil {
			return err
		}
		client.endpointBase = u
		return nil
	}
}

func (client *Client) url(endpoint string) string {
	u := *client.endpointBase
	u.Path = path.Join(u.Path, endpoint)
	return u.String()
}

func (client *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+client.channelToken)
	req.Header.Set("User-Agent", "LINE-BotSDK-Go/"+version)
	if ctx != nil {
		res, err := client.httpClient.Do(req.WithContext(ctx))
		if err != nil {
			select {
			case <-ctx.Done():
				err = ctx.Err()
			default:
			}
		}

		return res, err
	}
	return client.httpClient.Do(req)

}

func (client *Client) get(ctx context.Context, endpoint string, query url.Values) (*http.Response, error) {
	req, err := http.NewRequest("GET", client.url(endpoint), nil)
	if err != nil {
		return nil, err
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	return client.do(ctx, req)
}

func (client *Client) post(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", client.url(endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return client.do(ctx, req)
}

func (client *Client) put(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("PUT", client.url(endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return client.do(ctx, req)
}

func (client *Client) delete(ctx context.Context, endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", client.url(endpoint), nil)
	if err != nil {
		return nil, err
	}
	return client.do(ctx, req)
}
