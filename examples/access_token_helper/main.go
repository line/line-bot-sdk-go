// Copyright 2020 LINE Corporation
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
		mode            = flag.String("mode", "list", "mode of access token v2 helper [issue|list|revoke]")
		channelID       = flag.String("channel_id", os.Getenv("CHANNEL_ID"), "Channel ID on channel console")
		accessToken     = flag.String("access_token", os.Getenv("ACCESS_TOKEN"), "channel access token")
		clientAssertion = flag.String("client_assertion", os.Getenv("CLIENT_ASSERTION"), "A JSON Web Token the client needs to create and sign with the private key")
	)
	checkErr := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	flag.Parse()
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	checkErr(err)

	switch *mode {
	case "issue":
		res, err := bot.IssueAccessTokenV2(*clientAssertion).Do()
		checkErr(err)
		log.Printf("%v", res)
	case "list":
		res, err := bot.GetAccessTokensV2(*clientAssertion).Do()
		checkErr(err)
		log.Printf("%v", res)
	case "revoke":
		res, err := bot.RevokeAccessTokenV2(*channelID, os.Getenv("CHANNEL_SECRET"), *accessToken).Do()
		checkErr(err)
		log.Printf("%v", res)
	default:
		log.Fatal("implement me")
	}
}
