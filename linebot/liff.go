// Copyright 2018 LINE Corporation
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
	"fmt"
	"io"
)

// LIFFViewType type
type LIFFViewType string

// LIFFViewType constants
const (
	LIFFViewTypeCompact LIFFViewType = "compact"
	LIFFViewTypeTall    LIFFViewType = "tall"
	LIFFViewTypeFull    LIFFViewType = "full"
)

// LIFFViewScopeType type
type LIFFViewScopeType string

// LIFFViewScopeType constants
const (
	LIFFViewScopeTypeOpenID           LIFFViewScopeType = "openid"
	LIFFViewScopeTypeEmail            LIFFViewScopeType = "email"
	LIFFViewScopeTypeProfile          LIFFViewScopeType = "profile"
	LIFFViewScopeTypeChatMessageWrite LIFFViewScopeType = "chat_message.write"
)

// LIFFApp type
// Deprecated: Use OpenAPI based classes instead.
type LIFFApp struct {
	LIFFID               string              `json:"liffId"`
	View                 View                `json:"view"`
	Description          string              `json:"description,omitempty"`
	Features             *LIFFAppFeatures    `json:"features,omitempty"`
	PermanentLinkPattern string              `json:"permanentLinkPattern,omitempty"`
	Scope                []LIFFViewScopeType `json:"scope,omitempty"`
	BotPrompt            string              `json:"botprompt,omitempty"`
}

// View type
// Deprecated: Use OpenAPI based classes instead.
type View struct {
	Type       LIFFViewType `json:"type"`
	URL        string       `json:"url"`
	ModlueMode bool         `json:"moduleMode,omitempty"`
}

// LIFFAppFeatures type
// Deprecated: Use OpenAPI based classes instead.
type LIFFAppFeatures struct {
	BLE    bool `json:"ble,omitempty"`
	QRCode bool `json:"qrCode,omitempty"`
}

// GetLIFF method
func (client *Client) GetLIFF() *GetLIFFAllCall {
	return &GetLIFFAllCall{
		c: client,
	}
}

// GetLIFFAllCall type
// Deprecated: Use OpenAPI based classes instead.
type GetLIFFAllCall struct {
	c   *Client
	ctx context.Context
}

// WithContext method
func (call *GetLIFFAllCall) WithContext(ctx context.Context) *GetLIFFAllCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetLIFFAllCall) Do() (*LIFFAppsResponse, error) {
	res, err := call.c.get(call.ctx, call.c.endpointBase, APIEndpointGetAllLIFFApps, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToLIFFResponse(res)
}

// AddLIFF method
func (client *Client) AddLIFF(app LIFFApp) *AddLIFFCall {
	return &AddLIFFCall{
		c:   client,
		app: app,
	}
}

// AddLIFFCall type
// Deprecated: Use OpenAPI based classes instead.
type AddLIFFCall struct {
	c   *Client
	ctx context.Context

	app LIFFApp
}

// WithContext method
func (call *AddLIFFCall) WithContext(ctx context.Context) *AddLIFFCall {
	call.ctx = ctx
	return call
}

func (call *AddLIFFCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		View                 View                `json:"view"`
		Description          string              `json:"description,omitempty"`
		Features             *LIFFAppFeatures    `json:"features,omitempty"`
		PermanentLinkPattern string              `json:"permanentLinkPattern,omitempty"`
		Scope                []LIFFViewScopeType `json:"scope,omitempty"`
		BotPrompt            string              `json:"botPrompt,omitempty"`
	}{
		View:                 call.app.View,
		Description:          call.app.Description,
		Features:             call.app.Features,
		PermanentLinkPattern: call.app.PermanentLinkPattern,
		Scope:                call.app.Scope,
		BotPrompt:            call.app.BotPrompt,
	})
}

// Do method
func (call *AddLIFFCall) Do() (*LIFFIDResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointAddLIFFApp, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToLIFFIDResponse(res)
}

// UpdateLIFF method
func (client *Client) UpdateLIFF(liffID string, app LIFFApp) *UpdateLIFFCall {
	return &UpdateLIFFCall{
		c:      client,
		liffID: liffID,
		app:    app,
	}
}

// UpdateLIFFCall type
// Deprecated: Use OpenAPI based classes instead.
type UpdateLIFFCall struct {
	c   *Client
	ctx context.Context

	liffID string
	app    LIFFApp
}

// WithContext method
func (call *UpdateLIFFCall) WithContext(ctx context.Context) *UpdateLIFFCall {
	call.ctx = ctx
	return call
}

func (call *UpdateLIFFCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		View                 View                `json:"view"`
		Description          string              `json:"description,omitempty"`
		Features             *LIFFAppFeatures    `json:"features,omitempty"`
		PermanentLinkPattern string              `json:"permanentLinkPattern,omitempty"`
		Scope                []LIFFViewScopeType `json:"scope,omitempty"`
		BotPrompt            string              `json:"botPrompt,omitempty"`
	}{
		View:                 call.app.View,
		Description:          call.app.Description,
		Features:             call.app.Features,
		PermanentLinkPattern: call.app.PermanentLinkPattern,
		Scope:                call.app.Scope,
		BotPrompt:            call.app.BotPrompt,
	})
}

// Do method
func (call *UpdateLIFFCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(APIEndpointUpdateLIFFApp, call.liffID)
	res, err := call.c.put(call.ctx, endpoint, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// DeleteLIFF method
func (client *Client) DeleteLIFF(liffID string) *DeleteLIFFCall {
	return &DeleteLIFFCall{
		c:      client,
		liffID: liffID,
	}
}

// DeleteLIFFCall type
// Deprecated: Use OpenAPI based classes instead.
type DeleteLIFFCall struct {
	c   *Client
	ctx context.Context

	liffID string
}

// WithContext method
func (call *DeleteLIFFCall) WithContext(ctx context.Context) *DeleteLIFFCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *DeleteLIFFCall) Do() (*BasicResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointDeleteLIFFApp, call.liffID)
	res, err := call.c.delete(call.ctx, endpoint)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}
