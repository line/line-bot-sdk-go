package linebot

import (
	"bytes"
	"fmt"
)

// LeaveGroup method
func (client *Client) LeaveGroup(groupID string) (result *ResponseContent, err error) {
	endpoint := fmt.Sprintf(APIEndpointLeaveGroup, groupID)
	result, err = client.post(endpoint, bytes.NewReader([]byte{}))
	return
}

// LeaveRoom method
func (client *Client) LeaveRoom(roomID string) (result *ResponseContent, err error) {
	endpoint := fmt.Sprintf(APIEndpointLeaveRoom, roomID)
	result, err = client.post(endpoint, bytes.NewReader([]byte{}))
	return
}
