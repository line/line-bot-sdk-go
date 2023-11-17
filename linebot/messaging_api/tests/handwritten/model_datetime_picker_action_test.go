package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestDatetimePickerAction(t *testing.T) {
	msg := &messaging_api.DatetimePickerAction{}
	encodedMsg, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Failed to encode message: %v", err)
	}
	encodedMsgStr := string(encodedMsg)
	if strings.Contains(encodedMsgStr, "mode") {
		t.Errorf("Encoded message contains unexpected field 'mode': %s\nenum value must not included if the value is empty.", encodedMsgStr)
	}
}
