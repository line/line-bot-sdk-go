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
	"errors"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// EventsHandlerFunc type
type EventsHandlerFunc func([]*linebot.Event, *http.Request)

// ErrorHandlerFunc type
type ErrorHandlerFunc func(error, *http.Request)

// WebhookHandler type
// Deprecated: Use OpenAPI based classes instead.
type WebhookHandler struct {
	channelSecret string
	channelToken  string

	handleEvents EventsHandlerFunc
	handleError  ErrorHandlerFunc
}

// New returns a new WebhookHandler instance.
// Deprecated: Use OpenAPI based classes instead.
func New(channelSecret, channelToken string) (*WebhookHandler, error) {
	if channelSecret == "" {
		return nil, errors.New("missing channel secret")
	}
	if channelToken == "" {
		return nil, errors.New("missing channel access token")
	}
	h := &WebhookHandler{
		channelSecret: channelSecret,
		channelToken:  channelToken,
	}
	return h, nil
}

// HandleEvents method
func (wh *WebhookHandler) HandleEvents(f EventsHandlerFunc) {
	wh.handleEvents = f
}

// HandleError method
func (wh *WebhookHandler) HandleError(f ErrorHandlerFunc) {
	wh.handleError = f
}

// NewClient method
func (wh *WebhookHandler) NewClient(options ...linebot.ClientOption) (*linebot.Client, error) {
	return linebot.New(wh.channelSecret, wh.channelToken, options...)
}

func (wh *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := linebot.ParseRequest(wh.channelSecret, r)
	if err != nil {
		if wh.handleError != nil {
			wh.handleError(err, r)
		}
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	wh.handleEvents(events, r)
}
