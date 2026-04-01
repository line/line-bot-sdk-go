package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/channel_access_token"
)

func TestGetsAllValidChannelAccessTokenKeyIdsWithHttpInfo(t *testing.T) {
	tests := []struct {
		name                string
		clientAssertionType string
		clientAssertion     string
	}{
		{
			name:                "Uses snake_case query parameter keys",
			clientAssertionType: "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
			clientAssertion:     "test_jwt_assertion",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				query := r.URL.Query()

				if got := query.Get("client_assertion_type"); got != tt.clientAssertionType {
					t.Errorf("Expected query param client_assertion_type=%q, got %q", tt.clientAssertionType, got)
				}
				if got := query.Get("client_assertion"); got != tt.clientAssertion {
					t.Errorf("Expected query param client_assertion=%q, got %q", tt.clientAssertion, got)
				}

				if got := query.Get("clientAssertionType"); got != "" {
					t.Errorf("camelCase key clientAssertionType should not be present, got %q", got)
				}
				if got := query.Get("clientAssertion"); got != "" {
					t.Errorf("camelCase key clientAssertion should not be present, got %q", got)
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(channel_access_token.ChannelAccessTokenKeyIdsResponse{
					Kids: []string{"kid1", "kid2"},
				})
			}))
			defer mockServer.Close()

			client, err := channel_access_token.NewChannelAccessTokenAPI(
				channel_access_token.WithEndpoint(mockServer.URL),
			)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			_, result, err := client.GetsAllValidChannelAccessTokenKeyIdsWithHttpInfo(
				tt.clientAssertionType,
				tt.clientAssertion,
			)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if len(result.Kids) != 2 {
				t.Errorf("Expected 2 kids, got %d", len(result.Kids))
			}
		})
	}
}
