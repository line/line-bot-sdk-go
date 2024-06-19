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

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/line/line-bot-sdk-go/v8/util"
)

func main() {
	app, err := NewKitchenSink(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
		os.Getenv("APP_BASE_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// serve /static/** files
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	staticDir := filepath.Join(dir, "static")
	log.Printf("Serving static content from %s\n", staticDir)
	staticFileServer := http.FileServer(http.Dir(staticDir))
	http.HandleFunc("/static/", http.StripPrefix("/static/", staticFileServer).ServeHTTP)
	// serve /downloaded/** files
	downloadedFileServer := http.FileServer(http.Dir(app.downloadDir))
	http.HandleFunc("/downloaded/", http.StripPrefix("/downloaded/", downloadedFileServer).ServeHTTP)

	http.HandleFunc("/callback", app.Callback)
	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	fmt.Println("http://localhost:" + port + "/")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// KitchenSink app
type KitchenSink struct {
	channelSecret string
	bot           *messaging_api.MessagingApiAPI
	blob          *messaging_api.MessagingApiBlobAPI
	appBaseURL    string
	downloadDir   string
}

// NewKitchenSink function
func NewKitchenSink(channelSecret, channelToken, appBaseURL string) (*KitchenSink, error) {
	if appBaseURL == "" {
		return nil, fmt.Errorf("missing appBaseURL")
	}
	bot, err := messaging_api.NewMessagingApiAPI(channelToken)
	if err != nil {
		return nil, err
	}
	blob, err := messaging_api.NewMessagingApiBlobAPI(channelToken)
	if err != nil {
		return nil, err
	}
	downloadDir := filepath.Join(filepath.Dir(os.Args[0]), "line-bot")
	_, err = os.Stat(downloadDir)
	if err != nil {
		if err := os.Mkdir(downloadDir, 0777); err != nil {
			return nil, err
		}
	}
	return &KitchenSink{
		channelSecret: channelSecret,
		bot:           bot,
		blob:          blob,
		appBaseURL:    appBaseURL,
		downloadDir:   downloadDir,
	}, nil
}

// Callback function for http server
func (app *KitchenSink) Callback(w http.ResponseWriter, r *http.Request) {
	cb, err := webhook.ParseRequest(app.channelSecret, r)
	if err != nil {
		if errors.Is(err, webhook.ErrInvalidSignature) {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range cb.Events {
		log.Printf("Got event %v", event)
		switch e := event.(type) {
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				if err := app.handleText(&message, e.ReplyToken, e.Source); err != nil {
					log.Print(err)
				}
			case webhook.ImageMessageContent:
				if err := app.handleImage(&message, e.ReplyToken); err != nil {
					log.Print(err)
				}
			case webhook.VideoMessageContent:
				if err := app.handleVideo(&message, e.ReplyToken); err != nil {
					log.Print(err)
				}
			case webhook.AudioMessageContent:
				if err := app.handleAudio(&message, e.ReplyToken); err != nil {
					log.Print(err)
				}
			case webhook.FileMessageContent:
				if err := app.handleFile(&message, e.ReplyToken); err != nil {
					log.Print(err)
				}
			case webhook.LocationMessageContent: // TODO
				if err := app.handleLocation(&message, e.ReplyToken); err != nil {
					log.Print(err)
				}
			case webhook.StickerMessageContent:
				if err := app.handleSticker(&message, e.ReplyToken); err != nil {
					log.Print(err)
				}
			default:
				log.Printf("Unknown message: %v", message)
			}
		case webhook.FollowEvent:
			if err := app.replyText(e.ReplyToken, "Got followed event"); err != nil {
				log.Print(err)
			}
		case webhook.UnfollowEvent:
			log.Printf("Unfollowed this bot: %v", event)
		case webhook.JoinEvent:
			if err := app.replyText(e.ReplyToken, "Joined "+string(e.Source.GetType())); err != nil {
				log.Print(err)
			}
		case webhook.LeaveEvent:
			log.Printf("Left: %v", e)
		case webhook.PostbackEvent:
			data := e.Postback.Data
			if data == "DATE" || data == "TIME" || data == "DATETIME" {
				data += fmt.Sprintf("(%v)", e.Postback.Params)
			}
			if err := app.replyText(e.ReplyToken, "Got postback: "+data); err != nil {
				log.Print(err)
			}
		case webhook.BeaconEvent:
			if err := app.replyText(e.ReplyToken, "Got beacon: "+e.Beacon.Hwid); err != nil {
				log.Print(err)
			}
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}

func (app *KitchenSink) handleText(message *webhook.TextMessageContent, replyToken string, source webhook.SourceInterface) error {
	switch message.Text {
	case "profile":
		switch s := source.(type) {
		case webhook.UserSource:
			profile, err := app.bot.GetProfile(s.UserId)
			if err != nil {
				return app.replyText(replyToken, err.Error())
			}
			if _, err := app.bot.ReplyMessage(
				&messaging_api.ReplyMessageRequest{
					ReplyToken: replyToken,
					Messages: []messaging_api.MessageInterface{
						messaging_api.TextMessage{
							Text: "Display name: " + profile.DisplayName,
						},
						messaging_api.TextMessage{
							Text: "Status message: " + profile.StatusMessage,
						},
					},
				},
			); err != nil {
				return err
			}
		default:
			return app.replyText(replyToken, "Bot can't use profile API without user ID")
		}
	case "buttons":
		imageURL := app.appBaseURL + "/static/buttons/1040.jpg"
		template := &messaging_api.ButtonsTemplate{
			ThumbnailImageUrl: imageURL,
			Title:             "My button sample",
			Text:              "Hello",
			Actions: []messaging_api.ActionInterface{
				&messaging_api.UriAction{
					Label: "Go to line.me",
					Uri:   "https://line.me",
				},
				&messaging_api.PostbackAction{
					Label: "Say hello1",
					Data:  "hello こんにちは",
					Text:  "hello こんにちは",
				},
				&messaging_api.PostbackAction{
					Label: "言 hello2",
					Data:  "hello こんにちは",
					Text:  "hello こんにちは",
				},
				&messaging_api.MessageAction{
					Label: "Say message",
					Text:  "Rice=米",
				},
			},
		}
		if _, err := app.bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					&messaging_api.TemplateMessage{
						AltText:  "Buttons alt text",
						Template: template,
					},
				},
			},
		); err != nil {
			return err
		}
	case "confirm":
		if _, err := app.bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					&messaging_api.TemplateMessage{
						AltText: "Confirm alt text",
						Template: &messaging_api.ConfirmTemplate{
							Text: "Do it?",
							Actions: []messaging_api.ActionInterface{
								&messaging_api.MessageAction{
									Label: "Yes",
									Text:  "Yes!",
								},
								&messaging_api.MessageAction{
									Label: "No",
									Text:  "No!",
								},
							},
						},
					},
				},
			},
		); err != nil {
			return err
		}
	case "carousel":
		imageURL := app.appBaseURL + "/static/buttons/1040.jpg"
		template := messaging_api.CarouselTemplate{
			Columns: []messaging_api.CarouselColumn{
				{
					ThumbnailImageUrl: imageURL,
					Title:             "hoge",
					Text:              "fuga",
					DefaultAction: messaging_api.UriAction{
						Label: "Go to line.me",
						Uri:   "https://line.me",
					},
					Actions: []messaging_api.ActionInterface{
						messaging_api.PostbackAction{
							Label: "Say hello1",
							Data:  "hello こんにちは",
						},
					},
				},
				{
					ThumbnailImageUrl: imageURL,
					Title:             "hoge",
					Text:              "fuga",
					DefaultAction: messaging_api.PostbackAction{
						Label: "言 hello2",
						Text:  "hello こんにちは",
						Data:  "hello こんにちは",
					},
					Actions: []messaging_api.ActionInterface{
						messaging_api.MessageAction{
							Label: "Say message",
							Text:  "Rice=米",
						},
					},
				},
			},
		}
		if _, err := app.bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					&messaging_api.TemplateMessage{
						AltText:  "Carousel alt text",
						Template: template,
					},
				},
			},
		); err != nil {
			return err
		}
	case "image carousel":
		imageURL := app.appBaseURL + "/static/buttons/1040.jpg"
		_, err := app.bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					&messaging_api.TemplateMessage{
						AltText: "Image carousel alt text",
						Template: &messaging_api.ImageCarouselTemplate{
							Columns: []messaging_api.ImageCarouselColumn{
								{
									ImageUrl: imageURL,
									Action: messaging_api.UriAction{
										Label: "Go to LINE",
										Uri:   "https://line.me",
									},
								},
								{
									ImageUrl: imageURL,
									Action:   messaging_api.PostbackAction{Label: "Say hello1", Data: "hello こんにちは"},
								},
								{
									ImageUrl: imageURL,
									Action:   messaging_api.MessageAction{Label: "Say message", Text: "Rice=米"},
								},
								{
									ImageUrl: imageURL,
									Action: messaging_api.DatetimePickerAction{
										Label: "datetime",
										Data:  "DATETIME",
										Mode:  messaging_api.DatetimePickerActionMODE_DATETIME,
									},
								},
							},
						},
					},
				},
			},
		)
		return err
	case "datetime":
		result, err := app.bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					&messaging_api.TemplateMessage{
						AltText: "Datetime pickers alt text",
						Template: &messaging_api.ButtonsTemplate{
							Text: "Select date / time !",
							Actions: []messaging_api.ActionInterface{
								&messaging_api.DatetimePickerAction{
									Label: "date",
									Data:  "DATE",
									Mode:  messaging_api.DatetimePickerActionMODE_DATE,
								},
								&messaging_api.DatetimePickerAction{
									Label: "time",
									Data:  "TIME",
									Mode:  messaging_api.DatetimePickerActionMODE_TIME,
								},
								&messaging_api.DatetimePickerAction{
									Label: "datetime",
									Data:  "DATETIME",
									Mode:  messaging_api.DatetimePickerActionMODE_DATETIME,
								},
							},
						},
					},
				},
			},
		)
		if err == nil {
			log.Printf("Sent reply: %v", result)
		}
		log.Printf("Sent reply: %v %v", result, err)
		return err
	case "flex":
		// {
		//   "type": "bubble",
		//   "body": {
		//     "type": "box",
		//     "layout": "horizontal",
		//     "contents": [
		//       {
		//         "type": "text",
		//         "text": "Hello,"
		//       },
		//       {
		//         "type": "text",
		//         "text": "World!"
		//       }
		//     ]
		//   }
		// }
		if _, err := app.bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					&messaging_api.FlexMessage{
						AltText: "Flex message alt text",
						Contents: messaging_api.FlexBubble{
							Body: &messaging_api.FlexBox{
								Layout: messaging_api.FlexBoxLAYOUT_HORIZONTAL,
								Contents: []messaging_api.FlexComponentInterface{
									&messaging_api.FlexText{
										Text: "Hello,",
									},
									&messaging_api.FlexText{
										Text: "World!",
									},
								},
							},
						},
					},
				},
			},
		); err != nil {
			return err
		}
	case "flex carousel":
		err := handleFlexCarousel(app, replyToken)
		if err != nil {
			return err
		}
	case "flex json": // TODO
		err2 := handleFlexJson(app, replyToken)
		if err2 != nil {
			return err2
		}
	case "imagemap":
		if _, err := app.bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					messaging_api.ImagemapMessage{
						BaseUrl:  app.appBaseURL + "/static/rich",
						AltText:  "Imagemap alt text",
						BaseSize: &messaging_api.ImagemapBaseSize{Width: 1040, Height: 1040},
						Actions: []messaging_api.ImagemapActionInterface{
							&messaging_api.UriImagemapAction{
								Label:   "LINE Store Manga",
								LinkUri: "https://store.line.me/family/manga/en",
								Area:    &messaging_api.ImagemapArea{X: 0, Y: 0, Width: 520, Height: 520},
							},
							&messaging_api.UriImagemapAction{
								Label:   "LINE Store Music",
								LinkUri: "https://store.line.me/family/music/en",
								Area:    &messaging_api.ImagemapArea{X: 520, Y: 0, Width: 520, Height: 520},
							},
							&messaging_api.UriImagemapAction{
								Label:   "LINE Store Play",
								LinkUri: "https://store.line.me/family/play/en",
								Area:    &messaging_api.ImagemapArea{X: 0, Y: 520, Width: 520, Height: 520},
							},
							&messaging_api.MessageImagemapAction{
								Label: "URANAI!",
								Text:  "URANAI!",
								Area:  &messaging_api.ImagemapArea{X: 520, Y: 520, Width: 520, Height: 520},
							},
						},
					},
				},
			},
		); err != nil {
			return err
		}
	case "imagemap video":
		err := handleImagemapVideo(app, replyToken)
		if err != nil {
			return err
		}
	case "quick":
		err := handleQuickReply(app, replyToken)
		if err != nil {
			return err
		}
	case "bye":
		switch s := source.(type) {
		case webhook.UserSource:
			return app.replyText(replyToken, "Bot can't leave from 1:1 chat")
		case webhook.GroupSource:
			if err := app.replyText(replyToken, "Leaving group"); err != nil {
				return err
			}
			if _, err := app.bot.LeaveGroup(s.GroupId); err != nil {
				return app.replyText(replyToken, err.Error())
			}
		case webhook.RoomSource:
			if err := app.replyText(replyToken, "Leaving room"); err != nil {
				return err
			}
			if _, err := app.bot.LeaveRoom(s.RoomId); err != nil {
				return app.replyText(replyToken, err.Error())
			}
		}
	case "with http info":
		resp, _, _ := app.bot.ReplyMessageWithHttpInfo(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					messaging_api.TextMessage{
						Text: "Hello, world",
					},
				},
			},
		)
		log.Printf("status code: (%v), x-line-request-id: (%v)", resp.StatusCode, resp.Header.Get("x-line-request-id"))
	case "with http info error":
		resp, _, err := app.bot.ReplyMessageWithHttpInfo(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken + "invalid",
				Messages: []messaging_api.MessageInterface{
					messaging_api.TextMessage{
						Text: "Hello, world",
					},
				},
			},
		)
		if err != nil && resp.StatusCode >= 400 && resp.StatusCode < 500 {
			decoder := json.NewDecoder(resp.Body)
			errorResponse := &messaging_api.ErrorResponse{}
			if err := decoder.Decode(&errorResponse); err != nil {
				log.Fatal("failed to decode JSON: %w", err)
			}
			log.Printf("status code: (%v), x-line-request-id: (%v), error response: (%v)", resp.StatusCode, resp.Header.Get("x-line-request-id"), errorResponse)
		}
	case "emoji":
		message := "Hello, $ hello こんにちは $, สวัสดีครับ $"
		emojiIndexes := util.FindDollarSignIndexInUTF16Text(message)
		emojis := []messaging_api.Emoji{}
		for _, index := range emojiIndexes {
			emojis = append(emojis, messaging_api.Emoji{
				Index:     int32(index),
				ProductId: "5ac1bfd5040ab15980c9b435",
				EmojiId:   "001",
			})
		}
		result, _, err := app.bot.ReplyMessageWithHttpInfo(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					messaging_api.TextMessage{
						Text:   message,
						Emojis: emojis,
					},
				},
			},
		)
		if err == nil {
			log.Printf("Sent reply: %v", result)
		}
		log.Printf("Sent reply: %v %v", result, err)
		return err
	default:
		log.Printf("Echo message to %s: %s", replyToken, message.Text)
		if _, err := app.bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					messaging_api.TextMessage{
						Text: message.Text,
					},
				},
			},
		); err != nil {
			return err
		}
	}
	return nil
}

