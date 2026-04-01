package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/channel_access_token"
)

func TestVerifyChannelTokenByJWTWithHttpInfo(t *testing.T) {
	tests := []struct {
		name        string
		accessToken string
	}{
		{
			name:        "Uses snake_case query parameter key",
			accessToken: "test_channel_access_token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				query := r.URL.Query()

				if got := query.Get("access_token"); got != tt.accessToken {
					t.Errorf("Expected query param access_token=%q, got %q", tt.accessToken, got)
				}

				if got := query.Get("accessToken"); got != "" {
					t.Errorf("camelCase key accessToken should not be present, got %q", got)
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(channel_access_token.VerifyChannelAccessTokenResponse{
					ClientId:  "test_client_id",
					ExpiresIn: 3600,
				})
			}))
			defer mockServer.Close()

			client, err := channel_access_token.NewChannelAccessTokenAPI(
				channel_access_token.WithEndpoint(mockServer.URL),
			)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			_, result, err := client.VerifyChannelTokenByJWTWithHttpInfo(
				tt.accessToken,
			)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result.ClientId != "test_client_id" {
				t.Errorf("Expected ClientId: test_client_id, got: %s", result.ClientId)
			}
		})
	}
}
