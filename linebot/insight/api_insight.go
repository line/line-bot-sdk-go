/**
 * LINE Messaging API(Insight)
 * This document describes LINE Messaging API(Insight).
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py

package insight

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type InsightAPI struct {
	httpClient   *http.Client
	endpoint     *url.URL
	channelToken string
	ctx          context.Context
}

// InsightAPIOption type
type InsightAPIOption func(*InsightAPI) error

// New returns a new bot client instance.
func NewInsightAPI(channelToken string, options ...InsightAPIOption) (*InsightAPI, error) {
	if channelToken == "" {
		return nil, errors.New("missing channel access token")
	}

	c := &InsightAPI{
		channelToken: channelToken,
		httpClient:   http.DefaultClient,
	}
	err := c.SetEndpoint("https://api.line.me")
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// WithContext method
func (call *InsightAPI) WithContext(ctx context.Context) *InsightAPI {
	call.ctx = ctx
	return call
}

func (client *InsightAPI) Do(req *http.Request) (*http.Response, error) {
	if client.channelToken != "" {
		req.Header.Set("Authorization", "Bearer "+client.channelToken)
	}
	req.Header.Set("User-Agent", "LINE-BotSDK-Go/"+linebot.GetVersion())
	if client.ctx != nil {
		req = req.WithContext(client.ctx)
	}
	return client.httpClient.Do(req)
}

func (client *InsightAPI) Url(endpointPath string) string {
	u := client.endpoint
	u.Path = path.Join(u.Path, endpointPath)
	return u.String()
}

func (client *InsightAPI) SetEndpoint(endpoint string) error {
	u, err := url.ParseRequestURI(endpoint)
	if err != nil {
		return err
	}
	client.endpoint = u
	return nil
}

// GetFriendsDemographics
//
// Retrieves the demographic attributes for a LINE Official Account's friends.You can only retrieve information about friends for LINE Official Accounts created by users in Japan (JP), Thailand (TH), Taiwan (TW) and Indonesia (ID).
// Parameters:

// https://developers.line.biz/en/reference/messaging-api/#get-demographic
func (client *InsightAPI) GetFriendsDemographics() (*GetFriendsDemographicsResponse, error) {
	path := "/v2/bot/insight/demographic"

	log.Printf("Sending request: method=Get path=%s\n", path)
	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	log.Printf("Got response from '%s %s': status=%d, contentLength=%d", req.Method, req.URL, res.StatusCode, res.ContentLength)

	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		return nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(body))
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	result := GetFriendsDemographicsResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return &result, nil

}

// GetMessageEvent
// Get user interaction statistics
// Returns statistics about how users interact with narrowcast messages or broadcast messages sent from your LINE Official Account.
// Parameters:
//        requestId             Request ID of a narrowcast message or broadcast message. Each Messaging API request has a request ID.

// https://developers.line.biz/en/reference/messaging-api/#get-message-event
func (client *InsightAPI) GetMessageEvent(

	requestId string,

) (*GetMessageEventResponse, error) {
	path := "/v2/bot/insight/message/event"

	log.Printf("Sending request: method=Get path=%s\n", path)
	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, err
	}

	var query url.Values
	query = url.Values{"requestId": []string{string(requestId)}}

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	log.Printf("Got response from '%s %s': status=%d, contentLength=%d", req.Method, req.URL, res.StatusCode, res.ContentLength)

	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		return nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(body))
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	result := GetMessageEventResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return &result, nil

}

// GetNumberOfFollowers
// Get number of followers
// Returns the number of users who have added the LINE Official Account on or before a specified date.
// Parameters:
//        date             Date for which to retrieve the number of followers.  Format: yyyyMMdd (e.g. 20191231) Timezone: UTC+9

// https://developers.line.biz/en/reference/messaging-api/#get-number-of-followers
func (client *InsightAPI) GetNumberOfFollowers(

	date string,

) (*GetNumberOfFollowersResponse, error) {
	path := "/v2/bot/insight/followers"

	log.Printf("Sending request: method=Get path=%s\n", path)
	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, err
	}

	var query url.Values
	query = url.Values{"date": []string{string(date)}}

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	log.Printf("Got response from '%s %s': status=%d, contentLength=%d", req.Method, req.URL, res.StatusCode, res.ContentLength)

	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		return nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(body))
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	result := GetNumberOfFollowersResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return &result, nil

}

// GetNumberOfMessageDeliveries
// Get number of message deliveries
// Returns the number of messages sent from LINE Official Account on a specified day.
// Parameters:
//        date             Date for which to retrieve number of sent messages. - Format: yyyyMMdd (e.g. 20191231) - Timezone: UTC+9

// https://developers.line.biz/en/reference/messaging-api/#get-number-of-delivery-messages
func (client *InsightAPI) GetNumberOfMessageDeliveries(

	date string,

) (*GetNumberOfMessageDeliveriesResponse, error) {
	path := "/v2/bot/insight/message/delivery"

	log.Printf("Sending request: method=Get path=%s\n", path)
	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, err
	}

	var query url.Values
	query = url.Values{"date": []string{string(date)}}

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	log.Printf("Got response from '%s %s': status=%d, contentLength=%d", req.Method, req.URL, res.StatusCode, res.ContentLength)

	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		return nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(body))
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	result := GetNumberOfMessageDeliveriesResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return &result, nil

}

// GetStatisticsPerUnit
//
// You can check the per-unit statistics of how users interact with push messages and multicast messages sent from your LINE Official Account.
// Parameters:
//        customAggregationUnit             Name of aggregation unit specified when sending the message. Case-sensitive. For example, `Promotion_a` and `Promotion_A` are regarded as different unit names.
//        from             Start date of aggregation period.  Format: yyyyMMdd (e.g. 20210301) Time zone: UTC+9
//        to             End date of aggregation period. The end date can be specified for up to 30 days later. For example, if the start date is 20210301, the latest end date is 20210331.  Format: yyyyMMdd (e.g. 20210301) Time zone: UTC+9

// https://developers.line.biz/en/reference/messaging-api/#get-statistics-per-unit
func (client *InsightAPI) GetStatisticsPerUnit(

	customAggregationUnit string,

	from string,

	to string,

) (*GetStatisticsPerUnitResponse, error) {
	path := "/v2/bot/insight/message/event/aggregation"

	log.Printf("Sending request: method=Get path=%s\n", path)
	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, err
	}

	var query url.Values
	query = url.Values{"customAggregationUnit": []string{string(customAggregationUnit)}}
	query = url.Values{"from": []string{string(from)}}
	query = url.Values{"to": []string{string(to)}}

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	log.Printf("Got response from '%s %s': status=%d, contentLength=%d", req.Method, req.URL, res.StatusCode, res.ContentLength)

	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		return nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(body))
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	result := GetStatisticsPerUnitResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return &result, nil

}