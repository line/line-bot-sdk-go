package linebot

import (
	"encoding/json"
	"net/http"
)

// BasicResponse type
type BasicResponse struct {
	RequestID string `json:"requestId"`
}

// ErrorResponse type
type ErrorResponse struct {
	RequestID string `json:"requestId"`
	Message   string `json:"message"`
	Details   []struct {
		Message  string `json:"message"`
		Property string `json:"property"`
	} `json:"details"`
}

// UserProfileResponse type
type UserProfileResponse struct {
	RequestID     string `json:"requestId"`
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PicutureURL   string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

func decodeToBasicResponse(res *http.Response) (*BasicResponse, error) {
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
	result := BasicResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
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
