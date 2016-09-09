package linebot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

// UserProfileResponse type
type UserProfileResponse struct {
	RequestID     string `json:"requestId"`
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PicutureURL   string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

func decodeToUserProfileResponse(res *http.Response) (*UserProfileResponse, error) {
	decoder := json.NewDecoder(res.Body)
	if res.StatusCode != http.StatusOK {
		result := ErrorResponse{}
		if err := decoder.Decode(&result); err != nil {
			return nil, &APIError{
				Code: res.StatusCode,
			}
		}
		return nil, &APIError{
			Code:     res.StatusCode,
			Response: &result,
		}
	}
	result := UserProfileResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

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
	res, err := call.c.getCtx(call.ctx, endpoint)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeToUserProfileResponse(res)
}
