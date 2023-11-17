package tests

import (
	"encoding/json"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func TestStickerMessage(t *testing.T) {
	var cb webhook.CallbackRequest
	if err := json.Unmarshal([]byte(`{
		"destination": "Uaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"events": [
			{
				"type": "UNKNOWN",
				"great": "new-field"
			}
		]
	}`), &cb); err != nil {
		t.Fatalf("Failed to unmarshal callback request: %v", err)
	}

	_, ok := cb.Events[0].(webhook.UnknownEvent)
	if !ok {
		t.Fatalf("Failed to cast to UnknownEvent: %v", cb.Events[0])
	}
}
