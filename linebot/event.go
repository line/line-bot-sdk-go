package linebot

import (
	"encoding/json"
)

// EventType type
type EventType string

// EventType constants
const (
	EventTypeMessage  = "message"
	EventTypeFollow   = "follow"
	EventTypeUnfollow = "unfollow"
	EventTypeJoin     = "join"
	EventTypeLeave    = "leave"
	EventTypePostback = "postback"
	EventTypeBeacon   = "beacon"
)

// EventSourceType type
type EventSourceType string

// EventSourceType constants
const (
	EventSourceTypeUser  = "user"
	EventSourceTypeGroup = "group"
	EventSourceTypeRoom  = "room"
)

// EventSource type
type EventSource struct {
	Type    EventSourceType `json:"type"`
	UserID  string          `json:"userId"`
	GroupID string          `json:"groupId"`
	RoomID  string          `json:"roomId"`
}

// EventPostback type
type EventPostback struct {
	Data string `json:"data"`
}

// EventBeacon type
type EventBeacon struct {
	Hwid string `json:"hwid"`
	Type string `json:"type"`
}

// Event type
type Event struct {
	ReplyToken string
	Type       EventType
	Timestamp  int64
	Source     *EventSource
	Message    Message
	Postback   *EventPostback
	Beacon     *EventBeacon
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
			Duration  int         `json:"duration"`
			Title     string      `json:"title"`
			Address   string      `json:"address"`
			Latitude  float64     `json:"latitude"`
			Longitude float64     `json:"longitude"`
			PackageID string      `json:"packageId"`
			StickerID string      `json:"stickerId"`
		} `json:"message"`
		Postback EventPostback `json:"postback"`
		Beacon   EventBeacon   `json:"beacon"`
	}{}
	if err = json.Unmarshal(body, &rawEvent); err != nil {
		return
	}

	e.ReplyToken = rawEvent.ReplyToken
	e.Type = rawEvent.Type
	e.Timestamp = rawEvent.Timestamp
	e.Source = &rawEvent.Source

	switch rawEvent.Type {
	case EventTypeMessage:
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
		case MessageTypeVideo:
			e.Message = &VideoMessage{
				ID: rawEvent.Message.ID,
			}
		case MessageTypeAudio:
			e.Message = &AudioMessage{
				ID:       rawEvent.Message.ID,
				Duration: rawEvent.Message.Duration,
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
	case EventTypePostback:
		e.Postback = &rawEvent.Postback
	case EventTypeBeacon:
		e.Beacon = &rawEvent.Beacon
	}
	return
}
