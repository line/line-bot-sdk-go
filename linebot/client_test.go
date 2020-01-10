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
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func mockClient(server *httptest.Server, dataServer *httptest.Server) (*Client, error) {
	u, err := url.ParseRequestURI(server.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse server.URL: %v", err)
	}
	du, err := url.ParseRequestURI(dataServer.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse dataServer.URL: %v", err)
	}
	return &Client{
		channelSecret:    "testsecret",
		channelToken:     "testtoken",
		endpointBase:     u,
		endpointBaseData: du,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}, nil
}

func TestNewClient(t *testing.T) {
	secret := "testsecret"
	token := "testtoken"
	wantURL, _ := url.Parse(APIEndpointBase)
	wantDataURL, _ := url.Parse(APIEndpointBaseData)
	client, err := New(secret, token)
	if err != nil {
		t.Fatal(err)
	}
	if client.channelSecret != secret {
		t.Errorf("channelSecret %s; want %s", client.channelSecret, secret)
	}
	if client.channelToken != token {
		t.Errorf("channelToken %s; want %s", client.channelToken, token)
	}
	if !reflect.DeepEqual(client.endpointBase, wantURL) {
		t.Errorf("endpointBase %v; want %v", client.endpointBase, wantURL)
	}
	if !reflect.DeepEqual(client.endpointBaseData, wantDataURL) {
		t.Errorf("endpointBase %v; want %v", client.endpointBaseData, wantDataURL)
	}
	if client.httpClient != http.DefaultClient {
		t.Errorf("httpClient %p; want %p", client.httpClient, http.DefaultClient)
	}
}

func TestNewClientWithOptions(t *testing.T) {
	secret := "testsecret"
	token := "testtoken"
	endpoint := "https://example.test/"
	dataEndpoint := "https://example-data.test/"
	httpClient := http.Client{}
	wantURL, _ := url.Parse(endpoint)
	wantDataURL, _ := url.Parse(dataEndpoint)
	client, err := New(
		secret,
		token,
		WithHTTPClient(&httpClient),
		WithEndpointBase(endpoint),
		WithEndpointBaseData(dataEndpoint),
	)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(client.endpointBase, wantURL) {
		t.Errorf("endpointBase %v; want %v", client.endpointBase, wantURL)
	}
	if !reflect.DeepEqual(client.endpointBaseData, wantDataURL) {
		t.Errorf("endpointBaseData %v; want %v", client.endpointBaseData, wantDataURL)
	}
	if client.httpClient != &httpClient {
		t.Errorf("httpClient %p; want %p", client.httpClient, &httpClient)
	}
}

func expectCtxDeadlineExceed(ctx context.Context, err error, t *testing.T) {
	if err == nil || ctx.Err() != context.DeadlineExceeded {
		t.Errorf("err %v; want %v", err, context.DeadlineExceeded)
	}
}
