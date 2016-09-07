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

	EventTypeMessage  = "message"
	EventTypeFollow   = "follow"
	EventTypeUnfollow = "unfollow"
	EventTypeJoin     = "join"
	EventTypeLeave    = "leave"
	EventTypePostback = "postback"
	EventTypeBeacon   = "beacon"

	EventSourceTypeUser  = "user"
	EventSourceTypeGroup = "group"

	MessageTypeText     = "text"
	MessageTypeImage    = "image"
	MessageTypeVideo    = "video"
	MessageTypeAudio    = "audio"
	MessageTypeLocation = "location"
	MessageTypeSticker  = "sticker"
)

// EventSourceType type
type EventSourceType string

// MessageType type
type MessageType string

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

// Push method
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

// Reply method
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
	Type MessageType `json:"type"`
	Text string      `json:"text"`
}

// ImageMessage type
type ImageMessage struct {
	Type MessageType `json:"type"`
}

// NewTextMessage function
func NewTextMessage(content string) *TextMessage {
	return &TextMessage{
		Type: MessageTypeText,
		Text: content,
	}
}

// Event type
type Event struct {
	ReplyToken string      `json:"replyToken"`
	Type       EventType   `json:"type"`
	Timestamp  int64       `json:"timestamp"`
	Source     EventSource `json:"source"`
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

// EventSource type
type EventSource struct {
	Type    EventSourceType `json:"type"`
	UserID  string          `json:"userId"`
	GroupID string          `json:"groupId"`
}

// Message returns Message
func (e *Event) Message() (Message, error) {
	if e.Type != EventTypeMessage {
		return nil, ErrInvalidEventType
	}
	if e.RawMessage.Type == MessageTypeText {
		return NewTextMessage(e.RawMessage.Text), nil
	}
	return nil, ErrUnknown
}

// ParseRequest function
func (client *Client) ParseRequest(r *http.Request) (events []Event, err error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	if !client.validateSignature(r.Header.Get("X-LINE-Signature"), body) {
		return nil, ErrInvalidSignature
	}

	request := &struct {
		Events []Event `json:"events"`
	}{}
	if err = json.Unmarshal(body, request); err != nil {
		return nil, err
	}
	return request.Events, nil
}
