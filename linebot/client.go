package linebot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// errors
var (
	ErrInvalidSignature   = errors.New("Invalid Signature")
	ErrInvalidContentType = errors.New("Invalid ContentType")
	ErrInvalidEventType   = errors.New("Invalid EventType")
)

// Client type
type Client struct {
	channelSecret string
	channelToken  string
	endpointBase  string       // default APIEndpointBaseTrial
	httpClient    *http.Client // default http.DefaultClient
}

// ClientOption type
type ClientOption func(*Client) error

// NewClient function
func NewClient(channelSecret, channelToken string, options ...ClientOption) (*Client, error) {
	c := &Client{
		channelSecret: channelSecret,
		channelToken:  channelToken,
		endpointBase:  APIEndpointBaseTrial,
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

func (client *Client) sendSingleMessage(to []string, content SingleMessageContent) (result *ResponseContent, err error) {
	message := SingleMessage{
		To:        to,
		ToChannel: SendingMessageChannelID,
		EventType: EventTypeSendingMessage,
		Content:   content,
	}
	body, err := json.Marshal(message)
	if err != nil {
		return
	}
	result, err = client.post(APIEndpointEvents, bytes.NewReader(body))
	return
}

func (client *Client) sendMultipleMessage(to []string, content MultipleMessageContent) (result *ResponseContent, err error) {
	message := MultipleMessage{
		To:        to,
		ToChannel: SendingMessageChannelID,
		EventType: EventTypeSendingMultipleMessage,
		Content:   content,
	}
	body, err := json.Marshal(message)
	if err != nil {
		return
	}
	result, err = client.post(APIEndpointEvents, bytes.NewReader(body))
	return
}

func (client *Client) get(endpoint, rawQuery string) (res *http.Response, err error) {
	url, err := client.url(endpoint)
	if err != nil {
		return
	}
	url.RawQuery = rawQuery
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return
	}
	return client.do(req)
}

func (client *Client) post(endpoint string, body io.Reader) (result *ResponseContent, err error) {
	url, err := client.url(endpoint)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", url.String(), body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	res, err := client.do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	result = &ResponseContent{}
	err = decoder.Decode(result)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		// TODO: error details
		return nil, fmt.Errorf("%d: %s", res.StatusCode, result.Message)
	}

	return
}

func (client *Client) do(req *http.Request) (res *http.Response, err error) {
	req.Header.Set("X-LINE-ChannelToken", client.channelToken)
	req.Header.Set("Authorization", "Bearer "+client.channelToken)
	res, err = client.httpClient.Do(req)
	return
}

func (client *Client) url(endpoint string) (url *url.URL, err error) {
	url, err = url.Parse(client.endpointBase)
	if err != nil {
		return
	}
	url.Path = endpoint
	return
}
