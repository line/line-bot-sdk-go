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
			m, err := event.Message()
			if err != nil {
				log.Print(err)
				continue
			}
			switch message := m.(type) {
			case *linebot.TextMessage:
				if err := app.handleText(message, event.ReplyToken, event.Source); err != nil {
					log.Print(err)
					continue
				}
				break
			case *linebot.ImageMessage:
				// TODO
				break
			}
		}
	}
}

func (app *KitchenSink) handleText(message *linebot.TextMessage, replyToken string, source linebot.ReceivedEventSource) error {
	switch message.Text {
	case "profile":
		if source.UserID != "" {
			var profile *linebot.UserProfileResponse
			profile, err := app.bot.GetUserProfile(source.UserID)
			if err != nil {
				return err
			}
			messages := []linebot.Message{
				linebot.NewTextMessage("Display name: " + profile.DisplayName),
				linebot.NewTextMessage("Status message: " + profile.StatusMessage),
			}
			if _, err := app.bot.Reply(replyToken, messages); err != nil {
				return err
			}
		}
		break
	// TODO
	case "":
	default:
		log.Printf("echo message to %s: %s", replyToken, message.Text)
		messages := []linebot.Message{
			linebot.NewTextMessage(message.Text),
		}
		if _, err := app.bot.Reply(replyToken, messages); err != nil {
			return err
		}
	}
	return nil
}
