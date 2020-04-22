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

func TestGetProfile(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *UserProfileResponse
		Error       error
	}
	var testCases = []struct {
		Label        string
		UserID       string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			Label:        "Success",
			UserID:       "U0047556f2e40dba2456887320ba7c76d",
			ResponseCode: 200,
			Response:     []byte(`{"userId":"U0047556f2e40dba2456887320ba7c76d","displayName":"BOT API","pictureUrl":"https://obs.line-apps.com/abcdefghijklmn","statusMessage":"Hello, LINE!","language":"en"}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointGetProfile, "U0047556f2e40dba2456887320ba7c76d"),
				RequestBody: []byte(""),
				Response: &UserProfileResponse{
					UserID:        "U0047556f2e40dba2456887320ba7c76d",
					DisplayName:   "BOT API",
					PictureURL:    "https://obs.line-apps.com/abcdefghijklmn",
					StatusMessage: "Hello, LINE!",
					Language:      "en",
				},
			},
		},
		{
			Label:        "Internal server error",
			UserID:       "U0047556f2e40dba2456887320ba7c76d",
			ResponseCode: 500,
			Response:     []byte("500 Internal server error"),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointGetProfile, "U0047556f2e40dba2456887320ba7c76d"),
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
			res, err := client.GetProfile(tc.UserID).Do()
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

func TestGetProfileWithContext(t *testing.T) {
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
	_, err = client.GetProfile("U0047556f2e40dba2456887320ba7c76d").WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func BenchmarkGetProfile(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte(`{"userId":"U","displayName":"A","pictureUrl":"https://","statusMessage":"B"}`))
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
		client.GetProfile("U0047556f2e40dba2456887320ba7c76d").Do()
	}
}

func TestGetGroupMemberProfile(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *UserProfileResponse
		Error       error
	}
	var testCases = []struct {
		Label        string
		GroupID      string
		UserID       string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			Label:        "Success",
			GroupID:      "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			UserID:       "U0047556f2e40dba2456887320ba7c76d",
			ResponseCode: 200,
			Response:     []byte(`{"userId":"U0047556f2e40dba2456887320ba7c76d","displayName":"BOT API","pictureUrl":"https://obs.line-apps.com/abcdefghijklmn","statusMessage":"Hello, LINE!"}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointGetGroupMemberProfile, "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "U0047556f2e40dba2456887320ba7c76d"),
				RequestBody: []byte(""),
				Response: &UserProfileResponse{
					UserID:        "U0047556f2e40dba2456887320ba7c76d",
					DisplayName:   "BOT API",
					PictureURL:    "https://obs.line-apps.com/abcdefghijklmn",
					StatusMessage: "Hello, LINE!",
				},
			},
		},
		{
			Label:        "Internal server error",
			GroupID:      "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			UserID:       "U0047556f2e40dba2456887320ba7c76d",
			ResponseCode: 500,
			Response:     []byte("500 Internal server error"),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointGetGroupMemberProfile, "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "U0047556f2e40dba2456887320ba7c76d"),
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
			res, err := client.GetGroupMemberProfile(tc.GroupID, tc.UserID).Do()
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

func TestGetGroupMemberProfileWithContext(t *testing.T) {
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
	_, err = client.GetGroupMemberProfile("cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "U0047556f2e40dba2456887320ba7c76d").WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func BenchmarkGetGroupMemberProfile(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte(`{"userId":"U","displayName":"A","pictureUrl":"https://","statusMessage":"B"}`))
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
		client.GetGroupMemberProfile("cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "U0047556f2e40dba2456887320ba7c76d").Do()
	}
}

func TestGetRoomMemberProfile(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *UserProfileResponse
		Error       error
	}
	var testCases = []struct {
		Label        string
		RoomID       string
		UserID       string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			Label:        "Success",
			RoomID:       "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			UserID:       "U0047556f2e40dba2456887320ba7c76d",
			ResponseCode: 200,
			Response:     []byte(`{"userId":"U0047556f2e40dba2456887320ba7c76d","displayName":"BOT API","pictureUrl":"https://obs.line-apps.com/abcdefghijklmn","statusMessage":"Hello, LINE!"}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointGetRoomMemberProfile, "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "U0047556f2e40dba2456887320ba7c76d"),
				RequestBody: []byte(""),
				Response: &UserProfileResponse{
					UserID:        "U0047556f2e40dba2456887320ba7c76d",
					DisplayName:   "BOT API",
					PictureURL:    "https://obs.line-apps.com/abcdefghijklmn",
					StatusMessage: "Hello, LINE!",
				},
			},
		},
		{
			Label:        "Internal server error",
			RoomID:       "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			UserID:       "U0047556f2e40dba2456887320ba7c76d",
			ResponseCode: 500,
			Response:     []byte("500 Internal server error"),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointGetRoomMemberProfile, "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "U0047556f2e40dba2456887320ba7c76d"),
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
			res, err := client.GetRoomMemberProfile(tc.RoomID, tc.UserID).Do()
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

func TestGetRoomMemberProfileWithContext(t *testing.T) {
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
	_, err = client.GetRoomMemberProfile("cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "U0047556f2e40dba2456887320ba7c76d").WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func BenchmarkGetRoomMemberProfile(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte(`{"userId":"U","displayName":"A","pictureUrl":"https://","statusMessage":"B"}`))
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
		client.GetRoomMemberProfile("cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "U0047556f2e40dba2456887320ba7c76d").Do()
	}
}
