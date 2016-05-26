# line-bot-sdk-go

[![Build Status](https://travis-ci.org/line/line-bot-sdk-go.svg?branch=master)](https://travis-ci.org/line/line-bot-sdk-go)

SDK of the LINE BOT API Trial for Go

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
	bot, err := linebot.NewClient(<Channel ID>, "<Channel Secret>", "<MID>")
	...
}

```

### Configuration with http.Client ###

```go
	client := &http.Client{}
	bot, err := linebot.NewClient(<Channel ID>, "<Channel Secret>", "<MID>", linebot.WithHTTPClient(client))
	...
```

## Usage ##

### Sending messages ###

Send a text message, image, video, audio, location, or sticker to the mids.

- [https://developers.line.me/bot-api/api-reference#sending_message](https://developers.line.me/bot-api/api-reference#sending_message)

```go
	// send text
	res, err := bot.SendText([]string{"<target user's MID>"}, "Hello, world!")

	// send image
	res, err := bot.SendImage([]string{"<target user's MID>"}, "http://example.com/image.jpg", "http://example.com/image_preview.jpg")

	// send video
	res, err := bot.SendVideo([]string{"<target user's MID>"}, "http://example.com/video.mp4", "http://example.com/image_preview.jpg")

	// send audio
	res, err := bot.SendAudio([]string{"<target user's MID>"}, "http://example.com/audio.mp3", 2000)

	// send location
	res, err := bot.SendLocation([]string{"<target user's MID>"}, "location label", "tokyo shibuya-ku", 35.661777, 139.704051)

	// send sticker
	res, err := bot.SendSticker([]string{"<target user's MID>"}, 1, 1, 100)
```

### Sending multiple messages ###

The `multiple_message` method allows you to use the _Sending multiple messages API_.

- [https://developers.line.me/bot-api/api-reference#sending_multiple_messages](https://developers.line.me/bot-api/api-reference#sending_multiple_messages)

```go
	res, err := bot.NewMultipleMessage().
		AddText("Hello,").
		AddText("world!").
		AddImage("http://example.com/image.jpg", "http://example.com/image_preview.jpg")
		AddVideo("http://example.com/video.mp4", "http://example.com/image_preview.jpg")
		AddAudio("http://example.com/audio.mp3", 2000)
		AddLocation("Location label", "tokyo shibuya-ku", 35.61823286112982, 139.72824096679688).
		AddSticker(1, 1, 100).
		Send([]string{"<target user's MID>"})
```

### Sending rich messages ###

The `rich_message` method allows you to use the _Sending rich messages API_.

- [https://developers.line.me/bot-api/api-reference#sending_rich_content_message](https://developers.line.me/bot-api/api-reference#sending_rich_content_message)

```go
	res, err := bot.NewRichMessage(1040).
		SetAction("MANGA", "manga", "https://store.line.me/family/manga/en").
		SetListener("MANGA", 0, 0, 520, 520).
		SetAction("MUSIC", "music", "https://store.line.me/family/music/en").
		SetListener("MUSIC", 520, 0, 520, 520).
		Send([]string{"<target user's MID>"}, "https://example.com/rich-image/foo", "This is a alt text.")
```

### Receiving messages ###

The following utility method allows you to easily process messages sent from the BOT API platform via a Callback URL.

- [https://developers.line.me/bot-api/api-reference#receiving_messages](https://developers.line.me/bot-api/api-reference#receiving_messages)

```go
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		received, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, result := range received.Results {
			content := result.Content()
			if content != nil && content.IsMessage && content.ContentType == linebot.ContentTypeText {
				text, err := content.TextContent()
				_, err := bot.SendText([]string{content.From}, "OK " + text.Text)
				if err != nil {
					log.Println(err)
				}
			}
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
```

### Getting message content ###

Retrieve the original file which was sent by user.

- [https://developers.line.me/bot-api/api-reference#getting_message_content](https://developers.line.me/bot-api/api-reference#getting_message_content)

```go
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		...
		for _, result := range received.Results {
			content := result.Content()
			if content != nil && content.IsMessage && content.ContentType == linebot.ContentTypeImage {
				res, err := bot.GetMessageContent(content)
				if err != nil {
					return
				}
				defer res.Content.Close()
				image, err := jpeg.Decode(res.Content)
				if err != nil {
					return
				}
				log.Printf("image %v", image.Bounds())
			}
		}
	})
```

### Getting previews of message content ###

Retrieve the preview image file which was sent by user.

- [https://developers.line.me/bot-api/api-reference#getting_message_content_preview](https://developers.line.me/bot-api/api-reference#getting_message_content_preview)

```go
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		...
		for _, result := range received.Results {
			content := result.Content()
			if content != nil && content.IsMessage && content.ContentType == linebot.MessageContentTypeImage {
				res, err := bot.GetMessageContentPreview(content)
				if err != nil {
					return
				}
				defer res.Content.Close()
				image, err := jpeg.Decode(res.Content)
				if err != nil {
					return
				}
				log.Printf("image %v", image.Bounds())
			}
		}
	})
```

### Getting user profile information ###

You can retrieve the user profile information by specifying the mid.

- [https://developers.line.me/bot-api/api-reference#getting_user_profile_information](https://developers.line.me/bot-api/api-reference#getting_user_profile_information)

```go
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		...
		for _, result := range received.Results {
			content := result.Content()
			if content != nil {
				result, err := bot.GetUserProfile([]string{content.From})
				if err != nil {
					return
				}
				log.Printf("profile: %v", result)
			}
		}
	})
```
