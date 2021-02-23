// Copyright 2019 LINE Corporation
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
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"
)

// TestInsights tests GetNumberMessagesDelivery, GetNumberFollowers
// and GetFriendDemographics func

func TestGetNumberMessagesDelivery(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *MessagesNumberDeliveryResponse
		Error       error
	}
	testCases := []struct {
		Label        string
		Date         string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			Label:        "Success",
			Date:         "20190418",
			ResponseCode: 200,
			Response: []byte(`{
				"status": "ready",
				"broadcast": 5385,
				"targeting": 522
			}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointInsight, InsightTypeMessageDelivery),
				RequestBody: []byte(""),
				Response: &MessagesNumberDeliveryResponse{
					Status:    "ready",
					Broadcast: 5385,
					Targeting: 522,
				},
			},
		},
		{
			Label:        "Internal server error",
			Date:         "20190418",
			ResponseCode: 500,
			Response:     []byte("500 Internal server error"),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointInsight, InsightTypeMessageDelivery),
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 500,
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodGet {
			t.Errorf("Method %s; want %s", r.Method, http.MethodGet)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			res, err := client.GetNumberMessagesDelivery(tc.Date).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %v; want %v", res, tc.Want.Response)
			}
		})
	}
}

func TestGetNumberMessagesDeliveryContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.GetNumberMessagesDelivery("20190418").WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func BenchmarkGetNumberMessagesDelivery(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte(`{
			"status": "ready",
			"broadcast": 5385,
			"targeting": 522
		}`))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.GetNumberMessagesDelivery("20190418").Do()
	}
}

func TestGetNumberFollowers(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *MessagesNumberFollowersResponse
		Error       error
	}
	testCases := []struct {
		Label        string
		Date         string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			Label:        "Success",
			Date:         "20190418",
			ResponseCode: 200,
			Response: []byte(`{
				"status": "ready",
				"followers": 7620,
				"targetedReaches": 5848,
				"blocks": 237
			}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointInsight, InsightTypeFollowers),
				RequestBody: []byte(""),
				Response: &MessagesNumberFollowersResponse{
					Status:          "ready",
					Followers:       7620,
					TargetedReaches: 5848,
					Blocks:          237,
				},
			},
		},
		{
			Label:        "Internal server error",
			Date:         "20190418",
			ResponseCode: 500,
			Response:     []byte("500 Internal server error"),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointInsight, InsightTypeFollowers),
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 500,
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodGet {
			t.Errorf("Method %s; want %s", r.Method, http.MethodGet)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			res, err := client.GetNumberFollowers(tc.Date).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %v; want %v", res, tc.Want.Response)
			}
		})
	}
}

func TestGetNumberFollowersContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.GetNumberFollowers("20190418").WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func BenchmarkGetNumberFollowers(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte(`{
			"status": "ready",
			"followers": 7620,
			"targetedReaches": 5848,
			"blocks": 237
		}`))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.GetNumberFollowers("20190418").Do()
	}
}

