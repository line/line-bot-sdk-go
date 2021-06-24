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
	"strconv"
	"testing"
)

func TestMessageTypes(t *testing.T) {
	type want struct {
		Type MessageType
	}
	testCases := []struct {
		Label   string
		Message Message
		Want    want
	}{
		{
			Label:   "A text message",
			Message: NewTextMessage("Hello, world"),
			Want: want{
				Type: MessageTypeText,
			},
		},
		{
			Label:   "A location message",
			Message: NewLocationMessage("title", "address", 35.65910807942215, 139.70372892916203),
			Want: want{
				Type: MessageTypeLocation,
			},
		},
		{
			Label:   "A image message",
			Message: NewImageMessage("https://example.com/original.jpg", "https://example.com/preview.jpg"),
			Want: want{
				Type: MessageTypeImage,
			},
		},
		{
			Label:   "A sticker message",
			Message: NewStickerMessage("1", "1"),
			Want: want{
				Type: MessageTypeSticker,
			},
		},
		{
			Label:   "A audio message",
			Message: NewAudioMessage("https://example.com/original.m4a", 1000),
			Want: want{
				Type: MessageTypeAudio,
			},
		},
		{
			Label: "A template message",
			Message: NewTemplateMessage(
				"this is a buttons template",
				NewButtonsTemplate(
					"https://example.com/bot/images/image.jpg",
					"",
					"Please select",
					NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
				),
			),
			Want: want{
				Type: MessageTypeTemplate,
			},
		},
		{
			Label: "A flex message",
			Message: NewFlexMessage(
				"this is a flex message",
				&BubbleContainer{}),
			Want: want{
				Type: MessageTypeFlex,
			},
		},
		{
			Label: "A Imagemap message",
			Message: NewImagemapMessage(
				"https://example.com/bot/images/rm001",
				"this is an imagemap",
				ImagemapBaseSize{1040, 1040},
				NewURIImagemapAction("example", "https://example.com/", ImagemapArea{520, 0, 520, 1040}),
			),
			Want: want{
				Type: MessageTypeImagemap,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			if tc.Message.Type() != tc.Want.Type {
				t.Errorf("Message type mismatch: have %v; want %v", tc.Message.Type(), tc.Want.Type)
			}
		})
	}
}
