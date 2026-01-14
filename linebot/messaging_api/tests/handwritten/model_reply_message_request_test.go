package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestTemplateMessage(t *testing.T) {
	req := &messaging_api.ReplyMessageRequest{
		ReplyToken: "JKLJDSFhkljdsjfkla",
		Messages: []messaging_api.MessageInterface{
			&messaging_api.TemplateMessage{
				AltText: "Buttons alt text",
				Template: &messaging_api.ButtonsTemplate{
					ThumbnailImageUrl: "https://example.com/static/buttons/1040.jpg",
					Title:             "My button sample",
					Actions: []messaging_api.ActionInterface{
						&messaging_api.UriAction{
							Label: "Go to line.me",
							Uri:   "https://line.me",
						},
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
	if !strings.Contains(encodedMsgStr, `"type":"buttons"`) {
		t.Errorf("Encoded message doens't contains expected default value: %s", encodedMsgStr)
	}
}

func TestFlexBubble(t *testing.T) {
	req := &messaging_api.ReplyMessageRequest{
		ReplyToken: "KHJKLJSDFKLJSfudsifsjfakljfl",
		Messages: []messaging_api.MessageInterface{
			&messaging_api.FlexMessage{
				AltText: "Flex message alt text",
				Contents: messaging_api.FlexBubble{
					Body: &messaging_api.FlexBox{
						Layout: messaging_api.FlexBoxLAYOUT_HORIZONTAL,
						Contents: []messaging_api.FlexComponentInterface{
							&messaging_api.FlexText{
								Text: "Hello,",
							},
							&messaging_api.FlexText{
								Text: "World!",
							},
						},
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
	if !strings.Contains(encodedMsgStr, `"type":"box"`) {
		t.Errorf("Encoded message doens't contains expected default value: %s", encodedMsgStr)
	}
}

func TestFlexBoxUnmarshalJSONWithEmptyContents(t *testing.T) {
	jsonStr := `{"type":"box","layout":"horizontal","contents":[]}`
	var flexBox messaging_api.FlexBox
	err := json.Unmarshal([]byte(jsonStr), &flexBox)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if flexBox.Contents == nil {
		t.Error("FlexBox.Contents should not be nil when unmarshaling JSON with \"contents\":[]")
	}

	if len(flexBox.Contents) != 0 {
		t.Errorf("FlexBox.Contents should be empty, but got %d elements", len(flexBox.Contents))
	}
}
