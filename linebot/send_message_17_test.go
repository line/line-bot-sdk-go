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

// +build go1.7

package linebot

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPushMessagesWith17Context(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.PushMessage("U0cc15697597f61dd8b01cea8b027050e", NewTextMessage("Hello, world")).WithContext(ctx).Do()
	if err != context.DeadlineExceeded {
		t.Errorf("err %v; want %v", err, context.DeadlineExceeded)
	}
}
