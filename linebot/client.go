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
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"path"
	"time"
)

// APIEndpoint constants
const (
	APIEndpointBase     = "https://api.line.me"
	APIEndpointBaseData = "https://api-data.line.me"

	APIEndpointPushMessage                = "/v2/bot/message/push"
	APIEndpointBroadcastMessage           = "/v2/bot/message/broadcast"
	APIEndpointReplyMessage               = "/v2/bot/message/reply"
	APIEndpointMulticast                  = "/v2/bot/message/multicast"
	APIEndpointNarrowcast                 = "/v2/bot/message/narrowcast"
	APIEndpointValidatePushMessage        = "/v2/bot/message/validate/push"
	APIEndpointValidateBroadcastMessage   = "/v2/bot/message/validate/broadcast"
	APIEndpointValidateReplyMessage       = "/v2/bot/message/validate/reply"
	APIEndpointValidateMulticastMessage   = "/v2/bot/message/validate/multicast"
	APIEndpointValidateNarrowcastMessage  = "/v2/bot/message/validate/narrowcast"
	APIEndpointGetMessageContent          = "/v2/bot/message/%s/content"
	APIEndpointGetMessageQuota            = "/v2/bot/message/quota"
	APIEndpointGetMessageConsumption      = "/v2/bot/message/quota/consumption"
	APIEndpointGetMessageQuotaConsumption = "/v2/bot/message/quota/consumption"
	APIEndpointLeaveGroup                 = "/v2/bot/group/%s/leave"
	APIEndpointLeaveRoom                  = "/v2/bot/room/%s/leave"
	APIEndpointGetProfile                 = "/v2/bot/profile/%s"
	APIEndpointGetFollowerIDs             = "/v2/bot/followers/ids"
	APIEndpointGetGroupMemberProfile      = "/v2/bot/group/%s/member/%s"
	APIEndpointGetRoomMemberProfile       = "/v2/bot/room/%s/member/%s"
	APIEndpointGetGroupMemberIDs          = "/v2/bot/group/%s/members/ids"
	APIEndpointGetRoomMemberIDs           = "/v2/bot/room/%s/members/ids"
	APIEndpointGetGroupMemberCount        = "/v2/bot/group/%s/members/count"
	APIEndpointGetRoomMemberCount         = "/v2/bot/room/%s/members/count"
	APIEndpointGetGroupSummary            = "/v2/bot/group/%s/summary"
	APIEndpointCreateRichMenu             = "/v2/bot/richmenu"
	APIEndpointGetRichMenu                = "/v2/bot/richmenu/%s"
	APIEndpointListRichMenu               = "/v2/bot/richmenu/list"
	APIEndpointDeleteRichMenu             = "/v2/bot/richmenu/%s"
	APIEndpointGetUserRichMenu            = "/v2/bot/user/%s/richmenu"
	APIEndpointLinkUserRichMenu           = "/v2/bot/user/%s/richmenu/%s"
	APIEndpointUnlinkUserRichMenu         = "/v2/bot/user/%s/richmenu"
	APIEndpointSetDefaultRichMenu         = "/v2/bot/user/all/richmenu/%s"
	APIEndpointDefaultRichMenu            = "/v2/bot/user/all/richmenu"   // Get: GET / Delete: DELETE
	APIEndpointDownloadRichMenuImage      = "/v2/bot/richmenu/%s/content" // Download: GET / Upload: POST
	APIEndpointUploadRichMenuImage        = "/v2/bot/richmenu/%s/content" // Download: GET / Upload: POST
	APIEndpointBulkLinkRichMenu           = "/v2/bot/richmenu/bulk/link"
	APIEndpointBulkUnlinkRichMenu         = "/v2/bot/richmenu/bulk/unlink"
	APIEndpointValidateRichMenuObject     = "/v2/bot/richmenu/validate"

	APIEndpointCreateRichMenuAlias = "/v2/bot/richmenu/alias"
	APIEndpointGetRichMenuAlias    = "/v2/bot/richmenu/alias/%s"
	APIEndpointUpdateRichMenuAlias = "/v2/bot/richmenu/alias/%s"
	APIEndpointDeleteRichMenuAlias = "/v2/bot/richmenu/alias/%s"
	APIEndpointListRichMenuAlias   = "/v2/bot/richmenu/alias/list"

	APIEndpointGetAllLIFFApps = "/liff/v1/apps"
	APIEndpointAddLIFFApp     = "/liff/v1/apps"
	APIEndpointUpdateLIFFApp  = "/liff/v1/apps/%s/view"
	APIEndpointDeleteLIFFApp  = "/liff/v1/apps/%s"

	APIEndpointLinkToken = "/v2/bot/user/%s/linkToken"

	APIEndpointGetMessageDelivery = "/v2/bot/message/delivery/%s"
	APIEndpointGetMessageProgress = "/v2/bot/message/progress/%s"
	APIEndpointInsight            = "/v2/bot/insight/%s"
	APIEndpointGetBotInfo         = "/v2/bot/info"

	APIEndpointIssueAccessToken  = "/v2/oauth/accessToken"
	APIEndpointRevokeAccessToken = "/v2/oauth/revoke"
	APIEndpointVerifyAccessToken = "/v2/oauth/verify"

	APIEndpointIssueAccessTokenV2  = "/oauth2/v2.1/token"
	APIEndpointGetAccessTokensV2   = "/oauth2/v2.1/tokens/kid"
	APIEndpointRevokeAccessTokenV2 = "/oauth2/v2.1/revoke"

	APIEndpointGetWebhookInfo     = "/v2/bot/channel/webhook/endpoint"
	APIEndpointSetWebhookEndpoint = "/v2/bot/channel/webhook/endpoint"
	APIEndpointTestWebhook        = "/v2/bot/channel/webhook/test"

	APIAudienceGroupUpload            = "/v2/bot/audienceGroup/upload"
	APIAudienceGroupUploadByFile      = "/v2/bot/audienceGroup/upload/byFile"
	APIAudienceGroupClick             = "/v2/bot/audienceGroup/click"
	APIAudienceGroupIMP               = "/v2/bot/audienceGroup/imp"
	APIAudienceGroupUpdateDescription = "/v2/bot/audienceGroup/%d/updateDescription"
	APIAudienceGroupActivate          = "/v2/bot/audienceGroup/%d/activate"
	APIAudienceGroup                  = "/v2/bot/audienceGroup/%d"
	APIAudienceGroupList              = "/v2/bot/audienceGroup/list"
	APIAudienceGroupAuthorityLevel    = "/v2/bot/audienceGroup/authorityLevel"
)

