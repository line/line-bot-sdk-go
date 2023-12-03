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
	"fmt"
	"io"
	"net/http"
	"os"
)

// RichMenuActionType type
type RichMenuActionType string

// RichMenuActionType constants
const (
	RichMenuActionTypeURI            RichMenuActionType = "uri"
	RichMenuActionTypeMessage        RichMenuActionType = "message"
	RichMenuActionTypePostback       RichMenuActionType = "postback"
	RichMenuActionTypeDatetimePicker RichMenuActionType = "datetimepicker"
	RichMenuActionTypeRichMenuSwitch RichMenuActionType = "richmenuswitch"
)

// RichMenuSize type
// Deprecated: Use OpenAPI based classes instead.
type RichMenuSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// RichMenuBounds type
// Deprecated: Use OpenAPI based classes instead.
type RichMenuBounds struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

// RichMenuAction with type
// Deprecated: Use OpenAPI based classes instead.
type RichMenuAction struct {
	Type            RichMenuActionType `json:"type"`
	URI             string             `json:"uri,omitempty"`
	Text            string             `json:"text,omitempty"`
	DisplayText     string             `json:"displayText,omitempty"`
	Label           string             `json:"label,omitempty"`
	Data            string             `json:"data,omitempty"`
	Mode            string             `json:"mode,omitempty"`
	Initial         string             `json:"initial,omitempty"`
	Max             string             `json:"max,omitempty"`
	Min             string             `json:"min,omitempty"`
	RichMenuAliasID string             `json:"richMenuAliasId,omitempty"`
	InputOption     InputOption        `json:"inputOption,omitempty"`
	FillInText      string             `json:"fillInText,omitempty"`
}

// AreaDetail type for areas array
// Deprecated: Use OpenAPI based classes instead.
type AreaDetail struct {
	Bounds RichMenuBounds `json:"bounds"`
	Action RichMenuAction `json:"action"`
}

// RichMenu type
// Deprecated: Use OpenAPI based classes instead.
type RichMenu struct {
	Size        RichMenuSize
	Selected    bool
	Name        string
	ChatBarText string
	Areas       []AreaDetail
}

/*
{
  "richmenus": [
    {
      "richMenuId": "{richMenuId}",
      "size": {
        "width": 2500,
        "height": 1686
      },
      "selected": false,
      "areas": [
        {
          "bounds": {
            "x": 0,
            "y": 0,
            "width": 2500,
            "height": 1686
          },
          "action": {
            "type": "postback",
            "data": "action=buy&itemid=123",
            "label":"Buy",
            "displayText":"Buy"
          }
        }
      ]
    }
  ]
}
*/

// GetRichMenu method
func (client *Client) GetRichMenu(richMenuID string) *GetRichMenuCall {
	return &GetRichMenuCall{
		c:          client,
		richMenuID: richMenuID,
	}
}

// GetRichMenuCall type
// Deprecated: Use OpenAPI based classes instead.
type GetRichMenuCall struct {
	c   *Client
	ctx context.Context

	richMenuID string
}

// WithContext method
func (call *GetRichMenuCall) WithContext(ctx context.Context) *GetRichMenuCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetRichMenuCall) Do() (*RichMenuResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointGetRichMenu, call.richMenuID)
	res, err := call.c.get(call.ctx, call.c.endpointBase, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToRichMenuResponse(res)
}

// GetUserRichMenu method
func (client *Client) GetUserRichMenu(userID string) *GetUserRichMenuCall {
	return &GetUserRichMenuCall{
		c:      client,
		userID: userID,
	}
}

// GetUserRichMenuCall type
// Deprecated: Use OpenAPI based classes instead.
type GetUserRichMenuCall struct {
	c   *Client
	ctx context.Context

	userID string
}

// WithContext method
func (call *GetUserRichMenuCall) WithContext(ctx context.Context) *GetUserRichMenuCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetUserRichMenuCall) Do() (*RichMenuResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointGetUserRichMenu, call.userID)
	res, err := call.c.get(call.ctx, call.c.endpointBase, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToRichMenuResponse(res)
}

