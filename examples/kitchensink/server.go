package main

import (
	"github.com/line/line-bot-sdk-go-v2/linebot"
	"log"
	"net/http"
	"os"
)

func main() {
	app, err := NewKitchenSink(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
		os.Getenv("ENDPOINT_BASE"),
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", app.Callback)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

// KitchenSink app
type KitchenSink struct {
	bot *linebot.Client
}

// NewKitchenSink function
func NewKitchenSink(channelSecret, channelToken, apiEndpointBase string) (*KitchenSink, error) {
	bot, err := linebot.New(channelSecret, channelToken, linebot.WithEndpointBase(apiEndpointBase))
	if err != nil {
		return nil, err
	}
	return &KitchenSink{
		bot: bot,
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
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := app.handleText(message, event.ReplyToken, event.Source); err != nil {
					log.Print(err)
					continue
				}
				break
			case *linebot.LocationMessage:
				if err := app.handleLocation(message, event.ReplyToken, event.Source); err != nil {
					log.Print(err)
					continue
				}
				break
			case *linebot.StickerMessage:
				if err := app.handleSticker(message, event.ReplyToken, event.Source); err != nil {
					log.Print(err)
					continue
				}
			default:
				// TODO
				break
			}
		}
	}
}

func (app *KitchenSink) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	switch message.Text {
	case "profile":
		if source.UserID != "" {
			var profile *linebot.UserProfileResponse
			profile, err := app.bot.GetUserProfile(source.UserID).Do()
			if err != nil {
				return app.replyText(replyToken, err.Error())
			}
			messages := []linebot.Message{
				linebot.NewTextMessage("Display name: " + profile.DisplayName),
				linebot.NewTextMessage("Status message: " + profile.StatusMessage),
			}
			if _, err := app.bot.Reply(replyToken, messages).Do(); err != nil {
				return err
			}
		} else {
			return app.replyText(replyToken, "Bot can't use profile API without user ID")
		}
		break
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
			if err := app.replyText(replyToken, "Leaving group"); err != nil {
				return err
			}
			if _, err := app.bot.LeaveRoom(source.RoomID).Do(); err != nil {
				return app.replyText(replyToken, err.Error())
			}
		}
		break
	default:
		log.Printf("echo message to %s: %s", replyToken, message.Text)
		messages := []linebot.Message{
			linebot.NewTextMessage(message.Text),
		}
		if _, err := app.bot.Reply(replyToken, messages).Do(); err != nil {
			return err
		}
		break
	}
	return nil
}

func (app *KitchenSink) handleLocation(message *linebot.LocationMessage, replyToken string, source *linebot.EventSource) error {
	messages := []linebot.Message{
		linebot.NewLocationMessage(message.Title, message.Address, message.Latitude, message.Longitude),
	}
	if _, err := app.bot.Reply(replyToken, messages).Do(); err != nil {
		return err
	}
	return nil
}

func (app *KitchenSink) handleSticker(message *linebot.StickerMessage, replyToken string, source *linebot.EventSource) error {
	messages := []linebot.Message{
		linebot.NewStickerMessage(message.PackageID, message.StickerID),
	}
	if _, err := app.bot.Reply(replyToken, messages).Do(); err != nil {
		return err
	}
	return nil
}

func (app *KitchenSink) replyText(replyToken, text string) error {
	if _, err := app.bot.Reply(replyToken, []linebot.Message{linebot.NewTextMessage(text)}).Do(); err != nil {
		return err
	}
	return nil
}
