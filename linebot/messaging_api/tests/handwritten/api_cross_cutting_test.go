package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/channel_access_token"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/module_attach"
)

// B2: x-line-retry-key header
func TestRetryKey_WithUUID(t *testing.T) {
	retryKey := "123e4567-e89b-12d3-a456-426614174000"

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			got := r.Header.Get("X-Line-Retry-Key")
			if got != retryKey {
				t.Errorf("Expected X-Line-Retry-Key %q, got %q", retryKey, got)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"sentMessages":[{"id":"msg1"}]}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	result, err := client.PushMessage(
		&messaging_api.PushMessageRequest{
			To: "user123",
			Messages: []messaging_api.MessageInterface{
				&messaging_api.TextMessage{Text: "Hello"},
			},
		},
		retryKey,
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(result.SentMessages) != 1 || result.SentMessages[0].Id != "msg1" {
		t.Errorf("Unexpected result: %+v", result)
	}
}

func TestRetryKey_Empty(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			got := r.Header.Get("X-Line-Retry-Key")
			if got != "" {
				t.Errorf("Expected empty X-Line-Retry-Key, got %q", got)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"sentMessages":[{"id":"msg1"}]}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.PushMessage(
		&messaging_api.PushMessageRequest{
			To: "user123",
			Messages: []messaging_api.MessageInterface{
				&messaging_api.TextMessage{Text: "Hello"},
			},
		},
		"",
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

// B3: Empty response body (empty JSON object)
func TestBroadcast_EmptyResponseBody(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	result, err := client.Broadcast(
		&messaging_api.BroadcastRequest{
			Messages: []messaging_api.MessageInterface{
				&messaging_api.TextMessage{Text: "Hello"},
			},
		},
		"",
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
}

// B4: Nil parameter omission
func TestNilParamOmission(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				t.Fatalf("Failed to parse form: %v", err)
			}
			if r.PostForm.Get("grant_type") != "client_credentials" {
				t.Errorf("Expected grant_type=client_credentials, got %q", r.PostForm.Get("grant_type"))
			}
			if r.PostForm.Get("client_id") != "my_client_id" {
				t.Errorf("Expected client_id=my_client_id, got %q", r.PostForm.Get("client_id"))
			}
			if r.PostForm.Get("client_secret") != "my_client_secret" {
				t.Errorf("Expected client_secret=my_client_secret, got %q", r.PostForm.Get("client_secret"))
			}
			if _, ok := r.PostForm["client_assertion_type"]; ok {
				t.Error("client_assertion_type should not be present when empty")
			}
			if _, ok := r.PostForm["client_assertion"]; ok {
				t.Error("client_assertion should not be present when empty")
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(channel_access_token.IssueStatelessChannelAccessTokenResponse{
				AccessToken: "test_token",
				TokenType:   "Bearer",
				ExpiresIn:   900,
			})
		}),
	)
	defer server.Close()

	client, err := channel_access_token.NewChannelAccessTokenAPI(
		channel_access_token.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	result, err := client.IssueStatelessChannelToken(
		"client_credentials", "", "", "my_client_id", "my_client_secret",
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.AccessToken != "test_token" {
		t.Errorf("Expected AccessToken 'test_token', got %q", result.AccessToken)
	}
}

// B5: User-Agent header
func TestUserAgentHeader(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ua := r.Header.Get("User-Agent")
			if !strings.HasPrefix(ua, "LINE-BotSDK-Go/") {
				t.Errorf("Expected User-Agent starting with 'LINE-BotSDK-Go/', got %q", ua)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"displayName":"Test","userId":"user123","pictureUrl":"","statusMessage":""}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.GetProfile("user123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

// B6: Error handling (4xx/5xx)
func TestErrorHandling_400(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message":"invalid request"}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.PushMessage(
		&messaging_api.PushMessageRequest{
			To: "user123",
			Messages: []messaging_api.MessageInterface{
				&messaging_api.TextMessage{Text: "Hello"},
			},
		},
		"",
	)
	if err == nil {
		t.Fatal("Expected error for 400 response, got nil")
	}
	if !strings.Contains(err.Error(), "400") {
		t.Errorf("Expected error to contain '400', got %q", err.Error())
	}
}

func TestErrorHandling_500(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"internal server error"}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.PushMessage(
		&messaging_api.PushMessageRequest{
			To: "user123",
			Messages: []messaging_api.MessageInterface{
				&messaging_api.TextMessage{Text: "Hello"},
			},
		},
		"",
	)
	if err == nil {
		t.Fatal("Expected error for 500 response, got nil")
	}
	if !strings.Contains(err.Error(), "500") {
		t.Errorf("Expected error to contain '500', got %q", err.Error())
	}
}

// B7: base_url customization
func TestBaseUrlCustomization_MessagingApi(t *testing.T) {
	called := false
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"displayName":"Test","userId":"user123","pictureUrl":"","statusMessage":""}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.GetProfile("user123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !called {
		t.Error("Custom endpoint was not called")
	}
}

func TestBaseUrlCustomization_BlobApi(t *testing.T) {
	called := false
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
			w.Header().Set("Content-Type", "image/png")
			w.Write([]byte{0x89, 0x50, 0x4E, 0x47})
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiBlobAPI(
		"channelToken",
		messaging_api.WithBlobEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	resp, err := client.GetRichMenuImage("richmenu-abc")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer resp.Body.Close()
	io.ReadAll(resp.Body)

	if !called {
		t.Error("Custom blob endpoint was not called")
	}
}

func TestBaseUrlCustomization_ModuleAttach(t *testing.T) {
	called := false
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"bot_id":"bot123","scopes":["message"]}`))
		}),
	)
	defer server.Close()

	client, err := module_attach.NewLineModuleAttachAPI(
		"channelToken",
		module_attach.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.AttachModule(
		"authorization_code", "code123", "https://example.com/callback",
		"", "", "", "", "", "", "",
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !called {
		t.Error("Custom module_attach endpoint was not called")
	}
}

// B8: Module attach uses Bearer auth with custom base_url
func TestModuleAttach_BearerAuthAndCustomEndpoint(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/module/auth/v1/token" {
				t.Errorf("Expected path /module/auth/v1/token, got %s", r.URL.Path)
			}

			auth := r.Header.Get("Authorization")
			if auth != "Bearer myModuleToken" {
				t.Errorf("Expected Authorization 'Bearer myModuleToken', got %q", auth)
			}

			if ct := r.Header.Get("Content-Type"); ct != "application/x-www-form-urlencoded" {
				t.Errorf("Expected Content-Type 'application/x-www-form-urlencoded', got %q", ct)
			}

			if err := r.ParseForm(); err != nil {
				t.Fatalf("Failed to parse form: %v", err)
			}
			if r.PostForm.Get("grant_type") != "authorization_code" {
				t.Errorf("Expected grant_type=authorization_code, got %q", r.PostForm.Get("grant_type"))
			}
			if r.PostForm.Get("code") != "auth_code_123" {
				t.Errorf("Expected code=auth_code_123, got %q", r.PostForm.Get("code"))
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"bot_id":"U1234567890","scopes":["message","profile"]}`))
		}),
	)
	defer server.Close()

	client, err := module_attach.NewLineModuleAttachAPI(
		"myModuleToken",
		module_attach.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	result, err := client.AttachModule(
		"authorization_code", "auth_code_123", "https://example.com/callback",
		"", "", "", "", "", "", "",
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.BotId != "U1234567890" {
		t.Errorf("Expected BotId 'U1234567890', got %q", result.BotId)
	}
	if len(result.Scopes) != 2 {
		t.Errorf("Expected 2 scopes, got %d", len(result.Scopes))
	}
}

// B9: Unknown response fields don't crash deserialization
func TestUnknownResponseFields(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"displayName":"Test","userId":"user123","pictureUrl":"https://example.com/pic.jpg","statusMessage":"hi","unknownField":"value","anotherUnknown":42,"nested":{"deep":true}}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	result, err := client.GetProfile("user123")
	if err != nil {
		t.Fatalf("Unexpected error with unknown fields: %v", err)
	}
	if result.DisplayName != "Test" {
		t.Errorf("Expected DisplayName 'Test', got %q", result.DisplayName)
	}
	if result.UserId != "user123" {
		t.Errorf("Expected UserId 'user123', got %q", result.UserId)
	}
}
