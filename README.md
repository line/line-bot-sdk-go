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

## Requirements

This library requires Go 1.6 or later.

## LICENSE

See LICENSE.txt