// Client type
// Deprecated: use OpenAPI based instead.
// Deprecated: Use OpenAPI based classes instead.
type Client struct {
	channelSecret    string
	channelToken     string
	endpointBase     *url.URL     // default APIEndpointBase
	endpointBaseData *url.URL     // default APIEndpointBaseData
	httpClient       *http.Client // default http.DefaultClient
	retryKeyID       string       // X-Line-Retry-Key allows you to safely retry API requests without duplicating messages
}

// ClientOption type
type ClientOption func(*Client) error

// New returns a new bot client instance.
// Deprecated: Use OpenAPI based classes instead.
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
	if c.endpointBaseData == nil {
		u, err := url.ParseRequestURI(APIEndpointBaseData)
		if err != nil {
			return nil, err
		}
		c.endpointBaseData = u
	}
	return c, nil
}

// WithHTTPClient function
// Deprecated: Use OpenAPI based classes instead.
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = c
		return nil
	}
}

// WithEndpointBase function
// Deprecated: Use OpenAPI based classes instead.
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

// WithEndpointBaseData function
// Deprecated: Use OpenAPI based classes instead.
func WithEndpointBaseData(endpointBaseData string) ClientOption {
	return func(client *Client) error {
		u, err := url.ParseRequestURI(endpointBaseData)
		if err != nil {
			return err
		}
		client.endpointBaseData = u
		return nil
	}
}

func (client *Client) url(base *url.URL, endpoint string) string {
	u := *base
	u.Path = path.Join(u.Path, endpoint)
	return u.String()
}

func (client *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+client.channelToken)
	req.Header.Set("User-Agent", "LINE-BotSDK-Go/"+version)
	if len(client.retryKeyID) > 0 {
		req.Header.Set("X-Line-Retry-Key", client.retryKeyID)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return client.httpClient.Do(req)
}

func (client *Client) get(ctx context.Context, base *url.URL, endpoint string, query url.Values) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, client.url(base, endpoint), nil)
	if err != nil {
		return nil, err
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	return client.do(ctx, req)
}

func (client *Client) post(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, client.url(client.endpointBase, endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return client.do(ctx, req)
}

func (client *Client) postForm(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", client.url(client.endpointBase, endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return client.do(ctx, req)
}

func (client *Client) postFormFile(ctx context.Context, endpoint string, values map[string]io.Reader) (*http.Response, error) {
	b, contentType, err := uploadFile(values)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, client.url(client.endpointBaseData, endpoint), &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return client.do(ctx, req)
}

func (client *Client) put(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, client.url(client.endpointBase, endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return client.do(ctx, req)
}

func (client *Client) putFormFile(ctx context.Context, endpoint string, values map[string]io.Reader) (*http.Response, error) {
	b, contentType, err := uploadFile(values)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, client.url(client.endpointBaseData, endpoint), &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return client.do(ctx, req)
}

func (client *Client) delete(ctx context.Context, endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, client.url(client.endpointBase, endpoint), nil)
	if err != nil {
		return nil, err
	}
	return client.do(ctx, req)
}

func (client *Client) setRetryKey(retryKey string) {
	client.retryKeyID = retryKey
}

// Deprecated: Use OpenAPI based classes instead.
func closeResponse(res *http.Response) error {
	defer res.Body.Close()
	_, err := io.Copy(io.Discard, res.Body)
	return err
}

// Deprecated: Use OpenAPI based classes instead.
func uploadFile(values map[string]io.Reader) (bytes.Buffer, string, error) {
	var (
		b   bytes.Buffer
		err error
	)
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if _, ok := r.(*bytes.Buffer); ok {
			h := make(textproto.MIMEHeader)
			h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s.txt"`, key, time.Now().Format("20060102150405")))
			h.Set("Content-Type", "text/plain")
			if fw, err = w.CreatePart(h); err != nil {
				return b, "", err
			}
		} else {
			if fw, err = w.CreateFormField(key); err != nil {
				return b, "", err
			}
		}
		if _, err := io.Copy(fw, r); err != nil {
			return b, "", err
		}
	}
	w.Close()
	return b, w.FormDataContentType(), nil
}