// CreateRichMenu method
func (client *Client) CreateRichMenu(richMenu RichMenu) *CreateRichMenuCall {
	return &CreateRichMenuCall{
		c:        client,
		richMenu: richMenu,
	}
}

// CreateRichMenuCall type
// Deprecated: Use OpenAPI based classes instead.
type CreateRichMenuCall struct {
	c   *Client
	ctx context.Context

	richMenu RichMenu
}

// WithContext method
func (call *CreateRichMenuCall) WithContext(ctx context.Context) *CreateRichMenuCall {
	call.ctx = ctx
	return call
}

func (call *CreateRichMenuCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		Size        RichMenuSize `json:"size"`
		Selected    bool         `json:"selected"`
		Name        string       `json:"name"`
		ChatBarText string       `json:"chatBarText"`
		Areas       []AreaDetail `json:"areas"`
	}{
		Size:        call.richMenu.Size,
		Selected:    call.richMenu.Selected,
		Name:        call.richMenu.Name,
		ChatBarText: call.richMenu.ChatBarText,
		Areas:       call.richMenu.Areas,
	})
}

// Do method
func (call *CreateRichMenuCall) Do() (*RichMenuIDResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointCreateRichMenu, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToRichMenuIDResponse(res)
}

// DeleteRichMenu method
func (client *Client) DeleteRichMenu(richMenuID string) *DeleteRichMenuCall {
	return &DeleteRichMenuCall{
		c:          client,
		richMenuID: richMenuID,
	}
}

// DeleteRichMenuCall type
// Deprecated: Use OpenAPI based classes instead.
type DeleteRichMenuCall struct {
	c   *Client
	ctx context.Context

	richMenuID string
}

// WithContext method
func (call *DeleteRichMenuCall) WithContext(ctx context.Context) *DeleteRichMenuCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *DeleteRichMenuCall) Do() (*BasicResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointDeleteRichMenu, call.richMenuID)
	res, err := call.c.delete(call.ctx, endpoint)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// LinkUserRichMenu method
func (client *Client) LinkUserRichMenu(userID, richMenuID string) *LinkUserRichMenuCall {
	return &LinkUserRichMenuCall{
		c:          client,
		userID:     userID,
		richMenuID: richMenuID,
	}
}

// LinkUserRichMenuCall type
// Deprecated: Use OpenAPI based classes instead.
type LinkUserRichMenuCall struct {
	c   *Client
	ctx context.Context

	userID     string
	richMenuID string
}

// WithContext method
func (call *LinkUserRichMenuCall) WithContext(ctx context.Context) *LinkUserRichMenuCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *LinkUserRichMenuCall) Do() (*BasicResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointLinkUserRichMenu, call.userID, call.richMenuID)
	res, err := call.c.post(call.ctx, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// UnlinkUserRichMenu method
func (client *Client) UnlinkUserRichMenu(userID string) *UnlinkUserRichMenuCall {
	return &UnlinkUserRichMenuCall{
		c:      client,
		userID: userID,
	}
}

// UnlinkUserRichMenuCall type
// Deprecated: Use OpenAPI based classes instead.
type UnlinkUserRichMenuCall struct {
	c   *Client
	ctx context.Context

	userID string
}

// WithContext method
func (call *UnlinkUserRichMenuCall) WithContext(ctx context.Context) *UnlinkUserRichMenuCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *UnlinkUserRichMenuCall) Do() (*BasicResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointUnlinkUserRichMenu, call.userID)
	res, err := call.c.delete(call.ctx, endpoint)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// SetDefaultRichMenu method
func (client *Client) SetDefaultRichMenu(richMenuID string) *SetDefaultRichMenuCall {
	return &SetDefaultRichMenuCall{
		c:          client,
		richMenuID: richMenuID,
	}
}

// SetDefaultRichMenuCall type
// Deprecated: Use OpenAPI based classes instead.
type SetDefaultRichMenuCall struct {
	c   *Client
	ctx context.Context

	richMenuID string
}

