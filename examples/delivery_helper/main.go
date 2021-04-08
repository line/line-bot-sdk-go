// Copyright 2019 LINE Corporation
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

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	var (
		mode = flag.String("mode", "reply", "mode of delivery helper [multicast|reply|push]")
		date = flag.String("date", "", "date the messages were sent, format 'yyyyMMdd'")
	)
	flag.Parse()
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	var res *linebot.MessagesNumberResponse
	switch *mode {
	case "multicast":
		res, err = bot.GetNumberMulticastMessages(*date).Do()
	case "push":
		res, err = bot.GetNumberPushMessages(*date).Do()
	case "reply":
		res, err = bot.GetNumberReplyMessages(*date).Do()
	case "broadcast":
		res, err = bot.GetNumberBroadcastMessages(*date).Do()
	default:
		log.Fatal("implement me")
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", res)
}