func handleQuickReply(app *KitchenSink, replyToken string) error {
	msg := &messaging_api.TextMessage{
		Text: "Select your favorite food category or send me your location!",
		QuickReply: &messaging_api.QuickReply{
			Items: []messaging_api.QuickReplyItem{
				{
					ImageUrl: app.appBaseURL + "/static/quick/sushi.png",
					Action: &messaging_api.MessageAction{
						Label: "Sushi",
						Text:  "Sushi",
					},
				},
				{
					ImageUrl: app.appBaseURL + "/static/quick/tempura.png",
					Action: &messaging_api.MessageAction{
						Label: "Tempura",
						Text:  "Tempura",
					},
				},
				{
					Action: &messaging_api.LocationAction{
						Label: "Send location",
					},
				},
				{
					Action: &messaging_api.UriAction{
						Label: "LINE Developer",
						Uri:   "https://developers.line.biz/",
					},
				},
			},
		},
	}
	if _, err := app.bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages:   []messaging_api.MessageInterface{msg},
		},
	); err != nil {
		return err
	}
	return nil
}

func handleImagemapVideo(app *KitchenSink, replyToken string) error {
	if _, err := app.bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				messaging_api.ImagemapMessage{
					BaseUrl: app.appBaseURL + "/static/rich",
					AltText: "Imagemap with video alt text",
					BaseSize: &messaging_api.ImagemapBaseSize{
						Width: 1040, Height: 1040,
					},
					Actions: []messaging_api.ImagemapActionInterface{
						&messaging_api.UriImagemapAction{
							Label:   "LINE Store Manga",
							LinkUri: "https://store.line.me/family/manga/en",
							Area:    &messaging_api.ImagemapArea{X: 0, Y: 0, Width: 520, Height: 520},
						},
						&messaging_api.UriImagemapAction{
							Label:   "LINE Store Music",
							LinkUri: "https://store.line.me/family/music/en",
							Area:    &messaging_api.ImagemapArea{X: 520, Y: 0, Width: 520, Height: 520},
						},
						&messaging_api.UriImagemapAction{
							Label:   "LINE Store Play",
							LinkUri: "https://store.line.me/family/play/en",
							Area:    &messaging_api.ImagemapArea{X: 0, Y: 520, Width: 520, Height: 520},
						},
						&messaging_api.MessageImagemapAction{
							Text:  "URANAI!",
							Label: "URANAI!",
							Area:  &messaging_api.ImagemapArea{X: 520, Y: 520, Width: 520, Height: 520},
						},
					},
					Video: &messaging_api.ImagemapVideo{
						OriginalContentUrl: app.appBaseURL + "/static/imagemap/video.mp4",
						PreviewImageUrl:    app.appBaseURL + "/static/imagemap/preview.jpg",
						Area:               &messaging_api.ImagemapArea{X: 280, Y: 385, Width: 480, Height: 270},
						ExternalLink:       &messaging_api.ImagemapExternalLink{LinkUri: "https://line.me", Label: "LINE"},
					},
				},
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func handleFlexJson(app *KitchenSink, replyToken string) error {
	jsonString := `{
  "type": "bubble",
  "hero": {
    "type": "image",
    "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
    "size": "full",
    "aspectRatio": "20:13",
    "aspectMode": "cover",
    "action": {
      "type": "uri",
      "uri": "http://linecorp.com/"
    }
  },
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "text",
        "text": "Brown Cafe",
        "weight": "bold",
        "size": "xl"
      },
      {
        "type": "box",
        "layout": "baseline",
        "margin": "md",
        "contents": [
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gray_star_28.png"
          },
          {
            "type": "text",
            "text": "4.0",
            "size": "sm",
            "color": "#999999",
            "margin": "md",
            "flex": 0
          }
        ]
      },
      {
        "type": "box",
        "layout": "vertical",
        "margin": "lg",
        "spacing": "sm",
        "contents": [
          {
            "type": "box",
            "layout": "baseline",
            "spacing": "sm",
            "contents": [
              {
                "type": "text",
                "text": "Place",
                "color": "#aaaaaa",
                "size": "sm",
                "flex": 1
              },
              {
                "type": "text",
                "text": "Miraina Tower, 4-1-6 Shinjuku, Tokyo",
                "wrap": true,
                "color": "#666666",
                "size": "sm",
                "flex": 5
              }
            ]
          },
          {
            "type": "box",
            "layout": "baseline",
            "spacing": "sm",
            "contents": [
              {
                "type": "text",
                "text": "Time",
                "color": "#aaaaaa",
                "size": "sm",
                "flex": 1
              },
              {
                "type": "text",
                "text": "10:00 - 23:00",
                "wrap": true,
                "color": "#666666",
                "size": "sm",
                "flex": 5
              }
            ]
          }
        ]
      }
    ]
  },
  "footer": {
    "type": "box",
    "layout": "vertical",
    "spacing": "sm",
    "contents": [
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "CALL",
          "uri": "https://linecorp.com"
        }
      },
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "WEBSITE",
          "uri": "https://linecorp.com",
          "altUri": {
            "desktop": "https://line.me/ja/download"
          }
        }
      },
      {
        "type": "spacer",
        "size": "sm"
      }
    ],
    "flex": 0
  }
}`
	contents, err := messaging_api.UnmarshalFlexContainer([]byte(jsonString))
	if err != nil {
		return err
	}
	if _, err := app.bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				&messaging_api.FlexMessage{
					AltText:  "Flex message alt text",
					Contents: contents,
				},
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func handleFlexCarousel(app *KitchenSink, replyToken string) error {
	// {
	//   "type": "carousel",
	//   "contents": [
	//     {
	//       "type": "bubble",
	//       "body": {
	//         "type": "box",
	//         "layout": "vertical",
	//         "contents": [
	//           {
	//             "type": "text",
	//             "text": "First bubble"
	//           }
	//         ]
	//       }
	//     },
	//     {
	//       "type": "bubble",
	//       "body": {
	//         "type": "box",
	//         "layout": "vertical",
	//         "contents": [
	//           {
	//             "type": "text",
	//             "text": "Second bubble"
	//           }
	//         ]
	//       }
	//     }
	//   ]
	// }
	contents := &messaging_api.FlexCarousel{
		Contents: []messaging_api.FlexBubble{
			{
				Body: &messaging_api.FlexBox{
					Layout: messaging_api.FlexBoxLAYOUT_VERTICAL,
					Contents: []messaging_api.FlexComponentInterface{
						&messaging_api.FlexText{
							Text: "First bubble",
						},
					},
				},
			},
			{
				Body: &messaging_api.FlexBox{
					Layout: messaging_api.FlexBoxLAYOUT_VERTICAL,
					Contents: []messaging_api.FlexComponentInterface{
						&messaging_api.FlexText{
							Text: "Second bubble",
						},
					},
				},
			},
		},
	}
	if _, err := app.bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{&messaging_api.FlexMessage{
				Contents: contents,
				AltText:  "Flex message alt text",
			}},
		},
	); err != nil {
		return err
	}
	return nil
}

