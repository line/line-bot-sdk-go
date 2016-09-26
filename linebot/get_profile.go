package linebot

import (
	"fmt"

	"golang.org/x/net/context"
)

// GetProfile function
func (client *Client) GetProfile(userID string) *GetProfileCall {
	return &GetProfileCall{
		c:      client,
		userID: userID,
	}
}

// GetProfileCall type
type GetProfileCall struct {
	c   *Client
	ctx context.Context

	userID string
}

// WithContext method
func (call *GetProfileCall) WithContext(ctx context.Context) *GetProfileCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetProfileCall) Do() (*UserProfileResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointGetProfile, call.userID)
	res, err := call.c.get(call.ctx, endpoint)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToUserProfileResponse(res)
}
