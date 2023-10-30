// Copyright 2022 LINE Corporation
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

// ValidatePushMessage method
func (client *Client) ValidatePushMessage(messages ...SendingMessage) *ValidatePushMessageCall {
	return &ValidatePushMessageCall{
		c:        client,
		messages: messages,
	}
}

// ValidatePushMessageCall type
// Deprecated: Use OpenAPI based classes instead.
type ValidatePushMessageCall struct {
	c   *Client
	ctx context.Context

	messages []SendingMessage
}

// WithContext method
func (call *ValidatePushMessageCall) WithContext(ctx context.Context) *ValidatePushMessageCall {
	call.ctx = ctx
	return call
}

func (call *ValidatePushMessageCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		Messages []SendingMessage `json:"messages"`
	}{
		Messages: call.messages,
	})
}

// Do method
func (call *ValidatePushMessageCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointValidatePushMessage, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// ValidateReplyMessage method
func (client *Client) ValidateReplyMessage(messages ...SendingMessage) *ValidateReplyMessageCall {
	return &ValidateReplyMessageCall{
		c:        client,
		messages: messages,
	}
}

// ValidateReplyMessageCall type
// Deprecated: Use OpenAPI based classes instead.
type ValidateReplyMessageCall struct {
	c   *Client
	ctx context.Context

	messages []SendingMessage
}

// WithContext method
func (call *ValidateReplyMessageCall) WithContext(ctx context.Context) *ValidateReplyMessageCall {
	call.ctx = ctx
	return call
}

func (call *ValidateReplyMessageCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		Messages []SendingMessage `json:"messages"`
	}{
		Messages: call.messages,
	})
}

// Do method
func (call *ValidateReplyMessageCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointValidateReplyMessage, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// ValidateMulticastMessage method
func (client *Client) ValidateMulticastMessage(messages ...SendingMessage) *ValidateMulticastMessageCall {
	return &ValidateMulticastMessageCall{
		c:        client,
		messages: messages,
	}
}

// ValidateMulticastMessageCall type
// Deprecated: Use OpenAPI based classes instead.
type ValidateMulticastMessageCall struct {
	c   *Client
	ctx context.Context

	messages []SendingMessage
}

// WithContext method
func (call *ValidateMulticastMessageCall) WithContext(ctx context.Context) *ValidateMulticastMessageCall {
	call.ctx = ctx
	return call
}

func (call *ValidateMulticastMessageCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		Messages []SendingMessage `json:"messages"`
	}{
		Messages: call.messages,
	})
}

// Do method
func (call *ValidateMulticastMessageCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointValidateMulticastMessage, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// ValidateBroadcastMessage method
func (client *Client) ValidateBroadcastMessage(messages ...SendingMessage) *ValidateBroadcastMessageCall {
	return &ValidateBroadcastMessageCall{
		c:        client,
		messages: messages,
	}
}

// ValidateBroadcastMessageCall type
// Deprecated: Use OpenAPI based classes instead.
type ValidateBroadcastMessageCall struct {
	c   *Client
	ctx context.Context

	messages []SendingMessage
}

// WithContext method
func (call *ValidateBroadcastMessageCall) WithContext(ctx context.Context) *ValidateBroadcastMessageCall {
	call.ctx = ctx
	return call
}

func (call *ValidateBroadcastMessageCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		Messages []SendingMessage `json:"messages"`
	}{
		Messages: call.messages,
	})
}

// Do method
func (call *ValidateBroadcastMessageCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointValidateBroadcastMessage, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// ValidateNarrowcastMessage method
func (client *Client) ValidateNarrowcastMessage(messages ...SendingMessage) *ValidateNarrowcastMessageCall {
	return &ValidateNarrowcastMessageCall{
		c:        client,
		messages: messages,
	}
}

// ValidateNarrowcastMessageCall type
// Deprecated: Use OpenAPI based classes instead.
type ValidateNarrowcastMessageCall struct {
	c   *Client
	ctx context.Context

	messages []SendingMessage
}

// WithContext method
func (call *ValidateNarrowcastMessageCall) WithContext(ctx context.Context) *ValidateNarrowcastMessageCall {
	call.ctx = ctx
	return call
}

func (call *ValidateNarrowcastMessageCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		Messages []SendingMessage `json:"messages"`
	}{
		Messages: call.messages,
	})
}

// Do method
func (call *ValidateNarrowcastMessageCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointValidateNarrowcastMessage, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}