func handleDatetime(app *KitchenSink, replyToken string) error {
	result, err := app.bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				&messaging_api.TemplateMessage{
					AltText: "Datetime pickers alt text",
					Template: &messaging_api.ButtonsTemplate{
						Text: "Select date / time !",
						Actions: []messaging_api.ActionInterface{
							&messaging_api.DatetimePickerAction{
								Label: "date",
								Data:  "DATE",
								Mode:  messaging_api.DatetimePickerActionMODE_DATE,
							},
							&messaging_api.DatetimePickerAction{
								Label: "time",
								Data:  "TIME",
								Mode:  messaging_api.DatetimePickerActionMODE_TIME,
							},
							&messaging_api.DatetimePickerAction{
								Label: "datetime",
								Data:  "DATETIME",
								Mode:  messaging_api.DatetimePickerActionMODE_DATETIME,
							},
						},
					},
				},
			},
		},
	)
	if err == nil {
		log.Printf("Sent reply: %v", result)
	}
	log.Printf("Sent reply: %v %v", result, err)
	return err
}

func (app *KitchenSink) handleImage(message *webhook.ImageMessageContent, replyToken string) error {
	return app.handleHeavyContent(message.Id, func(originalContent *os.File) error {
		// You need to install ImageMagick.
		// And you should consider about security and scalability.
		previewImagePath := originalContent.Name() + "-preview"
		_, err := exec.Command("convert", "-resize", "240x", "jpeg:"+originalContent.Name(), "jpeg:"+previewImagePath).Output()
		if err != nil {
			return err
		}

		originalContentURL := app.appBaseURL + "/downloaded/" + filepath.Base(originalContent.Name())
		previewImageURL := app.appBaseURL + "/downloaded/" + filepath.Base(previewImagePath)
		if _, err := app.bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					&messaging_api.ImageMessage{
						OriginalContentUrl: originalContentURL,
						PreviewImageUrl:    previewImageURL,
					},
				},
			},
		); err != nil {
			return err
		}
		return nil
	})
}

