// Copyright 2018 LINE Corporation
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
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

func TestGetLIFF(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *LIFFAppsResponse
		Error       error
	}
	testCases := []struct {
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			ResponseCode: 200,
			Response:     []byte(`{"apps":[{"liffId":"testliffId1","view":{"type":"full","url":"https://example.com/myservice"}},{"liffId":"testliffId2","view":{"type":"tall","url":"https://example.com/myservice2"}}]}`),
			Want: want{
				RequestBody: []byte(``),
				Response: &LIFFAppsResponse{
					Apps: []LIFFApp{
						{
							LIFFID: "testliffId1",
							View: View{
								Type: LIFFViewTypeFull,
								URL:  "https://example.com/myservice",
							},
						},
						{
							LIFFID: "testliffId2",
							View: View{
								Type: LIFFViewTypeTall,
								URL:  "https://example.com/myservice2",
							},
						},
					},
				},
			},
		},
		{
			ResponseCode: 200,
			Response:     []byte(`{"apps":[{"liffId":"{liffId}","view":{"type":"full","url":"https://example.com/myservice"},"description":"Happy New York","permanentLinkPattern":"concat"},{"liffId":"{liffId}","view":{"type":"tall","url":"https://example.com/myservice2"},"features":{"ble":true,"qrCode":true},"permanentLinkPattern":"concat","scope":["profile","chat_message.write"],"botPrompt":"none"}]}`),
			Want: want{
				RequestBody: []byte(``),
				Response: &LIFFAppsResponse{
					Apps: []LIFFApp{
						{
							LIFFID: "{liffId}",
							View: View{
								Type: LIFFViewTypeFull,
								URL:  "https://example.com/myservice",
							},
							Description:          "Happy New York",
							PermanentLinkPattern: "concat",
						},
						{
							LIFFID: "{liffId}",
							View: View{
								Type: LIFFViewTypeTall,
								URL:  "https://example.com/myservice2",
							},
							Features:             &LIFFAppFeatures{BLE: true, QRCode: true},
							PermanentLinkPattern: "concat",
							Scope:                []LIFFViewScopeType{LIFFViewScopeTypeProfile, LIFFViewScopeTypeChatMessageWrite},
							BotPrompt:            "none",
						},
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
		endpoint := APIEndpointGetAllLIFFApps
		if r.URL.Path != endpoint {
			t.Errorf("URLPath %s; want %s", r.URL.Path, endpoint)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody\n %s; want\n %s", body, tc.Want.RequestBody)
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
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res, err := client.GetLIFF().Do()
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %v; want %v", res, tc.Want.Response)
			}
		})
	}
}

func TestAddLIFF(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *LIFFIDResponse
		Error       error
	}
	testCases := []struct {
		LIFFApp      LIFFApp
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			LIFFApp: LIFFApp{
				View: View{
					Type: LIFFViewTypeFull,
					URL:  "https://example.com/myservice",
				},
			},
			ResponseCode: 200,
			Response:     []byte(`{"liffId":"testliffId"}`),
			Want: want{
				RequestBody: []byte(`{"view":{"type":"full","url":"https://example.com/myservice"}}` + "\n"),
				Response:    &LIFFIDResponse{LIFFID: "testliffId"},
			},
		},
		{
			LIFFApp: LIFFApp{
				View: View{
					Type: LIFFViewTypeFull,
					URL:  "https://example.com/myservice2",
				},
				Scope:                []LIFFViewScopeType{LIFFViewScopeTypeProfile, LIFFViewScopeTypeChatMessageWrite},
				PermanentLinkPattern: "concat",
				Description:          "Service Example",
				Features:             &LIFFAppFeatures{BLE: true, QRCode: true},
				BotPrompt:            "none",
			},
			ResponseCode: 200,
			Response:     []byte(`{"liffId":"testliffId3"}`),
			Want: want{
				RequestBody: []byte(`{"view":{"type":"full","url":"https://example.com/myservice2"},"description":"Service Example","features":{"ble":true,"qrCode":true},"permanentLinkPattern":"concat","scope":["profile","chat_message.write"],"botPrompt":"none"}` + "\n"),
				Response:    &LIFFIDResponse{LIFFID: "testliffId3"},
			},
		},
	}
	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		endpoint := APIEndpointAddLIFFApp
		if r.URL.Path != endpoint {
			t.Errorf("URLPath %s; want %s", r.URL.Path, endpoint)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("%d, RequestBody\n %s; want\n %s", currentTestIdx, body, tc.Want.RequestBody)
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
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res, err := client.AddLIFF(tc.LIFFApp).Do()
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %v; want %v", res, tc.Want.Response)
			}
		})
	}
}

func TestUpdateLIFF(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		LIFFID       string
		LIFFApp      LIFFApp
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			LIFFID: "testliffId",
			LIFFApp: LIFFApp{
				View: View{
					Type: LIFFViewTypeFull,
					URL:  "https://example.com/myservice",
				},
			},
			ResponseCode: 200,
			Response:     []byte(``),
			Want: want{
				RequestBody: []byte(`{"view":{"type":"full","url":"https://example.com/myservice"}}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
	}
	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPut {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPut)
		}
		endpoint := fmt.Sprintf(APIEndpointUpdateLIFFApp, tc.LIFFID)
		if r.URL.Path != endpoint {
			t.Errorf("URLPath %s; want %s", r.URL.Path, endpoint)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody\n %s; want\n %s", body, tc.Want.RequestBody)
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
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res, err := client.UpdateLIFF(tc.LIFFID, tc.LIFFApp).Do()
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %v; want %v", res, tc.Want.Response)
			}
		})
	}
}

func TestDeleteLIFF(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		LIFFID       string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			LIFFID:       "testliffId",
			ResponseCode: 200,
			Response:     []byte(``),
			Want: want{
				RequestBody: []byte(``),
				Response:    &BasicResponse{},
			},
		},
	}
	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodDelete {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPut)
		}
		endpoint := fmt.Sprintf(APIEndpointDeleteLIFFApp, tc.LIFFID)
		if r.URL.Path != endpoint {
			t.Errorf("URLPath %s; want %s", r.URL.Path, endpoint)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody\n %s; want\n %s", body, tc.Want.RequestBody)
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
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res, err := client.DeleteLIFF(tc.LIFFID).Do()
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %v; want %v", res, tc.Want.Response)
			}
		})
	}
}
