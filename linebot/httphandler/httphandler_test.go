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

package httphandler

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var testRequestBody = `{
    "events": [
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "message",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "message": {
                "id": "325708",
                "type": "text",
                "text": "Hello, world"
            }
        }
    ]
}
`

const (
	testChannelSecret = "testsecret"
	testChannelToken  = "testtoken"
)

func TestWebhookHandler(t *testing.T) {
	handler, err := New(testChannelSecret, testChannelToken)
	if err != nil {
		t.Error(err)
	}
	handlerFunc := func(events []*linebot.Event, r *http.Request) {
		if events == nil {
			t.Errorf("events is nil")
		}
		if r == nil {
			t.Errorf("r is nil")
		}
		bot, err := handler.NewClient()
		if err != nil {
			t.Fatal(err)
		}
		if bot == nil {
			t.Errorf("bot is nil")
		}
	}
	handler.HandleEvents(handlerFunc)

	server := httptest.NewTLSServer(handler)
	defer server.Close()
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// valid signature
	{
		body := []byte(testRequestBody)
		req, err := http.NewRequest("POST", server.URL, bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		// generate signature
		mac := hmac.New(sha256.New, []byte(testChannelSecret))
		mac.Write(body)

		req.Header.Set("X-Line-Signature", base64.StdEncoding.EncodeToString(mac.Sum(nil)))
		res, err := httpClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if res == nil {
			t.Error("response is nil")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("status: %d", res.StatusCode)
		}
	}

	// invalid signature
	handler.HandleError(func(err error, r *http.Request) {
		if err != linebot.ErrInvalidSignature {
			t.Errorf("err %v; want %v", err, linebot.ErrInvalidSignature)
		}
		if r == nil {
			t.Errorf("r is nil")
		}
	})
	{
		body := []byte(testRequestBody)
		req, err := http.NewRequest("POST", server.URL, bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("X-LINE-Signature", "invalidSignature")
		res, err := httpClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if res == nil {
			t.Error("response is nil")
		}
		if res.StatusCode != 400 {
			t.Errorf("status: %d", 400)
		}
	}
}
