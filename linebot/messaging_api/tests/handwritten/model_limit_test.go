package tests

import (
	"encoding/json"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func assertJSONFields(t *testing.T, jsonData []byte, expected map[string]any, absentKeys []string) {
	var got map[string]any
	if err := json.Unmarshal(jsonData, &got); err != nil {
		t.Fatalf("Invalid JSON: %v", err)
	}

	for k, v := range expected {
		val, ok := got[k]
		if !ok {
			t.Errorf("Expected field %q to be present, but it's missing", k)
			continue
		}

		// If the deserialized value is float64, convert it to int before comparison
		if f, ok := val.(float64); ok {
			val = int(f)
		}

		if val != v {
			t.Errorf("Expected %q=%v, got %v", k, v, val)
		}
	}

	for _, k := range absentKeys {
		if _, ok := got[k]; ok {
			t.Errorf("Expected field %q to be absent, but it exists", k)
		}
	}
}

func TestLimitSerialization(t *testing.T) {
	tests := []struct {
		name       string
		limit      messaging_api.Limit
		expected   map[string]any
		absentKeys []string
	}{
		{
			name:     "Max is set",
			limit:    messaging_api.Limit{Max: 10, UpToRemainingQuota: true},
			expected: map[string]any{"max": 10, "upToRemainingQuota": true},
		},
		{
			name:       "Max is zero (omitempty)",
			limit:      messaging_api.Limit{Max: 0, UpToRemainingQuota: true},
			expected:   map[string]any{"upToRemainingQuota": true},
			absentKeys: []string{"max"},
		},
		{
			name:       "Max is zero and UpToRemainingQuota is false",
			limit:      messaging_api.Limit{Max: 0, UpToRemainingQuota: false},
			expected:   map[string]any{"upToRemainingQuota": false},
			absentKeys: []string{"max"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.limit)
			if err != nil {
				t.Fatalf("marshal error: %v", err)
			}
			assertJSONFields(t, data, tt.expected, tt.absentKeys)
		})
	}
}
