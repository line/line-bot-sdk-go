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
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func mockClient(server *httptest.Server) (*Client, error) {
	client, err := New(
		"testsecret",
		"testtoken",
		WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}),
		WithEndpointBase(server.URL),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TestNewClient(t *testing.T) {
	secret := "testsecret"
	token := "testtoken"
	wantURL, _ := url.Parse(APIEndpointBase)
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
		t.Errorf("endpointBase %q; want %q", client.endpointBase, wantURL)
	}
	if client.httpClient != http.DefaultClient {
		t.Errorf("httpClient %p; want %p", client.httpClient, http.DefaultClient)
	}
}

func TestNewClientWithOptions(t *testing.T) {
	secret := "testsecret"
	token := "testtoken"
	endpoint := "https://example.test/"
	httpClient := http.Client{}
	wantURL, _ := url.Parse(endpoint)
	client, err := New(
		secret,
		token,
		WithHTTPClient(&httpClient),
		WithEndpointBase(endpoint),
	)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(client.endpointBase, wantURL) {
		t.Errorf("endpointBase %q; want %q", client.endpointBase, wantURL)
	}
	if client.httpClient != &httpClient {
		t.Errorf("httpClient %p; want %p", client.httpClient, &httpClient)
	}
}
