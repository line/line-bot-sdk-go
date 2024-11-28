package tests

import (
	"encoding/json"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestLimitSerialization(t *testing.T) {
	tests := []struct {
		name     string
		limit    messaging_api.Limit
		expected string
	}{
		{
			name:     "Max is set",
			limit:    messaging_api.Limit{Max: 10, UpToRemainingQuota: true},
			expected: `{"max":10,"upToRemainingQuota":true}`,
		},
		{
			name:     "Max is zero (omitempty)",
			limit:    messaging_api.Limit{Max: 0, UpToRemainingQuota: true},
			expected: `{"upToRemainingQuota":true}`,
		},
		{
			name:     "Max is zero and UpToRemainingQuota is false",
			limit:    messaging_api.Limit{Max: 0, UpToRemainingQuota: false},
			expected: `{"upToRemainingQuota":false}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.limit)
			if err != nil {
				t.Fatalf("Failed to marshal Limit: %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("Expected JSON: %s, got: %s", tt.expected, string(data))
			}
		})
	}
}
