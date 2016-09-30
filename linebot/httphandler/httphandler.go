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

package httphandler

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

type FollowHandlerFunc func(*linebot.Client, *linebot.Event)
type UnfollowHandlerFunc func(*linebot.Client, *linebot.Event)
type JoinHandlerFunc func(*linebot.Client, *linebot.Event)
type LeaveHandlerFunc func(*linebot.Client, *linebot.Event)
type PostbackHandlerFunc func(*linebot.Client, *linebot.Event, *linebot.Postback)
type BeaconHandlerFunc func(*linebot.Client, *linebot.Event, *linebot.Beacon)
type TextMessageHandlerFunc func(*linebot.Client, *linebot.Event, *linebot.TextMessage)

type OptionNewClientFunc func() (*linebot.Client, error)

type EventHandler struct {
	channelSecret       string
	channelToken        string
	optionNewClientFunc OptionNewClientFunc

	handleFollow          FollowHandlerFunc
	handleUnfollow        UnfollowHandlerFunc
	handleJoin            JoinHandlerFunc
	handleLeave           LeaveHandlerFunc
	handlePostback        PostbackHandlerFunc
	handleBeacon          BeaconHandlerFunc
	handleTextMessage     TextMessageHandlerFunc
	handleImageMessage    func(*linebot.Client, *linebot.Event, *linebot.ImageMessage)
	handleVideoMessage    func(*linebot.Client, *linebot.Event, *linebot.VideoMessage)
	handleAudioMessage    func(*linebot.Client, *linebot.Event, *linebot.AudioMessage)
	handleLocationMessage func(*linebot.Client, *linebot.Event, *linebot.LocationMessage)
	handleStickerMessage  func(*linebot.Client, *linebot.Event, *linebot.StickerMessage)

	handleError func(error)
}

// EventHandlerOption type
type EventHandlerOption func(*EventHandler) error

// WithNewClientFunc function
func WithNewClientFunc(f OptionNewClientFunc) EventHandlerOption {
	return func(eh *EventHandler) error {
		eh.optionNewClientFunc = f
		return nil
	}
}

func New(channelSecret, channelToken string, options ...EventHandlerOption) *EventHandler {
	return &EventHandler{
		channelSecret: channelSecret,
		channelToken:  channelToken,
	}
}

func (eh *EventHandler) HandleFollow(f FollowHandlerFunc) {
	eh.handleFollow = f
}

func (eh *EventHandler) HandleUnfollow(f UnfollowHandlerFunc) {
	eh.handleUnfollow = f
}

func (eh *EventHandler) HandleJoin(f JoinHandlerFunc) {
	eh.handleJoin = f
}

func (eh *EventHandler) HandleLeave(f LeaveHandlerFunc) {
	eh.handleLeave = f
}

func (eh *EventHandler) HandlePostback(f PostbackHandlerFunc) {
	eh.handlePostback = f
}

func (eh *EventHandler) HandleBeacon(f BeaconHandlerFunc) {
	eh.handleBeacon = f
}

func (eh *EventHandler) HandleTextMessage(f TextMessageHandlerFunc) {
	eh.handleTextMessage = f
}

func (eh *EventHandler) HandleError(f func(error)) {
	eh.handleError = f
}

func (eh *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handleError := func(err error) {
		if eh.handleError != nil {
			eh.handleError(err)
		}
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
	}

	var (
		bot *linebot.Client
		err error
	)
	if eh.optionNewClientFunc != nil {
		bot, err = eh.optionNewClientFunc()
		if err != nil {
			handleError(err)
			return
		}
	} else {
		bot, err = linebot.New(eh.channelSecret, eh.channelToken)
		if err != nil {
			handleError(err)
			return
		}
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		handleError(err)
		return
	}
	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			switch m := event.Message.(type) {
			case *linebot.TextMessage:
				if eh.handleTextMessage != nil {
					eh.handleTextMessage(bot, event, m)
				}
			case *linebot.ImageMessage:
				if eh.handleImageMessage != nil {
					eh.handleImageMessage(bot, event, m)
				}
				// TODO: handle other messages
			}
		case linebot.EventTypeFollow:
			if eh.handleFollow != nil {
				eh.handleFollow(bot, event)
			}
		case linebot.EventTypeUnfollow:
			if eh.handleUnfollow != nil {
				eh.handleUnfollow(bot, event)
			}
		case linebot.EventTypeJoin:
			if eh.handleJoin != nil {
				eh.handleJoin(bot, event)
			}
		case linebot.EventTypeLeave:
			if eh.handleLeave != nil {
				eh.handleLeave(bot, event)
			}
			// TODO: handle other events
		}
	}
}
