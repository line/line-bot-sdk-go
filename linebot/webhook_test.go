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
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"
)

var webhookTestRequestBody = `{
    "events": [
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "message",
            "mode": "active",
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
            "mode": "active",
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
            "mode": "active",
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
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "message": {
                "id": "325708",
                "type": "file",
                "fileName": "file.txt",
                "fileSize": 2138
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "message",
            "mode": "active",
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
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "message": {
                "id": "325708",
                "type": "sticker",
                "packageId": "1",
                "stickerId": "1",
                "stickerResourceType": "STATIC"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "message",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "message": {
                "id": "325708",
                "type": "sticker",
                "packageId": "1",
                "stickerId": "3",
                "stickerResourceType": "ANIMATION_SOUND"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "message",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "message": {
                "id": "3257088",
                "type": "sticker",
                "packageId": "20",
                "stickerId": "3",
                "stickerResourceType": "PER_STICKER_TEXT",
				"keywords": ["cony","sally","Staring","hi","whatsup","line","howdy","HEY","Peeking","wave","peek","Hello","yo","greetings"]
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "follow",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            }
        },
        {
            "type": "unfollow",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "join",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "group",
                "groupId": "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
            }
        },
        {
            "type": "leave",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "group",
                "groupId": "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "postback",
            "mode": "active",
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
            "type": "postback",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "postback": {
                "data": "action=sel&only=date",
				"params": {
					"date": "2017-09-03"
				}
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "postback",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "postback": {
                "data": "action=sel&only=time",
				"params": {
					"time": "15:38"
				}
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "postback",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "postback": {
                "data": "action=sel",
				"params": {
					"datetime": "2017-09-03T15:38"
				}
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "beacon",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "U012345678901234567890123456789ab"
            },
            "beacon": {
                "hwid":"374591320",
                "type":"enter"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "beacon",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "U012345678901234567890123456789ab"
            },
            "beacon": {
                "hwid":"374591320",
                "type":"enter",
                "dm":"1234567890abcdef"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "beacon",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "U012345678901234567890123456789ab"
            },
            "beacon": {
                "hwid":"374591320",
                "type":"stay",
                "dm":"1234567890abcdef"
            }
        },
        {
          "type": "accountLink",
          "mode": "active",
          "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
          "source": {
            "userId": "U012345678901234567890123456789ab",
            "type": "user"
          },
          "timestamp": 1462629479859,
          "link": {
            "result": "ok",
            "nonce": "xxxxxxxxxxxxxxx"
          }
        },
        {
          "replyToken": "0f3779fba3b349968c5d07db31eabf65",
          "type": "memberJoined",
          "mode": "active",
          "timestamp": 1462629479859,
          "source": {
            "type": "group",
            "groupId": "C4af498062901234567890123456789ab"
          },
          "joined": {
            "members": [
              {
                "type": "user",
                "userId": "U4af498062901234567890123456789ab"
              },
              {
                "type": "user",
                "userId": "U91eeaf62d901234567890123456789ab"
              }
            ]
          }
        },
        {
          "type": "memberLeft",
          "mode": "active",
          "timestamp": 1462629479960,
          "source": {
            "type": "group",
            "groupId": "C4af498062901234567890123456789ab"
          },
          "left": {
            "members": [
              {
                "type": "user",
                "userId": "U4af498062901234567890123456789ab"
              },
              {
                "type": "user",
                "userId": "U91eeaf62d901234567890123456789ab"
              }
            ]
		  }
        },
        {
          "type": "things",
          "mode": "active",
          "timestamp": 1462629479859,
          "source": {
            "type": "user",
            "userId": "U91eeaf62d901234567890123456789ab"
          },
          "things": {
            "deviceId": "t2c449c9d1...",
            "type": "link"
          }
        },
        {
          "type": "things",
          "mode": "active",
          "timestamp": 1462629479859,
          "source": {
            "type": "user",
            "userId": "U91eeaf62d901234567890123456789ab"
          },
          "things": {
            "deviceId": "t2c449c9d1...",
            "type": "unlink"
          }
        },
        {
          "type": "things",
          "mode": "active",
          "timestamp": 1462629479859,
          "source": {
            "type": "user",
            "userId": "U91eeaf62d901234567890123456789ab"
          },
          "things": {
            "deviceId": "t016b8dc6...",
            "type": "scenarioResult",
            "result": {
              "scenarioId": "01DE9CH7H...",
              "revision": 3,
              "startTime": 1563511217095,
              "endTime": 1563511217097,
              "resultCode": "success",
              "bleNotificationPayload": "AQ==",
              "actionResults": [
                {
                  "type": "binary",
                  "data": "/w=="
                }
              ]
            }
          }
        },
        {
          "type":"things",
          "mode": "active",
          "replyToken":"f026a377...",
          "source":{
            "userId":"U91eeaf62d901234567890123456789ab",
            "type":"user"
          },
          "timestamp":1563511218376,
          "things":{
            "deviceId":"t016b8d...",
            "result":{
              "scenarioId":"01DE9CH7H...",
              "revision":3,
              "startTime":1563511217095,
              "endTime":1563511217097,
              "resultCode":"gatt_error",
              "errorReason":"c.l.h.D: No characteristic is found for the given UUID.",
              "actionResults":[

              ]
            },
            "type":"scenarioResult"
          }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "message",
            "mode": "standby",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "message": {
                "id": "325708",
                "type": "text",
                "text": "Stand by me"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "message",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "user",
                "userId": "u206d25c2ea6bd87c17655609a1c37cb8"
            },
            "message": {
                "id": "325708",
                "type": "text",
                "text": "Hello, world! (love)",
                "emojis": [
                    {
                        "index": 14,
                        "length": 6,
                        "productId": "5ac1bfd5040ab15980c9b435",
                        "emojiId": "001"
                    }
                ]
            }
        },
        {
            "type": "unsend",
            "mode": "active",
            "timestamp": 1462629479859,
            "source": {
                "type": "group",
                "groupId": "Ca56f94637c...",
                "userId": "U4af4980629..."
            },
            "unsend": {
                "messageId": "325708"
            }
        },
        {
            "replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
            "type": "videoPlayComplete",
            "timestamp": 1462629479859,
            "mode": "active",
            "source": {
                "type": "user",
                "userId": "U4af4980629..."
            },
            "videoPlayComplete": {
                "trackingId": "track_id"
            }
        }
    ]
}
`

