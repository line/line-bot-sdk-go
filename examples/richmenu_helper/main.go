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
	"flag"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	var (
		mode     = flag.String("mode", "list", "mode of richmenu helper [list|create|link|unlink|get|delete|upload|download]")
		uid      = flag.String("uid", "", "user id")
		rid      = flag.String("rid", "", "richmenu id")
		filePath = flag.String("image.path", "", "path to image, used in upload/download mode")
	)
	flag.Parse()
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	switch *mode {
	case "upload":
		if _, err = bot.UploadRichMenuImage(*rid, *filePath).Do(); err != nil {
			log.Fatal(err)
		}
	case "link":
		if _, err = bot.LinkUserRichMenu(*uid, *rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "unlink":
		if _, err = bot.UnlinkUserRichMenu(*uid).Do(); err != nil {
			log.Fatal(err)
		}
	case "list":
		res, err := bot.GetRichMenuList().Do()
		if err != nil {
			log.Fatal(err)
		}
		for _, richmenu := range res {
			log.Printf("%v\n", richmenu)
		}
	case "create":
		richMenu := linebot.RichMenu{
			Size:        linebot.RichMenuSize{Width: 2500, Height: 1686},
			Selected:    false,
			Name:        "Menu1",
			ChatBarText: "ChatText",
			Areas: []linebot.AreaDetail{
				linebot.AreaDetail{
					Bounds: linebot.RichMenuBounds{X: 0, Y: 0, Width: 1250, Height: 843},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypePostback,
						Data: "action=buy&itemid=123",
					},
				},
				linebot.AreaDetail{
					Bounds: linebot.RichMenuBounds{X: 1250, Y: 0, Width: 1250, Height: 843},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeURI,
						URI:  "https://developers.line.me/",
						Text: "click me",
					},
				},
				linebot.AreaDetail{
					Bounds: linebot.RichMenuBounds{X: 0, Y: 843, Width: 1250, Height: 843},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeMessage,
						Text: "hello world!",
					},
				},
				linebot.AreaDetail{
					Bounds: linebot.RichMenuBounds{X: 1250, Y: 843, Width: 1250, Height: 843},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeDatetimePicker,
						Data: "datetime picker!",
						Mode: "datetime",
					},
				},
			},
		}
		res, err := bot.CreateRichMenu(richMenu).Do()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.RichMenuID)
	default:
		log.Fatal("implement me")
	}
}
