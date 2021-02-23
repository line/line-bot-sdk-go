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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestGetFollowerIDs(t *testing.T) {
	type want struct {
		URLPath           string
		ContinuationToken string
		RequestBody       []byte
		Response          *UserIDsResponse
		Error             error
	}
	testCases := []struct {
		Label             string
		GroupID           string
		ContinuationToken string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:             "With ContinuationToken",
			ContinuationToken: "xxxxx",
			ResponseCode:      200,
			Response:          []byte(`{"userIds": ["U0047556f2e40dba2456887320ba7c76d", "U0047556f2e40dba2456887320ba7c76e"]}`),
			Want: want{
				URLPath:           APIEndpointGetFollowerIDs,
				ContinuationToken: "xxxxx",
				RequestBody:       []byte(""),
				Response: &UserIDsResponse{
					UserIDs: []string{
						"U0047556f2e40dba2456887320ba7c76d",
						"U0047556f2e40dba2456887320ba7c76e",
					},
				},
			},
		},
		{
			Label:        "Without ContinuationToken",
			GroupID:      "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ResponseCode: 200,
			Response:     []byte(`{"userIds": ["U0047556f2e40dba2456887320ba7c76d", "U0047556f2e40dba2456887320ba7c76e"], "next": "xxxxx"}`),
			Want: want{
				URLPath:     APIEndpointGetFollowerIDs,
				RequestBody: []byte(""),
				Response: &UserIDsResponse{
					UserIDs: []string{
						"U0047556f2e40dba2456887320ba7c76d",
						"U0047556f2e40dba2456887320ba7c76e",
					},
					Next: "xxxxx",
				},
			},
		},
		{
			Label:             "Internal server error",
			GroupID:           "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ContinuationToken: "xxxxx",
			ResponseCode:      500,
			Response:          []byte("500 Internal server error"),
			Want: want{
				URLPath:           APIEndpointGetFollowerIDs,
				ContinuationToken: "xxxxx",
				RequestBody:       []byte(""),
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
		q := r.URL.Query()
		if start, want := q.Get("start"), tc.Want.ContinuationToken; start != want {
			t.Errorf("ContinuationToken: %s; want %s", start, want)
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
			res, err := client.GetFollowerIDs(tc.ContinuationToken).Do()
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

func TestGetFollowerIDsWithContext(t *testing.T) {
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
	_, err = client.GetFollowerIDs("xxxxx").WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func BenchmarkGetFollowerIDs(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte(`{"userIds": ["U0047556f2e40dba2456887320ba7c76d", "U0047556f2e40dba2456887320ba7c76e"], "next": "xxxxx"}`))
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
		client.GetFollowerIDs("xxxxx").Do()
	}
}

func TestGetFollowerIDsScanner(t *testing.T) {
	res1 := &UserIDsResponse{
		UserIDs: []string{},
		Next:    "xxxxx",
	}

	res2 := &UserIDsResponse{
		UserIDs: []string{},
	}

	for i := 0; i < 450; i++ {
		if i < 300 {
			res1.UserIDs = append(res1.UserIDs, fmt.Sprintf("U%032d", i))
		} else {
			res2.UserIDs = append(res2.UserIDs, fmt.Sprintf("U%032d", i))
		}
	}

	under300Res := &UserIDsResponse{
		UserIDs: []string{},
	}

	for i := 0; i < 50; i++ {
		under300Res.UserIDs = append(under300Res.UserIDs, fmt.Sprintf("U%032d", i))
	}

	testCases := []struct {
		FirstResponse  *UserIDsResponse
		SecondResponse *UserIDsResponse
	}{
		{
			FirstResponse:  res1,
			SecondResponse: res2,
		},
		{
			FirstResponse: under300Res,
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodGet {
			t.Errorf("Method %s; want %s", r.Method, http.MethodGet)
		}
		q := r.URL.Query()
		w.WriteHeader(200)
		if q.Get("start") == res1.Next {
			if err := json.NewEncoder(w).Encode(tc.SecondResponse); err != nil {
				t.Fatal(err)
			}
			return
		}
		if err := json.NewEncoder(w).Encode(tc.FirstResponse); err != nil {
			t.Fatal(err)
		}
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

	for i := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := client.GetFollowerIDs("").NewScanner()
			for k := 0; s.Scan(); k++ {
				if got, want := s.ID(), fmt.Sprintf("U%032d", k); got != want {
					t.Fatalf("got = %s; want %s; scanner = %#v", got, want, s)
				}
			}
			if err := s.Err(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
