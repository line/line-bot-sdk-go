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

	"golang.org/x/net/context"

	"github.com/line/line-bot-sdk-go/linebot"
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
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "message",
            "timestamp": 1462629479859,
            "source": {
                "type": "group",
                "groupId": "u206d25c2ea6bd87c17655609a1c37cb8",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "message": {
                "id": "325708",
                "type": "text",
                "text": "Hello, world"
            }
        },
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
                "type": "image"
            }
        },
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
                "type": "video"
            }
        },
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
                "type": "audio",
                "duration": 1111
            }
        },
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
                "type": "location",
                "title": "hello",
                "address": "〒150-0002 東京都渋谷区渋谷２丁目２１−１",
                "latitude": 35.65910807942215,
                "longitude": 139.70372892916203
            }
        },
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
                "type": "sticker",
                "packageId": "1",
                "stickerId": "1"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "follow",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            }
        },
        {
            "type": "unfollow",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "join",
            "timestamp": 1462629479859,
            "source": {
                "type": "group",
                "groupId": "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
            }
        },
        {
            "type": "leave",
            "timestamp": 1462629479859,
            "source": {
                "type": "group",
                "groupId": "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "postback",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "postback": {
                "data": "action=buyItem&itemId=123123&color=red"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "beacon",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "U012345678901234567890123456789ab"
            },
            "beacon": {
                "hwid":"374591320",
                "type":"enter"
            }
        }
    ]
}
`

var testCases = []struct {
	Type linebot.EventType
}{
	{linebot.EventTypeMessage},
	{linebot.EventTypeMessage},
	{linebot.EventTypeMessage},
	{linebot.EventTypeMessage},
	{linebot.EventTypeMessage},
	{linebot.EventTypeMessage},
	{linebot.EventTypeMessage},
	{linebot.EventTypeFollow},
	{linebot.EventTypeUnfollow},
	{linebot.EventTypeJoin},
	{linebot.EventTypeLeave},
	{linebot.EventTypePostback},
	{linebot.EventTypeBeacon},
}

const (
	testChannelSecret = "testsecret"
	testChannelToken  = "testtoken"
)

func TestEventHandler(t *testing.T) {
	var currentTestIdx int

	handler, err := New(testChannelSecret, testChannelToken)
	if err != nil {
		t.Error(err)
	}

	eventHandlerFunc := func(ctx context.Context, c *linebot.Client, e *linebot.Event) {
		if ctx == nil {
			t.Errorf("eventHandlerFunc: ctx is nil")
		}
		if c == nil {
			t.Errorf("eventHandlerFunc: c is nil")
		}
		if e == nil {
			t.Errorf("eventHandlerFunc: e is nil")
		}
		testCase := testCases[currentTestIdx]
		if e.Type != testCase.Type {
			t.Errorf("eventHandlerFunc: e.Type %s; want %s", e.Type, testCase.Type)
		}
		currentTestIdx++
	}

	handler.HandleMessage(eventHandlerFunc)
	handler.HandleFollow(eventHandlerFunc)
	handler.HandleUnfollow(eventHandlerFunc)
	handler.HandleJoin(eventHandlerFunc)
	handler.HandleLeave(eventHandlerFunc)
	handler.HandlePostback(eventHandlerFunc)
	handler.HandleBeacon(eventHandlerFunc)

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
	if currentTestIdx != len(testCases) {
		t.Errorf("currentTestIdx %d; want %d", currentTestIdx, len(testCases))
	}

	// invalid signature
	handler.HandleError(func(ctx context.Context, err error) {
		if ctx == nil {
			t.Errorf("ctx is nil")
		}
		if err != linebot.ErrInvalidSignature {
			t.Errorf("err %q; want %q", err, linebot.ErrInvalidSignature)
		}
	})
	{
		body := []byte(testRequestBody)
		req, err := http.NewRequest("POST", server.URL, bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("X-LINE-Signature", "invalidsignatue")
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

func TestEventHandlerWithOptions(t *testing.T) {
	handler, err := New(
		testChannelSecret, testChannelToken,
		WithNewClientFunc(func(ctx context.Context, channelSecret, channelToken string) (*linebot.Client, error) {
			return linebot.New(channelSecret, channelToken)
		}),
		WithNewContextFunc(func(req *http.Request) (context.Context, error) {
			return context.Background(), nil
		}),
	)
	if err != nil {
		t.Error(err)
	}

	server := httptest.NewTLSServer(handler)
	defer server.Close()
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
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
}