var webhookTestWantEvents = []*Event{
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeMessage,
		Mode:       EventModeActive,
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
		Mode:       EventModeActive,
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
		Mode:       EventModeActive,
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
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &FileMessage{
			ID:       "325708",
			FileName: "file.txt",
			FileSize: 2138,
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeMessage,
		Mode:       EventModeActive,
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
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &StickerMessage{
			ID:                  "325708",
			PackageID:           "1",
			StickerID:           "1",
			StickerResourceType: StickerResourceTypeStatic,
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeMessage,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &StickerMessage{
			ID:                  "325708",
			PackageID:           "1",
			StickerID:           "3",
			StickerResourceType: StickerResourceTypeAnimationSound,
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeMessage,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &StickerMessage{
			ID:                  "3257088",
			PackageID:           "20",
			StickerID:           "3",
			StickerResourceType: StickerResourceTypePerStickerText,
			Keywords:            []string{"cony", "sally", "Staring", "hi", "whatsup", "line", "howdy", "HEY", "Peeking", "wave", "peek", "Hello", "yo", "greetings"},
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeFollow,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
	},
	{
		Type:      EventTypeUnfollow,
		Mode:      EventModeActive,
		Timestamp: time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeJoin,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:    EventSourceTypeGroup,
			GroupID: "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		},
	},
	{
		Type:      EventTypeLeave,
		Mode:      EventModeActive,
		Timestamp: time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:    EventSourceTypeGroup,
			GroupID: "cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypePostback,
		Mode:       EventModeActive,
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
		Type:       EventTypePostback,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Postback: &Postback{
			Data: "action=sel&only=date",
			Params: &Params{
				Date: "2017-09-03",
			},
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypePostback,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Postback: &Postback{
			Data: "action=sel&only=time",
			Params: &Params{
				Time: "15:38",
			},
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypePostback,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Postback: &Postback{
			Data: "action=sel",
			Params: &Params{
				Datetime: "2017-09-03T15:38",
			},
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeBeacon,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "U012345678901234567890123456789ab",
		},
		Beacon: &Beacon{
			Hwid:          "374591320",
			Type:          BeaconEventTypeEnter,
			DeviceMessage: []byte{},
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeBeacon,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "U012345678901234567890123456789ab",
		},
		Beacon: &Beacon{
			Hwid:          "374591320",
			Type:          BeaconEventTypeEnter,
			DeviceMessage: []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef},
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeBeacon,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "U012345678901234567890123456789ab",
		},
		Beacon: &Beacon{
			Hwid:          "374591320",
			Type:          BeaconEventTypeStay,
			DeviceMessage: []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef},
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeAccountLink,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "U012345678901234567890123456789ab",
		},
		AccountLink: &AccountLink{
			Result: AccountLinkResultOK,
			Nonce:  "xxxxxxxxxxxxxxx",
		},
	},
	{
		ReplyToken: "0f3779fba3b349968c5d07db31eabf65",
		Type:       EventTypeMemberJoined,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:    EventSourceTypeGroup,
			GroupID: "C4af498062901234567890123456789ab",
		},
		Members: []*EventSource{
			{
				Type:   EventSourceTypeUser,
				UserID: "U4af498062901234567890123456789ab",
			},
			{
				Type:   EventSourceTypeUser,
				UserID: "U91eeaf62d901234567890123456789ab",
			},
		},
	},
	{
		Type:      EventTypeMemberLeft,
		Mode:      EventModeActive,
		Timestamp: time.Date(2016, time.May, 7, 13, 57, 59, int(960*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:    EventSourceTypeGroup,
			GroupID: "C4af498062901234567890123456789ab",
		},
		Members: []*EventSource{
			{
				Type:   EventSourceTypeUser,
				UserID: "U4af498062901234567890123456789ab",
			},
			{
				Type:   EventSourceTypeUser,
				UserID: "U91eeaf62d901234567890123456789ab",
			},
		},
	},
	{
		Type:      EventTypeThings,
		Mode:      EventModeActive,
		Timestamp: time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "U91eeaf62d901234567890123456789ab",
		},
		Things: &Things{
			DeviceID: `t2c449c9d1...`,
			Type:     `link`,
		},
	},
	{
		Type:      EventTypeThings,
		Mode:      EventModeActive,
		Timestamp: time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "U91eeaf62d901234567890123456789ab",
		},
		Things: &Things{
			DeviceID: `t2c449c9d1...`,
			Type:     `unlink`,
		},
	},
	{
		Type:      EventTypeThings,
		Mode:      EventModeActive,
		Timestamp: time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "U91eeaf62d901234567890123456789ab",
		},
		Things: &Things{
			DeviceID: "t016b8dc6...",
			Type:     "scenarioResult",
			Result: &ThingsResult{
				ScenarioID:             "01DE9CH7H...",
				Revision:               3,
				StartTime:              1563511217095,
				EndTime:                1563511217097,
				ResultCode:             ThingsResultCodeSuccess,
				BLENotificationPayload: []byte(`AQ==`),
				ActionResults: []*ThingsActionResult{
					&ThingsActionResult{
						Type: ThingsActionResultTypeBinary,
						Data: []byte(`/w==`),
					},
				},
			},
		},
	},
	{
		Type:       EventTypeThings,
		Mode:       EventModeActive,
		ReplyToken: "f026a377...",
		Timestamp:  time.Date(2019, time.July, 19, 4, 40, 18, int(376*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "U91eeaf62d901234567890123456789ab",
		},
		Things: &Things{
			DeviceID: "t016b8d...",
			Type:     "scenarioResult",
			Result: &ThingsResult{
				ScenarioID:             "01DE9CH7H...",
				Revision:               3,
				StartTime:              1563511217095,
				EndTime:                1563511217097,
				ResultCode:             ThingsResultCodeGattError,
				ErrorReason:            "c.l.h.D: No characteristic is found for the given UUID.",
				ActionResults:          []*ThingsActionResult{},
				BLENotificationPayload: []byte{},
			},
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeMessage,
		Mode:       EventModeStandby,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &TextMessage{
			ID:   "325708",
			Text: "Stand by me",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeMessage,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "u206d25c2ea6bd87c17655609a1c37cb8",
		},
		Message: &TextMessage{
			ID:   "325708",
			Text: "Hello, world! (love)",
			Emojis: []*Emoji{
				&Emoji{Index: 14, Length: 6, ProductID: "5ac1bfd5040ab15980c9b435", EmojiID: "001"},
			},
		},
	},
	{
		Type:      EventTypeUnsend,
		Mode:      EventModeActive,
		Timestamp: time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:    EventSourceTypeGroup,
			GroupID: "Ca56f94637c...",
			UserID:  "U4af4980629...",
		},
		Unsend: &Unsend{
			MessageID: "325708",
		},
	},
	{
		ReplyToken: "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
		Type:       EventTypeVideoPlayComplete,
		Mode:       EventModeActive,
		Timestamp:  time.Date(2016, time.May, 7, 13, 57, 59, int(859*time.Millisecond), time.UTC),
		Source: &EventSource{
			Type:   EventSourceTypeUser,
			UserID: "U4af4980629...",
		},
		VideoPlayComplete: &VideoPlayComplete{
			TrackingID: "track_id",
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
				t.Errorf("Event %d %v; want %v", i, got, want)
				gota := got
				if !reflect.DeepEqual(
					gota.Things.Result.BLENotificationPayload,
					want.Things.Result.BLENotificationPayload,
				) {
					t.Log("this")
					t.Log(gota.Things.Result.BLENotificationPayload == nil)
					t.Log(want.Things.Result.BLENotificationPayload == nil)
				}
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
		req.Header.Set("X-Line-Signature", "invalidsignatue")
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

func TestEventMarshaling(t *testing.T) {
	testCases := &struct {
		Events []map[string]interface{} `json:"events"`
	}{}
	err := json.Unmarshal([]byte(webhookTestRequestBody), testCases)
	if err != nil {
		t.Fatal(err)
	}
	for i, want := range testCases.Events {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotJSON, err := json.Marshal(webhookTestWantEvents[i])
			if err != nil {
				t.Error(err)
				return
			}
			got := map[string]interface{}{}
			err = json.Unmarshal(gotJSON, &got)
			if err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Event marshal %v; want %v", got, want)
			}
		})
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
		req.Header.Set("X-Line-Signature", sign)
		client.ParseRequest(req)
	}
}

func TestGetWebhookInfo(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *WebhookInfoResponse
		Error       error
	}
	var testCases = []struct {
		Label        string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			Label:        "Success",
			ResponseCode: 200,
			Response:     []byte(`{"endpoint":"https://example.herokuapp.com/test","active":"true"}`),
			Want: want{
				URLPath:     APIEndpointGetWebhookInfo,
				RequestBody: []byte(""),
				Response: &WebhookInfoResponse{
					Endpoint: "https://example.herokuapp.com/test",
					Active:   "true",
				},
			},
		},
		{
			Label:        "Internal server error",
			ResponseCode: 500,
			Response:     []byte("500 Internal server error"),
			Want: want{
				URLPath:     APIEndpointGetWebhookInfo,
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 500,
				},
			},
		},
		{
			Label:        "Invalid channelAccessToken error",
			ResponseCode: 401,
			Response:     []byte(`{"message":"Authentication failed due to the following reason: invalid token. Confirm that the access token in the authorization header is valid."}`),
			Want: want{
				URLPath:     APIEndpointGetWebhookInfo,
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 401,
					Response: &ErrorResponse{
						Message: "Authentication failed due to the following reason: invalid token. Confirm that the access token in the authorization header is valid.",
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
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
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
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			res, err := client.GetWebhookInfo().Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %v; want %v", res, tc.Want.Response)
			}
		})
	}
}

func TestGetWebhookInfoWithContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.GetWebhookInfo().WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func BenchmarkGetWebhookInfo(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte(`{"endpoint":"https://example.herokuapp.com/test","active":"true"}`))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.GetWebhookInfo().Do()
	}
}

func TestTestWebhook(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *TestWebhookResponse
		Error       error
	}
	var testCases = []struct {
		Label        string
		ResponseCode int
		Response     []byte
		Want         want
	}{
		{
			Label:        "Success",
			ResponseCode: 200,
			Response: []byte(`{
				"success": true,
				"timestamp": "2020-09-30T05:38:20.031Z",
				"statusCode": 200,
				"reason": "OK",
				"detail": "200"
			}`),
			Want: want{
				URLPath:     APIEndpointTestWebhook,
				RequestBody: []byte(""),
				Response: &TestWebhookResponse{
					Success:    true,
					Timestamp:  time.Date(2020, time.September, 30, 05, 38, 20, int(31*time.Millisecond), time.UTC),
					StatusCode: 200,
					Reason:     "OK",
					Detail:     "200",
				},
			},
		},
		{
			Label:        "Internal server error",
			ResponseCode: 500,
			Response:     []byte("500 Internal server error"),
			Want: want{
				URLPath:     APIEndpointTestWebhook,
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 500,
				},
			},
		},
		{
			Label:        "Invalid channelAccessToken error",
			ResponseCode: 401,
			Response:     []byte(`{"message":"Authentication failed due to the following reason: invalid token. Confirm that the access token in the authorization header is valid."}`),
			Want: want{
				URLPath:     APIEndpointTestWebhook,
				RequestBody: []byte(""),
				Error: &APIError{
					Code: 401,
					Response: &ErrorResponse{
						Message: "Authentication failed due to the following reason: invalid token. Confirm that the access token in the authorization header is valid.",
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
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
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
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			res, err := client.TestWebhook().Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %v; want %v", res, tc.Want.Response)
			}
		})
	}
}

func TestTestWebhookWithContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.TestWebhook().WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func TestSetWebhookEndpointURL(t *testing.T) {
	type want struct {
		URLPath     string
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	var testCases = []struct {
		Label        string
		Endpoint     string
		ResponseCode int
		RequestID    string
		Response     []byte
		Want         want
	}{
		{
			Label:        "Success",
			Endpoint:     "https://example.com/abcdefghijklmn",
			ResponseCode: 200,
			RequestID:    "f70dd685-499a-4231-a441-f24b8d4fba21",
			Response:     []byte(`{}`),
			Want: want{
				URLPath:     APIEndpointSetWebhookEndpoint,
				RequestBody: []byte(`{"endpoint":"https://example.com/abcdefghijklmn"}` + "\n"),
				Response: &BasicResponse{
					RequestID: "f70dd685-499a-4231-a441-f24b8d4fba21",
				},
			},
		},
		{
			Label:        "Internal server error",
			Endpoint:     "https://example.com/abcdefghijklmn",
			ResponseCode: 500,
			RequestID:    "f70dd685-499a-4231-a441-f24b8d4fba21",
			Response:     []byte("500 Internal server error"),
			Want: want{
				URLPath:     APIEndpointSetWebhookEndpoint,
				RequestBody: []byte(`{"endpoint":"https://example.com/abcdefghijklmn"}` + "\n"),
				Error: &APIError{
					Code: 500,
				},
			},
		},
		{
			Label:        "Invalid webhook URL error:not https",
			Endpoint:     "http://example.com/not/https",
			ResponseCode: 400,
			RequestID:    "f70dd685-499a-4231-a441-f24b8d4fba21",
			Response:     []byte(`{"message":"Invalid webhook endpoint URL"}`),
			Want: want{
				URLPath:     APIEndpointSetWebhookEndpoint,
				RequestBody: []byte(`{"endpoint":"http://example.com/not/https"}` + "\n"),
				Error: &APIError{
					Code: 400,
					Response: &ErrorResponse{
						Message: "Invalid webhook endpoint URL",
					},
				},
			},
		},
		{
			Label:        "Invalid webhook URL error:more 500 characters",
			Endpoint:     "https://example.com/exceed/500/characters",
			ResponseCode: 500,
			RequestID:    "f70dd685-499a-4231-a441-f24b8d4fba21",
			Response:     []byte(`{"message":"The request body has 1 error(s)","details":[{"message":"Size must be between 0 and 500","property":"endpoint"}]}`),
			Want: want{
				URLPath:     APIEndpointSetWebhookEndpoint,
				RequestBody: []byte(`{"endpoint":"https://example.com/exceed/500/characters"}` + "\n"),
				Error: &APIError{
					Code: 500,
					Response: &ErrorResponse{
						Message: "The request body has 1 error(s)",
						Details: []errorResponseDetail{
							{
								Message:  "Size must be between 0 and 500",
								Property: "endpoint",
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
		if r.Method != http.MethodPut {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPut)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
		}
		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
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
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			res, err := client.SetWebhookEndpointURL(tc.Endpoint).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if !reflect.DeepEqual(res, tc.Want.Response) {
				t.Errorf("Response %v; want %v", res, tc.Want.Response)
			}
		})
	}
}

func TestSetWebhookEndpointURLWithContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.SetWebhookEndpointURL("https://example.com/abcdefghijklmn").WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func BenchmarkSetWebhookEndpointURL(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b.Error("Unexpected Data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.SetWebhookEndpointURL("https://example.com/abcdefghijklmn").Do()
	}
}

func BenchmarkTestWebhook(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b.Error("Unexpected data API call")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.TestWebhook().Do()
	}
}
