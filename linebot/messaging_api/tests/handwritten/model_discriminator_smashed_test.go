package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// C1c: ImagemapMessage type="imagemap" (lowercase smashed, not "Imagemap")
func TestImagemapMessage_DiscriminatorType(t *testing.T) {
	msg := &messaging_api.ImagemapMessage{
		BaseUrl: "https://example.com/imagemap",
		AltText: "This is an imagemap",
		BaseSize: &messaging_api.ImagemapBaseSize{
			Width:  1040,
			Height: 1040,
		},
		Actions: []messaging_api.ImagemapActionInterface{
			&messaging_api.UriImagemapAction{
				LinkUri: "https://example.com/",
				Area: &messaging_api.ImagemapArea{
					X:      0,
					Y:      0,
					Width:  520,
					Height: 1040,
				},
			},
		},
	}

	data, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Failed to marshal ImagemapMessage: %v", err)
	}

	jsonStr := string(data)
	if !strings.Contains(jsonStr, `"type":"imagemap"`) {
		t.Errorf("Expected type 'imagemap' (lowercase smashed), got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"type":"Imagemap"`) || strings.Contains(jsonStr, `"type":"ImagemapMessage"`) {
		t.Errorf("Type should be 'imagemap', not class name: %s", jsonStr)
	}
}

// C1e: RichMenuSwitchAction type="richmenuswitch" (lowercase smashed, not "RichMenuSwitch")
func TestRichMenuSwitchAction_DiscriminatorType(t *testing.T) {
	action := &messaging_api.RichMenuSwitchAction{
		RichMenuAliasId: "richmenu-alias-001",
		Data:            "action=switch",
	}

	data, err := json.Marshal(action)
	if err != nil {
		t.Fatalf("Failed to marshal RichMenuSwitchAction: %v", err)
	}

	jsonStr := string(data)
	if !strings.Contains(jsonStr, `"type":"richmenuswitch"`) {
		t.Errorf("Expected type 'richmenuswitch' (lowercase smashed), got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"type":"RichMenuSwitch"`) || strings.Contains(jsonStr, `"type":"richMenuSwitch"`) {
		t.Errorf("Type should be 'richmenuswitch', not camelCase or PascalCase: %s", jsonStr)
	}
}
