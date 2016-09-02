package linebot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UserProfileResponse type
type UserProfileResponse struct {
	RequestID     string `json:"requestId"`
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PicutureURL   string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

// GetUserProfile function
func (client *Client) GetUserProfile(uid string) (result *UserProfileResponse, err error) {
	res, err := client.get("/v2/bot/profile/" + uid)
	if err != nil {
		return
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if res.StatusCode != http.StatusOK {
		var content ResponseContent
		if err = decoder.Decode(&content); err != nil {
			return
		}
		return nil, fmt.Errorf("%d: %s", res.StatusCode, content.Message)
	}

	result = &UserProfileResponse{}
	err = decoder.Decode(result)
	if err != nil {
		return
	}
	return
}
