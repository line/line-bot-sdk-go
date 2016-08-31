package linebot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// constants
const (
	APIEndpointEventsPush  = "/v2/bot/message/push"
	APIEndpointEventsReply = "/v2/bot/message/reply"
)

// Push function
func (client *Client) Push(to []string, messages []Message) (result *ResponseContent, err error) {
	body, err := json.Marshal(PushMessage{
		To:       to,
		Messages: messages,
	})
	if err != nil {
		return
	}
	println(string(body))
	result, err = client.post(APIEndpointEventsPush, bytes.NewReader(body))
	return
}

// ResponseContent type
type ResponseContent struct {
	RequestID string `json:"requestId"`
	Message   string `json:"message"`
	Details   []struct {
		Message  string `json:"message"`
		Property string `json:"property"`
	} `json:"details"`
}

// PushMessage type
type PushMessage struct {
	To       []string  `json:"to"`
	Messages []Message `json:"messages"`
}

// Message inteface
type Message interface {
}

// TextMessage type
type TextMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// NewTextMessage function
func NewTextMessage(content string) *TextMessage {
	return &TextMessage{
		Type: "text",
		Text: content,
	}
}

// ReceivedEvents type
type ReceivedEvents struct {
	Events []interface{} `json:"events"`
}

// ParseRequest function
func (client *Client) ParseRequest(r *http.Request) (events *ReceivedEvents, err error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	if !client.validateSignature(r.Header.Get("X-LINE-ChannelSignature"), body) {
		return nil, ErrInvalidSignature
	}

	events = &ReceivedEvents{}
	if err = json.Unmarshal(body, events); err != nil {
		return nil, err
	}
	return
}
