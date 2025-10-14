# LINE Messaging API SDK for Go

[![Build Status](https://github.com/line/line-bot-sdk-go/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/line/line-bot-sdk-go/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/line/line-bot-sdk-go/v8/linebot.svg)](https://pkg.go.dev/github.com/line/line-bot-sdk-go/v8/linebot)
[![Go Report Card](https://goreportcard.com/badge/github.com/line/line-bot-sdk-go)](https://goreportcard.com/report/github.com/line/line-bot-sdk-go)


## Introduction
The LINE Messaging API SDK for Go makes it easy to develop bots using LINE Messaging API, and you can create a sample bot within minutes.

## Documentation

See the official API documentation for more information.

- English: https://developers.line.biz/en/docs/messaging-api/overview/
- Japanese: https://developers.line.biz/ja/docs/messaging-api/overview/

## Requirements

This library requires Go 1.24 or later.

## Installation ##

```sh
$ go get -u github.com/line/line-bot-sdk-go/v8/linebot
```

## Import all packages in your code ##
```go
import (
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/channel_access_token"
	"github.com/line/line-bot-sdk-go/v8/linebot/insight"
	"github.com/line/line-bot-sdk-go/v8/linebot/liff"
	"github.com/line/line-bot-sdk-go/v8/linebot/manage_audience"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/module"
	"github.com/line/line-bot-sdk-go/v8/linebot/module_attach"
	"github.com/line/line-bot-sdk-go/v8/linebot/shop"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

```

## Configuration ##

```go
import (
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func main() {
	bot, err := messaging_api.NewMessagingApiAPI(
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	...
}

```

### Configuration with http.Client ###

Every client application allows configuration with WithHTTPClient and WithEndpoint.
(For Blob client, configurations WithBlobHTTPClient and WithBlobEndpoint are also available.)

```go
client := &http.Client{}
bot, err := messaging_api.NewMessagingApiAPI(
	os.Getenv("LINE_CHANNEL_TOKEN"),
	messaging_api.WithHTTPClient(client),
)
...
```

## Getting Started ##

The LINE Messaging API primarily utilizes the JSON data format. To parse the incoming HTTP requests, the `webhook.ParseRequest()` method is provided. This method reads the `*http.Request` content and returns a slice of pointers to Event Objects.

```go
import (
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

cb, err := webhook.ParseRequest(os.Getenv("LINE_CHANNEL_SECRET"), req)
if err != nil {
	// Handle any errors that occur.
}
```

The LINE Messaging API is capable of handling various event types. The Messaging API SDK automatically unmarshals these events into respective classes like `webhook.MessageEvent`, `webhook.FollowEvent`, and so on. You can easily check the type of the event and respond accordingly using a switch statement as shown below:


```go
for _, event := range cb.Events {
	switch e := event.(type) {
		case webhook.MessageEvent:
			// Do Something...
		case webhook.StickerMessageContent:
			// Do Something...
	}
}
```

We provide code [examples](./examples).
- [EchoBot](./examples/echo_bot/server.go)
  - a simple echo bot
- [KitchenSink](./examples/kitchensink/server.go)
  - a bot that handles many types of events
- [EchoBotHandler](./examples/echo_bot_handler/server.go)
  - A simple bot that automatically verifies signatures and only handles Webhook events
- [DeliveryHelper](./examples/delivery_helper/main.go)
- [InsightHelper](./examples/insight_helper/main.go)
- [RichmenuHelper](./examples/richmenu_helper/main.go)

### Receiver ###

To send a message to a user, group, or room, you need either an ID

```go
userID := event.Source.UserId
groupID := event.Source.GroupId
RoomID := event.Source.RoomId
```

or a reply token.

```go
replyToken := event.ReplyToken
```

### Create message ###

The LINE Messaging API provides various types of message.

```go
bot.ReplyMessage(
	&messaging_api.ReplyMessageRequest{
		ReplyToken: e.ReplyToken,
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: replyMessage,
			},
		},
	},
)
```

### Send message ###

With an ID, you can send message using ```PushMessage()```

```go
bot.PushMessage(
	&messaging_api.PushMessageRequest{
		To: "U.......",
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: pushMessage,
			},
		},
	},
	"", // x-line-retry-key
)
```

With a reply token, you can reply to messages using ```ReplyMessage()```

```go
bot.ReplyMessage(
	&messaging_api.ReplyMessageRequest{
		ReplyToken: e.ReplyToken,
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: replyMessage,
			},
		},
	},
)
```

### How to get response header and error message ###
You may need to store the ```x-line-request-id``` header obtained as a response from several APIs. In this case, please use ```~WithHttpInfo```. You can get headers and status codes. The ```x-line-accepted-request-id``` or ```content-type``` header can also be obtained in the same way.

```go
resp, _, _ := app.bot.ReplyMessageWithHttpInfo(
	&messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: "Hello, world",
			},
		},
	}, 
)
log.Printf("status code: (%v), x-line-request-id: (%v)", resp.StatusCode, resp.Header.Get("x-line-request-id"))

```

Similarly, you can get specific error messages by using ```~WithHttpInfo```.

```go
resp, _, err := app.bot.ReplyMessageWithHttpInfo(
    &messaging_api.ReplyMessageRequest{
        ReplyToken: replyToken + "invalid",
        Messages: []messaging_api.MessageInterface{
            messaging_api.TextMessage{
                Text: "Hello, world",
            },
        },
    },
)
if err != nil && resp.StatusCode >= 400 && resp.StatusCode < 500 {
    decoder := json.NewDecoder(resp.Body)
    errorResponse := &messaging_api.ErrorResponse{}
    if err := decoder.Decode(&errorResponse); err != nil {
        log.Fatal("failed to decode JSON: %w", err)
    }
    log.Printf("status code: (%v), x-line-request-id: (%v), error response: (%v)", resp.StatusCode, resp.Header.Get("x-line-request-id"), errorResponse)
}
```

## Help and media

FAQ: https://developers.line.biz/en/faq/

News: https://developers.line.biz/en/news/


## Versioning

This project respects semantic versioning.
- See https://semver.org/

However, if a feature that was publicly released is discontinued for business reasons and becomes completely unusable, we will release changes as a patch release.


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
