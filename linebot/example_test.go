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

package linebot_test

import (
	"fmt"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func ExampleIDsScanner() {
	bot, err := linebot.New("secret", "token")
	if err != nil {
		fmt.Fprintln(os.Stderr, "linebot.New:", err)
	}
	s := bot.GetGroupMemberIDs("cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "").NewScanner()
	for s.Scan() {
		fmt.Fprintln(os.Stdout, s.ID())
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "GetGroupMemberIDs:", err)
	}
}

func ExampleUnmarshalFlexMessageJSON() {
	container, err := linebot.UnmarshalFlexMessageJSON([]byte(`{
    "type": "bubble",
    "body": {
      "type": "box",
      "layout": "vertical",
      "contents": [
        {
          "type": "text",
          "text": "hello"
        },
        {
          "type": "text",
          "text": "world"
        }
      ]
    }
  }`))
	if err != nil {
		panic(err)
	}

	linebot.NewFlexMessage("alt text", container)
}
