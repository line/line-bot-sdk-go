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

// EventsHandlerFunc type
type EventsHandlerFunc func([]*linebot.Event, *http.Request)

// ErrorHandlerFunc type
type ErrorHandlerFunc func(error, *http.Request)

// WebhookHandler type
type WebhookHandler struct {
	channelSecret string
	channelToken  string

	handleEvents EventsHandlerFunc
	handleError  ErrorHandlerFunc
}

// New returns a new WebhookHandler instance.
func New(channelSecret, channelToken string) (*WebhookHandler, error) {
	h := &WebhookHandler{
		channelSecret: channelSecret,
		channelToken:  channelToken,
	}
	return h, nil
}

// HandleEvents method
func (eh *WebhookHandler) HandleEvents(f EventsHandlerFunc) {
	eh.handleEvents = f
}

// HandleError method
func (eh *WebhookHandler) HandleError(f ErrorHandlerFunc) {
	eh.handleError = f
}

// NewClient method
func (eh *WebhookHandler) NewClient(options ...linebot.ClientOption) (*linebot.Client, error) {
	return linebot.New(eh.channelSecret, eh.channelToken, options...)
}

func (eh *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := linebot.ParseRequest(eh.channelSecret, r)
	if err != nil {
		if eh.handleError != nil {
			eh.handleError(err, r)
		}
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	eh.handleEvents(events, r)
}
