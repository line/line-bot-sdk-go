package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestTextMessageV2(t *testing.T) {
	req := &messaging_api.ReplyMessageRequest{
		ReplyToken: "JKLJDSFhkljdsjfkla",
		Messages: []messaging_api.MessageInterface{
			&messaging_api.TextMessageV2{
				Text: "Hello, {newbee}! {smile}",
				Substitution: map[string]messaging_api.SubstitutionObjectInterface{
					"newbee": &messaging_api.MentionSubstitutionObject{
						Mentionee: &messaging_api.UserMentionTarget{
							UserId: "U1234567890abcdef1234567890abcdef",
						},
					},
					"smile": &messaging_api.EmojiSubstitutionObject{
						ProductId: "5ac1bfd5040ab15980c9b435",
						EmojiId: "002",
					},
				},
			},
		},
	}
	encodedMsg, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to encode message: %v", err)
	}
	encodedMsgStr := string(encodedMsg)
	if !strings.Contains(encodedMsgStr, `"type":"textV2"`) {
		t.Errorf("Encoded message doens't contains expected default value: %s", encodedMsgStr)
	}
	if !strings.Contains(encodedMsgStr, `"type":"mention"`) {
		t.Errorf("Encoded message doens't contains expected default value: %s", encodedMsgStr)
	}
	if !strings.Contains(encodedMsgStr, `"type":"emoji"`) {
		t.Errorf("Encoded message doens't contains expected default value: %s", encodedMsgStr)
	}
}
