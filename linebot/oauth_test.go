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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestIssueAccessToken(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *AccessTokenResponse
		Error       error
	}
	testCases := []struct {
		ClientID     string
		ClientSecret string
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			ClientID:     "testid",
			ClientSecret: "testsecret",
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte("client_id=testid&client_secret=testsecret&grant_type=client_credentials"),
				Response:    &AccessTokenResponse{},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != APIEndpointIssueAccessToken {
			t.Errorf("URLPath %s; want %s", r.URL.Path, APIEndpointIssueAccessToken)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		tc := testCases[currentTestIdx]
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
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
		res, err := client.IssueAccessToken(tc.ClientID, tc.ClientSecret).Do()
		if tc.Want.Error != nil {
			if !reflect.DeepEqual(err, tc.Want.Error) {
				t.Errorf("Error %d %q; want %q", i, err, tc.Want.Error)
			}
		} else {
			if err != nil {
				t.Error(err)
			}
		}
		if tc.Want.Response != nil {
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %d %q; want %q", i, res, tc.Want.Response)
			}
		}
	}
}

func TestIssueAccessTokenCall_WithContext(t *testing.T) {
	type fields struct {
		c             *Client
		ctx           context.Context
		channelID     string
		channelSecret string
	}
	type args struct {
		ctx context.Context
	}

	oldCtx := context.Background()
	type key string
	newCtx := context.WithValue(oldCtx, key("foo"), "bar")

	tests := []struct {
		name   string
		fields fields
		args   args
		want   context.Context
	}{
		{
			name: "replace context",
			fields: fields{
				ctx: oldCtx,
			},
			args: args{
				ctx: newCtx,
			},
			want: newCtx,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := &IssueAccessTokenCall{
				c:             tt.fields.c,
				ctx:           tt.fields.ctx,
				channelID:     tt.fields.channelID,
				channelSecret: tt.fields.channelSecret,
			}
			call = call.WithContext(tt.args.ctx)
			got := call.ctx
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IssueAccessTokenCall.WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevokeAccessToken(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		AccessToken  string
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			AccessToken:  "testtoken",
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte("access_token=testtoken"),
				Response:    &BasicResponse{},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != APIEndpointRevokeAccessToken {
			t.Errorf("URLPath %s; want %s", r.URL.Path, APIEndpointIssueAccessToken)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		tc := testCases[currentTestIdx]
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
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
		res, err := client.RevokeAccessToken(tc.AccessToken).Do()
		if tc.Want.Error != nil {
			if !reflect.DeepEqual(err, tc.Want.Error) {
				t.Errorf("Error %d %q; want %q", i, err, tc.Want.Error)
			}
		} else {
			if err != nil {
				t.Error(err)
			}
		}
		if tc.Want.Response != nil {
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %d %q; want %q", i, res, tc.Want.Response)
			}
		}
	}
}

func TestRevokeAccessTokenCall_WithContext(t *testing.T) {
	type fields struct {
		c           *Client
		ctx         context.Context
		accessToken string
	}
	type args struct {
		ctx context.Context
	}

	oldCtx := context.Background()
	type key string
	newCtx := context.WithValue(oldCtx, key("foo"), "bar")

	tests := []struct {
		name   string
		fields fields
		args   args
		want   context.Context
	}{
		{
			name: "replace context",
			fields: fields{
				ctx: oldCtx,
			},
			args: args{
				ctx: newCtx,
			},
			want: newCtx,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := &RevokeAccessTokenCall{
				c:           tt.fields.c,
				ctx:         tt.fields.ctx,
				accessToken: tt.fields.accessToken,
			}
			call = call.WithContext(tt.args.ctx)
			got := call.ctx
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RevokeAccessTokenCall.WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
