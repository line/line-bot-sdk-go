package linebot

import (
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"strconv"
)

// BasicResponse type
type BasicResponse struct {
}

type errorResponseDetail struct {
	Message  string `json:"message"`
	Property string `json:"property"`
}

// ErrorResponse type
type ErrorResponse struct {
	Message string                `json:"message"`
	Details []errorResponseDetail `json:"details"`
}

// UserProfileResponse type
type UserProfileResponse struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PicutureURL   string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

// MessageContentResponse type
type MessageContentResponse struct {
	Content       io.ReadCloser
	ContentLength int64
	ContentType   string
	FileName      string
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

func decodeToMessageContentResponse(res *http.Response) (*MessageContentResponse, error) {
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
	_, params, err := mime.ParseMediaType(res.Header.Get("Content-Disposition"))
	if err != nil {
		return nil, err
	}

	result := MessageContentResponse{
		Content:       res.Body,
		ContentType:   res.Header.Get("Content-Type"),
		ContentLength: res.ContentLength,
		FileName:      params["filename"],
	}
	return &result, nil
}
