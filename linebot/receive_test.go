package linebot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReceiveMessage(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client, err := NewClient(1000000000, "testsecret", "TEST_MID")
		if err != nil {
			t.Error(err)
			return
		}
		messages, err := client.ParseRequest(r)
		if err != nil {
			if err == ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
				t.Error(err)
			}
			return
		}
		if messages == nil {
			t.Error("nil messages")
			return
		}
		if len(messages.Results) != 9 {
			t.Errorf("messages: %d", len(messages.Results))
			return
		}
		if messages.Results[0].Content() == nil {
			t.Error("could not retrieve content")
			return
		}
		{
			// Text Message Content
			message := messages.Results[0].Content()
			if !message.IsMessage {
				t.Error("IsMessage: false")
				return
			}
			text, err := message.TextContent()
			if err != nil {
				t.Error(err)
				return
			}
			if text.ID != "325708" {
				t.Errorf("message: %v", text)
			}
			if text.Text != "hello" {
				t.Errorf("message: %v", text)
			}
		}
		{
			// Image Message Content
			message := messages.Results[1].Content()
			if !message.IsMessage {
				t.Error("IsMessage: false")
				return
			}
			_, err := message.ImageContent()
			if err != nil {
				t.Error(err)
				return
			}
		}
		{
			// Video Message Content
			message := messages.Results[2].Content()
			if !message.IsMessage {
				t.Error("IsMessage: false")
				return
			}
			_, err := message.VideoContent()
			if err != nil {
				t.Error(err)
				return
			}
		}
		{
			// Audio Message Content
			message := messages.Results[3].Content()
			if !message.IsMessage {
				t.Error("IsMessage: false")
				return
			}
			audio, err := message.AudioContent()
			if err != nil {
				t.Error(err)
				return
			}
			if audio.Duration != 1234 {
				t.Errorf("message: %v", audio)
			}
		}
		{
			// Location Message Content
			message := messages.Results[4].Content()
			if !message.IsMessage {
				t.Error("IsMessage: false")
				return
			}
			location, err := message.LocationContent()
			if err != nil {
				t.Error(err)
				return
			}
			if location.Text != "位置情報" {
				t.Errorf("message: %v", location)
			}
			if location.Title != "位置情報" {
				t.Errorf("message: %v", location)
			}
			if location.Address != "日本 〒150-0002 東京都渋谷区渋谷" {
				t.Errorf("message: %v", location)
			}
			if location.Latitude != 35.658977 {
				t.Errorf("message: %v", location)
			}
			if location.Longitude != 139.703811 {
				t.Errorf("message: %v", location)
			}
		}
		{
			// Sticker Message Content
			message := messages.Results[5].Content()
			if !message.IsMessage {
				t.Error("IsMessage: false")
				return
			}
			sticker, err := message.StickerContent()
			if err != nil {
				t.Error(err)
				return
			}
			if sticker.PackageID != 1234 {
				t.Errorf("message: %v", sticker)
			}
			if sticker.ID != 5678 {
				t.Errorf("message: %v", sticker)
			}
			if sticker.Version != 1 {
				t.Errorf("message: %v", sticker)
			}
		}
		{
			// Contact Message Content
			message := messages.Results[6].Content()
			if !message.IsMessage {
				t.Error("IsMessage: false")
				return
			}
			contact, err := message.ContactContent()
			if err != nil {
				t.Error(err)
				return
			}
			if contact.Mid != "u0cc15697597f61dd8b01cea8b027050e" {
				t.Errorf("message: %v", contact)
			}
			if contact.DisplayName != "USER NAME" {
				t.Errorf("message: %v", contact)
			}
		}
		{
			// Operation 1
			content := messages.Results[7].Content()
			if !content.IsOperation {
				t.Error("IsOperation: false")
				return
			}
			operation, err := content.OperationContent()
			if err != nil {
				t.Error(err)
				return
			}
			if operation.OpType != OpTypeAddedAsFriend {
				t.Errorf("opType: %v", operation.OpType)
			}
		}
		{
			// Operation 2
			content := messages.Results[8].Content()
			if !content.IsOperation {
				t.Error("IsOperation: false")
				return
			}
			operation, err := content.OperationContent()
			if err != nil {
				t.Error(err)
				return
			}
			if operation.OpType != OpTypeBlocked {
				t.Errorf("opType: %v", operation.OpType)
			}
		}
		w.Write([]byte(`OK`))
	}))
	defer server.Close()

	json := `{
  "result":[
    {
      "from":"u206d25c2ea6bd87c17655609a1c37cb8",
      "fromChannel":1341301815,
      "to":["u0cc15697597f61dd8b01cea8b027050e"],
      "toChannel":1441301333,
      "eventType":"138311609000106303",
      "id":"ABCDEF-12345678901",
      "content":{
        "id":"325708",
        "createdTime":1332394961610,
        "from":"uff2aec188e58752ee1fb0f9507c6529a",
        "to":["u0a556cffd4da0dd89c94fb36e36e1cdc"],
        "toType":1,
        "contentType":1,
        "text":"hello"
      }
    },
    {
      "from":"u206d25c2ea6bd87c17655609a1c37cb8",
      "fromChannel":1341301815,
      "to":["u0cc15697597f61dd8b01cea8b027050e"],
      "toChannel":1441301333,
      "eventType":"138311609000106303",
      "id":"ABCDEF-12345678901",
      "content":{
        "id":"325708",
        "createdTime":1332394961610,
        "from":"uff2aec188e58752ee1fb0f9507c6529a",
        "to":["u0a556cffd4da0dd89c94fb36e36e1cdc"],
        "toType":1,
        "contentType":2
      }
    },
    {
      "from":"u206d25c2ea6bd87c17655609a1c37cb8",
      "fromChannel":1341301815,
      "to":["u0cc15697597f61dd8b01cea8b027050e"],
      "toChannel":1441301333,
      "eventType":"138311609000106303",
      "id":"ABCDEF-12345678901",
      "content":{
        "id":"325708",
        "createdTime":1332394961610,
        "from":"uff2aec188e58752ee1fb0f9507c6529a",
        "to":["u0a556cffd4da0dd89c94fb36e36e1cdc"],
        "toType":1,
        "contentType":3
      }
    },
    {
      "from":"u206d25c2ea6bd87c17655609a1c37cb8",
      "fromChannel":1341301815,
      "to":["u0cc15697597f61dd8b01cea8b027050e"],
      "toChannel":1441301333,
      "eventType":"138311609000106303",
      "id":"ABCDEF-12345678901",
      "content":{
        "id":"325708",
        "createdTime":1332394961610,
        "from":"uff2aec188e58752ee1fb0f9507c6529a",
        "to":["u0a556cffd4da0dd89c94fb36e36e1cdc"],
        "toType":1,
        "contentType":4,
        "contentMetadata":{
            "AUDLEN": "1234"
        }
      }
    },
    {
      "from":"u206d25c2ea6bd87c17655609a1c37cb8",
      "fromChannel":1341301815,
      "to":["u0cc15697597f61dd8b01cea8b027050e"],
      "toChannel":1441301333,
      "eventType":"138311609000106303",
      "id":"ABCDEF-12345678901",
      "content":{
        "id":"325708",
        "createdTime":1332394961610,
        "from":"uff2aec188e58752ee1fb0f9507c6529a",
        "to":["u0a556cffd4da0dd89c94fb36e36e1cdc"],
        "toType":1,
        "contentType":7,
        "text": "位置情報",
        "location":{
            "title":"位置情報",
            "address":"日本 〒150-0002 東京都渋谷区渋谷",
            "latitude":35.658977,
            "longitude":139.703811,
            "phone":null
        }
      }
    },
    {
      "from":"u206d25c2ea6bd87c17655609a1c37cb8",
      "fromChannel":1341301815,
      "to":["u0cc15697597f61dd8b01cea8b027050e"],
      "toChannel":1441301333,
      "eventType":"138311609000106303",
      "id":"ABCDEF-12345678901",
      "content":{
        "id":"325708",
        "createdTime":1332394961610,
        "from":"uff2aec188e58752ee1fb0f9507c6529a",
        "to":["u0a556cffd4da0dd89c94fb36e36e1cdc"],
        "toType":1,
        "contentType":8,
        "contentMetadata":{
            "STKTXT":"[]",
            "STKVER":"1",
            "STKID":"5678",
            "STKPKGID":"1234"
        }
      }
    },
    {
      "from":"u206d25c2ea6bd87c17655609a1c37cb8",
      "fromChannel":1341301815,
      "to":["u0cc15697597f61dd8b01cea8b027050e"],
      "toChannel":1441301333,
      "eventType":"138311609000106303",
      "id":"ABCDEF-12345678901",
      "content":{
        "id":"325708",
        "createdTime":1332394961610,
        "from":"uff2aec188e58752ee1fb0f9507c6529a",
        "to":["u0a556cffd4da0dd89c94fb36e36e1cdc"],
        "toType":1,
        "contentType":10,
        "contentMetadata":{
            "mid":"u0cc15697597f61dd8b01cea8b027050e",
            "displayName":"USER NAME"
        }
      }
    },
    {
      "from":"u206d25c2ea6bd87c17655609a1c37cb8",
      "fromChannel":1341301815,
      "to":["u0cc15697597f61dd8b01cea8b027050e"],
      "toChannel":1441301333,
      "eventType":"138311609100106403",
      "id":"ABCDEF-12345678902",
      "content":{
        "revision":2469,
        "opType":4,
        "params":[
          "u0f3bfc598b061eba02183bfc5280886a",
          null,
          null
        ]
      }
    },
    {
      "from":"u206d25c2ea6bd87c17655609a1c37cb8",
      "fromChannel":1341301815,
      "to":["u0cc15697597f61dd8b01cea8b027050e"],
      "toChannel":1441301333,
      "eventType":"138311609100106403",
      "id":"ABCDEF-12345678903",
      "content":{
        "revision":2470,
        "opType":8,
        "params":[
          "u0f3bfc598b061eba02183bfc5280886a",
          null,
          null
        ]
      }
    }
  ]
}
`
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	// no signature
	{
		req, err := http.NewRequest("POST", server.URL, bytes.NewReader([]byte(json)))
		if err != nil {
			t.Error(err)
			return
		}
		res, err := httpClient.Do(req)
		if err != nil {
			t.Error(err)
			return
		}
		if res == nil {
			t.Error("response is nil")
			return
		}
		if res.StatusCode != 400 {
			t.Error("status should not be 400")
			return
		}
	}
	// valid signature
	{
		req, err := http.NewRequest("POST", server.URL, bytes.NewReader([]byte(json)))
		if err != nil {
			t.Error(err)
			return
		}
		// generate signature
		mac := hmac.New(sha256.New, []byte("testsecret"))
		mac.Write([]byte(json))

		req.Header.Set("X-LINE-ChannelSignature", base64.StdEncoding.EncodeToString(mac.Sum(nil)))
		res, err := httpClient.Do(req)
		if err != nil {
			t.Error(err)
			return
		}
		if res == nil {
			t.Error("response is nil")
			return
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("status: %d", res.StatusCode)
			return
		}
	}
}
