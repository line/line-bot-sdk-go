# line-bot-sdk-go

[![Build Status](https://travis-ci.org/line/line-bot-sdk-go.svg?branch=master)](https://travis-ci.org/line/line-bot-sdk-go)

Go SDK for the LINE Messaging API


## About LINE Messaging API

See the official API documentation for more information.

English: https://devdocs.line.me/en/<br>
Japanese: https://devdocs.line.me/ja/

## Installation ##

```sh
$ go get github.com/line/line-bot-sdk-go/linebot
```

## Configuration ##

```go
import (
	"github.com/line/line-bot-sdk-go/linebot"
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
	leftBtn := linebot.NewMessageTemplateAction("left", "left clicked")
	rightBtn := linebot.NewMessageTemplateAction("right", "right clicked")

	template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)

	messgage := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
```

### Send message ###

With an ID, you can send message using ```PushMessage()```

```go
	var messages []linebot.Message

	// append some message to messages

	_, err := bot.PushMessage(ID, messages...).Do()
	if err != nil {
		// Do something when some bad happened
	}
```

With a reply token, you can reply to messages using ```ReplyMessage()```

```go
	var messages []linebot.Message

	// append some message to messages

	_, err := bot.ReplyMessage(replyToken, messages...).Do()
	if err != nil {
		// Do something when some bad happened
	}
```

## Requirements

This library requires Go 1.6 or later.

## LICENSE

See LICENSE.txt
