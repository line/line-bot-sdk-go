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

// TestGetRichMenu tests GetRichMenu, GetUserRichMenu
func TestGetRichMenu(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *RichMenuResponse
		Error       error
	}
	var testCases = []struct {
		UserID       string
		RichMenuID   string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			RichMenuID:   "123456",
			ResponseCode: 200,
			Response:     []byte(`{"richMenuId":"123456","size":{"width":2500,"height":1686},"selected":false,"areas":[{"bounds":{"x":0,"y":0,"width":2500,"height":1686},"action":{"type":"postback","data":"action=buy&itemid=123"}}]}`),
			Want: want{
				URLPath: fmt.Sprintf(APIEndpointGetRichMenu, "123456"),
				Response: &RichMenuResponse{
					RichMenuID:  "123456",
					Size:        RichMenuSize{Width: 2500, Height: 1686},
					Selected:    false,
					ChatBarText: "",
					Areas: []AreaDetail{
						AreaDetail{
							Bounds: RichMenuBounds{X: 0, Y: 0, Width: 2500, Height: 1686},
							Action: RichMenuAction{Type: RichMenuActionTypePostback, Data: "action=buy&itemid=123"},
						},
					},
				},
			},
		},
		{
			RichMenuID:   "654321",
			UserID:       "user1",
			ResponseCode: 200,
			Response:     []byte(`{"richMenuId":"654321","size":{"width":2500,"height":1686},"selected":false,"areas":[{"bounds":{"x":0,"y":0,"width":2500,"height":1686},"action":{"type":"postback","data":"action=buy&itemid=123"}}]}`),
			Want: want{
				URLPath: fmt.Sprintf(APIEndpointGetUserRichMenu, "user1"),
				Response: &RichMenuResponse{
					RichMenuID:  "654321",
					Size:        RichMenuSize{Width: 2500, Height: 1686},
					Selected:    false,
					ChatBarText: "",
					Areas: []AreaDetail{
						AreaDetail{
							Bounds: RichMenuBounds{X: 0, Y: 0, Width: 2500, Height: 1686},
							Action: RichMenuAction{Type: RichMenuActionTypePostback, Data: "action=buy&itemid=123"},
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
	client, err := mockClient(server)
	if err != nil {
		t.Fatal(err)
	}
	var res *RichMenuResponse
	for i, tc := range testCases {
		currentTestIdx = i
		if tc.UserID != "" { // test get user
			res, err = client.GetUserRichMenu(tc.UserID).Do()
		} else {
			res, err = client.GetRichMenu(tc.RichMenuID).Do()
		}
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
				t.Errorf("Response %d %v; want %v", i, res, tc.Want.Response)
			}
		}
	}
}

func TestCreateRichMenu(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *RichMenuIDResponse
		Error       error
	}
	var testCases = []struct {
		Request      RichMenu
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			Request: RichMenu{
				Size:        RichMenuSize{Width: 2500, Height: 1686},
				Selected:    false,
				Name:        "Menu1",
				ChatBarText: "ChatText",
				Areas: []AreaDetail{
					AreaDetail{
						Bounds: RichMenuBounds{X: 0, Y: 0, Width: 2500, Height: 1686},
						Action: RichMenuAction{Type: RichMenuActionTypePostback, Data: "action=buy&itemid=123"},
					},
				},
			},
			ResponseCode: 200,
			Response:     []byte(`{"richMenuId":"abcefg"}`),
			Want: want{
				RequestBody: []byte(`{"size":{"width":2500,"height":1686},"selected":false,"name":"Menu1","chatBarText":"ChatText","areas":[{"bounds":{"x":0,"y":0,"width":2500,"height":1686},"action":{"type":"postback","data":"action=buy\u0026itemid=123"}}]}` + "\n"),
				Response:    &RichMenuIDResponse{RichMenuID: "abcefg"},
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
		if r.URL.Path != APIEndpointCreateRichMenu {
			t.Errorf("URLPath %s; want %s", r.URL.Path, APIEndpointCreateRichMenu)
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
		res, err := client.CreateRichMenu(tc.Request).Do()
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

// TestLinkRichMenu tests LinkUserRichMenu
func TestLinkRichMenu(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	var testCases = []struct {
		RichMenuID   string
		UserID       string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			RichMenuID:   "654321",
			UserID:       "userId1",
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				URLPath:  fmt.Sprintf(APIEndpointLinkUserRichMenu, "userId1", "654321"),
				Response: &BasicResponse{},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPost {
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
	client, err := mockClient(server)
	if err != nil {
		t.Fatal(err)
	}
	var res *BasicResponse
	for i, tc := range testCases {
		currentTestIdx = i
		res, err = client.LinkUserRichMenu(tc.UserID, tc.RichMenuID).Do()
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

// TestListRichMenu
func TestListRichMenu(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    []*RichMenuResponse
		Error       error
	}
	var testCases = []struct {
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			ResponseCode: 200,
			Response: []byte(`{"richmenus":[` +
				`{"richMenuId":"123","size":{"width":2500,"height":1686},"selected":false,"areas":[{"bounds":{"x":0,"y":0,"width":2500,"height":1686},"action":{"type":"postback","data":"action=buy&itemid=123"}}]},` +
				`{"richMenuId":"456","size":{"width":2500,"height":1686},"selected":false,"chatBarText":"hello","areas":[{"bounds":{"x":0,"y":0,"width":2500,"height":1686},"action":{"type":"message","text":"text"}}]},` +
				`{"richMenuId":"789","size":{"width":2500,"height":1686},"selected":false,"chatBarText":"line.me","areas":[{"bounds":{"x":0,"y":0,"width":2500,"height":1686},"action":{"type":"uri","uri":"http://line.me/"}}]}` +
				`]}`),
			Want: want{
				URLPath: APIEndpointListRichMenu,
				Response: []*RichMenuResponse{
					&RichMenuResponse{
						RichMenuID:  "123",
						Size:        RichMenuSize{Width: 2500, Height: 1686},
						Selected:    false,
						ChatBarText: "",
						Areas: []AreaDetail{
							AreaDetail{
								Bounds: RichMenuBounds{X: 0, Y: 0, Width: 2500, Height: 1686},
								Action: RichMenuAction{Type: RichMenuActionTypePostback, Data: "action=buy&itemid=123"},
							},
						},
					},
					&RichMenuResponse{
						RichMenuID:  "456",
						Size:        RichMenuSize{Width: 2500, Height: 1686},
						Selected:    false,
						ChatBarText: "hello",
						Areas: []AreaDetail{
							AreaDetail{
								Bounds: RichMenuBounds{X: 0, Y: 0, Width: 2500, Height: 1686},
								Action: RichMenuAction{Type: RichMenuActionTypeMessage, Text: "text"},
							},
						},
					},
					&RichMenuResponse{
						RichMenuID:  "789",
						Size:        RichMenuSize{Width: 2500, Height: 1686},
						Selected:    false,
						ChatBarText: "line.me",
						Areas: []AreaDetail{
							AreaDetail{
								Bounds: RichMenuBounds{X: 0, Y: 0, Width: 2500, Height: 1686},
								Action: RichMenuAction{Type: RichMenuActionTypeURI, URI: "http://line.me/"},
							},
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
	client, err := mockClient(server)
	if err != nil {
		t.Fatal(err)
	}
	for i, tc := range testCases {
		currentTestIdx = i
		res, err := client.GetRichMenuList().Do()
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
				t.Errorf("Response %d\n %v; want\n %v", i, res, tc.Want.Response)
			}
		}
	}
}
