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
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"strconv"
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
		Label        string
		UserID       string
		RichMenuID   string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			Label:        "Without UserID",
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
						{
							Bounds: RichMenuBounds{X: 0, Y: 0, Width: 2500, Height: 1686},
							Action: RichMenuAction{Type: RichMenuActionTypePostback, Data: "action=buy&itemid=123"},
						},
					},
				},
			},
		},
		{
			Label:        "With UserID",
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
						{
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

	var res *RichMenuResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			if tc.UserID != "" { // test get user
				res, err = client.GetUserRichMenu(tc.UserID).Do()
			} else {
				res, err = client.GetRichMenu(tc.RichMenuID).Do()
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
					{
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
			res, err := client.CreateRichMenu(tc.Request).Do()
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
	var res *BasicResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res, err = client.LinkUserRichMenu(tc.UserID, tc.RichMenuID).Do()
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

// Test method for GetDefaultRichMenu
func (call *GetDefaultRichMenuCall) Test() {
}

// Test method for CancelDefaultRichMenu
func (call *CancelDefaultRichMenuCall) Test() {
}

// Test method for SetDefaultRichMenu
func (call *SetDefaultRichMenuCall) Test() {
}

// TestDefaultRichMenu tests SetDefaultRichMenu, CancelDefaultRichMenu, GetDefaultRichMenu
func TestDefaultRichMenu(t *testing.T) {
	type testMethod interface {
		Test()
	}
	type want struct {
		URLPath     string
		HTTPMethod  string
		RequestBody []byte
		Response    interface{}
		Error       error
	}
	var testCases = []struct {
		TestMethod   testMethod
		RichMenuID   string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			RichMenuID:   "654321",
			TestMethod:   new(SetDefaultRichMenuCall),
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				HTTPMethod: http.MethodPost,
				URLPath:    fmt.Sprintf(APIEndpointSetDefaultRichMenu, "654321"),
				Response:   &BasicResponse{},
			},
		},
		{
			RichMenuID:   "N/A",
			TestMethod:   new(GetDefaultRichMenuCall),
			ResponseCode: 200,
			Response:     []byte(`{"richMenuId": "654321"}`),
			Want: want{
				HTTPMethod: http.MethodGet,
				URLPath:    APIEndpointDefaultRichMenu,
				Response:   &RichMenuIDResponse{RichMenuID: "654321"},
			},
		},
		{
			RichMenuID:   "N/A",
			TestMethod:   new(CancelDefaultRichMenuCall),
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				HTTPMethod: http.MethodDelete,
				URLPath:    APIEndpointDefaultRichMenu,
				Response:   &BasicResponse{},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != tc.Want.HTTPMethod {
			t.Errorf("Method %s; want %s", r.Method, tc.Want.HTTPMethod)
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
		t.Run(strconv.Itoa(i)+"/"+tc.RichMenuID+"."+string(tc.Response), func(t *testing.T) {
			switch tc.TestMethod.(type) {
			case *SetDefaultRichMenuCall:
				res, err = client.SetDefaultRichMenu(tc.RichMenuID).Do()
			case *CancelDefaultRichMenuCall:
				res, err = client.CancelDefaultRichMenu().Do()
			case *GetDefaultRichMenuCall:
				res, err = client.GetDefaultRichMenu().Do()
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
				`{"richMenuId":"789","size":{"width":2500,"height":1686},"selected":false,"chatBarText":"line.me","areas":[{"bounds":{"x":0,"y":0,"width":2500,"height":1686},"action":{"type":"uri","uri":"https://line.me/"}}]}` +
				`]}`),
			Want: want{
				URLPath: APIEndpointListRichMenu,
				Response: []*RichMenuResponse{
					{
						RichMenuID:  "123",
						Size:        RichMenuSize{Width: 2500, Height: 1686},
						Selected:    false,
						ChatBarText: "",
						Areas: []AreaDetail{
							{
								Bounds: RichMenuBounds{X: 0, Y: 0, Width: 2500, Height: 1686},
								Action: RichMenuAction{Type: RichMenuActionTypePostback, Data: "action=buy&itemid=123"},
							},
						},
					},
					{
						RichMenuID:  "456",
						Size:        RichMenuSize{Width: 2500, Height: 1686},
						Selected:    false,
						ChatBarText: "hello",
						Areas: []AreaDetail{
							{
								Bounds: RichMenuBounds{X: 0, Y: 0, Width: 2500, Height: 1686},
								Action: RichMenuAction{Type: RichMenuActionTypeMessage, Text: "text"},
							},
						},
					},
					{
						RichMenuID:  "789",
						Size:        RichMenuSize{Width: 2500, Height: 1686},
						Selected:    false,
						ChatBarText: "line.me",
						Areas: []AreaDetail{
							{
								Bounds: RichMenuBounds{X: 0, Y: 0, Width: 2500, Height: 1686},
								Action: RichMenuAction{Type: RichMenuActionTypeURI, URI: "https://line.me/"},
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
			res, err := client.GetRichMenuList().Do()
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
					t.Errorf("Response\n %v; want\n %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestBulkRichMenu tests BulkLinkRichMenu, BulkUnlinkRichMenu
func TestBulkRichMenu(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	var testCases = []struct {
		UserIDs      []string
		RichMenuID   string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			ResponseCode: 202,
			Response:     []byte(`{}`),
			UserIDs:      []string{"userId1", "userId2"},
			RichMenuID:   "richMenuId",
			Want: want{
				RequestBody: []byte(`{"richMenuId":"richMenuId","userIds":["userId1","userId2"]}` + "\n"),
				URLPath:     APIEndpointBulkLinkRichMenu,
				Response:    &BasicResponse{},
			},
		},
		{
			ResponseCode: 202,
			Response:     []byte(`{}`),
			UserIDs:      []string{"userId1", "userId2"},
			RichMenuID:   "", // bulk unlink has no richmenuid
			Want: want{
				RequestBody: []byte(`{"userIds":["userId1","userId2"]}` + "\n"),
				URLPath:     APIEndpointBulkUnlinkRichMenu,
				Response:    &BasicResponse{},
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
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
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
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if tc.RichMenuID == "" { // unlink
				res, err = client.BulkUnlinkRichMenu(tc.UserIDs...).Do()
			} else { // bulk link
				res, err = client.BulkLinkRichMenu(tc.RichMenuID, tc.UserIDs...).Do()
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
					t.Errorf("Response\n %v; want\n %v", res, tc.Want.Response)
				}
			}
		})
	}
}

func TestUploadRichMenuImage(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	var testCases = []struct {
		RichMenuID   string
		ImagePath    string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			RichMenuID:   "123456",
			ImagePath:    filepath.Join("..", "testdata", "img", "richmenu.png"),
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				URLPath:  fmt.Sprintf(APIEndpointUploadRichMenuImage, "123456"),
				Response: &BasicResponse{},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer server.Close()

	var currentTestIdx int
	dataServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		imgBody, err := ioutil.ReadFile(tc.ImagePath)
		if err != nil {
			t.Fatal(err)
		}
		wantImgSize := len(imgBody)
		if len(body) != wantImgSize {
			t.Errorf("ContentLength %d; want %d", len(body), wantImgSize)
		}
		if r.ContentLength != int64(wantImgSize) {
			t.Errorf("ContentLength %d; want %d", r.ContentLength, wantImgSize)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer dataServer.Close()
	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res, err := client.UploadRichMenuImage(tc.RichMenuID, tc.ImagePath).Do()
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
					t.Errorf("Response\n %v; want\n %v", res, tc.Want.Response)
				}
			}
		})
	}
}

func TestDownloadRichMenuImage(t *testing.T) {
	filePath := filepath.Join("..", "testdata", "img", "richmenu.png")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}

	type want struct {
		URLPath  string
		Response *MessageContentResponse
		Error    error
	}
	var testCases = []struct {
		RichMenuID     string
		ImagePath      string
		Response       []byte
		ResponseCode   int
		ResponseHeader map[string]string
		Want           want
	}{
		{
			RichMenuID:   "123456",
			Response:     file,
			ResponseCode: 200,
			ResponseHeader: map[string]string{
				"Content-Type":   "image/png",
				"Content-Length": strconv.Itoa(len(file)),
			},
			Want: want{
				URLPath: fmt.Sprintf(APIEndpointUploadRichMenuImage, "123456"),
				Response: &MessageContentResponse{
					ContentType:   "image/png",
					ContentLength: int64(len(file)),
					Content:       ioutil.NopCloser(bytes.NewReader(file)),
				},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer server.Close()

	var currentTestIdx int
	dataServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodGet {
			t.Errorf("Method %s; want %s", r.Method, http.MethodGet)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		if err != nil {
			t.Fatal(err)
		}
		for k, v := range tc.ResponseHeader {
			w.Header().Add(k, v)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res, err := client.DownloadRichMenuImage(tc.RichMenuID).Do()
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
				body := res.Content
				defer body.Close()
				resImage, err := ioutil.ReadAll(body)
				if err != nil {
					t.Fatal(err)
				}
				if !bytes.Equal(resImage, tc.Response) {
					t.Error("Expected image content is not returned")
				}
				res.Content = nil
				tc.Want.Response.Content = nil
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response\n %+v; want\n %+v", res, tc.Want.Response)
				}
			}
		})
	}
}
