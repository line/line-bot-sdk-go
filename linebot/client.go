package linebot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// errors
var (
	ErrInvalidSignature   = errors.New("Invalid Signature")
	ErrInvalidContentType = errors.New("Invalid ContentType")
	ErrInvalidEventType   = errors.New("Invalid EventType")
)

// Client type
type Client struct {
	channelID     int64
	channelSecret string
	mid           string
	endpointBase  string       // default APIEndpointBaseTrial
	httpClient    *http.Client // default http.DefaultClient
}

// ClientOption type
type ClientOption func(*Client) error

// NewClient function
func NewClient(channelID int64, channelSecret, mid string, options ...ClientOption) (*Client, error) {
	c := &Client{
		channelID:     channelID,
		channelSecret: channelSecret,
		mid:           mid,
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

	if res.StatusCode != http.StatusOK {
		var content ErrorResponseContent
		if err = decoder.Decode(&content); err != nil {
			return
		}
		return nil, fmt.Errorf("%s: %s", content.Code, content.Message)
	}

	result = &ResponseContent{}
	err = decoder.Decode(result)
	if err != nil {
		return
	}
	return
}

func (client *Client) do(req *http.Request) (res *http.Response, err error) {
	req.Header.Set("X-Line-ChannelID", strconv.FormatInt(client.channelID, 10))
	req.Header.Set("X-Line-ChannelSecret", client.channelSecret)
	req.Header.Set("X-Line-Trusted-User-With-ACL", client.mid)
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
