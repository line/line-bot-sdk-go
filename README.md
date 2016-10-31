# line-bot-sdk-go

[![Build Status](https://travis-ci.org/line/line-bot-sdk-go.svg?branch=master)](https://travis-ci.org/line/line-bot-sdk-go)

Go SDK for the LINE Messaging API


## About LINE Messaging API

Please refer to the official api documents for details.

en:  https://devdocs.line.me/en/

ja:  https://devdocs.line.me/ja/


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

## How to Start ##

Line Messaging API use JSON to format data.
```ParseRequest()``` will help you to parse the ```*http.Request``` content and return a slice of Event Object.

```go
	events := bot.ParseRequest(req)
```

Line Messaging API define 7 types of event - ```EventTypeMessage```, ```EventTypeFollow```, ```EventTypeUnfollow```, ```EventTypeJoin``` , ```EventTypeLeave``` , ```EventTypePostback```, ```EventTypeBeacon```. You can check this by  ```event.Type```

```go
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			// Do Something...
		}
	}
```

### Receiver ###

To send a message to user/group/room, you might need an ID

```go
	userID := event.Source.UserID
	groupID := event.Source.GroupID
	RoomID := event.Source.RoomID
```

or reply token

```go
	replyToken := event.ReplyToken
```

### Create Message ###

Line Messaging API supply many types of message, just use the ```New<Type>Message()``` to create it.

```go
	leftBtn := linebot.NewMessageTemplateAction("left", "left clicked")
	rightBtn := linebot.NewMessageTemplateAction("right", "right clicked")

	template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)

	messgage := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
```
### Send Message ###

With ID, you can send message via ```PushMessage```

```go
	var messages []linebot.Message

	// append some message to messages

	_, err := bot.PushMessage(ID, messages... ).Do()
	if err != nil {
		// Do something when some bad happened
	}
```

With reply token, you can reply message via ```ReplyMessage()```

```go
	var messages []linebot.Message

	// append some message to messages

	_, err := bot.PushMessage( replyToken, messages... ).Do()
	if err != nil {
		// Do something when some bad happened
	}
```

## Requirements

This library requires Go 1.6 or later.

## LICENSE

See LICENSE.txt