func TestGetFriendDemographics(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *MessagesFriendDemographicsResponse
		Error       error
	}
	testCases := []struct {
		Label        string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			Label:        "Success",
			ResponseCode: 200,
			Response: []byte(`{
				"available": true,
				"genders": [
					{
						"gender": "unknown",
						"percentage": 37.6
					}
				],
				"ages": [
					{
						"age": "unknown",
						"percentage": 37.6
					}
				],
				"areas": [
					{
						"area": "unknown",
						"percentage": 42.9
					}
				],
				"appTypes": [
					{
						"appType": "ios",
						"percentage": 62.4
					}
				],
				"subscriptionPeriods": [
					{
						"subscriptionPeriod": "over365days",
						"percentage": 96.4
					}
				]
			}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointInsight, InsightTypeDemographic),
				RequestBody: []byte(""),
				Response: &MessagesFriendDemographicsResponse{
					Available: true,
					Genders: []GenderDetail{
						{Gender: "unknown", Percentage: 37.6},
					},
					Ages: []AgeDetail{
						{Age: "unknown", Percentage: 37.6},
					},
					Areas: []AreasDetail{
						{Area: "unknown", Percentage: 42.9},
					},
					AppTypes: []AppTypeDetail{
						{AppType: "ios", Percentage: 62.4},
					},
					SubscriptionPeriods: []SubscriptionPeriodDetail{
						{SubscriptionPeriod: "over365days", Percentage: 96.4},
					},
				},
			},
		},
		{
			Label:        "Internal server error",
			ResponseCode: 500,
			Response:     []byte("500 Internal server error"),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointInsight, InsightTypeDemographic),
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 500,
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodGet {
			t.Errorf("Method %s; want %s", r.Method, http.MethodGet)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			res, err := client.GetFriendDemographics().Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %v; want %v", res, tc.Want.Response)
			}
		})
	}
}

func TestGetFriendDemographicsContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.GetFriendDemographics().WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func BenchmarkGetFriendDemographics(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte(`{
			"available": true,
			"genders": [
				{
					"gender": "unknown",
					"percentage": 37.6
				}
			],
			"ages": [
				{
					"age": "unknown",
					"percentage": 37.6
				}
			],
			"areas": [
				{
					"area": "unknown",
					"percentage": 42.9
				}
			],
			"appTypes": [
				{
					"appType": "ios",
					"percentage": 62.4
				}
			],
			"subscriptionPeriods": [
				{
					"subscriptionPeriod": "over365days",
					"percentage": 96.4
				}
			]
		}`))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.GetFriendDemographics().Do()
	}
}

func TestGetUserInteractionStats(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *MessagesUserInteractionStatsResponse
		Error       error
	}
	testCases := []struct {
		Label        string
		RequestID    string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			Label:        "Success",
			RequestID:    "f70dd685-499a-4231-a441-f24b8d4fba21",
			ResponseCode: 200,
			Response: []byte(`{
				"overview": {
					"requestId": "f70dd685-499a-4231-a441-f24b8d4fba21",
					"timestamp": 1568214000,
					"delivered": 32,
					"uniqueImpression": 4,
					"uniqueClick": null,
					"uniqueMediaPlayed": 2,
					"uniqueMediaPlayed100Percent": -1
				},
				"messages": [
					{
						"seq": 1,
						"impression": 18,
						"mediaPlayed": 11,
						"mediaPlayed25Percent": -1,
						"mediaPlayed50Percent": -1,
						"mediaPlayed75Percent": -1,
						"mediaPlayed100Percent": -1,
						"uniqueMediaPlayed": 2,
						"uniqueMediaPlayed25Percent": -1,
						"uniqueMediaPlayed50Percent": -1,
						"uniqueMediaPlayed75Percent": -1,
						"uniqueMediaPlayed100Percent": -1
					}
				],
				"clicks": [
					{
						"seq": 1,
						"url": "https://www.yahoo.co.jp/",
						"click": -1,
						"uniqueClick": -1,
						"uniqueClickOfRequest": -1
					},
					{
						"seq": 1,
						"url": "https://www.google.com/?hl=ja",
						"click": -1,
						"uniqueClick": -1,
						"uniqueClickOfRequest": -1
					}
				]
			}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointInsight, InsightTypeUserInteractionStats),
				RequestBody: []byte(""),
				Response: &MessagesUserInteractionStatsResponse{
					Overview: OverviewDetail{
						RequestID:                   "f70dd685-499a-4231-a441-f24b8d4fba21",
						Timestamp:                   1568214000,
						Delivered:                   32,
						UniqueImpression:            4,
						UniqueClick:                 0,
						UniqueMediaPlayed:           2,
						UniqueMediaPlayed100Percent: -1,
					},
					Messages: []MessageDetail{
						{
							Seq:                         1,
							Impression:                  18,
							MediaPlayed:                 11,
							MediaPlayed25Percent:        -1,
							MediaPlayed50Percent:        -1,
							MediaPlayed75Percent:        -1,
							MediaPlayed100Percent:       -1,
							UniqueMediaPlayed:           2,
							UniqueMediaPlayed25Percent:  -1,
							UniqueMediaPlayed50Percent:  -1,
							UniqueMediaPlayed75Percent:  -1,
							UniqueMediaPlayed100Percent: -1,
						},
					},
					Clicks: []ClickDetail{
						{
							Seq:                  1,
							URL:                  "https://www.yahoo.co.jp/",
							Click:                -1,
							UniqueClick:          -1,
							UniqueClickOfRequest: -1,
						},
						{
							Seq:                  1,
							URL:                  "https://www.google.com/?hl=ja",
							Click:                -1,
							UniqueClick:          -1,
							UniqueClickOfRequest: -1,
						},
					},
				},
			},
		},
		{
			Label:        "Internal server error",
			ResponseCode: 500,
			Response:     []byte("500 Internal server error"),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointInsight, InsightTypeUserInteractionStats),
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 500,
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodGet {
			t.Errorf("Method %s; want %s", r.Method, http.MethodGet)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			res, err := client.GetUserInteractionStats(tc.RequestID).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %v; want %v", res, tc.Want.Response)
			}
		})
	}
}

func TestGetUserInteractionStatsContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.GetUserInteractionStats("f70dd685-499a-4231-a441-f24b8d4fba21").WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}
