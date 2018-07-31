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
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	app, err := NewKitchenSink(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
		os.Getenv("APP_BASE_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// serve /static/** files
	staticFileServer := http.FileServer(http.Dir("static"))
	http.HandleFunc("/static/", http.StripPrefix("/static/", staticFileServer).ServeHTTP)
	// serve /downloaded/** files
	downloadedFileServer := http.FileServer(http.Dir(app.downloadDir))
	http.HandleFunc("/downloaded/", http.StripPrefix("/downloaded/", downloadedFileServer).ServeHTTP)

	http.HandleFunc("/callback", app.Callback)
	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

// KitchenSink app
type KitchenSink struct {
	bot         *linebot.Client
	appBaseURL  string
	downloadDir string
}

// NewKitchenSink function
func NewKitchenSink(channelSecret, channelToken, appBaseURL string) (*KitchenSink, error) {
	apiEndpointBase := os.Getenv("ENDPOINT_BASE")
	if apiEndpointBase == "" {
		apiEndpointBase = linebot.APIEndpointBase
	}
	bot, err := linebot.New(
		channelSecret,
		channelToken,
		linebot.WithEndpointBase(apiEndpointBase), // Usually you omit this.
	)
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
		bot:         bot,
		appBaseURL:  appBaseURL,
		downloadDir: downloadDir,
	}, nil
}

// Callback function for http server
func (app *KitchenSink) Callback(w http.ResponseWriter, r *http.Request) {
	events, err := app.bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		log.Printf("Got event %v", event)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := app.handleText(message, event.ReplyToken, event.Source); err != nil {
					log.Print(err)
				}
			case *linebot.ImageMessage:
				if err := app.handleImage(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.VideoMessage:
				if err := app.handleVideo(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.AudioMessage:
				if err := app.handleAudio(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.FileMessage:
				if err := app.handleFile(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				if err := app.handleLocation(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.StickerMessage:
				if err := app.handleSticker(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			default:
				log.Printf("Unknown message: %v", message)
			}
		case linebot.EventTypeFollow:
			if err := app.replyText(event.ReplyToken, "Got followed event"); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeUnfollow:
			log.Printf("Unfollowed this bot: %v", event)
		case linebot.EventTypeJoin:
			if err := app.replyText(event.ReplyToken, "Joined "+string(event.Source.Type)); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeLeave:
			log.Printf("Left: %v", event)
		case linebot.EventTypePostback:
			data := event.Postback.Data
			if data == "DATE" || data == "TIME" || data == "DATETIME" {
				data += fmt.Sprintf("(%v)", *event.Postback.Params)
			}
			if err := app.replyText(event.ReplyToken, "Got postback: "+data); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeBeacon:
			if err := app.replyText(event.ReplyToken, "Got beacon: "+event.Beacon.Hwid); err != nil {
				log.Print(err)
			}
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}

func (app *KitchenSink) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	switch message.Text {
	case "profile":
		if source.UserID != "" {
			profile, err := app.bot.GetProfile(source.UserID).Do()
			if err != nil {
				return app.replyText(replyToken, err.Error())
			}
			if _, err := app.bot.ReplyMessage(
				replyToken,
				linebot.NewTextMessage("Display name: "+profile.DisplayName),
				linebot.NewTextMessage("Status message: "+profile.StatusMessage),
			).Do(); err != nil {
				return err
			}
		} else {
			return app.replyText(replyToken, "Bot can't use profile API without user ID")
		}
	case "buttons":
		imageURL := app.appBaseURL + "/static/buttons/1040.jpg"
		template := linebot.NewButtonsTemplate(
			imageURL, "My button sample", "Hello, my button",
			linebot.NewURIAction("Go to line.me", "https://line.me"),
			linebot.NewPostbackAction("Say hello1", "hello こんにちは", "", "hello こんにちは"),
			linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
			linebot.NewMessageAction("Say message", "Rice=米"),
		)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Buttons alt text", template),
		).Do(); err != nil {
			return err
		}
	case "confirm":
		template := linebot.NewConfirmTemplate(
			"Do it?",
			linebot.NewMessageAction("Yes", "Yes!"),
			linebot.NewMessageAction("No", "No!"),
		)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Confirm alt text", template),
		).Do(); err != nil {
			return err
		}
	case "carousel":
		imageURL := app.appBaseURL + "/static/buttons/1040.jpg"
		template := linebot.NewCarouselTemplate(
			linebot.NewCarouselColumn(
				imageURL, "hoge", "fuga",
				linebot.NewURIAction("Go to line.me", "https://line.me"),
				linebot.NewPostbackAction("Say hello1", "hello こんにちは", "", ""),
			),
			linebot.NewCarouselColumn(
				imageURL, "hoge", "fuga",
				linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
				linebot.NewMessageAction("Say message", "Rice=米"),
			),
		)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Carousel alt text", template),
		).Do(); err != nil {
			return err
		}
	case "image carousel":
		imageURL := app.appBaseURL + "/static/buttons/1040.jpg"
		template := linebot.NewImageCarouselTemplate(
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewURIAction("Go to LINE", "https://line.me"),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewPostbackAction("Say hello1", "hello こんにちは", "", ""),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewMessageAction("Say message", "Rice=米"),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewDatetimePickerAction("datetime", "DATETIME", "datetime", "", "", ""),
			),
		)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Image carousel alt text", template),
		).Do(); err != nil {
			return err
		}
	case "datetime":
		template := linebot.NewButtonsTemplate(
			"", "", "Select date / time !",
			linebot.NewDatetimePickerAction("date", "DATE", "date", "", "", ""),
			linebot.NewDatetimePickerAction("time", "TIME", "time", "", "", ""),
			linebot.NewDatetimePickerAction("datetime", "DATETIME", "datetime", "", "", ""),
		)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Datetime pickers alt text", template),
		).Do(); err != nil {
			return err
		}
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
		contents := &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Body: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "Hello,",
					},
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "World!",
					},
				},
			},
		}
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewFlexMessage("Flex message alt text", contents),
		).Do(); err != nil {
			return err
		}
	case "flex carousel":
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
		contents := &linebot.CarouselContainer{
			Type: linebot.FlexContainerTypeCarousel,
			Contents: []*linebot.BubbleContainer{
				&linebot.BubbleContainer{
					Type: linebot.FlexContainerTypeBubble,
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "First bubble",
							},
						},
					},
				},
				&linebot.BubbleContainer{
					Type: linebot.FlexContainerTypeBubble,
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "Second bubble",
							},
						},
					},
				},
			},
		}
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewFlexMessage("Flex message alt text", contents),
		).Do(); err != nil {
			return err
		}
	case "flex json":
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
          "uri": "https://linecorp.com"
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
		contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))
		if err != nil {
			return err
		}
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewFlexMessage("Flex message alt text", contents),
		).Do(); err != nil {
			return err
		}
	case "imagemap":
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewImagemapMessage(
				app.appBaseURL+"/static/rich",
				"Imagemap alt text",
				linebot.ImagemapBaseSize{1040, 1040},
				linebot.NewURIImagemapAction("https://store.line.me/family/manga/en", linebot.ImagemapArea{0, 0, 520, 520}),
				linebot.NewURIImagemapAction("https://store.line.me/family/music/en", linebot.ImagemapArea{520, 0, 520, 520}),
				linebot.NewURIImagemapAction("https://store.line.me/family/play/en", linebot.ImagemapArea{0, 520, 520, 520}),
				linebot.NewMessageImagemapAction("URANAI!", linebot.ImagemapArea{520, 520, 520, 520}),
			),
		).Do(); err != nil {
			return err
		}
	case "quick":
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage("Select your favorite food category or send me your location!").
				WithQuickReplies(linebot.NewQuickReplyItems(
					linebot.NewQuickReplyButton(
						app.appBaseURL+"/static/quick/sushi.png",
						linebot.NewMessageAction("Sushi", "Sushi")),
					linebot.NewQuickReplyButton(
						app.appBaseURL+"/static/quick/tempura.png",
						linebot.NewMessageAction("Tempura", "Tempura")),
					linebot.NewQuickReplyButton(
						"",
						linebot.NewLocationAction("Send location")),
				)),
		).Do(); err != nil {
			return err
		}
	case "bye":
		switch source.Type {
		case linebot.EventSourceTypeUser:
			return app.replyText(replyToken, "Bot can't leave from 1:1 chat")
		case linebot.EventSourceTypeGroup:
			if err := app.replyText(replyToken, "Leaving group"); err != nil {
				return err
			}
			if _, err := app.bot.LeaveGroup(source.GroupID).Do(); err != nil {
				return app.replyText(replyToken, err.Error())
			}
		case linebot.EventSourceTypeRoom:
			if err := app.replyText(replyToken, "Leaving room"); err != nil {
				return err
			}
			if _, err := app.bot.LeaveRoom(source.RoomID).Do(); err != nil {
				return app.replyText(replyToken, err.Error())
			}
		}
	default:
		log.Printf("Echo message to %s: %s", replyToken, message.Text)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(message.Text),
		).Do(); err != nil {
			return err
		}
	}
	return nil
}

