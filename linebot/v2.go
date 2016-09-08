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
	APIEndpointLeaveGroup  = "/v2/bot/group/%s/leave"
	APIEndpointLeaveRoom   = "/v2/bot/room/%s/leave"

	EventTypeMessage  = "message"
	EventTypeFollow   = "follow"
	EventTypeUnfollow = "unfollow"
	EventTypeJoin     = "join"
	EventTypeLeave    = "leave"
	EventTypePostback = "postback"
	EventTypeBeacon   = "beacon"

	EventSourceTypeUser  = "user"
	EventSourceTypeGroup = "group"
	EventSourceTypeRoom  = "room"

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

// Push method
func (client *Client) Push(to string, messages []Message) (result *ResponseContent, err error) {
	body, err := json.Marshal(&struct {
		To       string    `json:"to"`
		Messages []Message `json:"messages"`
	}{
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
	body, err := json.Marshal(&struct {
		ReplyToken string    `json:"replyToken"`
		Messages   []Message `json:"messages"`
	}{
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
type Message interface {
	MarshalJSON() ([]byte, error)
}

// TextMessage type
type TextMessage struct {
	ID   string
	Text string
}

// MarshalJSON method of TextMessage
func (m *TextMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type MessageType `json:"type"`
		Text string      `json:"text"`
	}{
		Type: MessageTypeText,
		Text: m.Text,
	})
}

// ImageMessage type
type ImageMessage struct {
	ID                 string
	OriginalContentURL string
	PreviewImageURL    string
}

// MarshalJSON method of ImageMessage
func (m *ImageMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type               MessageType `json:"type"`
		OriginalContentURL string      `json:"originalContentUrl"`
		PreviewImageURL    string      `json:"previewImageUrl"`
	}{
		Type:               MessageTypeImage,
		OriginalContentURL: m.OriginalContentURL,
		PreviewImageURL:    m.PreviewImageURL,
	})
}

// LocationMessage type
type LocationMessage struct {
	ID        string
	Title     string
	Address   string
	Latitude  float64
	Longitude float64
}

// MarshalJSON method of LocationMessage
func (m *LocationMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      MessageType `json:"type"`
		Title     string      `json:"title"`
		Address   string      `json:"address"`
		Latitude  float64     `json:"latitude"`
		Longitude float64     `json:"longitude"`
	}{
		Type:      MessageTypeLocation,
		Title:     m.Title,
		Address:   m.Address,
		Latitude:  m.Latitude,
		Longitude: m.Longitude,
	})
}

// StickerMessage type
type StickerMessage struct {
	ID        string
	PackageID string
	StickerID string
}

// MarshalJSON method of StickerMessage
func (m *StickerMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      MessageType `json:"type"`
		PackageID string      `json:"packageId"`
		StickerID string      `json:"stickerId"`
	}{
		Type:      MessageTypeSticker,
		PackageID: m.PackageID,
		StickerID: m.StickerID,
	})
}

// NewTextMessage function
func NewTextMessage(content string) *TextMessage {
	return &TextMessage{
		Text: content,
	}
}

// NewLocationMessage function
func NewLocationMessage(title, address string, latitude, longitude float64) *LocationMessage {
	return &LocationMessage{
		Title:     title,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewStickerMessage function
func NewStickerMessage(packageID, stickerID string) *StickerMessage {
	return &StickerMessage{
		PackageID: packageID,
		StickerID: stickerID,
	}
}

// EventSource type
type EventSource struct {
	Type    EventSourceType `json:"type"`
	UserID  string          `json:"userId"`
	GroupID string          `json:"groupId"`
	RoomID  string          `json:"roomId"`
}

// Event type
type Event struct {
	ReplyToken string
	Type       EventType
	Timestamp  int64
	Source     *EventSource
	Message    Message
}

// UnmarshalJSON returns a Event from JSON-encoded data.
func (e *Event) UnmarshalJSON(body []byte) (err error) {
	rawEvent := struct {
		ReplyToken string      `json:"replyToken"`
		Type       EventType   `json:"type"`
		Timestamp  int64       `json:"timestamp"`
		Source     EventSource `json:"source"`
		Message    struct {
			ID        string      `json:"id"`
			Type      MessageType `json:"type"`
			Text      string      `json:"text"`
			Title     string      `json:"title"`
			Address   string      `json:"address"`
			Latitude  float64     `json:"latitude"`
			Longitude float64     `json:"longitude"`
			PackageID string      `json:"packageId"`
			StickerID string      `json:"stickerId"`
		} `json:"message"`
	}{}
	if err = json.Unmarshal(body, &rawEvent); err != nil {
		return
	}

	e.ReplyToken = rawEvent.ReplyToken
	e.Type = rawEvent.Type
	e.Timestamp = rawEvent.Timestamp
	e.Source = &rawEvent.Source

	if rawEvent.Type == EventTypeMessage {
		switch rawEvent.Message.Type {
		case MessageTypeText:
			e.Message = &TextMessage{
				ID:   rawEvent.Message.ID,
				Text: rawEvent.Message.Text,
			}
		case MessageTypeImage:
			e.Message = &ImageMessage{
				ID: rawEvent.Message.ID,
			}
		case MessageTypeLocation:
			e.Message = &LocationMessage{
				ID:        rawEvent.Message.ID,
				Title:     rawEvent.Message.Title,
				Address:   rawEvent.Message.Address,
				Latitude:  rawEvent.Message.Latitude,
				Longitude: rawEvent.Message.Longitude,
			}
		case MessageTypeSticker:
			e.Message = &StickerMessage{
				ID:        rawEvent.Message.ID,
				PackageID: rawEvent.Message.PackageID,
				StickerID: rawEvent.Message.StickerID,
			}
		}
	}
	return
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
