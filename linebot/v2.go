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

	EventMessageTypeText     = "text"
	EventMessageTypeImage    = "image"
	EventMessageTypeVideo    = "video"
	EventMessageTypeAudio    = "audio"
	EventMessageTypeLocation = "location"
	EventMessageTypeSticker  = "sticker"
)

// PushMessage type
type PushMessage struct {
	To       string    `json:"to"`
	Messages []Message `json:"messages"`
}

// ReplyMessage type
type ReplyMessage struct {
	ReplyToken string    `json:"replyToken"`
	Messages   []Message `json:"messages"`
}

// Push function
func (client *Client) Push(to string, messages []Message) (result *ResponseContent, err error) {
	body, err := json.Marshal(PushMessage{
		To:       to,
		Messages: messages,
	})
	if err != nil {
		return
	}
	result, err = client.post(APIEndpointEventsPush, bytes.NewReader(body))
	return
}

// Reply function
func (client *Client) Reply(token string, messages []Message) (result *ResponseContent, err error) {
	body, err := json.Marshal(ReplyMessage{
		ReplyToken: token,
		Messages:   messages,
	})
	if err != nil {
		return
	}
	result, err = client.post(APIEndpointEventsReply, bytes.NewReader(body))
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

// Message inteface
type Message interface{}

// TextMessage type
type TextMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// NewTextMessage function
func NewTextMessage(content string) *TextMessage {
	return &TextMessage{
		Type: EventMessageTypeText,
		Text: content,
	}
}

// ReceivedEvent type
type ReceivedEvent struct {
	ReplyToken string `json:"replyToken"`
	Type       string `json:"type"`
	Timestamp  int64  `json:"timestamp"`
	RawSource  struct {
		Type    string `json:"type"`
		UserID  string `json:"userId"`
		GroupID string `json:"groupId"`
	} `json:"source"`
	RawMessage struct {
		ID        string  `json:"id"`
		Type      string  `json:"type"`
		Text      string  `json:"text"`
		Title     string  `json:"title"`
		Address   string  `json:"address"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		PackageID string  `json:"packageId"`
		StickerID string  `json:"stickerId"`
	} `json:"message"`
}

// TextMessage function
func (e *ReceivedEvent) TextMessage() (*TextMessage, error) {
	if e.RawMessage.Type != EventMessageTypeText {
		return nil, ErrInvalidContentType
	}
	return NewTextMessage(e.RawMessage.Text), nil
}

// ParseRequest function
func (client *Client) ParseRequest(r *http.Request) (events []ReceivedEvent, err error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	if !client.validateSignature(r.Header.Get("X-LINE-Signature"), body) {
		return nil, ErrInvalidSignature
	}

	request := &struct {
		Events []ReceivedEvent `json:"events"`
	}{}
	if err = json.Unmarshal(body, request); err != nil {
		return nil, err
	}
	return request.Events, nil
}
