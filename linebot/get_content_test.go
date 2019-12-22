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

func TestGetMessageContent(t *testing.T) {
	type want struct {
		URLPath         string
		RequestBody     []byte
		Response        *MessageContentResponse
		ResponseContent []byte
		Error           error
	}
	var testCases = []struct {
		Label          string
		MessageID      string
		ResponseCode   int
		Response       []byte
		ResponseHeader map[string]string
		Want           want
	}{
		{
			Label:        "Success",
			MessageID:    "325708",
			ResponseCode: 200,
			Response:     []byte{0xff, 0xd8, 0xff, 0xe0, 0x00, 0x10},
			ResponseHeader: map[string]string{
				"Content-Type":   "image/jpeg",
				"Content-Length": "6",
			},
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointGetMessageContent, "325708"),
				RequestBody: []byte(""),
				Response: &MessageContentResponse{
					ContentType:   "image/jpeg",
					ContentLength: 6,
				},
				ResponseContent: []byte{0xff, 0xd8, 0xff, 0xe0, 0x00, 0x10},
			},
		},
		{
			Label:        "503 Service Unavailable",
			MessageID:    "325708",
			ResponseCode: 503,
			Response:     []byte("Service Unavailable"),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointGetMessageContent, "325708"),
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 503,
				},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer server.Close()

	var currentTestIdx int
	dataServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		for k, v := range tc.ResponseHeader {
			w.Header().Add(k, v)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			res, err := client.GetMessageContent(tc.MessageID).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				body := res.Content
				defer body.Close()
				res.Content = nil // Set nil because streams aren't comparable.
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
				bodyGot, err := ioutil.ReadAll(body)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(bodyGot, tc.Want.ResponseContent) {
					t.Errorf("ResponseContent %X; want %X", bodyGot, tc.Want.ResponseContent)
				}
			}
		})
	}
}

func TestGetMessageContentWithContext(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer server.Close()

	dataServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte{0xff, 0xd8, 0xff, 0xe0, 0x00, 0x10})
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.GetMessageContent("325708A").WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func BenchmarkGetMessageContent(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b.Error("Unexpected API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer server.Close()

	dataServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte{0xff, 0xd8, 0xff, 0xe0, 0x00, 0x10})
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, _ := client.GetMessageContent("325708A").Do()
		defer res.Content.Close()
		ioutil.ReadAll(res.Content)
	}
}
