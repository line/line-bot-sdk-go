// Copyright 2021 LINE Corporation
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
	"context"
	"io"
	"net/http"
)

// NewRawCall method
func (client *Client) NewRawCall(method string, endpoint string) (*RawCall, error) {
	req, err := http.NewRequest(method, client.url(client.endpointBase, endpoint), nil)
	if err != nil {
		return nil, err
	}
	return &RawCall{
		c:   client,
		req: req,
	}, nil
}

// NewRawCallWithBody method
func (client *Client) NewRawCallWithBody(method string, endpoint string, body io.Reader) (*RawCall, error) {
	req, err := http.NewRequest(method, client.url(client.endpointBase, endpoint), body)
	if err != nil {
		return nil, err
	}
	return &RawCall{
		c:   client,
		req: req,
	}, nil
}

// RawCall type
// Deprecated: Use OpenAPI based classes instead.
type RawCall struct {
	c   *Client
	ctx context.Context

	req *http.Request
}

// AddHeader method
func (call *RawCall) AddHeader(key string, value string) {
	call.req.Header.Add(key, value)
}

// WithContext method
func (call *RawCall) WithContext(ctx context.Context) *RawCall {
	call.ctx = ctx
	return call
}

// Do method. Callee must close response object.
func (call *RawCall) Do() (*http.Response, error) {
	return call.c.do(call.ctx, call.req)
}