// WithContext method
func (call *SetDefaultRichMenuCall) WithContext(ctx context.Context) *SetDefaultRichMenuCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *SetDefaultRichMenuCall) Do() (*BasicResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointSetDefaultRichMenu, call.richMenuID)
	res, err := call.c.post(call.ctx, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// CancelDefaultRichMenu method
func (client *Client) CancelDefaultRichMenu() *CancelDefaultRichMenuCall {
	return &CancelDefaultRichMenuCall{
		c: client,
	}
}

// CancelDefaultRichMenuCall type
// Deprecated: Use OpenAPI based classes instead.
type CancelDefaultRichMenuCall struct {
	c   *Client
	ctx context.Context
}

// WithContext method
func (call *CancelDefaultRichMenuCall) WithContext(ctx context.Context) *CancelDefaultRichMenuCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *CancelDefaultRichMenuCall) Do() (*BasicResponse, error) {
	res, err := call.c.delete(call.ctx, APIEndpointDefaultRichMenu)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// GetDefaultRichMenu method
func (client *Client) GetDefaultRichMenu() *GetDefaultRichMenuCall {
	return &GetDefaultRichMenuCall{
		c: client,
	}
}

// GetDefaultRichMenuCall type
// Deprecated: Use OpenAPI based classes instead.
type GetDefaultRichMenuCall struct {
	c   *Client
	ctx context.Context
}

// WithContext method
func (call *GetDefaultRichMenuCall) WithContext(ctx context.Context) *GetDefaultRichMenuCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetDefaultRichMenuCall) Do() (*RichMenuIDResponse, error) {
	res, err := call.c.get(call.ctx, call.c.endpointBase, APIEndpointDefaultRichMenu, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToRichMenuIDResponse(res)
}

// GetRichMenuList method
func (client *Client) GetRichMenuList() *GetRichMenuListCall {
	return &GetRichMenuListCall{
		c: client,
	}
}

// GetRichMenuListCall type
// Deprecated: Use OpenAPI based classes instead.
type GetRichMenuListCall struct {
	c   *Client
	ctx context.Context
}

// WithContext method
func (call *GetRichMenuListCall) WithContext(ctx context.Context) *GetRichMenuListCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetRichMenuListCall) Do() ([]*RichMenuResponse, error) {
	res, err := call.c.get(call.ctx, call.c.endpointBase, APIEndpointListRichMenu, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToRichMenuListResponse(res)
}

// DownloadRichMenuImage method
func (client *Client) DownloadRichMenuImage(richMenuID string) *DownloadRichMenuImageCall {
	return &DownloadRichMenuImageCall{
		c:          client,
		richMenuID: richMenuID,
	}
}

// DownloadRichMenuImageCall type
// Deprecated: Use OpenAPI based classes instead.
type DownloadRichMenuImageCall struct {
	c   *Client
	ctx context.Context

	richMenuID string
}

// WithContext method
func (call *DownloadRichMenuImageCall) WithContext(ctx context.Context) *DownloadRichMenuImageCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *DownloadRichMenuImageCall) Do() (*MessageContentResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointDownloadRichMenuImage, call.richMenuID)
	res, err := call.c.get(call.ctx, call.c.endpointBaseData, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToMessageContentResponse(res)
}

// UploadRichMenuImage method
func (client *Client) UploadRichMenuImage(richMenuID, imgPath string) *UploadRichMenuImageCall {
	return &UploadRichMenuImageCall{
		c:          client,
		richMenuID: richMenuID,
		imgPath:    imgPath,
	}
}

// UploadRichMenuImageCall type
// Deprecated: Use OpenAPI based classes instead.
type UploadRichMenuImageCall struct {
	c   *Client
	ctx context.Context

	richMenuID string
	imgPath    string
}

