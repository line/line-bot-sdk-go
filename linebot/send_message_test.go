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
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestPushMessages(t *testing.T) {
	var toUserID = "U0cc15697597f61dd8b01cea8b027050e"
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	var testCases = []struct {
		Messages     []SendingMessage
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			// A text message
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"text","text":"Hello, world"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A image message
			Messages:     []SendingMessage{NewImageMessage("http://example.com/original.jpg", "http://example.com/preview.jpg")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"image","originalContentUrl":"http://example.com/original.jpg","previewImageUrl":"http://example.com/preview.jpg"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A video message
			Messages:     []SendingMessage{NewVideoMessage("http://example.com/original.mp4", "http://example.com/preview.jpg")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"video","originalContentUrl":"http://example.com/original.mp4","previewImageUrl":"http://example.com/preview.jpg"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A audio message
			Messages:     []SendingMessage{NewAudioMessage("http://example.com/original.m4a", 1000)},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"audio","originalContentUrl":"http://example.com/original.m4a","duration":1000}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A location message
			Messages:     []SendingMessage{NewLocationMessage("title", "address", 35.65910807942215, 139.70372892916203)},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"location","title":"title","address":"address","latitude":35.65910807942215,"longitude":139.70372892916203}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A sticker message
			Messages:     []SendingMessage{NewStickerMessage("1", "1")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"sticker","packageId":"1","stickerId":"1"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A buttons template message
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"https://example.com/bot/images/image.jpg",
						"Menu",
						"Please select",
						NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
						NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
						NewURIAction("View detail", "http://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","thumbnailImageUrl":"https://example.com/bot/images/image.jpg","title":"Menu","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"http://example.com/page/123"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A buttons template message with datetimepicker action
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"https://example.com/bot/images/image.jpg",
						"Menu",
						"Please select a date, time or datetime",
						NewDatetimePickerAction("Date", "action=sel&only=date", "date", "2017-09-01", "2017-09-03", ""),
						NewDatetimePickerAction("Time", "action=sel&only=time", "time", "", "23:59", "00:00"),
						NewDatetimePickerAction("DateTime", "action=sel", "datetime", "2017-09-01T12:00", "", ""),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","thumbnailImageUrl":"https://example.com/bot/images/image.jpg","title":"Menu","text":"Please select a date, time or datetime","actions":[{"type":"datetimepicker","label":"Date","data":"action=sel\u0026only=date","mode":"date","initial":"2017-09-01","max":"2017-09-03"},{"type":"datetimepicker","label":"Time","data":"action=sel\u0026only=time","mode":"time","max":"23:59","min":"00:00"},{"type":"datetimepicker","label":"DateTime","data":"action=sel","mode":"datetime","initial":"2017-09-01T12:00"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A buttons template message without thumbnailImageURL
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"",
						"Menu",
						"Please select",
						NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
						NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
						NewURIAction("View detail", "http://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","title":"Menu","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"http://example.com/page/123"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A buttons template message without title
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"https://example.com/bot/images/image.jpg",
						"",
						"Please select",
						NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
						NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
						NewURIAction("View detail", "http://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","thumbnailImageUrl":"https://example.com/bot/images/image.jpg","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"http://example.com/page/123"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A buttons template message without title, with image options
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"https://example.com/bot/images/image.jpg",
						"",
						"Please select",
						NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
						NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
						NewURIAction("View detail", "http://example.com/page/123"),
					).WithImageOptions("rectangle", "cover", "#FFFFFF"),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","thumbnailImageUrl":"https://example.com/bot/images/image.jpg","imageAspectRatio":"rectangle","imageSize":"cover","imageBackgroundColor":"#FFFFFF","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"http://example.com/page/123"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A buttons template message without thumbnailImageURL and title
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"",
						"",
						"Please select",
						NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
						NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
						NewURIAction("View detail", "http://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"http://example.com/page/123"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A confirm template message
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a confirm template",
					NewConfirmTemplate(
						"Are you sure?",
						NewMessageAction("Yes", "yes"),
						NewMessageAction("No", "no"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a confirm template","template":{"type":"confirm","text":"Are you sure?","actions":[{"type":"message","label":"Yes","text":"yes"},{"type":"message","label":"No","text":"no"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A carousel template message
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a carousel template",
					NewCarouselTemplate(
						NewCarouselColumn(
							"https://example.com/bot/images/item1.jpg",
							"this is menu",
							"description",
							NewPostbackAction("Buy", "action=buy&itemid=111", "", ""),
							NewPostbackAction("Add to cart", "action=add&itemid=111", "", ""),
							NewURIAction("View detail", "http://example.com/page/111"),
						),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a carousel template","template":{"type":"carousel","columns":[{"thumbnailImageUrl":"https://example.com/bot/images/item1.jpg","title":"this is menu","text":"description","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=111"},{"type":"postback","label":"Add to cart","data":"action=add\u0026itemid=111"},{"type":"uri","label":"View detail","uri":"http://example.com/page/111"}]}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A carousel template message, with new image options
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a carousel template with imageAspectRatio, imageSize and imageBackgroundColor",
					NewCarouselTemplate(
						NewCarouselColumn(
							"https://example.com/bot/images/item1.jpg",
							"this is menu",
							"description",
							NewPostbackAction("Buy", "action=buy&itemid=111", "", ""),
							NewPostbackAction("Add to cart", "action=add&itemid=111", "", ""),
							NewURIAction("View detail", "http://example.com/page/111"),
						).WithImageOptions("#FFFFFF"),
					).WithImageOptions("rectangle", "cover"),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a carousel template with imageAspectRatio, imageSize and imageBackgroundColor","template":{"type":"carousel","columns":[{"thumbnailImageUrl":"https://example.com/bot/images/item1.jpg","imageBackgroundColor":"#FFFFFF","title":"this is menu","text":"description","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=111"},{"type":"postback","label":"Add to cart","data":"action=add\u0026itemid=111"},{"type":"uri","label":"View detail","uri":"http://example.com/page/111"}]}],"imageAspectRatio":"rectangle","imageSize":"cover"}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A imagecarousel template message
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a image carousel template",
					NewImageCarouselTemplate(
						NewImageCarouselColumn(
							"https://example.com/bot/images/item1.jpg",
							NewURIAction("View detail", "http://example.com/page/111"),
						),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a image carousel template","template":{"type":"image_carousel","columns":[{"imageUrl":"https://example.com/bot/images/item1.jpg","action":{"type":"uri","label":"View detail","uri":"http://example.com/page/111"}}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A imagemap message
			Messages: []SendingMessage{
				NewImagemapMessage(
					"https://example.com/bot/images/rm001",
					"this is an imagemap",
					ImagemapBaseSize{1040, 1040},
					NewURIImagemapAction("https://example.com/", ImagemapArea{520, 0, 520, 1040}),
					NewMessageImagemapAction("hello", ImagemapArea{520, 0, 520, 1040}),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"imagemap","baseUrl":"https://example.com/bot/images/rm001","altText":"this is an imagemap","baseSize":{"width":1040,"height":1040},"actions":[{"type":"uri","linkUri":"https://example.com/","area":{"x":520,"y":0,"width":520,"height":1040}},{"type":"message","text":"hello","area":{"x":520,"y":0,"width":520,"height":1040}}]}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A flex message
			Messages: []SendingMessage{
				NewFlexMessage(
					"this is a flex message",
					&BubbleContainer{
						Type: FlexContainerTypeBubble,
						Body: &BoxComponent{
							Type:   FlexComponentTypeBox,
							Layout: FlexBoxLayoutTypeVertical,
							Contents: []FlexComponent{
								&TextComponent{
									Type: FlexComponentTypeText,
									Text: "hello",
								},
								&TextComponent{
									Type: FlexComponentTypeText,
									Text: "world",
								},
							},
						},
					},
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"flex","altText":"this is a flex message","contents":{"type":"bubble","body":{"type":"box","layout":"vertical","contents":[{"type":"text","text":"hello"},{"type":"text","text":"world"}]}}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A text message with quick replies
			Messages: []SendingMessage{
				NewTextMessage(
					"Select your favorite food category or send me your location!",
				).WithQuickReplies(
					NewQuickReplyItems(
						NewQuickReplyButton("https://example.com/sushi.png", NewMessageAction("Sushi", "Sushi")),
						NewQuickReplyButton("https://example.com/tempura.png", NewMessageAction("Tempura", "Tempura")),
						NewQuickReplyButton("", NewLocationAction("Send location")),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"text","text":"Select your favorite food category or send me your location!","quickReply":{"items":[{"type":"action","imageUrl":"https://example.com/sushi.png","action":{"type":"message","label":"Sushi","text":"Sushi"}},{"type":"action","imageUrl":"https://example.com/tempura.png","action":{"type":"message","label":"Tempura","text":"Tempura"}},{"type":"action","action":{"type":"location","label":"Send location"}}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// Multiple messages
			Messages:     []SendingMessage{NewTextMessage("Hello, world1"), NewTextMessage("Hello, world2")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"text","text":"Hello, world1"},{"type":"text","text":"Hello, world2"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// Bad request
			Messages:     []SendingMessage{NewTextMessage(""), NewTextMessage("")},
			ResponseCode: 400,
			Response:     []byte(`{"message":"Request body has 2 error(s).","details":[{"message":"may not be empty","property":"messages[0].text"},{"message":"may not be empty","property":"messages[1].text"}]}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"text","text":""},{"type":"text","text":""}]}` + "\n"),
				Error: &APIError{
					Code: 400,
					Response: &ErrorResponse{
						Message: "Request body has 2 error(s).",
						Details: []errorResponseDetail{
							{
								Message:  "may not be empty",
								Property: "messages[0].text",
							},
							{
								Message:  "may not be empty",
								Property: "messages[1].text",
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
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != APIEndpointPushMessage {
			t.Errorf("URLPath %s; want %s", r.URL.Path, APIEndpointPushMessage)
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
	client, err := mockClient(server)
	if err != nil {
		t.Fatal(err)
	}
	for i, tc := range testCases {
		currentTestIdx = i
		res, err := client.PushMessage(toUserID, tc.Messages...).Do()
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

func TestPushMessagesWithContext(t *testing.T) {
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

func TestReplyMessages(t *testing.T) {
	var replyToken = "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA"
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	var testCases = []struct {
		Messages     []SendingMessage
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			// A text message
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"replyToken":"nHuyWiB7yP5Zw52FIkcQobQuGDXCTA","messages":[{"type":"text","text":"Hello, world"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A location message
			Messages:     []SendingMessage{NewLocationMessage("title", "address", 35.65910807942215, 139.70372892916203)},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"replyToken":"nHuyWiB7yP5Zw52FIkcQobQuGDXCTA","messages":[{"type":"location","title":"title","address":"address","latitude":35.65910807942215,"longitude":139.70372892916203}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A image message
			Messages:     []SendingMessage{NewImageMessage("http://example.com/original.jpg", "http://example.com/preview.jpg")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"replyToken":"nHuyWiB7yP5Zw52FIkcQobQuGDXCTA","messages":[{"type":"image","originalContentUrl":"http://example.com/original.jpg","previewImageUrl":"http://example.com/preview.jpg"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A sticker message
			Messages:     []SendingMessage{NewStickerMessage("1", "1")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"replyToken":"nHuyWiB7yP5Zw52FIkcQobQuGDXCTA","messages":[{"type":"sticker","packageId":"1","stickerId":"1"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// Bad request
			Messages:     []SendingMessage{NewTextMessage(""), NewTextMessage("")},
			ResponseCode: 400,
			Response:     []byte(`{"message":"Request body has 2 error(s).","details":[{"message":"may not be empty","property":"messages[0].text"},{"message":"may not be empty","property":"messages[1].text"}]}`),
			Want: want{
				RequestBody: []byte(`{"replyToken":"nHuyWiB7yP5Zw52FIkcQobQuGDXCTA","messages":[{"type":"text","text":""},{"type":"text","text":""}]}` + "\n"),
				Error: &APIError{
					Code: 400,
					Response: &ErrorResponse{
						Message: "Request body has 2 error(s).",
						Details: []errorResponseDetail{
							{
								Message:  "may not be empty",
								Property: "messages[0].text",
							},
							{
								Message:  "may not be empty",
								Property: "messages[1].text",
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
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != APIEndpointReplyMessage {
			t.Errorf("URLPath %s; want %s", r.URL.Path, APIEndpointReplyMessage)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		tc := testCases[currentTestIdx]
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
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
		res, err := client.ReplyMessage(replyToken, tc.Messages...).Do()
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

func TestReplyMessagesWithContext(t *testing.T) {
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
	_, err = client.ReplyMessage("nHuyWiB7yP5Zw52FIkcQobQuGDXCTA", NewTextMessage("Hello, world")).WithContext(ctx).Do()
	if err != context.DeadlineExceeded {
		t.Errorf("err %v; want %v", err, context.DeadlineExceeded)
	}
}

func TestMulticastMessages(t *testing.T) {
	var toUserIDs = []string{
		"U0cc15697597f61dd8b01cea8b027050e",
		"U38ecbecfade326557b6971140741a4a6",
	}
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	var testCases = []struct {
		Messages     []SendingMessage
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			// A text message
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":["U0cc15697597f61dd8b01cea8b027050e","U38ecbecfade326557b6971140741a4a6"],"messages":[{"type":"text","text":"Hello, world"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A location message
			Messages:     []SendingMessage{NewLocationMessage("title", "address", 35.65910807942215, 139.70372892916203)},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":["U0cc15697597f61dd8b01cea8b027050e","U38ecbecfade326557b6971140741a4a6"],"messages":[{"type":"location","title":"title","address":"address","latitude":35.65910807942215,"longitude":139.70372892916203}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A image message
			Messages:     []SendingMessage{NewImageMessage("http://example.com/original.jpg", "http://example.com/preview.jpg")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":["U0cc15697597f61dd8b01cea8b027050e","U38ecbecfade326557b6971140741a4a6"],"messages":[{"type":"image","originalContentUrl":"http://example.com/original.jpg","previewImageUrl":"http://example.com/preview.jpg"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A sticker message
			Messages:     []SendingMessage{NewStickerMessage("1", "1")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":["U0cc15697597f61dd8b01cea8b027050e","U38ecbecfade326557b6971140741a4a6"],"messages":[{"type":"sticker","packageId":"1","stickerId":"1"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// Bad request
			Messages:     []SendingMessage{NewTextMessage(""), NewTextMessage("")},
			ResponseCode: 400,
			Response:     []byte(`{"message":"Request body has 2 error(s).","details":[{"message":"may not be empty","property":"messages[0].text"},{"message":"may not be empty","property":"messages[1].text"}]}`),
			Want: want{
				RequestBody: []byte(`{"to":["U0cc15697597f61dd8b01cea8b027050e","U38ecbecfade326557b6971140741a4a6"],"messages":[{"type":"text","text":""},{"type":"text","text":""}]}` + "\n"),
				Error: &APIError{
					Code: 400,
					Response: &ErrorResponse{
						Message: "Request body has 2 error(s).",
						Details: []errorResponseDetail{
							{
								Message:  "may not be empty",
								Property: "messages[0].text",
							},
							{
								Message:  "may not be empty",
								Property: "messages[1].text",
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
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != APIEndpointMulticast {
			t.Errorf("URLPath %s; want %s", r.URL.Path, APIEndpointMulticast)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		tc := testCases[currentTestIdx]
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody %s; want %s", body, tc.Want.RequestBody)
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
		res, err := client.Multicast(toUserIDs, tc.Messages...).Do()
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

func TestMulticastMessagesWithContext(t *testing.T) {
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
	_, err = client.Multicast([]string{"U0cc15697597f61dd8b01cea8b027050e", "U38ecbecfade326557b6971140741a4a6"}, NewTextMessage("Hello, world")).WithContext(ctx).Do()
	if err != context.DeadlineExceeded {
		t.Errorf("err %v; want %v", err, context.DeadlineExceeded)
	}
}

func BenchmarkPushMessages(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte("{}"))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.PushMessage("U0cc15697597f61dd8b01cea8b027050e", NewTextMessage("Hello, world")).Do()
	}
}

func BenchmarkReplyMessages(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte("{}"))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.ReplyMessage("nHuyWiB7yP5Zw52FIkcQobQuGDXCTA", NewTextMessage("Hello, world")).Do()
	}
}

func BenchmarkMulticast(b *testing.B) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte("{}"))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.Multicast([]string{"U0cc15697597f61dd8b01cea8b027050e", "U38ecbecfade326557b6971140741a4a6"}, NewTextMessage("Hello, world")).Do()
	}
}
