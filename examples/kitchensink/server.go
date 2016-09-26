package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/line/line-bot-sdk-go-v2/linebot"
)

func main() {
	app, err := NewKitchenSink(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
		os.Getenv("ENDPOINT_BASE"),
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
func NewKitchenSink(channelSecret, channelToken, apiEndpointBase, appBaseURL string) (*KitchenSink, error) {
	bot, err := linebot.New(channelSecret, channelToken, linebot.WithEndpointBase(apiEndpointBase))
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
	log.Printf("Got events %v", events)
	for _, event := range events {
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
			// TODO
		case linebot.EventTypeUnfollow:
			// TODO
		case linebot.EventTypeJoin:
			// TODO
		case linebot.EventTypeLeave:
			// TODO
		case linebot.EventTypePostback:
			if err := app.replyText(event.ReplyToken, "Got postback: "+event.Postback.Data); err != nil {
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
			if _, err := app.bot.Reply(
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
			linebot.NewPostbackAction("Say hello1", "hello こんにちは", ""),
			linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは"),
			linebot.NewMessageAction("Say message", "Rice=米"),
		)
		if _, err := app.bot.Reply(
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
		if _, err := app.bot.Reply(
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
				linebot.NewPostbackAction("Say hello1", "hello こんにちは", ""),
			),
			linebot.NewCarouselColumn(
				imageURL, "hoge", "fuga",
				linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは"),
				linebot.NewMessageAction("Say message", "Rice=米"),
			),
		)
		if _, err := app.bot.Reply(
			replyToken,
			linebot.NewTemplateMessage("Carousel alt text", template),
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
		if _, err := app.bot.Reply(
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
		previewImagePath := originalContent.Name() + "-preview.jpg"
		_, err := exec.Command("convert", "-resize", "240x", originalContent.Name(), previewImagePath).Output()
		if err != nil {
			return err
		}

		originalContentURL := app.appBaseURL + "/downloaded/" + filepath.Base(originalContent.Name())
		previewImageURL := app.appBaseURL + "/downloaded/" + filepath.Base(previewImagePath)
		if _, err := app.bot.Reply(
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
		originalVideoPath := originalContent.Name() + ".mp4"
		if err := os.Rename(originalContent.Name(), originalVideoPath); err != nil {
			return err
		}
		previewImagePath := originalContent.Name() + "-preview.jpg"
		_, err := exec.Command("convert", originalVideoPath+"[0]", previewImagePath).Output()
		if err != nil {
			return err
		}

		originalContentURL := app.appBaseURL + "/downloaded/" + filepath.Base(originalVideoPath)
		previewImageURL := app.appBaseURL + "/downloaded/" + filepath.Base(previewImagePath)
		if _, err := app.bot.Reply(
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
		if _, err := app.bot.Reply(
			replyToken,
			linebot.NewAudioMessage(originalContentURL, 100),
		).Do(); err != nil {
			return err
		}
		return nil
	})
}

func (app *KitchenSink) handleLocation(message *linebot.LocationMessage, replyToken string) error {
	if _, err := app.bot.Reply(
		replyToken,
		linebot.NewLocationMessage(message.Title, message.Address, message.Latitude, message.Longitude),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *KitchenSink) handleSticker(message *linebot.StickerMessage, replyToken string) error {
	if _, err := app.bot.Reply(
		replyToken,
		linebot.NewStickerMessage(message.PackageID, message.StickerID),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *KitchenSink) replyText(replyToken, text string) error {
	if _, err := app.bot.Reply(
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
	log.Printf("Got filename: %s", content.FileName)
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
