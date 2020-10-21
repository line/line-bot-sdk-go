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
	"net/url"
	"reflect"
	"testing"
)

func TestIssueAccessTokenV2(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *AccessTokenResponse
		Error       error
	}
	var testCases = []struct {
		ClientAssertion string
		Response        []byte
		ResponseCode    int
		Want            want
	}{
		{
			ClientAssertion: "testclientassertion",
			ResponseCode:    200,
			Response:        []byte(`{"access_token":"eyJhbGciOiJIUz.....","token_type":"Bearer","expires_in": 2592000,"key_id":"sDTOzw5wIfxxxxPEzcmeQA"}`),
			Want: want{
				RequestBody: []byte("client_assertion=testclientassertion&client_assertion_type=urn%3Aietf%3Aparams%3Aoauth%3Aclient-assertion-type%3Ajwt-bearer&grant_type=client_credentials"),
				Response: &AccessTokenResponse{
					AccessToken: "eyJhbGciOiJIUz.....",
					ExpiresIn:   2592000,
					TokenType:   "Bearer",
					KeyID:       "sDTOzw5wIfxxxxPEzcmeQA",
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != APIEndpointIssueAccessTokenV2 {
			t.Errorf("URLPath %s; want %s", r.URL.Path, APIEndpointIssueAccessTokenV2)
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
		res, err := client.IssueAccessTokenV2(tc.ClientAssertion).Do()
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

func TestIssueAccessTokenV2Call_WithContext(t *testing.T) {
	type fields struct {
		c               *Client
		ctx             context.Context
		clientAssertion string
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
			call := &IssueAccessTokenV2Call{
				c:               tt.fields.c,
				ctx:             tt.fields.ctx,
				clientAssertion: tt.fields.clientAssertion,
			}
			call = call.WithContext(tt.args.ctx)
			got := call.ctx
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IssueAccessTokenV2Call.WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAccessTokensV2(t *testing.T) {
	type want struct {
		RequestParams url.Values
		Response      *AccessTokensResponse
		Error         error
	}
	var testCases = []struct {
		ClientAssertion string
		Response        []byte
		ResponseCode    int
		Want            want
	}{
		{
			ClientAssertion: "testclientassertion",
			ResponseCode:    200,
			Response:        []byte(`{"kids":["U_gdnFYKTWRxxxxDVZexGg", "sDTOzw5wIfWxxxxzcmeQA", "73hDyp3PxGfxxxxD6U5qYA"]}`),
			Want: want{
				RequestParams: url.Values{
					"client_assertion_type": []string{clientAssertionTypeJWT},
					"client_assertion":      []string{"testclientassertion"},
				},
				Response: &AccessTokensResponse{
					KeyIDs: []string{"U_gdnFYKTWRxxxxDVZexGg", "sDTOzw5wIfWxxxxzcmeQA", "73hDyp3PxGfxxxxD6U5qYA"},
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodGet {
			t.Errorf("Method %s; want %s", r.Method, http.MethodGet)
		}
		if r.URL.Path != APIEndpointGetAccessTokensV2 {
			t.Errorf("URLPath %s; want %s", r.URL.Path, APIEndpointGetAccessTokensV2)
		}
		tc := testCases[currentTestIdx]
		if !reflect.DeepEqual(r.URL.Query(), tc.Want.RequestParams) {
			t.Errorf("RequestParams %v; want %v", r.URL.Query(), tc.Want.RequestParams)
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
		res, err := client.GetAccessTokensV2(tc.ClientAssertion).Do()
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

func TestGetAccessTokensV2Call_WithContext(t *testing.T) {
	type fields struct {
		c               *Client
		ctx             context.Context
		clientAssertion string
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
			call := &GetAccessTokensV2Call{
				c:               tt.fields.c,
				ctx:             tt.fields.ctx,
				clientAssertion: tt.fields.clientAssertion,
			}
			call = call.WithContext(tt.args.ctx)
			got := call.ctx
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccessTokensV2Call.WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevokeAccessTokenV2(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	var testCases = []struct {
		AccessToken  string
		ClientID     string
		ClientSecret string
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			AccessToken:  "testtoken",
			ClientID:     "testid",
			ClientSecret: "testsecret",
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte("access_token=testtoken&client_id=testid&client_secret=testsecret"),
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
		if r.URL.Path != APIEndpointRevokeAccessTokenV2 {
			t.Errorf("URLPath %s; want %s", r.URL.Path, APIEndpointIssueAccessTokenV2)
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
		res, err := client.RevokeAccessTokenV2(tc.ClientID, tc.ClientSecret, tc.AccessToken).Do()
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

func TestRevokeAccessTokenV2Call_WithContext(t *testing.T) {
	type fields struct {
		c             *Client
		ctx           context.Context
		accessToken   string
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
			call := &RevokeAccessTokenV2Call{
				c:             tt.fields.c,
				ctx:           tt.fields.ctx,
				accessToken:   tt.fields.accessToken,
				channelID:     tt.fields.channelID,
				channelSecret: tt.fields.channelSecret,
			}
			call = call.WithContext(tt.args.ctx)
			got := call.ctx
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RevokeAccessTokenV2Call.WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