// WithContext method
func (call *UploadRichMenuImageCall) WithContext(ctx context.Context) *UploadRichMenuImageCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *UploadRichMenuImageCall) Do() (*BasicResponse, error) {
	body, err := os.Open(call.imgPath)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	fi, err := body.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 512)
	n, err := body.Read(buf) // n, in case the file size < 512
	if err != nil && err != io.EOF {
		return nil, err
	}
	body.Seek(0, 0)
	endpoint := fmt.Sprintf(APIEndpointUploadRichMenuImage, call.richMenuID)
	req, err := http.NewRequest("POST", call.c.url(call.c.endpointBaseData, endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", http.DetectContentType(buf[:n]))
	req.ContentLength = fi.Size()
	res, err := call.c.do(call.ctx, req)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// BulkLinkRichMenu method
func (client *Client) BulkLinkRichMenu(richMenuID string, userIDs ...string) *BulkLinkRichMenuCall {
	return &BulkLinkRichMenuCall{
		c:          client,
		userIDs:    userIDs,
		richMenuID: richMenuID,
	}
}

// BulkLinkRichMenuCall type
// Deprecated: Use OpenAPI based classes instead.
type BulkLinkRichMenuCall struct {
	c   *Client
	ctx context.Context

	userIDs    []string
	richMenuID string
}

// WithContext method
func (call *BulkLinkRichMenuCall) WithContext(ctx context.Context) *BulkLinkRichMenuCall {
	call.ctx = ctx
	return call
}

func (call *BulkLinkRichMenuCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		RichMenuID string   `json:"richMenuId"`
		UserIDs    []string `json:"userIds"`
	}{
		RichMenuID: call.richMenuID,
		UserIDs:    call.userIDs,
	})
}

// Do method
func (call *BulkLinkRichMenuCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointBulkLinkRichMenu, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// BulkUnlinkRichMenu method
func (client *Client) BulkUnlinkRichMenu(userIDs ...string) *BulkUnlinkRichMenuCall {
	return &BulkUnlinkRichMenuCall{
		c:       client,
		userIDs: userIDs,
	}
}

// BulkUnlinkRichMenuCall type
// Deprecated: Use OpenAPI based classes instead.
type BulkUnlinkRichMenuCall struct {
	c   *Client
	ctx context.Context

	userIDs []string
}

// WithContext method
func (call *BulkUnlinkRichMenuCall) WithContext(ctx context.Context) *BulkUnlinkRichMenuCall {
	call.ctx = ctx
	return call
}

func (call *BulkUnlinkRichMenuCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		UserIDs []string `json:"userIds"`
	}{
		UserIDs: call.userIDs,
	})
}

// Do method
func (call *BulkUnlinkRichMenuCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointBulkUnlinkRichMenu, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// CreateRichMenuAlias method
func (client *Client) CreateRichMenuAlias(richMenuAliasID, richMenuID string) *CreateRichMenuAliasCall {
	return &CreateRichMenuAliasCall{
		c:               client,
		richMenuAliasID: richMenuAliasID,
		richMenuID:      richMenuID,
	}
}

// CreateRichMenuAliasCall type
// Deprecated: Use OpenAPI based classes instead.
type CreateRichMenuAliasCall struct {
	c   *Client
	ctx context.Context

	richMenuAliasID string
	richMenuID      string
}

// WithContext method
func (call *CreateRichMenuAliasCall) WithContext(ctx context.Context) *CreateRichMenuAliasCall {
	call.ctx = ctx
	return call
}

func (call *CreateRichMenuAliasCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		RichMenuAliasID string `json:"richMenuAliasId"`
		RichMenuID      string `json:"richMenuId"`
	}{
		RichMenuAliasID: call.richMenuAliasID,
		RichMenuID:      call.richMenuID,
	})
}

// Do method
func (call *CreateRichMenuAliasCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointCreateRichMenuAlias, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// UpdateRichMenuAlias method
func (client *Client) UpdateRichMenuAlias(richMenuAliasID, richMenuID string) *UpdateRichMenuAliasCall {
	return &UpdateRichMenuAliasCall{
		c:               client,
		richMenuAliasID: richMenuAliasID,
		richMenuID:      richMenuID,
	}
}

// UpdateRichMenuAliasCall type
// Deprecated: Use OpenAPI based classes instead.
type UpdateRichMenuAliasCall struct {
	c   *Client
	ctx context.Context

	richMenuAliasID string
	richMenuID      string
}

// WithContext method
func (call *UpdateRichMenuAliasCall) WithContext(ctx context.Context) *UpdateRichMenuAliasCall {
	call.ctx = ctx
	return call
}

func (call *UpdateRichMenuAliasCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		RichMenuID string `json:"richMenuId"`
	}{
		RichMenuID: call.richMenuID,
	})
}

