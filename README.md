# LINE Messaging API SDK for Go

[![Build Status](https://travis-ci.org/line/line-bot-sdk-go.svg?branch=master)](https://travis-ci.org/line/line-bot-sdk-go)
[![codecov](https://codecov.io/gh/line/line-bot-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/line/line-bot-sdk-go)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/line/line-bot-sdk-go/linebot)
[![Go Report Card](https://goreportcard.com/badge/github.com/line/line-bot-sdk-go)](https://goreportcard.com/report/github.com/line/line-bot-sdk-go)


## Introduction
The LINE Messaging API SDK for Go makes it easy to develop bots using LINE Messaging API, and you can create a sample bot within minutes.

## Documentation

See the official API documentation for more information.

- English: https://developers.line.biz/en/docs/messaging-api/overview/
- Japanese: https://developers.line.biz/ja/docs/messaging-api/overview/

## Requirements

This library requires Go 1.11 or later.

## Installation ##

```sh
$ go get -u github.com/line/line-bot-sdk-go/v7/linebot
```

## Configuration ##

```go
import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	bot, err := linebot.New("<channel secret>", "<channel access token>")
	...
}

```

### Configuration with http.Client ###

```go
client := &http.Client{}
bot, err := linebot.New("<channel secret>", "<channel accsss token>", linebot.WithHTTPClient(client))
...
```

## How to start ##

The LINE Messaging API uses the JSON data format.
```ParseRequest()``` will help you to parse the ```*http.Request``` content and return a slice of Pointer point to Event Object.

```go
events, err := bot.ParseRequest(req)
if err != nil {
	// Do something when something bad happened.
}
```

The LINE Messaging API defines 7 types of event - ```EventTypeMessage```, ```EventTypeFollow```, ```EventTypeUnfollow```, ```EventTypeJoin```, ```EventTypeLeave```, ```EventTypePostback```, ```EventTypeBeacon```. You can check the event type by using ```event.Type```

```go
for _, event := range events {
	if event.Type == linebot.EventTypeMessage {
		// Do Something...
	}
}
```

### Receiver ###

To send a message to a user, group, or room, you need either an ID

```go
userID := event.Source.UserID
groupID := event.Source.GroupID
RoomID := event.Source.RoomID
```

or a reply token.

```go
replyToken := event.ReplyToken
```

### Create message ###

The LINE Messaging API provides various types of message. To create a message, use ```New<Type>Message()```.

```go
leftBtn := linebot.NewMessageAction("left", "left clicked")
rightBtn := linebot.NewMessageAction("right", "right clicked")

template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)

message := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
```

### Send message ###

With an ID, you can send message using ```PushMessage()```

```go
var messages []linebot.SendingMessage

// append some message to messages

_, err := bot.PushMessage(ID, messages...).Do()
if err != nil {
	// Do something when some bad happened
}
```

With a reply token, you can reply to messages using ```ReplyMessage()```

```go
var messages []linebot.SendingMessage

// append some message to messages

_, err := bot.ReplyMessage(replyToken, messages...).Do()
if err != nil {
	// Do something when some bad happened
}
```

## Help and media

FAQ: https://developers.line.biz/en/faq/

Community Q&A: https://www.line-community.me/questions

News: https://developers.line.biz/en/news/

Twitter: @LINE_DEV


## Versioning
This project respects semantic versioning.

See http://semver.org/


## Contributing

Please check [CONTRIBUTING](CONTRIBUTING.md) before making a contribution.


## License

```
Copyright (C) 2016 LINE Corp.
 
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 
   http://www.apache.org/licenses/LICENSE-2.0
 
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
