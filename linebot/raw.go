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
	"net/url"
)

// GetRaw method
func (client *Client) GetRaw(endpoint string, query url.Values) *GetRawCall {
	return &GetRawCall{
		c: client,
		endpoint: endpoint,
		query: query,
	}
}

// GetRawCall type
type GetRawCall struct {
	c   *Client
	ctx context.Context

	endpoint string
	query url.Values
}

// WithContext method
func (call *GetRawCall) WithContext(ctx context.Context) *GetRawCall {
	call.ctx = ctx
	return call
}

// Do method. Callee must close response object.
func (call *GetRawCall) Do() (*http.Response, error) {
	return call.c.get(call.ctx, call.c.endpointBase, call.endpoint, call.query)
}

// PostRaw method
func (client *Client) PostRaw(endpoint string, body io.Reader) *PostRawCall {
	return &PostRawCall{
		c: client,
		endpoint: endpoint,
		body: body,
	}
}

// PostRawCall type
type PostRawCall struct {
	c   *Client
	ctx context.Context

	endpoint string
	body io.Reader
}

// WithContext method
func (call *PostRawCall) WithContext(ctx context.Context) *PostRawCall {
	call.ctx = ctx
	return call
}

// Do method. Callee must close response object.
func (call *PostRawCall) Do() (*http.Response, error) {
	return call.c.post(call.ctx, call.endpoint, call.body)
}

// PostFormRaw method
func (client *Client) PostFormRaw(endpoint string, body io.Reader) *PostFormRawCall {
	return &PostFormRawCall{
		c: client,

		endpoint: endpoint,
		body: body,
	}
}

// PostFormRawCall type
type PostFormRawCall struct {
	c   *Client
	ctx context.Context

	endpoint string
	body io.Reader
}

// WithContext method
func (call *PostFormRawCall) WithContext(ctx context.Context) *PostFormRawCall {
	call.ctx = ctx
	return call
}

// Do method. Callee must close response object.
func (call *PostFormRawCall) Do() (*http.Response, error) {
	return call.c.postform(call.ctx, call.endpoint, call.body)
}

// PutRaw method
func (client *Client) PutRaw(endpoint string, body io.Reader) *PutRawCall {
	return &PutRawCall{
		c:        client,
		endpoint: endpoint,
		body: body,
	}
}

// PutRawCall type
type PutRawCall struct {
	c   *Client
	ctx context.Context

	endpoint string
	body     io.Reader
}

// WithContext method
func (call *PutRawCall) WithContext(ctx context.Context) *PutRawCall {
	call.ctx = ctx
	return call
}

// Do method. Callee must close response object.
func (call *PutRawCall) Do() (*http.Response, error) {
	return call.c.put(call.ctx, call.endpoint, call.body)
}

// DeleteRaw method
func (client *Client) DeleteRaw(endpoint string) *DeleteRawCall {
	return &DeleteRawCall{
		c:        client,
		endpoint: endpoint,
	}
}

// DeleteRawCall type
type DeleteRawCall struct {
	c   *Client
	ctx context.Context

	endpoint string
}

// WithContext method
func (call *DeleteRawCall) WithContext(ctx context.Context) *DeleteRawCall {
	call.ctx = ctx
	return call
}

// Do method. Callee must close response object.
func (call *DeleteRawCall) Do() (*http.Response, error) {
	return call.c.delete(call.ctx, call.endpoint)
}
