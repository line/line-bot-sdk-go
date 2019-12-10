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
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
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

// MessagesNumberResponse type
type MessagesNumberResponse struct {
	Status  string
	Success int64
}

// MessageQuotaResponse type
type MessageQuotaResponse struct {
	Type       string
	Value      int64
	TotalUsage int64 `json:"totalUsage"`
}

// MessageConsumptionResponse type
type MessageConsumptionResponse struct {
	TotalUsage int64
}

// MessagesNumberDeliveryResponse type
type MessagesNumberDeliveryResponse struct {
	Status          string `json:"status"`
	Broadcast       int64  `json:"broadcast"`
	Targeting       int64  `json:"targeting"`
	AutoResponse    int64  `json:"autoResponse"`
	WelcomeResponse int64  `json:"welcomeResponse"`
	Chat            int64  `json:"chat"`
	APIBroadcast    int64  `json:"apiBroadcast"`
	APIPush         int64  `json:"apiPush"`
	APIMulticast    int64  `json:"apiMulticast"`
	APIReply        int64  `json:"apiReply"`
}

// MessagesNumberFollowersResponse type
type MessagesNumberFollowersResponse struct {
	Status          string `json:"status"`
	Followers       int64  `json:"followers"`
	TargetedReaches int64  `json:"targetedReaches"`
	Blocks          int64  `json:"blocks"`
}

// MessagesFriendDemographicsResponse type
type MessagesFriendDemographicsResponse struct {
	Available           bool                       `json:"available"`
	Genders             []GenderDetail             `json:"genders"`
	Ages                []AgeDetail                `json:"ages"`
	Areas               []AreasDetail          `json:"areas"`
	AppTypes            []AppTypeDetail            `json:"appTypes"`
	SubscriptionPeriods []SubscriptionPeriodDetail `json:"subscriptionPeriods"`
}

// GenderDetail type
type GenderDetail struct {
	Gender     string  `json:"gender"`
	Percentage float64 `json:"percentage"`
}

// AgeDetail type
type AgeDetail struct {
	Age        string  `json:"age"`
	Percentage float64 `json:"percentage"`
}

// AreasDetail type
type AreasDetail struct {
	Area       string  `json:"area"`
	Percentage float64 `json:"percentage"`
}

// AppTypeDetail type
type AppTypeDetail struct {
	AppType    string  `json:"appType"`
	Percentage float64 `json:"percentage"`
}

// SubscriptionPeriodDetail type
type SubscriptionPeriodDetail struct {
	SubscriptionPeriod string  `json:"subscriptionPeriod"`
	Percentage         float64 `json:"percentage"`
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

// isSuccess checks if status code is 2xx: The action was successfully received,
// understood, and accepted.
func isSuccess(code int) bool {
	return code/100 == 2
}

// AccessTokenResponse type
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func checkResponse(res *http.Response) error {
	if isSuccess(res.StatusCode) {
		return nil
	}
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
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return nil, err
	}
	result := MessageContentResponse{
		Content:       ioutil.NopCloser(&buf),
		ContentType:   res.Header.Get("Content-Type"),
		ContentLength: res.ContentLength,
	}
	return &result, nil
}

func decodeToMessageQuotaResponse(res *http.Response) (*MessageQuotaResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &MessageQuotaResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

func decodeToMessageConsumptionResponse(res *http.Response) (*MessageConsumptionResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := &MessageConsumptionResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
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

func decodeToMessagesNumberResponse(res *http.Response) (*MessagesNumberResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := MessagesNumberResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeToMessagesNumberDeliveryResponse(res *http.Response) (*MessagesNumberDeliveryResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := MessagesNumberDeliveryResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeToMessagesNumberFollowersResponse(res *http.Response) (*MessagesNumberFollowersResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := MessagesNumberFollowersResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeToMessagesFriendDemographicsResponse(res *http.Response) (*MessagesFriendDemographicsResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := MessagesFriendDemographicsResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeToAccessTokenResponse(res *http.Response) (*AccessTokenResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := AccessTokenResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
