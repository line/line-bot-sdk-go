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
	"encoding/hex"
	"encoding/json"
	"time"
)

// EventType type
type EventType string

// EventType constants
const (
	EventTypeMessage  EventType = "message"
	EventTypeFollow   EventType = "follow"
	EventTypeUnfollow EventType = "unfollow"
	EventTypeJoin     EventType = "join"
	EventTypeLeave    EventType = "leave"
	EventTypePostback EventType = "postback"
	EventTypeBeacon   EventType = "beacon"
)

// EventSourceType type
type EventSourceType string

// EventSourceType constants
const (
	EventSourceTypeUser  EventSourceType = "user"
	EventSourceTypeGroup EventSourceType = "group"
	EventSourceTypeRoom  EventSourceType = "room"
)

// EventSource type
type EventSource struct {
	Type    EventSourceType `json:"type"`
	UserID  string          `json:"userId,omitempty"`
	GroupID string          `json:"groupId,omitempty"`
	RoomID  string          `json:"roomId,omitempty"`
}

// Postback type
type Postback struct {
	Data string `json:"data"`
}

// BeaconEventType type
type BeaconEventType string

// BeaconEventType constants
const (
	BeaconEventTypeEnter  BeaconEventType = "enter"
	BeaconEventTypeLeave  BeaconEventType = "leave"
	BeaconEventTypeBanner BeaconEventType = "banner"
)

// Beacon type
type Beacon struct {
	Hwid          string
	Type          BeaconEventType
	DeviceMessage []byte
}

// Event type
type Event struct {
	ReplyToken string
	Type       EventType
	Timestamp  time.Time
	Source     *EventSource
	Message    Message
	Postback   *Postback
	Beacon     *Beacon
}

type rawEvent struct {
	ReplyToken string           `json:"replyToken,omitempty"`
	Type       EventType        `json:"type"`
	Timestamp  int64            `json:"timestamp"`
	Source     *EventSource     `json:"source"`
	Message    *rawEventMessage `json:"message,omitempty"`
	*Postback  `json:"postback,omitempty"`
	Beacon     *rawBeaconEvent `json:"beacon,omitempty"`
}

type rawEventMessage struct {
	ID        string      `json:"id"`
	Type      MessageType `json:"type"`
	Text      string      `json:"text,omitempty"`
	Duration  int         `json:"duration,omitempty"`
	Title     string      `json:"title,omitempty"`
	Address   string      `json:"address,omitempty"`
	Latitude  float64     `json:"latitude,omitempty"`
	Longitude float64     `json:"longitude,omitempty"`
	PackageID string      `json:"packageId,omitempty"`
	StickerID string      `json:"stickerId,omitempty"`
}

type rawBeaconEvent struct {
	Hwid string          `json:"hwid"`
	Type BeaconEventType `json:"type"`
	DM   string          `json:"dm,omitempty"`
}

const (
	millisecPerSec     = int64(time.Second / time.Millisecond)
	nanosecPerMillisec = int64(time.Millisecond / time.Nanosecond)
)

// MarshalJSON method of Event
func (e *Event) MarshalJSON() ([]byte, error) {
	raw := rawEvent{
		ReplyToken: e.ReplyToken,
		Type:       e.Type,
		Timestamp:  e.Timestamp.Unix()*millisecPerSec + int64(e.Timestamp.Nanosecond())/int64(time.Millisecond),
		Source:     e.Source,
		Postback:   e.Postback,
	}
	if e.Beacon != nil {
		raw.Beacon = &rawBeaconEvent{
			Hwid: e.Beacon.Hwid,
			Type: e.Beacon.Type,
			DM:   hex.EncodeToString(e.Beacon.DeviceMessage),
		}
	}

	switch m := e.Message.(type) {
	case *TextMessage:
		raw.Message = &rawEventMessage{
			Type: MessageTypeText,
			ID:   m.ID,
			Text: m.Text,
		}
	case *ImageMessage:
		raw.Message = &rawEventMessage{
			Type: MessageTypeImage,
			ID:   m.ID,
		}
	case *VideoMessage:
		raw.Message = &rawEventMessage{
			Type: MessageTypeVideo,
			ID:   m.ID,
		}
	case *AudioMessage:
		raw.Message = &rawEventMessage{
			Type:     MessageTypeAudio,
			ID:       m.ID,
			Duration: m.Duration,
		}
	case *LocationMessage:
		raw.Message = &rawEventMessage{
			Type:      MessageTypeLocation,
			ID:        m.ID,
			Title:     m.Title,
			Address:   m.Address,
			Latitude:  m.Latitude,
			Longitude: m.Longitude,
		}
	case *StickerMessage:
		raw.Message = &rawEventMessage{
			Type:      MessageTypeSticker,
			ID:        m.ID,
			PackageID: m.PackageID,
			StickerID: m.StickerID,
		}
	}
	return json.Marshal(&raw)
}

// UnmarshalJSON method of Event
func (e *Event) UnmarshalJSON(body []byte) (err error) {
	rawEvent := rawEvent{}
	if err = json.Unmarshal(body, &rawEvent); err != nil {
		return
	}

	e.ReplyToken = rawEvent.ReplyToken
	e.Type = rawEvent.Type
	e.Timestamp = time.Unix(rawEvent.Timestamp/millisecPerSec, (rawEvent.Timestamp%millisecPerSec)*nanosecPerMillisec).UTC()
	e.Source = rawEvent.Source

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
		e.Postback = rawEvent.Postback
	case EventTypeBeacon:
		var deviceMessage []byte
		deviceMessage, err = hex.DecodeString(rawEvent.Beacon.DM)
		if err != nil {
			return
		}
		e.Beacon = &Beacon{
			Hwid:          rawEvent.Beacon.Hwid,
			Type:          rawEvent.Beacon.Type,
			DeviceMessage: deviceMessage,
		}
	}
	return
}
