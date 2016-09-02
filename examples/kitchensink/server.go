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
	bot, err := linebot.NewClient(channelSecret, channelToken, linebot.WithEndpointBase(apiEndpointBase))
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
			var message linebot.Message
			if message, err = event.Message(); err != nil {
				log.Print(err)
				continue
			}
			if message.Type() == linebot.EventMessageTypeText {
				// TODO
				app.handleText(message, event.ReplyToken)
			}
			if message.Type() == linebot.EventMessageTypeImage {
				// TODO
			}
		}
	}

}

func (app *KitchenSink) handleText(message linebot.Message, replyToken string) {
	if tm, ok := message.(*linebot.TextMessage); ok {
		switch tm.Text {
		// TODO
		case "":
			break
		default:
			log.Printf("echo message to %s: %s", replyToken, tm.Text)
			app.bot.Reply(replyToken, []linebot.Message{
				linebot.NewTextMessage(tm.Text),
			})
		}
	}
}
