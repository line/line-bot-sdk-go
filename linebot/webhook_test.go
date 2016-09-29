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
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

var webhookTestRequestBody = `{
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

var webhookTestWantEvents = []Event{
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeMessage,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &TextMessage{
			ID:   "325708",
			Text: "Hello, world",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeMessage,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:    EventSourceTypeGroup,
			UserID:  "u206d25c2ea6bd87c17655609a1c37cb8",
			GroupID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &TextMessage{
			ID:   "325708",
			Text: "Hello, world",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeMessage,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &ImageMessage{
			ID: "325708",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeMessage,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &LocationMessage{
			ID:        "325708",
			Title:     "hello",
			Address:   "〒150-0002 東京都渋谷区渋谷２丁目２１−１",
			Latitude:  35.65910807942215,
			Longitude: 139.70372892916203,
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeMessage,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &StickerMessage{
			ID:        "325708",
			PackageID: "1",
			StickerID: "1",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeFollow,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
	},
	{
		Type:      EventTypeUnfollow,
		Timestamp: time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeJoin,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:    EventSourceTypeGroup,
			GroupID: "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		},
	},
	{
		Type:      EventTypeLeave,
		Timestamp: time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:    EventSourceTypeGroup,
			GroupID: "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypePostback,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Postback: &Postback{
			Data: "action=buyItem&itemId=123123&color=red",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeBeacon,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "U012345678901234567890123456789ab",
		},
		Beacon: &Beacon{
			Hwid: "374591320",
			Type: BeaconEventTypeEnter,
		},
	},
}

func TestParseRequest(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client, err := New("testsecret", "testtoken")
		if err != nil {
			t.Error(err)
		}
		gotEvents, err := client.ParseRequest(r)
		if err != nil {
			if err == ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
				t.Error(err)
			}
			return
		}
		if len(gotEvents) != len(webhookTestWantEvents) {
			t.Errorf("Event length %d; want %d", len(gotEvents), len(webhookTestWantEvents))
		}
		for i, got := range gotEvents {
			want := webhookTestWantEvents[i]
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Event %d %q; want %q", i, got, want)
			}
		}
	}))
	defer server.Close()
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// invalid signature
	{
		body := []byte(webhookTestRequestBody)
		req, err := http.NewRequest("POST", server.URL, bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("X-LINE-Signature", "invalidsignatue")
		res, err := httpClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != 400 {
			t.Errorf("StatusCode %d; want %d", res.StatusCode, 400)
		}
	}

	// valid signature
	{
		body := []byte(webhookTestRequestBody)
		req, err := http.NewRequest("POST", server.URL, bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		// generate signature
		mac := hmac.New(sha256.New, []byte("testsecret"))
		mac.Write(body)

		req.Header.Set("X-LINE-Signature", base64.StdEncoding.EncodeToString(mac.Sum(nil)))
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

func BenchmarkParseRequest(b *testing.B) {
	body := []byte(webhookTestRequestBody)
	client, err := New("testsecret", "testtoken")
	if err != nil {
		b.Fatal(err)
	}
	mac := hmac.New(sha256.New, []byte("testsecret"))
	mac.Write(body)
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", "", bytes.NewReader(body))
		req.Header.Set("X-LINE-Signature", sign)
		client.ParseRequest(req)
	}
}
