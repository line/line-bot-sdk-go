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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestIssueLinkToken(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *LinkTokenResponse
		Error       error
	}
	var testCases = []struct {
		UserID       string
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			UserID:       "u206d25c2ea6bd87c17655609a1c37cb8",
			ResponseCode: 200,
			Response:     []byte(`{"linkToken":"NMZTNuVrPTqlr2IF8Bnymkb7rXfYv5EY"}`),
			Want: want{
				RequestBody: []byte(""),
				Response:    &LinkTokenResponse{LinkToken: "NMZTNuVrPTqlr2IF8Bnymkb7rXfYv5EY"},
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
		endpoint := fmt.Sprintf(APIEndpointLinkToken, tc.UserID)
		if r.URL.Path != endpoint {
			t.Errorf("URLPath %s; want %s", r.URL.Path, endpoint)
		}
		body, err := ioutil.ReadAll(r.Body)
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
	client, err := mockClient(server)
	if err != nil {
		t.Fatal(err)
	}
	for i, tc := range testCases {
		currentTestIdx = i
		res, err := client.IssueLinkToken(tc.UserID).Do()
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
