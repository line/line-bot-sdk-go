package linebot

import (
	"fmt"
	"io"
	"mime"

	"golang.org/x/net/context"
)

// GetMessageContent function
func (client *Client) GetMessageContent(messageID string) *GetMessageContentCall {
	return &GetMessageContentCall{
		c:         client,
		messageID: messageID,
	}
}

// GetMessageContentCall type
type GetMessageContentCall struct {
	c   *Client
	ctx context.Context

	messageID string
}

// WithContext method
func (call *GetMessageContentCall) WithContext(ctx context.Context) *GetMessageContentCall {
	call.ctx = ctx
	return call
}

// MessageContentResponse type
type MessageContentResponse struct {
	Content  io.ReadCloser
	FileName string
}

// Do method
func (call *GetMessageContentCall) Do() (*MessageContentResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointMessageContent, call.messageID)
	res, err := call.c.get(call.ctx, endpoint)
	mc := &MessageContentResponse{
		Content: res.Body,
	}
	_, params, err := mime.ParseMediaType(res.Header.Get("Content-Disposition"))
	if err != nil {
		return nil, err
	}
	mc.FileName = params["filename"]
	return mc, nil
}
