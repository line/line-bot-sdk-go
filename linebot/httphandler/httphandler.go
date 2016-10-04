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

	"golang.org/x/net/context"

	"github.com/line/line-bot-sdk-go/linebot"
)

type (
	// EventHandlerFunc type
	EventHandlerFunc func(context.Context, *linebot.Client, *linebot.Event)
	// ErrorHandlerFunc tyoe
	ErrorHandlerFunc func(context.Context, error)
)

type (
	// OptionNewContextFunc type
	OptionNewContextFunc func(*http.Request) (context.Context, error)
	// OptionNewClientFunc type
	OptionNewClientFunc func(context.Context, string, string) (*linebot.Client, error)
)

// EventHandler type
type EventHandler struct {
	channelSecret        string
	channelToken         string
	optionNewContextFunc OptionNewContextFunc
	optionNewClientFunc  OptionNewClientFunc

	handleFollow   EventHandlerFunc
	handleUnfollow EventHandlerFunc
	handleJoin     EventHandlerFunc
	handleLeave    EventHandlerFunc
	handlePostback EventHandlerFunc
	handleBeacon   EventHandlerFunc
	handleMessage  EventHandlerFunc
	handleUnknown  EventHandlerFunc

	handleError ErrorHandlerFunc
}

// EventHandlerOption type
type EventHandlerOption func(*EventHandler) error

// WithNewContextFunc function
func WithNewContextFunc(f OptionNewContextFunc) EventHandlerOption {
	return func(eh *EventHandler) error {
		eh.optionNewContextFunc = f
		return nil
	}
}

// WithNewClientFunc function
func WithNewClientFunc(f OptionNewClientFunc) EventHandlerOption {
	return func(eh *EventHandler) error {
		eh.optionNewClientFunc = f
		return nil
	}
}

// New returns a new EventHandler instance.
func New(channelSecret, channelToken string, options ...EventHandlerOption) (*EventHandler, error) {
	h := &EventHandler{
		channelSecret: channelSecret,
		channelToken:  channelToken,
	}
	for _, option := range options {
		err := option(h)
		if err != nil {
			return nil, err
		}
	}
	return h, nil
}

// HandleFollow method
func (eh *EventHandler) HandleFollow(f EventHandlerFunc) {
	eh.handleFollow = f
}

// HandleUnfollow method
func (eh *EventHandler) HandleUnfollow(f EventHandlerFunc) {
	eh.handleUnfollow = f
}

// HandleJoin method
func (eh *EventHandler) HandleJoin(f EventHandlerFunc) {
	eh.handleJoin = f
}

// HandleLeave method
func (eh *EventHandler) HandleLeave(f EventHandlerFunc) {
	eh.handleLeave = f
}

// HandlePostback method
func (eh *EventHandler) HandlePostback(f EventHandlerFunc) {
	eh.handlePostback = f
}

// HandleBeacon method
func (eh *EventHandler) HandleBeacon(f EventHandlerFunc) {
	eh.handleBeacon = f
}

// HandleMessage method
func (eh *EventHandler) HandleMessage(f EventHandlerFunc) {
	eh.handleMessage = f
}

// HandleUnknown method
func (eh *EventHandler) HandleUnknown(f EventHandlerFunc) {
	eh.handleUnknown = f
}

// HandleError method
func (eh *EventHandler) HandleError(f ErrorHandlerFunc) {
	eh.handleError = f
}

func (eh *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handleError := func(c context.Context, e error) {
		if eh.handleError != nil {
			eh.handleError(c, e)
		}
		if e == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
	}

	var (
		ctx context.Context
		bot *linebot.Client
		err error
	)
	if eh.optionNewContextFunc != nil {
		ctx, err = eh.optionNewContextFunc(r)
	} else {
		ctx, err = eh.DefaultNewContext(r)
	}
	if err != nil {
		handleError(nil, err)
		return
	}
	if eh.optionNewClientFunc != nil {
		bot, err = eh.optionNewClientFunc(ctx, eh.channelSecret, eh.channelToken)
	} else {
		bot, err = eh.DefaultNewClient(ctx, eh.channelSecret, eh.channelToken)
	}
	if err != nil {
		handleError(ctx, err)
		return
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		handleError(ctx, err)
		return
	}
	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeFollow:
			if eh.handleFollow != nil {
				eh.handleFollow(ctx, bot, event)
			}
		case linebot.EventTypeUnfollow:
			if eh.handleUnfollow != nil {
				eh.handleUnfollow(ctx, bot, event)
			}
		case linebot.EventTypeJoin:
			if eh.handleJoin != nil {
				eh.handleJoin(ctx, bot, event)
			}
		case linebot.EventTypeLeave:
			if eh.handleLeave != nil {
				eh.handleLeave(ctx, bot, event)
			}
		case linebot.EventTypePostback:
			if eh.handlePostback != nil {
				eh.handlePostback(ctx, bot, event)
			}
		case linebot.EventTypeBeacon:
			if eh.handleBeacon != nil {
				eh.handleBeacon(ctx, bot, event)
			}
		case linebot.EventTypeMessage:
			if eh.handleMessage != nil {
				eh.handleMessage(ctx, bot, event)
			}
		default:
			if eh.handleUnknown != nil {
				eh.handleUnknown(ctx, bot, event)
			}
		}
	}
}
