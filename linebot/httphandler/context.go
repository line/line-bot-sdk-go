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

// +build !appengine

package httphandler

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"

	"golang.org/x/net/context"
)

// DefaultNewContext method
func (eh *EventHandler) DefaultNewContext(req *http.Request) (context.Context, error) {
	return context.Background(), nil
}

// DefaultNewClient method
func (eh *EventHandler) DefaultNewClient(ctx context.Context, secret string, token string) (*linebot.Client, error) {
	return linebot.New(secret, token)
}
