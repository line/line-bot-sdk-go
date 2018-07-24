// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package linebot

import (
	"encoding/json"
	"io"
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

// UserProfileResponse type
type UserProfileResponse struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

// MemberIDsResponse type
type MemberIDsResponse struct {
	MemberIDs []string `json:"memberIds"`
	Next      string   `json:"next"`
}

// MessageContentResponse type
type MessageContentResponse struct {
	Content       io.ReadCloser
	ContentLength int64
	ContentType   string
}

// RichMenuIDResponse type
type RichMenuIDResponse struct {
	RichMenuID string `json:"richMenuId"`
}

// RichMenuResponse type
type RichMenuResponse struct {
	RichMenuID  string       `json:"richMenuId"`
	Size        RichMenuSize `json:"size"`
	Selected    bool         `json:"selected"`
	Name        string       `json:"name"`
	ChatBarText string       `json:"chatBarText"`
	Areas       []AreaDetail `json:"areas"`
}

// LIFFAppsResponse type
type LIFFAppsResponse struct {
	Apps []LIFFApp `json:"apps"`
}

// LIFFIDResponse type
type LIFFIDResponse struct {
	LIFFID string `json:"liffId"`
}

// LinkTokenResponse type
type LinkTokenResponse struct {
	LinkToken string `json:"linkToken"`
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
		if err == io.EOF {
			return &result, nil
		}
		return nil, err
	}
	return &result, nil
}

func decodeToUserProfileResponse(res *http.Response) (*UserProfileResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := UserProfileResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeToMemberIDsResponse(res *http.Response) (*MemberIDsResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &MemberIDsResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

func decodeToMessageContentResponse(res *http.Response) (*MessageContentResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	result := MessageContentResponse{
		Content:       res.Body,
		ContentType:   res.Header.Get("Content-Type"),
		ContentLength: res.ContentLength,
	}
	return &result, nil
}

func decodeToRichMenuResponse(res *http.Response) (*RichMenuResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := RichMenuResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeToRichMenuListResponse(res *http.Response) ([]*RichMenuResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	var result = struct {
		RichMenus []*RichMenuResponse `json:"richmenus"`
	}{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return result.RichMenus, nil
}

func decodeToRichMenuIDResponse(res *http.Response) (*RichMenuIDResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := RichMenuIDResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeToLIFFResponse(res *http.Response) (*LIFFAppsResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &LIFFAppsResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

func decodeToLIFFIDResponse(res *http.Response) (*LIFFIDResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := LIFFIDResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeToLinkTokenResponse(res *http.Response) (*LinkTokenResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := LinkTokenResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
