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

package webhook

import (
	"errors"
	"net/http"
)

// EventsHandlerFunc type
type EventsHandlerFunc func(*CallbackRequest, *http.Request)

// ErrorHandlerFunc type
type ErrorHandlerFunc func(error, *http.Request)

// WebhookHandler type
type WebhookHandler struct {
	channelSecret string

	handleEvents EventsHandlerFunc
	handleError  ErrorHandlerFunc
}

// New returns a new WebhookHandler instance.
func NewWebhookHandler(channelSecret string) (*WebhookHandler, error) {
	if channelSecret == "" {
		return nil, errors.New("missing channel secret")
	}
	h := &WebhookHandler{
		channelSecret: channelSecret,
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

func (wh *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := ParseRequest(wh.channelSecret, r)
	if err != nil {
		if wh.handleError != nil {
			wh.handleError(err, r)
		}
		if err == ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	wh.handleEvents(events, r)
}
