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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

// TestGetNumberMessages tests GetNumberReplyMessages, GetNumberPushMessages, GetNumberMulticastMessages
// and GetNumberBroadcastMessages func
func TestGetNumberMessages(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *MessagesNumberResponse
		Error       error
	}
	testCases := []struct {
		TestType     DeliveryType
		Date         string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			Date:         "20190403",
			TestType:     DeliveryTypeReply,
			ResponseCode: 200,
			Response:     []byte(`{"status":"ready","success":123}`),
			Want: want{
				URLPath: fmt.Sprintf(APIEndpointGetMessageDelivery, DeliveryTypeReply),
				Response: &MessagesNumberResponse{
					Status:  "ready",
					Success: 123,
				},
			},
		},
		{
			Date:         "20180330",
			TestType:     DeliveryTypeReply,
			ResponseCode: 200,
			Response:     []byte(`{"status":"out_of_service"}`),
			Want: want{
				URLPath: fmt.Sprintf(APIEndpointGetMessageDelivery, DeliveryTypeReply),
				Response: &MessagesNumberResponse{
					Status: "out_of_service",
				},
			},
		},
		{
			Date:         "20290403",
			TestType:     DeliveryTypePush,
			ResponseCode: 200,
			Response:     []byte(`{"status":"unready"}`),
			Want: want{
				URLPath: fmt.Sprintf(APIEndpointGetMessageDelivery, DeliveryTypePush),
				Response: &MessagesNumberResponse{
					Status: "unready",
				},
			},
		},
		{
			Date:         "20190401",
			TestType:     DeliveryTypeMulticast,
			ResponseCode: 200,
			Response:     []byte(`{"status":"ready","success":456}`),
			Want: want{
				URLPath: fmt.Sprintf(APIEndpointGetMessageDelivery, DeliveryTypeMulticast),
				Response: &MessagesNumberResponse{
					Status:  "ready",
					Success: 456,
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
		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
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
	var res interface{}
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+string(tc.TestType)+"."+tc.Date, func(t *testing.T) {
			switch tc.TestType {
			case DeliveryTypeMulticast:
				res, err = client.GetNumberMulticastMessages(tc.Date).Do()
			case DeliveryTypePush:
				res, err = client.GetNumberPushMessages(tc.Date).Do()
			case DeliveryTypeReply:
				res, err = client.GetNumberReplyMessages(tc.Date).Do()
			case DeliveryTypeBroadcast:
				res, err = client.GetNumberBroadcastMessages(tc.Date).Do()
			}
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
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}