func (app *KitchenSink) handleVideo(message *webhook.VideoMessageContent, replyToken string) error {
	return app.handleHeavyContent(message.Id, func(originalContent *os.File) error {
		// You need to install FFmpeg and ImageMagick.
		// And you should consider about security and scalability.
		previewImagePath := originalContent.Name() + "-preview"
		_, err := exec.Command("convert", "mp4:"+originalContent.Name()+"[0]", "jpeg:"+previewImagePath).Output()
		if err != nil {
			return err
		}

		originalContentURL := app.appBaseURL + "/downloaded/" + filepath.Base(originalContent.Name())
		previewImageURL := app.appBaseURL + "/downloaded/" + filepath.Base(previewImagePath)
		if _, err := app.bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					&messaging_api.VideoMessage{
						OriginalContentUrl: originalContentURL,
						PreviewImageUrl:    previewImageURL,
					},
				},
			},
		); err != nil {
			return err
		}
		return nil
	})
}

func (app *KitchenSink) handleAudio(message *webhook.AudioMessageContent, replyToken string) error {
	return app.handleHeavyContent(message.Id, func(originalContent *os.File) error {
		originalContentURL := app.appBaseURL + "/downloaded/" + filepath.Base(originalContent.Name())

		if _, err := app.bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					&messaging_api.AudioMessage{
						OriginalContentUrl: originalContentURL,
						Duration:           100,
					},
				},
			},
		); err != nil {
			return err
		}
		return nil
	})
}

