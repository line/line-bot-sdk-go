package linebot

import (
	"fmt"

	"golang.org/x/net/context"
)

// LeaveGroup method
func (client *Client) LeaveGroup(groupID string) *LeaveGroupCall {
	return &LeaveGroupCall{
		c:       client,
		groupID: groupID,
	}
}

// LeaveGroupCall type
type LeaveGroupCall struct {
	c   *Client
	ctx context.Context

	groupID string
}

// WithContext method
func (call *LeaveGroupCall) WithContext(ctx context.Context) *LeaveGroupCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *LeaveGroupCall) Do() (*BasicResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointLeaveGroup, call.groupID)
	res, err := call.c.post(call.ctx, endpoint, nil)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToBasicResponse(res)
}

// LeaveRoom method
func (client *Client) LeaveRoom(roomID string) *LeaveRoomCall {
	return &LeaveRoomCall{
		c:      client,
		roomID: roomID,
	}
}

// LeaveRoomCall type
type LeaveRoomCall struct {
	c   *Client
	ctx context.Context

	roomID string
}

// WithContext method
func (call *LeaveRoomCall) WithContext(ctx context.Context) *LeaveRoomCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *LeaveRoomCall) Do() (*BasicResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointLeaveRoom, call.roomID)
	res, err := call.c.post(call.ctx, endpoint, nil)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToBasicResponse(res)
}
