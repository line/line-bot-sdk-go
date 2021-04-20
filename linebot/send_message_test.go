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
	"strconv"
	"testing"
	"time"
)

func TestPushMessages(t *testing.T) {
	toUserID := "U0cc15697597f61dd8b01cea8b027050e"
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label        string
		Messages     []SendingMessage
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			Label:        "A text message",
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"text","text":"Hello, world"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A image message",
			Messages:     []SendingMessage{NewImageMessage("https://example.com/original.jpg", "https://example.com/preview.jpg")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"image","originalContentUrl":"https://example.com/original.jpg","previewImageUrl":"https://example.com/preview.jpg"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A video message",
			Messages:     []SendingMessage{NewVideoMessage("https://example.com/original.mp4", "https://example.com/preview.jpg")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"video","originalContentUrl":"https://example.com/original.mp4","previewImageUrl":"https://example.com/preview.jpg"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A audio message",
			Messages:     []SendingMessage{NewAudioMessage("https://example.com/original.m4a", 1000)},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"audio","originalContentUrl":"https://example.com/original.m4a","duration":1000}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A location message",
			Messages:     []SendingMessage{NewLocationMessage("title", "address", 35.65910807942215, 139.70372892916203)},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"location","title":"title","address":"address","latitude":35.65910807942215,"longitude":139.70372892916203}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A sticker message",
			Messages:     []SendingMessage{NewStickerMessage("1", "1")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"sticker","packageId":"1","stickerId":"1"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A buttons template message",
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"https://example.com/bot/images/image.jpg",
						"Menu",
						"Please select",
						NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
						NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
						NewURIAction("View detail", "https://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","thumbnailImageUrl":"https://example.com/bot/images/image.jpg","title":"Menu","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"https://example.com/page/123"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A buttons template message with datetimepicker action",
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
			Label: "A buttons template message without thumbnailImageURL",
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"",
						"Menu",
						"Please select",
						NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
						NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
						NewURIAction("View detail", "https://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","title":"Menu","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"https://example.com/page/123"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A buttons template message without title",
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"https://example.com/bot/images/image.jpg",
						"",
						"Please select",
						NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
						NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
						NewURIAction("View detail", "https://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","thumbnailImageUrl":"https://example.com/bot/images/image.jpg","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"https://example.com/page/123"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A buttons template message without title, with image options",
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"https://example.com/bot/images/image.jpg",
						"",
						"Please select",
						NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
						NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
						NewURIAction("View detail", "https://example.com/page/123"),
					).WithImageOptions("rectangle", "cover", "#FFFFFF"),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","thumbnailImageUrl":"https://example.com/bot/images/image.jpg","imageAspectRatio":"rectangle","imageSize":"cover","imageBackgroundColor":"#FFFFFF","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"https://example.com/page/123"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A buttons template message without thumbnailImageURL and title",
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"",
						"",
						"Please select",
						NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
						NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
						NewURIAction("View detail", "https://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"https://example.com/page/123"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A confirm template message",
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
			Label: "A carousel template message",
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
							NewURIAction("View detail", "https://example.com/page/111"),
						),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a carousel template","template":{"type":"carousel","columns":[{"thumbnailImageUrl":"https://example.com/bot/images/item1.jpg","title":"this is menu","text":"description","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=111"},{"type":"postback","label":"Add to cart","data":"action=add\u0026itemid=111"},{"type":"uri","label":"View detail","uri":"https://example.com/page/111"}]}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A carousel template message, with new image options",
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
							NewURIAction("View detail", "https://example.com/page/111"),
						).WithImageOptions("#FFFFFF"),
					).WithImageOptions("rectangle", "cover"),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a carousel template with imageAspectRatio, imageSize and imageBackgroundColor","template":{"type":"carousel","columns":[{"thumbnailImageUrl":"https://example.com/bot/images/item1.jpg","imageBackgroundColor":"#FFFFFF","title":"this is menu","text":"description","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=111"},{"type":"postback","label":"Add to cart","data":"action=add\u0026itemid=111"},{"type":"uri","label":"View detail","uri":"https://example.com/page/111"}]}],"imageAspectRatio":"rectangle","imageSize":"cover"}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A imagecarousel template message",
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a image carousel template",
					NewImageCarouselTemplate(
						NewImageCarouselColumn(
							"https://example.com/bot/images/item1.jpg",
							NewURIAction("View detail", "https://example.com/page/111"),
						),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"template","altText":"this is a image carousel template","template":{"type":"image_carousel","columns":[{"imageUrl":"https://example.com/bot/images/item1.jpg","action":{"type":"uri","label":"View detail","uri":"https://example.com/page/111"}}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A imagemap message",
			Messages: []SendingMessage{
				NewImagemapMessage(
					"https://example.com/bot/images/rm001",
					"this is an imagemap",
					ImagemapBaseSize{1040, 1040},
					NewURIImagemapAction("example", "https://example.com/", ImagemapArea{520, 0, 520, 1040}),
					NewMessageImagemapAction("hello", "hello", ImagemapArea{520, 0, 520, 1040}),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"imagemap","baseUrl":"https://example.com/bot/images/rm001","altText":"this is an imagemap","baseSize":{"width":1040,"height":1040},"actions":[{"type":"uri","label":"example","linkUri":"https://example.com/","area":{"x":520,"y":0,"width":520,"height":1040}},{"type":"message","label":"hello","text":"hello","area":{"x":520,"y":0,"width":520,"height":1040}}]}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A imagemap messages with video 1",
			Messages: []SendingMessage{
				NewImagemapMessage(
					"https://example.com/bot/images/rm001",
					"this is an imagemap with video",
					ImagemapBaseSize{1040, 1040},
					NewURIImagemapAction("example", "https://example.com/", ImagemapArea{520, 0, 520, 1040}),
					NewMessageImagemapAction("hello", "hello", ImagemapArea{520, 0, 520, 1040}),
				).WithVideo(&ImagemapVideo{
					OriginalContentURL: "https://example.com/original.mp4",
					PreviewImageURL:    "https://example.com/preview.jpg",
					Area:               ImagemapArea{10, 10, 100, 200},
				}),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"imagemap","baseUrl":"https://example.com/bot/images/rm001","altText":"this is an imagemap with video","baseSize":{"width":1040,"height":1040},"actions":[{"type":"uri","label":"example","linkUri":"https://example.com/","area":{"x":520,"y":0,"width":520,"height":1040}},{"type":"message","label":"hello","text":"hello","area":{"x":520,"y":0,"width":520,"height":1040}}],"video":{"originalContentUrl":"https://example.com/original.mp4","previewImageUrl":"https://example.com/preview.jpg","area":{"x":10,"y":10,"width":100,"height":200}}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A imagemap messages with video 2",
			Messages: []SendingMessage{
				NewImagemapMessage(
					"https://example.com/bot/images/rm001",
					"this is an imagemap with video and external link",
					ImagemapBaseSize{1040, 1040},
					NewURIImagemapAction("example", "https://example.com/", ImagemapArea{520, 0, 520, 1040}),
					NewMessageImagemapAction("hello", "hello", ImagemapArea{520, 0, 520, 1040}),
				).WithVideo(&ImagemapVideo{
					OriginalContentURL: "https://example.com/original.mp4",
					PreviewImageURL:    "https://example.com/preview.jpg",
					Area:               ImagemapArea{10, 10, 100, 200},
					ExternalLink: &ImagemapVideoExternalLink{
						LinkURI: "https://example.com/",
						Label:   "external link",
					},
				}),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"imagemap","baseUrl":"https://example.com/bot/images/rm001","altText":"this is an imagemap with video and external link","baseSize":{"width":1040,"height":1040},"actions":[{"type":"uri","label":"example","linkUri":"https://example.com/","area":{"x":520,"y":0,"width":520,"height":1040}},{"type":"message","label":"hello","text":"hello","area":{"x":520,"y":0,"width":520,"height":1040}}],"video":{"originalContentUrl":"https://example.com/original.mp4","previewImageUrl":"https://example.com/preview.jpg","area":{"x":10,"y":10,"width":100,"height":200},"externalLink":{"linkUri":"https://example.com/","label":"external link"}}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A flex message",
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
									Flex: IntPtr(0),
								},
								&SpacerComponent{
									Type: FlexComponentTypeSpacer,
								},
							},
						},
					},
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"flex","altText":"this is a flex message","contents":{"type":"bubble","body":{"type":"box","layout":"vertical","contents":[{"type":"text","text":"hello"},{"type":"text","text":"world","flex":0},{"type":"spacer"}]}}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label: "A text message with quick replies",
			Messages: []SendingMessage{
				NewTextMessage(
					"Select your favorite food category or send me your location!",
				).WithQuickReplies(
					NewQuickReplyItems(
						NewQuickReplyButton("https://example.com/sushi.png", NewMessageAction("Sushi", "Sushi")),
						NewQuickReplyButton("https://example.com/tempura.png", NewMessageAction("Tempura", "Tempura")),
						NewQuickReplyButton("", NewLocationAction("Send location")),
						NewQuickReplyButton("", NewCameraAction("Open camera")),
						NewQuickReplyButton("", NewURIAction("Phone order", "tel:09001234567")),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"text","text":"Select your favorite food category or send me your location!","quickReply":{"items":[{"type":"action","imageUrl":"https://example.com/sushi.png","action":{"type":"message","label":"Sushi","text":"Sushi"}},{"type":"action","imageUrl":"https://example.com/tempura.png","action":{"type":"message","label":"Tempura","text":"Tempura"}},{"type":"action","action":{"type":"location","label":"Send location"}},{"type":"action","action":{"type":"camera","label":"Open camera"}},{"type":"action","action":{"type":"uri","label":"Phone order","uri":"tel:09001234567"}}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "Multiple messages",
			Messages:     []SendingMessage{NewTextMessage("Hello, world1"), NewTextMessage("Hello, world2")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"text","text":"Hello, world1"},{"type":"text","text":"Hello, world2"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "Bad request",
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
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			res, err := client.PushMessage(toUserID, tc.Messages...).Do()
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

func TestPushMessagesWithContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.PushMessage("U0cc15697597f61dd8b01cea8b027050e", NewTextMessage("Hello, world")).WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func TestReplyMessages(t *testing.T) {
	replyToken := "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA"
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label        string
		Messages     []SendingMessage
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			Label:        "A text message",
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"replyToken":"nHuyWiB7yP5Zw52FIkcQobQuGDXCTA","messages":[{"type":"text","text":"Hello, world"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A location message",
			Messages:     []SendingMessage{NewLocationMessage("title", "address", 35.65910807942215, 139.70372892916203)},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"replyToken":"nHuyWiB7yP5Zw52FIkcQobQuGDXCTA","messages":[{"type":"location","title":"title","address":"address","latitude":35.65910807942215,"longitude":139.70372892916203}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A image message",
			Messages:     []SendingMessage{NewImageMessage("https://example.com/original.jpg", "https://example.com/preview.jpg")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"replyToken":"nHuyWiB7yP5Zw52FIkcQobQuGDXCTA","messages":[{"type":"image","originalContentUrl":"https://example.com/original.jpg","previewImageUrl":"https://example.com/preview.jpg"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A sticker message",
			Messages:     []SendingMessage{NewStickerMessage("1", "1")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"replyToken":"nHuyWiB7yP5Zw52FIkcQobQuGDXCTA","messages":[{"type":"sticker","packageId":"1","stickerId":"1"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "Bad request",
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
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			res, err := client.ReplyMessage(replyToken, tc.Messages...).Do()
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

func TestReplyMessagesWithContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.ReplyMessage("nHuyWiB7yP5Zw52FIkcQobQuGDXCTA", NewTextMessage("Hello, world")).WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func TestMulticastMessages(t *testing.T) {
	toUserIDs := []string{
		"U0cc15697597f61dd8b01cea8b027050e",
		"U38ecbecfade326557b6971140741a4a6",
	}
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label        string
		Messages     []SendingMessage
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			Label:        "A text message",
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":["U0cc15697597f61dd8b01cea8b027050e","U38ecbecfade326557b6971140741a4a6"],"messages":[{"type":"text","text":"Hello, world"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A location message",
			Messages:     []SendingMessage{NewLocationMessage("title", "address", 35.65910807942215, 139.70372892916203)},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":["U0cc15697597f61dd8b01cea8b027050e","U38ecbecfade326557b6971140741a4a6"],"messages":[{"type":"location","title":"title","address":"address","latitude":35.65910807942215,"longitude":139.70372892916203}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A image message",
			Messages:     []SendingMessage{NewImageMessage("https://example.com/original.jpg", "https://example.com/preview.jpg")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":["U0cc15697597f61dd8b01cea8b027050e","U38ecbecfade326557b6971140741a4a6"],"messages":[{"type":"image","originalContentUrl":"https://example.com/original.jpg","previewImageUrl":"https://example.com/preview.jpg"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A sticker message",
			Messages:     []SendingMessage{NewStickerMessage("1", "1")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":["U0cc15697597f61dd8b01cea8b027050e","U38ecbecfade326557b6971140741a4a6"],"messages":[{"type":"sticker","packageId":"1","stickerId":"1"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "Bad request",
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
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			res, err := client.Multicast(toUserIDs, tc.Messages...).Do()
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

func TestMulticastMessagesWithContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.Multicast([]string{"U0cc15697597f61dd8b01cea8b027050e", "U38ecbecfade326557b6971140741a4a6"}, NewTextMessage("Hello, world")).WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func TestBroadcastMessages(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
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
				RequestBody: []byte(`{"messages":[{"type":"text","text":"Hello, world"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A image message
			Messages:     []SendingMessage{NewImageMessage("https://example.com/original.jpg", "https://example.com/preview.jpg")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"image","originalContentUrl":"https://example.com/original.jpg","previewImageUrl":"https://example.com/preview.jpg"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A video message
			Messages:     []SendingMessage{NewVideoMessage("https://example.com/original.mp4", "https://example.com/preview.jpg")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"video","originalContentUrl":"https://example.com/original.mp4","previewImageUrl":"https://example.com/preview.jpg"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A audio message
			Messages:     []SendingMessage{NewAudioMessage("https://example.com/original.m4a", 1000)},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"audio","originalContentUrl":"https://example.com/original.m4a","duration":1000}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A location message
			Messages:     []SendingMessage{NewLocationMessage("title", "address", 35.65910807942215, 139.70372892916203)},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"location","title":"title","address":"address","latitude":35.65910807942215,"longitude":139.70372892916203}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A sticker message
			Messages:     []SendingMessage{NewStickerMessage("1", "1")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"sticker","packageId":"1","stickerId":"1"}]}` + "\n"),
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
						NewURIAction("View detail", "https://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","thumbnailImageUrl":"https://example.com/bot/images/image.jpg","title":"Menu","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"https://example.com/page/123"}]}}]}` + "\n"),
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
				RequestBody: []byte(`{"messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","thumbnailImageUrl":"https://example.com/bot/images/image.jpg","title":"Menu","text":"Please select a date, time or datetime","actions":[{"type":"datetimepicker","label":"Date","data":"action=sel\u0026only=date","mode":"date","initial":"2017-09-01","max":"2017-09-03"},{"type":"datetimepicker","label":"Time","data":"action=sel\u0026only=time","mode":"time","max":"23:59","min":"00:00"},{"type":"datetimepicker","label":"DateTime","data":"action=sel","mode":"datetime","initial":"2017-09-01T12:00"}]}}]}` + "\n"),
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
						NewURIAction("View detail", "https://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","title":"Menu","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"https://example.com/page/123"}]}}]}` + "\n"),
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
						NewURIAction("View detail", "https://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","thumbnailImageUrl":"https://example.com/bot/images/image.jpg","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"https://example.com/page/123"}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A buttons template message without title, with image options and default action
			Messages: []SendingMessage{
				NewTemplateMessage(
					"this is a buttons template",
					NewButtonsTemplate(
						"https://example.com/bot/images/image.jpg",
						"",
						"Please select",
						NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
						NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
						NewURIAction("View detail", "https://example.com/page/123"),
					).WithImageOptions(
						"rectangle",
						"cover",
						"#FFFFFF",
					).WithDefaultAction(
						NewURIAction("View detail", "https://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","thumbnailImageUrl":"https://example.com/bot/images/image.jpg","imageAspectRatio":"rectangle","imageSize":"cover","imageBackgroundColor":"#FFFFFF","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"https://example.com/page/123"}],"defaultAction":{"type":"uri","label":"View detail","uri":"https://example.com/page/123"}}}]}` + "\n"),
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
						NewURIAction("View detail", "https://example.com/page/123"),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"template","altText":"this is a buttons template","template":{"type":"buttons","text":"Please select","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","displayText":"displayText"},{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=123","text":"text"},{"type":"uri","label":"View detail","uri":"https://example.com/page/123"}]}}]}` + "\n"),
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
				RequestBody: []byte(`{"messages":[{"type":"template","altText":"this is a confirm template","template":{"type":"confirm","text":"Are you sure?","actions":[{"type":"message","label":"Yes","text":"yes"},{"type":"message","label":"No","text":"no"}]}}]}` + "\n"),
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
							NewURIAction("View detail", "https://example.com/page/111"),
						),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"template","altText":"this is a carousel template","template":{"type":"carousel","columns":[{"thumbnailImageUrl":"https://example.com/bot/images/item1.jpg","title":"this is menu","text":"description","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=111"},{"type":"postback","label":"Add to cart","data":"action=add\u0026itemid=111"},{"type":"uri","label":"View detail","uri":"https://example.com/page/111"}]}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A carousel template message, with new image options and default action
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
							NewURIAction("View detail", "https://example.com/page/111"),
						).WithImageOptions(
							"#FFFFFF",
						).WithDefaultAction(
							NewURIAction("View detail", "http://example.com/page/123"),
						),
					).WithImageOptions("rectangle", "cover"),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"template","altText":"this is a carousel template with imageAspectRatio, imageSize and imageBackgroundColor","template":{"type":"carousel","columns":[{"thumbnailImageUrl":"https://example.com/bot/images/item1.jpg","imageBackgroundColor":"#FFFFFF","title":"this is menu","text":"description","actions":[{"type":"postback","label":"Buy","data":"action=buy\u0026itemid=111"},{"type":"postback","label":"Add to cart","data":"action=add\u0026itemid=111"},{"type":"uri","label":"View detail","uri":"https://example.com/page/111"}],"defaultAction":{"type":"uri","label":"View detail","uri":"http://example.com/page/123"}}],"imageAspectRatio":"rectangle","imageSize":"cover"}}]}` + "\n"),
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
							NewURIAction("View detail", "https://example.com/page/111"),
						),
					),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"template","altText":"this is a image carousel template","template":{"type":"image_carousel","columns":[{"imageUrl":"https://example.com/bot/images/item1.jpg","action":{"type":"uri","label":"View detail","uri":"https://example.com/page/111"}}]}}]}` + "\n"),
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
					NewURIImagemapAction("example", "https://example.com/", ImagemapArea{520, 0, 520, 1040}),
					NewMessageImagemapAction("hello", "hello", ImagemapArea{520, 0, 520, 1040}),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"imagemap","baseUrl":"https://example.com/bot/images/rm001","altText":"this is an imagemap","baseSize":{"width":1040,"height":1040},"actions":[{"type":"uri","label":"example","linkUri":"https://example.com/","area":{"x":520,"y":0,"width":520,"height":1040}},{"type":"message","label":"hello","text":"hello","area":{"x":520,"y":0,"width":520,"height":1040}}]}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A imagemap messages with video 1
			Messages: []SendingMessage{
				NewImagemapMessage(
					"https://example.com/bot/images/rm001",
					"this is an imagemap with video",
					ImagemapBaseSize{1040, 1040},
					NewURIImagemapAction("example", "https://example.com/", ImagemapArea{520, 0, 520, 1040}),
					NewMessageImagemapAction("hello", "hello", ImagemapArea{520, 0, 520, 1040}),
				).WithVideo(&ImagemapVideo{
					OriginalContentURL: "https://example.com/original.mp4",
					PreviewImageURL:    "https://example.com/preview.jpg",
					Area:               ImagemapArea{10, 10, 100, 200},
				}),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"imagemap","baseUrl":"https://example.com/bot/images/rm001","altText":"this is an imagemap with video","baseSize":{"width":1040,"height":1040},"actions":[{"type":"uri","label":"example","linkUri":"https://example.com/","area":{"x":520,"y":0,"width":520,"height":1040}},{"type":"message","label":"hello","text":"hello","area":{"x":520,"y":0,"width":520,"height":1040}}],"video":{"originalContentUrl":"https://example.com/original.mp4","previewImageUrl":"https://example.com/preview.jpg","area":{"x":10,"y":10,"width":100,"height":200}}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A imagemap messages with video 2
			Messages: []SendingMessage{
				NewImagemapMessage(
					"https://example.com/bot/images/rm001",
					"this is an imagemap with video and external link",
					ImagemapBaseSize{1040, 1040},
					NewURIImagemapAction("example", "https://example.com/", ImagemapArea{520, 0, 520, 1040}),
					NewMessageImagemapAction("hello", "hello", ImagemapArea{520, 0, 520, 1040}),
				).WithVideo(&ImagemapVideo{
					OriginalContentURL: "https://example.com/original.mp4",
					PreviewImageURL:    "https://example.com/preview.jpg",
					Area:               ImagemapArea{10, 10, 100, 200},
					ExternalLink: &ImagemapVideoExternalLink{
						LinkURI: "https://example.com/",
						Label:   "external link",
					},
				}),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"imagemap","baseUrl":"https://example.com/bot/images/rm001","altText":"this is an imagemap with video and external link","baseSize":{"width":1040,"height":1040},"actions":[{"type":"uri","label":"example","linkUri":"https://example.com/","area":{"x":520,"y":0,"width":520,"height":1040}},{"type":"message","label":"hello","text":"hello","area":{"x":520,"y":0,"width":520,"height":1040}}],"video":{"originalContentUrl":"https://example.com/original.mp4","previewImageUrl":"https://example.com/preview.jpg","area":{"x":10,"y":10,"width":100,"height":200},"externalLink":{"linkUri":"https://example.com/","label":"external link"}}}]}` + "\n"),
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
									Flex: IntPtr(0),
								},
								&SpacerComponent{
									Type: FlexComponentTypeSpacer,
								},
							},
						},
					},
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"flex","altText":"this is a flex message","contents":{"type":"bubble","body":{"type":"box","layout":"vertical","contents":[{"type":"text","text":"hello"},{"type":"text","text":"world","flex":0},{"type":"spacer"}]}}}]}` + "\n"),
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
				RequestBody: []byte(`{"messages":[{"type":"text","text":"Select your favorite food category or send me your location!","quickReply":{"items":[{"type":"action","imageUrl":"https://example.com/sushi.png","action":{"type":"message","label":"Sushi","text":"Sushi"}},{"type":"action","imageUrl":"https://example.com/tempura.png","action":{"type":"message","label":"Tempura","text":"Tempura"}},{"type":"action","action":{"type":"location","label":"Send location"}}]}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// Multiple messages
			Messages:     []SendingMessage{NewTextMessage("Hello, world1"), NewTextMessage("Hello, world2")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":"Hello, world1"},{"type":"text","text":"Hello, world2"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		// Single message with sender
		{
			Messages: []SendingMessage{
				NewAudioMessage("https://example.com/original.m4a", 1000).WithSender(
					NewSender("Cony", "https://line.me/conyprof"),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"audio","originalContentUrl":"https://example.com/original.m4a","duration":1000,"sender":{"name":"Cony","iconUrl":"https://line.me/conyprof"}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		// Single location message with sender
		{
			Messages: []SendingMessage{
				NewLocationMessage("title", "address", 35.65910807942215, 139.70372892916203).WithSender(
					NewSender("Cony", "https://line.me/conyprof"),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"location","title":"title","address":"address","latitude":35.65910807942215,"longitude":139.70372892916203,"sender":{"name":"Cony","iconUrl":"https://line.me/conyprof"}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A imagemap message with sender
			Messages: []SendingMessage{
				NewImagemapMessage(
					"https://example.com/bot/images/rm001",
					"this is an imagemap",
					ImagemapBaseSize{1040, 1040},
					NewURIImagemapAction("example", "https://example.com/", ImagemapArea{520, 0, 520, 1040}),
					NewMessageImagemapAction("hello", "hello", ImagemapArea{520, 0, 520, 1040}),
				).WithSender(
					NewSender("Cony", "https://line.me/conyprof"),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"imagemap","baseUrl":"https://example.com/bot/images/rm001","altText":"this is an imagemap","baseSize":{"width":1040,"height":1040},"actions":[{"type":"uri","label":"example","linkUri":"https://example.com/","area":{"x":520,"y":0,"width":520,"height":1040}},{"type":"message","label":"hello","text":"hello","area":{"x":520,"y":0,"width":520,"height":1040}}],"sender":{"name":"Cony","iconUrl":"https://line.me/conyprof"}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Messages: []SendingMessage{
				NewTextMessage("Hello, I am Cony!!").WithSender(
					NewSender("Cony", "https://line.me/conyprof"),
				),
				NewStickerMessage("1", "1").WithSender(
					NewSender("Cony", "https://line.me/conyprof"),
				),
				NewVideoMessage("https://example.com/original.mp4", "https://example.com/preview.jpg").WithSender(
					NewSender("Cony", "https://line.me/conyprof"),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":"Hello, I am Cony!!","sender":{"name":"Cony","iconUrl":"https://line.me/conyprof"}},{"type":"sticker","packageId":"1","stickerId":"1","sender":{"name":"Cony","iconUrl":"https://line.me/conyprof"}},{"type":"video","originalContentUrl":"https://example.com/original.mp4","previewImageUrl":"https://example.com/preview.jpg","sender":{"name":"Cony","iconUrl":"https://line.me/conyprof"}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// Multiple messages with sender
			Messages: []SendingMessage{
				NewTextMessage("Hello, I am Cony!!").WithSender(
					NewSender("Cony", "https://line.me/conyprof"),
				),
				NewStickerMessage("1", "1").WithSender(
					NewSender("Cony", "https://line.me/conyprof"),
				),
				NewVideoMessage("https://example.com/original.mp4", "https://example.com/preview.jpg").WithSender(
					NewSender("Cony", "https://line.me/conyprof"),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":"Hello, I am Cony!!","sender":{"name":"Cony","iconUrl":"https://line.me/conyprof"}},{"type":"sticker","packageId":"1","stickerId":"1","sender":{"name":"Cony","iconUrl":"https://line.me/conyprof"}},{"type":"video","originalContentUrl":"https://example.com/original.mp4","previewImageUrl":"https://example.com/preview.jpg","sender":{"name":"Cony","iconUrl":"https://line.me/conyprof"}}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// A text message with emojis
			Messages: []SendingMessage{
				NewTextMessage("$ Look at this: $ It's a LINE emoji!").AddEmoji(
					NewEmoji(0, "5ac1bfd5040ab15980c9b435", "086")).AddEmoji(
					NewEmoji(14, "5ac1bfd5040ab15980c9b435", "001"),
				),
			},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":"$ Look at this: $ It's a LINE emoji!","emojis":[{"index":0,"productId":"5ac1bfd5040ab15980c9b435","emojiId":"086"},{"index":14,"productId":"5ac1bfd5040ab15980c9b435","emojiId":"001"}]}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			// Bad request
			Messages:     []SendingMessage{NewTextMessage(""), NewTextMessage("")},
			ResponseCode: 400,
			Response:     []byte(`{"message":"Request body has 2 error(s).","details":[{"message":"may not be empty","property":"messages[0].text"},{"message":"may not be empty","property":"messages[1].text"}]}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":""},{"type":"text","text":""}]}` + "\n"),
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
		if r.URL.Path != APIEndpointBroadcastMessage {
			t.Errorf("URLPath %s; want %s", r.URL.Path, APIEndpointBroadcastMessage)
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
		res, err := client.BroadcastMessage(tc.Messages...).Do()
		if tc.Want.Error != nil {
			if !reflect.DeepEqual(err, tc.Want.Error) {
				t.Errorf("Error %d %v; want %v", i, err, tc.Want.Error)
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

func TestBroadcastMessagesWithContext(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("{}"))
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = client.BroadcastMessage(NewTextMessage("Hello, world")).WithContext(ctx).Do()
	expectCtxDeadlineExceed(ctx, err, t)
}

func TestMessagesWithNotificationDisabled(t *testing.T) {
	type testMethod interface {
		Do() (*BasicResponse, error)
	}
	toUserIDs := []string{
		"U0cc15697597f61dd8b01cea8b027050e",
		"U38ecbecfade326557b6971140741a4a6",
	}
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label        string
		TestMethod   testMethod
		Messages     []SendingMessage
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			Label:        "A text message for Push Message",
			TestMethod:   new(PushMessageCall),
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"text","text":"Hello, world"}],"notificationDisabled":true}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A text message for Reply Message",
			TestMethod:   new(ReplyMessageCall),
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"replyToken":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"text","text":"Hello, world"}],"notificationDisabled":true}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A text message for Multicast",
			TestMethod:   new(MulticastCall),
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":["U0cc15697597f61dd8b01cea8b027050e","U38ecbecfade326557b6971140741a4a6"],"messages":[{"type":"text","text":"Hello, world"}],"notificationDisabled":true}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
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
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			switch tc.TestMethod.(type) {
			case *PushMessageCall:
				res, err = client.PushMessage(toUserIDs[0], tc.Messages...).WithNotificationDisabled().Do()
			case *ReplyMessageCall: // use toUserIDs as replyToken because it doesn't matter
				res, err = client.ReplyMessage(toUserIDs[0], tc.Messages...).WithNotificationDisabled().Do()
			case *MulticastCall:
				res, err = client.Multicast(toUserIDs, tc.Messages...).WithNotificationDisabled().Do()
			}
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %d %v; want %v", i, err, tc.Want.Error)
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
		})
	}
}

func TestNarrowcastMessages(t *testing.T) {
	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label              string
		Messages           []SendingMessage
		Recipient          Recipient
		Demographic        DemographicFilter
		Max                int
		UpToRemainingQuota bool
		RequestID          string
		Response           []byte
		ResponseCode       int
		Want               want
	}{
		{
			Label:    "A text message for Narrowcast Message with Audience",
			Messages: []SendingMessage{NewTextMessage("Hello, world")},
			Recipient: RecipientOperatorAnd(
				NewAudienceObject(5614991017776),
				RecipientOperatorNot(
					NewAudienceObject(4389303728991),
				),
			),
			Demographic:  nil,
			Max:          0,
			RequestID:    "12222",
			Response:     []byte(`{}`),
			ResponseCode: 202,
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":"Hello, world"}],"recipient":{"type":"operator","and":[{"type":"audience","audienceGroupId":5614991017776},{"type":"operator","not":{"type":"audience","audienceGroupId":4389303728991}}]}}` + "\n"),
				Response:    &BasicResponse{RequestID: "12222"},
			},
		},
		{
			Label:    "A text message for Narrowcast Message with Redelivery",
			Messages: []SendingMessage{NewTextMessage("Hello, world")},
			Recipient: RecipientOperatorAnd(
				NewRedeliveryObject("f70dd685-499a-4231-a441-f24b8d4fba21"),
				RecipientOperatorNot(
					NewRedeliveryObject("x88dd9k2-gdd4-7fs0-a441-va668d4fb0x9"),
				),
			),
			Demographic:  nil,
			Max:          0,
			RequestID:    "12222",
			Response:     []byte(`{}`),
			ResponseCode: 202,
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":"Hello, world"}],"recipient":{"type":"operator","and":[{"type":"redelivery","requestId":"f70dd685-499a-4231-a441-f24b8d4fba21"},{"type":"operator","not":{"type":"redelivery","requestId":"x88dd9k2-gdd4-7fs0-a441-va668d4fb0x9"}}]}}` + "\n"),
				Response:    &BasicResponse{RequestID: "12222"},
			},
		},
		{
			Label:        "A text message for Narrowcast Message for android",
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			Recipient:    nil,
			Demographic:  NewAppTypeFilter(AppTypeAndroid),
			Max:          0,
			RequestID:    "22222",
			Response:     []byte(`{}`),
			ResponseCode: 202,
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":"Hello, world"}],"filter":{"demographic":{"type":"appType","oneOf":["android"]}}}` + "\n"),
				Response:    &BasicResponse{RequestID: "22222"},
			},
		},
		{
			Label:              "A text message for Narrowcast Message for male and age >= 30 and limit max to 10",
			Messages:           []SendingMessage{NewTextMessage("Hello, world")},
			Recipient:          nil,
			Demographic:        DemographicFilterOperatorAnd(NewGenderFilter(GenderMale), NewAgeFilter(Age30, AgeEmpty)),
			Max:                10,
			UpToRemainingQuota: true,
			RequestID:          "32222",
			Response:           []byte(`{}`),
			ResponseCode:       202,
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":"Hello, world"}],"filter":{"demographic":{"type":"operator","and":[{"type":"gender","oneOf":["male"]},{"type":"age","gte":"age_30"}]}},"limit":{"max":10,"upToRemainingQuota":true}}` + "\n"),
				Response:    &BasicResponse{RequestID: "32222"},
			},
		},
		{
			Label:    "An example message for sending narrowcast message based on official documentation",
			Messages: []SendingMessage{NewTextMessage("test message")},
			Recipient: RecipientOperatorAnd(
				NewAudienceObject(5614991017776),
				RecipientOperatorNot(
					NewAudienceObject(4389303728991),
				),
			),
			Demographic: DemographicFilterOperatorOr(
				DemographicFilterOperatorAnd(
					NewGenderFilter(GenderMale, GenderFemale),
					NewAgeFilter(Age20, Age25),
					NewAppTypeFilter(AppTypeAndroid, AppTypeIOS),
					NewAreaFilter(AreaJPAichi, AreaJPAkita),
					NewSubscriptionPeriodFilter(PeriodDay7, PeriodDay30),
				),
				DemographicFilterOperatorAnd(
					NewAgeFilter(Age35, Age40),
					DemographicFilterOperatorNot(NewGenderFilter(GenderMale)),
				),
			),
			Max:          100,
			RequestID:    "32222",
			Response:     []byte(`{}`),
			ResponseCode: 202,
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":"test message"}],"recipient":{"type":"operator","and":[{"type":"audience","audienceGroupId":5614991017776},{"type":"operator","not":{"type":"audience","audienceGroupId":4389303728991}}]},"filter":{"demographic":{"type":"operator","or":[{"type":"operator","and":[{"type":"gender","oneOf":["male","female"]},{"type":"age","gte":"age_20","lt":"age_25"},{"type":"appType","oneOf":["android","ios"]},{"type":"area","oneOf":["jp_23","jp_05"]},{"type":"subscriptionPeriod","gte":"day_7","lt":"day_30"}]},{"type":"operator","and":[{"type":"age","gte":"age_35","lt":"age_40"},{"type":"operator","not":{"type":"gender","oneOf":["male"]}}]}]}},"limit":{"max":100}}` + "\n"),
				Response:    &BasicResponse{RequestID: "32222"},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		tc := testCases[currentTestIdx]
		if !reflect.DeepEqual(body, tc.Want.RequestBody) {
			t.Errorf("RequestBody \n%s; want \n%s", body, tc.Want.RequestBody)
		}
		w.Header().Set("X-Line-Request-Id", tc.RequestID)
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
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			narrowCast := client.Narrowcast(tc.Messages...)
			if tc.Recipient != nil {
				narrowCast = narrowCast.WithRecipient(tc.Recipient)
			}
			if tc.Demographic != nil {
				narrowCast = narrowCast.WithDemographic(tc.Demographic)
			}
			if tc.Max > 0 {
				narrowCast = narrowCast.WithLimitMax(tc.Max)
			}
			if tc.UpToRemainingQuota {
				narrowCast = narrowCast.WithLimitMaxUpToRemainingQuota(tc.Max, tc.UpToRemainingQuota)
			}
			res, err := narrowCast.Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %d %v; want %v", i, err, tc.Want.Error)
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
		})
	}
}

func BenchmarkPushMessages(b *testing.B) {
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
		client.PushMessage("U0cc15697597f61dd8b01cea8b027050e", NewTextMessage("Hello, world")).Do()
	}
}

func TestMessagesWithRetryKey(t *testing.T) {
	type testMethod interface {
		Do() (*BasicResponse, error)
	}
	toUserIDs := []string{
		"U0cc15697597f61dd8b01cea8b027050e",
		"U38ecbecfade326557b6971140741a4a6",
	}
	var msgUUIDs = "123e4567-e89b-12d3-a456-426655440002"

	type want struct {
		RequestBody []byte
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label        string
		TestMethod   testMethod
		Messages     []SendingMessage
		Response     []byte
		ResponseCode int
		Want         want
	}{
		{
			Label:        "A text message for Push Message",
			TestMethod:   new(PushMessageCall),
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":"U0cc15697597f61dd8b01cea8b027050e","messages":[{"type":"text","text":"Hello, world"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A text message for Multicast",
			TestMethod:   new(MulticastCall),
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"to":["U0cc15697597f61dd8b01cea8b027050e","U38ecbecfade326557b6971140741a4a6"],"messages":[{"type":"text","text":"Hello, world"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A text message for Narrowcast",
			TestMethod:   new(NarrowcastCall),
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":"Hello, world"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
		{
			Label:        "A text message for Broadcast",
			TestMethod:   new(BroadcastMessageCall),
			Messages:     []SendingMessage{NewTextMessage("Hello, world")},
			ResponseCode: 200,
			Response:     []byte(`{}`),
			Want: want{
				RequestBody: []byte(`{"messages":[{"type":"text","text":"Hello, world"}]}` + "\n"),
				Response:    &BasicResponse{},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
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
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			switch tc.TestMethod.(type) {
			case *PushMessageCall:
				res, err = client.PushMessage(toUserIDs[0], tc.Messages...).WithRetryKey(msgUUIDs).Do()
			case *MulticastCall:
				res, err = client.Multicast(toUserIDs, tc.Messages...).WithRetryKey(msgUUIDs).Do()
			case *NarrowcastCall:
				res, err = client.Narrowcast(tc.Messages...).WithRetryKey(msgUUIDs).Do()
			case *BroadcastMessageCall:
				res, err = client.BroadcastMessage(tc.Messages...).WithRetryKey(msgUUIDs).Do()
			}
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %d %v; want %v", i, err, tc.Want.Error)
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
		})
	}
}

func BenchmarkReplyMessages(b *testing.B) {
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
		client.ReplyMessage("nHuyWiB7yP5Zw52FIkcQobQuGDXCTA", NewTextMessage("Hello, world")).Do()
	}
}

func BenchmarkMulticast(b *testing.B) {
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
		client.Multicast([]string{"U0cc15697597f61dd8b01cea8b027050e", "U38ecbecfade326557b6971140741a4a6"}, NewTextMessage("Hello, world")).Do()
	}
}

func BenchmarkBroadcastMessage(b *testing.B) {
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
		client.BroadcastMessage(NewTextMessage("Hello, world")).Do()
	}
}
