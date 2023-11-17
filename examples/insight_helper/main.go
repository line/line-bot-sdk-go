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

	"github.com/line/line-bot-sdk-go/v8/linebot/insight"
)

func main() {
	var (
		mode = flag.String("mode", "reply", "mode of insight helper [messages|followers|demographics]")
		date = flag.String("date", "", "date the messages were sent, format 'yyyyMMdd'")
	)
	flag.Parse()
	client, err := insight.NewInsightAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	switch *mode {
	case "messages":
		res, err := client.GetNumberOfMessageDeliveries(*date)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v", res)
	case "followers":
		res, err := client.GetNumberOfFollowers(*date)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v", res)
	case "demographics":
		res, err := client.GetFriendsDemographics()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v", res)

	default:
		log.Fatal("implement me")
	}
}
