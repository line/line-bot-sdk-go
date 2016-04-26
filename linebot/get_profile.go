package linebot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// UserProfile type
type UserProfile struct {
	Contacts []ContactInfo `json:"contacts"`
	Count    int           `json:"count"`
	Start    int           `json:"start"`
	Display  int           `json:"display"`
}

// ContactInfo type
type ContactInfo struct {
	DisplayName   string `json:"displayName"`
	MID           string `json:"mid"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

// GetUserProfile function
func (client *Client) GetUserProfile(mids []string) (result *UserProfile, err error) {
	query := url.Values{}
	query.Set("mids", strings.Join(mids, ","))
	res, err := client.get(APIEndpointProfiles, query.Encode())
	if err != nil {
		return
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if res.StatusCode != http.StatusOK {
		var content ErrorResponseContent
		if err = decoder.Decode(&content); err != nil {
			return
		}
		return nil, fmt.Errorf("%s: %s", content.Code, content.Message)
	}

	result = &UserProfile{}
	err = decoder.Decode(result)
	if err != nil {
		return
	}
	return
}
