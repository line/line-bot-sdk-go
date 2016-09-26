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
	"bytes"
	"encoding/json"
	"io"

	"golang.org/x/net/context"
)

// Push method
func (client *Client) Push(to string, messages ...Message) *PushCall {
	return &PushCall{
		c:        client,
		to:       to,
		messages: messages,
	}
}

// Reply method
func (client *Client) Reply(replyToken string, messages ...Message) *ReplyCall {
	return &ReplyCall{
		c:          client,
		replyToken: replyToken,
		messages:   messages,
	}
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

// Do method
func (call *PushCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointEventsPush, &buf)
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
	res, err := call.c.post(call.ctx, APIEndpointEventsReply, &buf)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToBasicResponse(res)
}