func (app *KitchenSink) handleImage(message *linebot.ImageMessage, replyToken string) error {
	return app.handleHeavyContent(message.ID, func(originalContent *os.File) error {
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
			replyToken,
			linebot.NewImageMessage(originalContentURL, previewImageURL),
		).Do(); err != nil {
			return err
		}
		return nil
	})
}

func (app *KitchenSink) handleVideo(message *linebot.VideoMessage, replyToken string) error {
	return app.handleHeavyContent(message.ID, func(originalContent *os.File) error {
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
			replyToken,
			linebot.NewVideoMessage(originalContentURL, previewImageURL),
		).Do(); err != nil {
			return err
		}
		return nil
	})
}

func (app *KitchenSink) handleAudio(message *linebot.AudioMessage, replyToken string) error {
	return app.handleHeavyContent(message.ID, func(originalContent *os.File) error {
		originalContentURL := app.appBaseURL + "/downloaded/" + filepath.Base(originalContent.Name())
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewAudioMessage(originalContentURL, 100),
		).Do(); err != nil {
			return err
		}
		return nil
	})
}

func (app *KitchenSink) handleFile(message *linebot.FileMessage, replyToken string) error {
	return app.replyText(replyToken, fmt.Sprintf("File `%s` (%d bytes) received.", message.FileName, message.FileSize))
}

func (app *KitchenSink) handleLocation(message *linebot.LocationMessage, replyToken string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewLocationMessage(message.Title, message.Address, message.Latitude, message.Longitude),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *KitchenSink) handleSticker(message *linebot.StickerMessage, replyToken string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewStickerMessage(message.PackageID, message.StickerID),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *KitchenSink) replyText(replyToken, text string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *KitchenSink) handleHeavyContent(messageID string, callback func(*os.File) error) error {
	content, err := app.bot.GetMessageContent(messageID).Do()
	if err != nil {
		return err
	}
	defer content.Content.Close()
	log.Printf("Got file: %s", content.ContentType)
	originalConent, err := app.saveContent(content.Content)
	if err != nil {
		return err
	}
	return callback(originalConent)
}

func (app *KitchenSink) saveContent(content io.ReadCloser) (*os.File, error) {
	file, err := ioutil.TempFile(app.downloadDir, "")
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
