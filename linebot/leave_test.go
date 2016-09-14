package linebot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func TestLeaveGroup(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	var testCases = []struct {
		GroupID      string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			GroupID:      "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointLeaveGroup, "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				RequestBody: []byte(""),
				Response:    &BasicResponse{},
			},
		},
		{
			// Too Many Requests
			GroupID:      "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ResponseCode: 429,
			Response:     []byte(`{"message":"Too Many Requests"}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointLeaveGroup, "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 429,
					Response: &ErrorResponse{
						Message: "Too Many Requests",
					},
				},
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
			t.Error(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
	}
	for i, tc := range testCases {
		currentTestIdx = i
		res, err := client.LeaveGroup(tc.GroupID).Do()
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

func TestLeaveGroupWithContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.LeaveGroup("cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").WithContext(ctx).Do()
	if err != context.DeadlineExceeded {
		t.Errorf("err %v; want %v", err, context.DeadlineExceeded)
	}
}

func TestLeaveRoom(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	var testCases = []struct {
		RoomID       string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			RoomID:       "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointLeaveRoom, "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				RequestBody: []byte(""),
				Response:    &BasicResponse{},
			},
		},
		{
			// Too Many Requests
			RoomID:       "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ResponseCode: 429,
			Response:     []byte(`{"message":"Too Many Requests"}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIEndpointLeaveRoom, "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 429,
					Response: &ErrorResponse{
						Message: "Too Many Requests",
					},
				},
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
			t.Error(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
	}
	for i, tc := range testCases {
		currentTestIdx = i
		res, err := client.LeaveRoom(tc.RoomID).Do()
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

func TestTestLeaveRoomWithContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.LeaveRoom("cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").WithContext(ctx).Do()
	if err != context.DeadlineExceeded {
		t.Errorf("err %v; want %v", err, context.DeadlineExceeded)
	}
}

func BenchmarkLeaveGroup(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte("{}"))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.LeaveGroup("cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").Do()
	}
}

func BenchmarkLeaveRoom(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte("{}"))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.LeaveRoom("cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").Do()
	}
}
