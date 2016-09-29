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
	UserID  string          `json:"userId"`
	GroupID string          `json:"groupId"`
	RoomID  string          `json:"roomId"`
}

// Postback type
type Postback struct {
	Data string `json:"data"`
}

// BeaconEventType type
type BeaconEventType string

// BeaconEventType constants
const (
	BeaconEventTypeEnter BeaconEventType = "enter"
)

// Beacon type
type Beacon struct {
	Hwid string          `json:"hwid"`
	Type BeaconEventType `json:"type"`
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

const (
	millisecPerSec     = int64(time.Second / time.Millisecond)
	nanosecPerMillisec = int64(time.Millisecond / time.Nanosecond)
)

// UnmarshalJSON method of Event
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
		Postback `json:"postback"`
		Beacon   `json:"beacon"`
	}{}
	if err = json.Unmarshal(body, &rawEvent); err != nil {
		return
	}

	e.ReplyToken = rawEvent.ReplyToken
	e.Type = rawEvent.Type
	e.Timestamp = time.Unix(rawEvent.Timestamp/millisecPerSec, (rawEvent.Timestamp%millisecPerSec)*nanosecPerMillisec).UTC()
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
