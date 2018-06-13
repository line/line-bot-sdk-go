package linebot

import (
	"context"
	"io"
	"encoding/json"
	"bytes"
	"fmt"
)
// LIFFViewType type
type LIFFViewType string

// LIFFViewType constants
const (
	LIFFViewTypeCompact     LIFFViewType = "compact"
	LIFFViewTypeTail     	LIFFViewType = "tail"
	LIFFViewTypeFull	    LIFFViewType = "full"
)

// LIFFIDResponse type
type LIFFIDResponse struct {
	LIFFID string `json:"liffId"`
}

type LIFFAPP struct {
	LIFFID string `json:"liffId"`
	View   View   `json:"view"`
}

type View struct {
	Type LIFFViewType `json:"type"`
	Url  string 	  `json:"url"`
}

// GetRichMenu method
func (client *Client) GetLIFF() *GetLIFFAllCall {
	return &GetLIFFAllCall{
		c:          client,
	}
}

//GetLIFFAllCall type
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
func (call *GetLIFFAllCall) Do() (*LIFFResponse, error) {
	res, err := call.c.get(call.ctx, APIEndpointGetLIFFAPP, nil)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToLIFFResponse(res)
}


// AddLIFFCall method
func (client *Client) AddLIFF(view View) *AddLIFFCall {
	return &AddLIFFCall{
		c:          client,
		View:		view,
	}
}

//AddLIFFCall type
type AddLIFFCall struct {
	c   *Client
	ctx context.Context

	View View
}

// WithContext method
func (call *AddLIFFCall) WithContext(ctx context.Context) *AddLIFFCall {
	call.ctx = ctx
	return call
}

func (call *AddLIFFCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		Type    LIFFViewType   `json:"type"`
		Url		string		   `json:"url"`
	}{
		Type: call.View.Type,
		Url: call.View.Url,
	})
}

// Do method
func (call *AddLIFFCall) Do() (*LIFFIDResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIEndpointAddLIFFAPP, &buf)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToLIFFIDResponse(res)
}

// DeleteRichMenu method
func (client *Client) UpdateLIFFCall(liffId string, view View) *UpdateLIFFCall {
	return &UpdateLIFFCall{
		c:          client,
		LIFFID: 	liffId,
		View:       view,
	}
}

//UpdateLIFFCall type
type UpdateLIFFCall struct {
	c   *Client
	ctx context.Context

	LIFFID string
	View View
}

// WithContext method
func (call *UpdateLIFFCall) WithContext(ctx context.Context) *UpdateLIFFCall {
	call.ctx = ctx
	return call
}

func (call *UpdateLIFFCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&struct {
		Type    LIFFViewType   `json:"type"`
		Url		string		   `json:"url"`
	}{
		Type: call.View.Type,
		Url: call.View.Url,
	})
}

// Do method
func (call *UpdateLIFFCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(APIEndpointUpdateLIFFAPP, call.LIFFID)
	res, err := call.c.put(call.ctx, endpoint, &buf)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToBasicResponse(res)
}

//DeleteLIFFCall type
type DeleteLIFFCall struct {
	c   *Client
	ctx context.Context

	LIFFID string
}

// WithContext method
func (call *DeleteLIFFCall) WithContext(ctx context.Context) *DeleteLIFFCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *DeleteLIFFCall) Do() (*BasicResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointDeleteLIFFAPP, call.LIFFID)
	res, err := call.c.delete(call.ctx, endpoint)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToBasicResponse(res)
}