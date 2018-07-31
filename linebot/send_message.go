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
	"context"
	"encoding/json"
	"io"
)

// PushMessage method
func (client *Client) PushMessage(to string, messages ...SendingMessage) *PushMessageCall {
	return &PushMessageCall{
		c:        client,
		to:       to,
		messages: messages,
	}
}

// PushMessageCall type
type PushMessageCall struct {
	c   *Client
	ctx context.Context

	to       string
	messages []SendingMessage
}

// WithContext method
func (call *PushMessageCall) WithContext(ctx context.Context) *PushMessageCall {
	call.ctx = ctx
	return call
}

func (call *PushMessageCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		To       string           `json:"to"`
		Messages []SendingMessage `json:"messages"`
	}{
		To:       call.to,
		Messages: call.messages,
	})
}

// Do method
func (call *PushMessageCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointPushMessage, &buf)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToBasicResponse(res)
}

// ReplyMessage method
func (client *Client) ReplyMessage(replyToken string, messages ...SendingMessage) *ReplyMessageCall {
	return &ReplyMessageCall{
		c:          client,
		replyToken: replyToken,
		messages:   messages,
	}
}

// ReplyMessageCall type
type ReplyMessageCall struct {
	c   *Client
	ctx context.Context

	replyToken string
	messages   []SendingMessage
}

// WithContext method
func (call *ReplyMessageCall) WithContext(ctx context.Context) *ReplyMessageCall {
	call.ctx = ctx
	return call
}

func (call *ReplyMessageCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		ReplyToken string           `json:"replyToken"`
		Messages   []SendingMessage `json:"messages"`
	}{
		ReplyToken: call.replyToken,
		Messages:   call.messages,
	})
}

// Do method
func (call *ReplyMessageCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointReplyMessage, &buf)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToBasicResponse(res)
}

// Multicast method
func (client *Client) Multicast(to []string, messages ...SendingMessage) *MulticastCall {
	return &MulticastCall{
		c:        client,
		to:       to,
		messages: messages,
	}
}

// MulticastCall type
type MulticastCall struct {
	c   *Client
	ctx context.Context

	to       []string
	messages []SendingMessage
}

// WithContext method
func (call *MulticastCall) WithContext(ctx context.Context) *MulticastCall {
	call.ctx = ctx
	return call
}

func (call *MulticastCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		To       []string         `json:"to"`
		Messages []SendingMessage `json:"messages"`
	}{
		To:       call.to,
		Messages: call.messages,
	})
}

// Do method
func (call *MulticastCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointMulticast, &buf)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToBasicResponse(res)
}
