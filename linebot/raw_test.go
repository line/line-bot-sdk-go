// Copyright 2021 LINE Corporation
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
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

type requestWant struct {
	Path        string
	Method      string
	ContentType string
	RequestBody []byte
}

type responseWant struct {
	Status       int
	ResponseBody []byte
}
type testcase struct {
	ResponseCode int
	Response     []byte
	ResponseWant responseWant
	RequestWant  requestWant
	Call         func(client *Client) (*http.Response, error)
}

func TestRaw(t *testing.T) {
	testcases := []testcase{
		newTestCase(
			http.MethodGet,
			"",
			[]byte(""),
			func(client *Client) (*http.Response, error) {
				call, err := client.NewRawCall(http.MethodGet, "/abcdefg")
				if err != nil {
					panic(err)
				}
				return call.Do()
			},
		),
		newTestCase(
			http.MethodPost,
			"application/json; charset=UTF-8",
			[]byte(`RRRREQUEST_BODY`),
			func(client *Client) (*http.Response, error) {
				call, err := client.NewRawCallWithBody(http.MethodPost, "/abcdefg", bytes.NewReader([]byte(`RRRREQUEST_BODY`)))
				if err != nil {
					panic(err)
				}
				call.AddHeader("content-type", "application/json; charset=UTF-8")
				return call.Do()
			},
		),
		newTestCase(
			http.MethodPost,
			"application/x-www-form-urlencoded",
			[]byte(`RRRREQUEST_BODY`),
			func(client *Client) (*http.Response, error) {
				call, err := client.NewRawCallWithBody(http.MethodPost, "/abcdefg", bytes.NewReader([]byte(`RRRREQUEST_BODY`)))
				if err != nil {
					panic(err)
				}
				call.AddHeader("content-type", "application/x-www-form-urlencoded")
				return call.Do()
			},
		),
		newTestCase(
			http.MethodPut,
			"application/json; charset=UTF-8",
			[]byte(`RRRREQUEST_BODY`),
			func(client *Client) (*http.Response, error) {
				call, err := client.NewRawCallWithBody(http.MethodPut, "/abcdefg", bytes.NewReader([]byte(`RRRREQUEST_BODY`)))
				if err != nil {
					panic(err)
				}
				call.AddHeader("content-type", "application/json; charset=UTF-8")
				return call.Do()
			},
		),
		newTestCase(
			http.MethodDelete,
			"",
			[]byte(``),
			func(client *Client) (*http.Response, error) {
				call, err := client.NewRawCall(http.MethodDelete, "/abcdefg")
				if err != nil {
					panic(err)
				}
				return call.Do()
			},
		),
	}

	for i, tc := range testcases {
		t.Run(strconv.Itoa(i)+"_"+tc.RequestWant.Method, func(t *testing.T) {
			runTestCase(t, tc)
		})
	}
}

func newTestCase(expectedRequestMethod string,
	expectedRequestContentType string,
	expectedRequestBody []byte,
	call func(client *Client) (*http.Response, error)) testcase {
	return testcase{
		ResponseCode: 200,
		Response:     []byte(`TESTDATA`),
		Call:         call,
		RequestWant: requestWant{
			Method:      expectedRequestMethod,
			Path:        "/abcdefg",
			RequestBody: expectedRequestBody,
			ContentType: expectedRequestContentType,
		},
		ResponseWant: responseWant{
			Status:       200,
			ResponseBody: []byte(`TESTDATA`),
		},
	}
}

func runTestCase(t *testing.T, tc testcase) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != tc.RequestWant.Method {
			t.Errorf("Method %s; want %s", r.Method, tc.RequestWant.Method)
		}
		if r.URL.Path != tc.RequestWant.Path {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.RequestWant.Path)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		contentType := r.Header.Get("content-type")
		if !reflect.DeepEqual(contentType, tc.RequestWant.ContentType) {
			t.Errorf("Content\n %v; want\n %v", contentType, tc.RequestWant.ContentType)
		}

		if tc.RequestWant.RequestBody != nil {
			if !reflect.DeepEqual(body, tc.RequestWant.RequestBody) {
				t.Errorf("RequestBody\n %s; want\n %s", body, tc.RequestWant.RequestBody)
			}
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
	res, err := tc.Call(client)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(body, tc.ResponseWant.ResponseBody) {
		t.Errorf("Response %v; want %v", body, tc.ResponseWant.ResponseBody)
	}
	if res.StatusCode != tc.ResponseWant.Status {
		t.Errorf("Response %v; want %v", res.StatusCode, tc.ResponseWant.Status)
	}
}
