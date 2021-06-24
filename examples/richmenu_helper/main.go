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
	"io"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	var (
		mode     = flag.String("mode", "list", "mode of richmenu helper [list|create|link|unlink|bulklink|bulkunlink|get|delete|upload|download|alias_create|alias_get|alias_update|alias_delete|alias_list]")
		aid      = flag.String("aid", "", "alias id")
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
	case "download":
		res, err := bot.DownloadRichMenuImage(*rid).Do()
		if err != nil {
			log.Fatal(err)
		}
		defer res.Content.Close()
		f, err := os.OpenFile(*filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(f, res.Content)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Image is written to %s", *filePath)
	case "alias_create":
		if _, err = bot.CreateRichMenuAlias(*aid, *rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "alias_get":
		if res, err := bot.GetRichMenuAlias(*aid).Do(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%v\n", res)
		}
	case "alias_update":
		if _, err = bot.UpdateRichMenuAlias(*aid, *rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "alias_delete":
		if _, err = bot.DeleteRichMenuAlias(*aid).Do(); err != nil {
			log.Fatal(err)
		}
	case "alias_list":
		res, err := bot.GetRichMenuAliasList().Do()
		if err != nil {
			log.Fatal(err)
		}
		for _, alias := range res {
			log.Printf("%v\n", alias)
		}
	case "link":
		if _, err = bot.LinkUserRichMenu(*uid, *rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "unlink":
		if _, err = bot.UnlinkUserRichMenu(*uid).Do(); err != nil {
			log.Fatal(err)
		}
	case "bulklink":
		if _, err = bot.BulkLinkRichMenu(*rid, *uid).Do(); err != nil {
			log.Fatal(err)
		}
	case "bulkunlink":
		if _, err = bot.BulkUnlinkRichMenu(*uid).Do(); err != nil {
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
	case "get_default":
		res, err := bot.GetDefaultRichMenu().Do()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v\n", res)
	case "set_default":
		if _, err = bot.SetDefaultRichMenu(*rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "cancel_default":
		if _, err = bot.CancelDefaultRichMenu().Do(); err != nil {
			log.Fatal(err)
		}
	case "create":
		richMenu := linebot.RichMenu{
			Size:        linebot.RichMenuSize{Width: 2500, Height: 1686},
			Selected:    false,
			Name:        "Menu1",
			ChatBarText: "ChatText",
			Areas: []linebot.AreaDetail{
				{
					Bounds: linebot.RichMenuBounds{X: 0, Y: 0, Width: 1250, Height: 843},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypePostback,
						Data: "action=buy&itemid=123",
					},
				},
				{
					Bounds: linebot.RichMenuBounds{X: 1250, Y: 0, Width: 1250, Height: 843},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeURI,
						URI:  "https://developers.line.me/",
						Text: "click me",
					},
				},
				{
					Bounds: linebot.RichMenuBounds{X: 0, Y: 843, Width: 1250, Height: 843},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeMessage,
						Text: "hello world!",
					},
				},
				{
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