func (app *KitchenSink) handleFile(message *webhook.FileMessageContent, replyToken string) error {
	return app.replyText(replyToken, fmt.Sprintf("File `%s` (%d bytes) received.", message.FileName, message.FileSize))
}

func (app *KitchenSink) handleLocation(message *webhook.LocationMessageContent, replyToken string) error {
	log.Printf("Got location: %#v", message)

	_, err := app.bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				&messaging_api.LocationMessage{
					Title:     message.Title,
					Address:   message.Address,
					Latitude:  message.Latitude,
					Longitude: message.Longitude,
				},
			},
		},
	)
	return err
}

func (app *KitchenSink) handleSticker(message *webhook.StickerMessageContent, replyToken string) error {
	if _, err := app.bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				&messaging_api.StickerMessage{
					PackageId: message.PackageId,
					StickerId: message.StickerId,
				},
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func (app *KitchenSink) replyText(replyToken, text string) error {
	if _, err := app.bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				&messaging_api.TextMessage{
					Text: text,
				},
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func (app *KitchenSink) handleHeavyContent(messageID string, callback func(*os.File) error) error {
	content, err := app.blob.GetMessageContent(messageID)
	if err != nil {
		return err
	}
	defer content.Body.Close()
	originalContent, err := app.saveContent(content.Body)
	if err != nil {
		return err
	}
	return callback(originalContent)
}

func (app *KitchenSink) saveContent(content io.ReadCloser) (*os.File, error) {
	file, err := os.CreateTemp(app.downloadDir, "")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return nil, err
	}
	log.Printf("Saved %s", file.Name())
	return file, nil
}
