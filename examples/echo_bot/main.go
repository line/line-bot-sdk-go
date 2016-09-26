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
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	os.Exit(Main())
}

// Main function for echo bot
func Main() int {
	var (
		channelID     int64
		channelSecret = os.Getenv("CHANNEL_SECRET")
		channelMID    = os.Getenv("CHANNEL_MID")
		err           error
	)

	// Setup bot client
	channelID, err = strconv.ParseInt(os.Getenv("CHANNEL_ID"), 10, 64)
	if err != nil {
		log.Print(err)
		return 1
	}
	bot, err := linebot.NewClient(channelID, channelSecret, channelMID)
	if err != nil {
		log.Print(err)
		return 1
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		received, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, result := range received.Results {
			content := result.Content()
			if content != nil && content.IsMessage && content.ContentType == linebot.ContentTypeText {
				text, err := content.TextContent()
				_, err = bot.SendText([]string{content.From}, text.Text)
				if err != nil {
					log.Print(err)
				}
			}
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Print(err)
		return 1
	}
	return 0
}
