package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/channel_access_token"
)

func TestIssueChannelToken_FormWithoutBearer(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, got %s", r.Method)
		}

		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			t.Errorf("Expected Content-Type starting with 'application/x-www-form-urlencoded', got '%s'", r.Header.Get("Content-Type"))
		}

		if auth := r.Header.Get("Authorization"); auth != "" {
			t.Errorf("Expected no Authorization header, got '%s'", auth)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatalf("Failed to parse form: %v", err)
		}

		if v := r.PostForm.Get("grant_type"); v != "client_credentials" {
			t.Errorf("Expected grant_type 'client_credentials', got '%s'", v)
		}
		if v := r.PostForm.Get("client_id"); v != "my_client_id" {
			t.Errorf("Expected client_id 'my_client_id', got '%s'", v)
		}
		if v := r.PostForm.Get("client_secret"); v != "my_client_secret" {
			t.Errorf("Expected client_secret 'my_client_secret', got '%s'", v)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(channel_access_token.IssueShortLivedChannelAccessTokenResponse{
			AccessToken: "short_lived_token",
			ExpiresIn:   2592000,
			TokenType:   "Bearer",
		})
	}))
	defer mockServer.Close()

	client, err := channel_access_token.NewChannelAccessTokenAPI(
		channel_access_token.WithEndpoint(mockServer.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	result, err := client.IssueChannelToken("client_credentials", "my_client_id", "my_client_secret")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result.AccessToken != "short_lived_token" {
		t.Errorf("Expected AccessToken 'short_lived_token', got '%s'", result.AccessToken)
	}
	if result.ExpiresIn != 2592000 {
		t.Errorf("Expected ExpiresIn 2592000, got %d", result.ExpiresIn)
	}
	if result.TokenType != "Bearer" {
		t.Errorf("Expected TokenType 'Bearer', got '%s'", result.TokenType)
	}
}
