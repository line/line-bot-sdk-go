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
