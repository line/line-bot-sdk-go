// Copyright 2020 LINE Corporation
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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestGetBotInfo(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *BotInfoResponse
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
			Response:     []byte(`{"userId":"u206d25c2ea6bd87c17655609a1c37cb8","basicId":"@012abcde","premiumId":"PremiumId","displayName":"BotTest","pictureUrl":"https://example.com/abcdefghijklmn","chatMode":"chat","markAsReadMode":"manual"}`),
			Want: want{
				URLPath:     APIEndpointGetBotInfo,
				RequestBody: []byte(""),
				Response: &BotInfoResponse{
					UserID:         "u206d25c2ea6bd87c17655609a1c37cb8",
					BasicID:        "@012abcde",
					PremiumID:      "PremiumId",
					DisplayName:    "BotTest",
					PictureURL:     "https://example.com/abcdefghijklmn",
					ChatMode:       ChatModeChat,
					MarkAsReadMode: MarkAsReadModeManual,
				},
			},
		},
		{
			Label:        "No premiumID Success",
			ResponseCode: 200,
			Response:     []byte(`{"userId":"u206d25c2ea6bd87c17655609a1c37cb8","basicId":"@012abcde","displayName":"BotTest","pictureUrl":"https://example.com/abcdefghijklmn","chatMode":"bot","markAsReadMode":"auto"}`),
			Want: want{
				URLPath:     APIEndpointGetBotInfo,
				RequestBody: []byte(""),
				Response: &BotInfoResponse{
					UserID:         "u206d25c2ea6bd87c17655609a1c37cb8",
					BasicID:        "@012abcde",
					PremiumID:      "",
					DisplayName:    "BotTest",
					PictureURL:     "https://example.com/abcdefghijklmn",
					ChatMode:       ChatModeBot,
					MarkAsReadMode: MarkAsReadModeAuto,
				},
			},
		},
		{
			Label:        "Internal server error",
			ResponseCode: 500,
			Response:     []byte("500 Internal server error"),
			Want: want{
				URLPath:     APIEndpointGetBotInfo,
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 500,
				},
			},
		},
		{
			Label:        "Invalid channelAccessToken error",
			ResponseCode: 401,
			Response:     []byte(`{"message":"Authentication failed due to the following reason: invalid token. Confirm that the access token in the authorization header is valid."}`),
			Want: want{
				URLPath:     APIEndpointGetBotInfo,
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 401,
					Response: &ErrorResponse{
						Message: "Authentication failed due to the following reason: invalid token. Confirm that the access token in the authorization header is valid.",
					},
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
			res, err := client.GetBotInfo().Do()
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

func TestGetBotInfoWithContext(t *testing.T) {
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
	_, err = client.GetBotInfo().WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func BenchmarkGetBotInfo(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte(`{"userId":"U","basicId":"@B","premiumId":"P","displayName":"BotTest","pictureUrl":"https://","chatMode":"chat","markAsReadMode":"manual"}`))
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
		client.GetBotInfo().Do()
	}
}
