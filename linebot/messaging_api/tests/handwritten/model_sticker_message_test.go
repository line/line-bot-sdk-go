package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestStickerMessage(t *testing.T) {
	req := &messaging_api.ReplyMessageRequest{
		ReplyToken: "JKLJDSFhkljdsjfkla",
		Messages: []messaging_api.MessageInterface{
			&messaging_api.StickerMessage{
				PackageId: "1",
				StickerId: "2",
			},
		},
	}
	encodedMsg, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to encode message: %v", err)
	}
	encodedMsgStr := string(encodedMsg)
	if !strings.Contains(encodedMsgStr, `"type":"sticker"`) {
		t.Errorf("Encoded message doens't contains expected default value: %s", encodedMsgStr)
	}
}
