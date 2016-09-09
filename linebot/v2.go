package linebot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
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
func (client *Client) Push(to string, messages []Message) *PushCall {
	return &PushCall{
		c:        client,
		to:       to,
		messages: messages,
	}
}

// Reply method
func (client *Client) Reply(replyToken string, messages []Message) *ReplyCall {
	return &ReplyCall{
		c:          client,
		replyToken: replyToken,
		messages:   messages,
	}
}

// ResponseContent type
// Duplicated
type ResponseContent struct {
	RequestID string `json:"requestId"`
	Message   string `json:"message"`
	Details   []struct {
		Message  string `json:"message"`
		Property string `json:"property"`
	} `json:"details"`
}

// BasicResponse type
type BasicResponse struct {
	RequestID string `json:"requestId"`
}

// ErrorResponse type
type ErrorResponse struct {
	RequestID string `json:"requestId"`
	Message   string `json:"message"`
	Details   []struct {
		Message  string `json:"message"`
		Property string `json:"property"`
	} `json:"details"`
}

// APIError type
type APIError struct {
	Code     int
	Response *ErrorResponse
}

func (e *APIError) Error() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "linebot: APIError %d ", e.Code)
	if e.Response != nil {
		fmt.Fprintf(&buf, "%s", e.Response.Message)
		for _, d := range e.Response.Details {
			fmt.Fprintf(&buf, "\n[%s] %s", d.Property, d.Message)
		}
	}
	return buf.String()
}

// PushCall type
type PushCall struct {
	c   *Client
	ctx context.Context

	to       string
	messages []Message
}

// WithContext method
func (call *PushCall) WithContext(ctx context.Context) *PushCall {
	call.ctx = ctx
	return call
}

func (call *PushCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		To       string    `json:"to"`
		Messages []Message `json:"messages"`
	}{
		To:       call.to,
		Messages: call.messages,
	})
}

func decodeToBasicResponse(res *http.Response) (*BasicResponse, error) {
	decoder := json.NewDecoder(res.Body)
	if res.StatusCode != http.StatusOK {
		result := ErrorResponse{}
		if err := decoder.Decode(&result); err != nil {
			return nil, &APIError{
				Code: res.StatusCode,
			}
		}
		return nil, &APIError{
			Code:     res.StatusCode,
			Response: &result,
		}
	}
	result := BasicResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Do method
func (call *PushCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.postCtx(call.ctx, APIEndpointEventsPush, &buf)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToBasicResponse(res)
}

// ReplyCall type
type ReplyCall struct {
	c   *Client
	ctx context.Context

	replyToken string
	messages   []Message
}

// WithContext method
func (call *ReplyCall) WithContext(ctx context.Context) *ReplyCall {
	call.ctx = ctx
	return call
}

func (call *ReplyCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		ReplyToken string    `json:"replyToken"`
		Messages   []Message `json:"messages"`
	}{
		ReplyToken: call.replyToken,
		Messages:   call.messages,
	})
}

// Do method
func (call *ReplyCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.postCtx(call.ctx, APIEndpointEventsReply, &buf)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToBasicResponse(res)
}

func (client *Client) doCtx(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("X-LINE-ChannelToken", client.channelToken)
	req.Header.Set("Authorization", "Bearer "+client.channelToken)
	if ctx == nil {
		return client.httpClient.Do(req)
	}
	return ctxhttp.Do(ctx, client.httpClient, req)
}

func (client *Client) postCtx(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	url, err := client.url(endpoint)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return client.doCtx(ctx, req)
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

// UnmarshalJSON constructs a Event from JSON-encoded data.
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