// Do method
func (call *UpdateRichMenuAliasCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf(APIEndpointUpdateRichMenuAlias, call.richMenuAliasID)
	res, err := call.c.post(call.ctx, endpoint, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// DeleteRichMenuAlias method
func (client *Client) DeleteRichMenuAlias(richMenuAliasID string) *DeleteRichMenuAliasCall {
	return &DeleteRichMenuAliasCall{
		c:               client,
		richMenuAliasID: richMenuAliasID,
	}
}

// DeleteRichMenuAliasCall type
// Deprecated: Use OpenAPI based classes instead.
type DeleteRichMenuAliasCall struct {
	c   *Client
	ctx context.Context

	richMenuAliasID string
}

// WithContext method
func (call *DeleteRichMenuAliasCall) WithContext(ctx context.Context) *DeleteRichMenuAliasCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *DeleteRichMenuAliasCall) Do() (*BasicResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointDeleteRichMenuAlias, call.richMenuAliasID)
	res, err := call.c.delete(call.ctx, endpoint)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// GetRichMenuAlias method
func (client *Client) GetRichMenuAlias(richMenuAliasID string) *GetRichMenuAliasCall {
	return &GetRichMenuAliasCall{
		c:               client,
		richMenuAliasID: richMenuAliasID,
	}
}

// GetRichMenuAliasCall type
// Deprecated: Use OpenAPI based classes instead.
type GetRichMenuAliasCall struct {
	c   *Client
	ctx context.Context

	richMenuAliasID string
}

// WithContext method
func (call *GetRichMenuAliasCall) WithContext(ctx context.Context) *GetRichMenuAliasCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetRichMenuAliasCall) Do() (*RichMenuAliasResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointGetRichMenuAlias, call.richMenuAliasID)
	res, err := call.c.get(call.ctx, call.c.endpointBase, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToRichMenuAliasResponse(res)
}

// GetRichMenuAliasList method
func (client *Client) GetRichMenuAliasList() *GetRichMenuAliasListCall {
	return &GetRichMenuAliasListCall{
		c: client,
	}
}

// GetRichMenuAliasListCall type
// Deprecated: Use OpenAPI based classes instead.
type GetRichMenuAliasListCall struct {
	c   *Client
	ctx context.Context
}

// WithContext method
func (call *GetRichMenuAliasListCall) WithContext(ctx context.Context) *GetRichMenuAliasListCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetRichMenuAliasListCall) Do() ([]*RichMenuAliasResponse, error) {
	res, err := call.c.get(call.ctx, call.c.endpointBase, APIEndpointListRichMenuAlias, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToRichMenuAliasListResponse(res)
}

// ValidateRichMenuObject method
func (client *Client) ValidateRichMenuObject(richMenu RichMenu) *ValidateRichMenuObjectCall {
	return &ValidateRichMenuObjectCall{
		c:        client,
		richMenu: richMenu,
	}
}

// ValidateRichMenuObjectCall type
// Deprecated: Use OpenAPI based classes instead.
type ValidateRichMenuObjectCall struct {
	c   *Client
	ctx context.Context

	richMenu RichMenu
}

// WithContext method
func (call *ValidateRichMenuObjectCall) WithContext(ctx context.Context) *ValidateRichMenuObjectCall {
	call.ctx = ctx
	return call
}

func (call *ValidateRichMenuObjectCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		Size        RichMenuSize `json:"size"`
		Selected    bool         `json:"selected"`
		Name        string       `json:"name"`
		ChatBarText string       `json:"chatBarText"`
		Areas       []AreaDetail `json:"areas"`
	}{
		Size:        call.richMenu.Size,
		Selected:    call.richMenu.Selected,
		Name:        call.richMenu.Name,
		ChatBarText: call.richMenu.ChatBarText,
		Areas:       call.richMenu.Areas,
	})
}

// Do method
func (call *ValidateRichMenuObjectCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointValidateRichMenuObject, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}
