package linebot

import (
	"io"
	"net/http"
	"net/url"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

// APIEndpoint constants
const (
	APIEndpointBase = "https://trialbot-api.line.me"

	APIEndpointEventsPush     = "/v2/bot/message/push"
	APIEndpointEventsReply    = "/v2/bot/message/reply"
	APIEndpointMessageContent = "/v2/bot/message/%s/content"
	APIEndpointLeaveGroup     = "/v2/bot/group/%s/leave"
	APIEndpointLeaveRoom      = "/v2/bot/room/%s/leave"
	APIEndpointGetProfile     = "/v2/bot/profile/%s"
)

// Client type
type Client struct {
	channelSecret string
	channelToken  string
	endpointBase  string       // default APIEndpointBase
	httpClient    *http.Client // default http.DefaultClient
}

// ClientOption type
type ClientOption func(*Client) error

// New returns a new bot client instance.
func New(channelSecret, channelToken string, options ...ClientOption) (*Client, error) {
	c := &Client{
		channelSecret: channelSecret,
		channelToken:  channelToken,
		endpointBase:  APIEndpointBase,
		httpClient:    http.DefaultClient,
	}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
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
		client.endpointBase = endpointBase
		return nil
	}
}

func (client *Client) url(endpoint string) (url *url.URL, err error) {
	url, err = url.Parse(client.endpointBase)
	if err != nil {
		return
	}
	url.Path = endpoint
	return
}

func (client *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("X-LINE-ChannelToken", client.channelToken)
	req.Header.Set("Authorization", "Bearer "+client.channelToken)
	req.Header.Set("User-Agent", "LINE-BotSDK-Go/"+version)
	if ctx == nil {
		return client.httpClient.Do(req)
	}
	return ctxhttp.Do(ctx, client.httpClient, req)
}

func (client *Client) get(ctx context.Context, endpoint string) (res *http.Response, err error) {
	url, err := client.url(endpoint)
	if err != nil {
		return
	}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return
	}
	return client.do(ctx, req)
}

func (client *Client) post(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	url, err := client.url(endpoint)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return client.do(ctx, req)
}
