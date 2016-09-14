package linebot

import (
	"fmt"

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

// Do method
func (call *GetMessageContentCall) Do() (*MessageContentResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointMessageContent, call.messageID)
	res, err := call.c.get(call.ctx, endpoint)
	if err != nil {
		return nil, err
	}
	return decodeToMessageContentResponse(res)
}
