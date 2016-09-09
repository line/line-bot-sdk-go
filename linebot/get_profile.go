package linebot

import (
	"fmt"

	"golang.org/x/net/context"
)

// GetUserProfile function
func (client *Client) GetUserProfile(userID string) *GetUserProfileCall {
	return &GetUserProfileCall{
		c:      client,
		userID: userID,
	}
}

// GetUserProfileCall type
type GetUserProfileCall struct {
	c   *Client
	ctx context.Context

	userID string
}

// WithContext method
func (call *GetUserProfileCall) WithContext(ctx context.Context) *GetUserProfileCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetUserProfileCall) Do() (*UserProfileResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointGetUserProfile, call.userID)
	res, err := call.c.get(call.ctx, endpoint)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToUserProfileResponse(res)
}
