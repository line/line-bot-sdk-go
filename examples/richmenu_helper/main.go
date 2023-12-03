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
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func main() {
	var (
		mode       = flag.String("mode", "list", "mode of richmenu helper [list|create|link|unlink|bulklink|bulkunlink|get|delete|upload|download|alias_create|alias_get|alias_update|alias_delete|alias_list]")
		aid        = flag.String("aid", "", "alias id")
		uid        = flag.String("uid", "", "user id")
		richMenuId = flag.String("richMenuId", "", "richmenu id")
		filePath   = flag.String("image.path", "", "path to image, used in upload/download mode")
	)
	flag.Parse()

	client, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	blob_client, err := messaging_api.NewMessagingApiBlobAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	switch *mode {
	case "upload":
		if *richMenuId == "" {
			log.Fatal("richMenuId is required")
		}

		file, err := os.Open(*filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var contentType string
		if strings.HasSuffix(*filePath, ".jpeg") || strings.HasSuffix(*filePath, ".jpg") {
			contentType = "image/jpeg"
		} else if strings.HasSuffix(*filePath, "png") {
			contentType = "image/png"
		} else {
			log.Fatal("image file must be jpeg or png... but got ", *filePath)
		}

		if _, err = blob_client.SetRichMenuImage(*richMenuId, contentType, file); err != nil {
			log.Fatal(err)
		}
	case "download":
		if *richMenuId == "" {
			log.Fatal("richMenuId is required")
		}
		if *filePath == "" {
			log.Fatal("filePath is required")
		}

		res, err := blob_client.GetRichMenuImage(*richMenuId)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		f, err := os.OpenFile(*filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(f, res.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Image is written to %s", *filePath)
	case "alias_create":
		if _, err = client.CreateRichMenuAlias(
			&messaging_api.CreateRichMenuAliasRequest{
				RichMenuAliasId: *aid,
				RichMenuId:      *richMenuId,
			},
		); err != nil {
			log.Fatal(err)
		}
	case "alias_get":
		if res, err := client.GetRichMenuAlias(*aid); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%v\n", res)
		}
	case "alias_update":
		if _, err = client.UpdateRichMenuAlias(*aid, &messaging_api.UpdateRichMenuAliasRequest{
			RichMenuId: *richMenuId,
		}); err != nil {
			log.Fatal(err)
		}
	case "alias_delete":
		if _, err = client.DeleteRichMenuAlias(*aid); err != nil {
			log.Fatal(err)
		}
	case "alias_list":
		res, err := client.GetRichMenuAliasList()
		if err != nil {
			log.Fatal(err)
		}
		for _, alias := range res.Aliases {
			log.Printf("%v\n", alias)
		}
	case "link":
		if _, err = client.LinkRichMenuIdToUser(*uid, *richMenuId); err != nil {
			log.Fatal(err)
		}
	case "unlink":
		if _, err = client.UnlinkRichMenuIdFromUser(*uid); err != nil {
			log.Fatal(err)
		}
	case "bulklink":
		if _, err = client.LinkRichMenuIdToUsers(
			&messaging_api.RichMenuBulkLinkRequest{
				RichMenuId: *richMenuId,
				UserIds:    []string{*uid},
			}); err != nil {
			log.Fatal(err)
		}
	case "bulkunlink":
		if _, err = client.UnlinkRichMenuIdFromUsers(
			&messaging_api.RichMenuBulkUnlinkRequest{
				UserIds: []string{*uid},
			},
		); err != nil {
			log.Fatal(err)
		}
	case "list":
		res, err := client.GetRichMenuList()
		if err != nil {
			log.Fatal(err)
		}
		for _, richmenu := range res.Richmenus {
			log.Printf("%v\n", richmenu)
		}
	case "get_default":
		if *richMenuId == "" {
			log.Fatal("richMenuId is required")
		}
		res, err := client.GetDefaultRichMenuId()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v\n", res)
	case "set_default":
		if _, err = client.SetDefaultRichMenu(*richMenuId); err != nil {
			log.Fatal(err)
		}
	case "cancel_default": // TODO
		if _, err = client.CancelDefaultRichMenu(); err != nil {
			log.Fatal(err)
		}
	case "create":
		richMenu := &messaging_api.RichMenuRequest{
			Size:        &messaging_api.RichMenuSize{Width: 2500, Height: 1686},
			Selected:    false,
			Name:        "Menu1",
			ChatBarText: "ChatText",
			Areas: []messaging_api.RichMenuArea{
				{
					Bounds: &messaging_api.RichMenuBounds{X: 0, Y: 0, Width: 1250, Height: 212},
					Action: &messaging_api.RichMenuSwitchAction{
						RichMenuAliasId: "richmenu-alias-a",
						Data:            "action=richmenu-changed-to-a",
					},
				},
				{
					Bounds: &messaging_api.RichMenuBounds{X: 1250, Y: 0, Width: 1250, Height: 212},
					Action: &messaging_api.RichMenuSwitchAction{
						RichMenuAliasId: "richmenu-alias-b",
						Data:            "action=richmenu-changed-to-b",
					},
				},
				{
					Bounds: &messaging_api.RichMenuBounds{X: 0, Y: 212, Width: 1250, Height: 737},
					Action: &messaging_api.PostbackAction{
						Data: "action=buy&itemid=123",
					},
				},
				{
					Bounds: &messaging_api.RichMenuBounds{X: 1250, Y: 212, Width: 1250, Height: 737},
					Action: &messaging_api.UriAction{
						Uri:   "https://developers.line.me/",
						Label: "click me",
					},
				},
				{
					Bounds: &messaging_api.RichMenuBounds{X: 0, Y: 949, Width: 1250, Height: 737},
					Action: &messaging_api.MessageAction{
						Text: "hello world!",
					},
				},
				{
					Bounds: &messaging_api.RichMenuBounds{X: 1250, Y: 949, Width: 1250, Height: 737},
					Action: &messaging_api.DatetimePickerAction{
						Data: "datetime picker!",
						Mode: "datetime",
					},
				},
			},
		}
		res, err := client.CreateRichMenu(richMenu)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.RichMenuId)
	default:
		log.Fatal("implement me")
	}
}
