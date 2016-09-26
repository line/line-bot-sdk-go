package linebot

import (
	"encoding/json"
	"io"
	"mime"
	"net/http"
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

// ProfileResponse type
type ProfileResponse struct {
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

func checkResponse(res *http.Response) error {
	if res.StatusCode != http.StatusOK {
		decoder := json.NewDecoder(res.Body)
		result := ErrorResponse{}
		if err := decoder.Decode(&result); err != nil {
			return &APIError{
				Code: res.StatusCode,
			}
		}
		return &APIError{
			Code:     res.StatusCode,
			Response: &result,
		}
	}
	return nil
}

func decodeToBasicResponse(res *http.Response) (*BasicResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := BasicResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeToProfileResponse(res *http.Response) (*ProfileResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := ProfileResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeToMessageContentResponse(res *http.Response) (*MessageContentResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
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
