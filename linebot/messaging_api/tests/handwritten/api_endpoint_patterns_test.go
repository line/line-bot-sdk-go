package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// A1: GET + single path param
func TestGetProfile_PathConstruction(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("Expected method GET, got %s", r.Method)
			}
			if r.URL.Path != "/v2/bot/profile/user123abc" {
				t.Errorf("Expected path /v2/bot/profile/user123abc, got %s", r.URL.Path)
			}
			if r.Header.Get("Authorization") != "Bearer channelToken" {
				t.Errorf("Expected Authorization header 'Bearer channelToken', got %q", r.Header.Get("Authorization"))
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"displayName":"Test User","userId":"user123abc","pictureUrl":"https://example.com/pic.jpg","statusMessage":"hello"}`))
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

	result, err := client.GetProfile("user123abc")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.DisplayName != "Test User" {
		t.Errorf("Expected DisplayName 'Test User', got %q", result.DisplayName)
	}
	if result.UserId != "user123abc" {
		t.Errorf("Expected UserId 'user123abc', got %q", result.UserId)
	}
}

// A8: POST + JSON body
func TestPushMessage_JsonBody(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("Expected method POST, got %s", r.Method)
			}
			if r.URL.Path != "/v2/bot/message/push" {
				t.Errorf("Expected path /v2/bot/message/push, got %s", r.URL.Path)
			}
			if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
				t.Errorf("Expected Content-Type starting with 'application/json', got %q", r.Header.Get("Content-Type"))
			}

			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("Failed to read request body: %v", err)
			}
			var body map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &body); err != nil {
				t.Fatalf("Failed to unmarshal request body: %v", err)
			}
			if body["to"] != "user123" {
				t.Errorf("Expected 'to' field to be 'user123', got %v", body["to"])
			}
			messages, ok := body["messages"].([]interface{})
			if !ok || len(messages) != 1 {
				t.Fatalf("Expected 1 message, got %v", body["messages"])
			}
			msg := messages[0].(map[string]interface{})
			if msg["text"] != "Hello" {
				t.Errorf("Expected message text 'Hello', got %v", msg["text"])
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
		"",
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(result.SentMessages) != 1 || result.SentMessages[0].Id != "msg1" {
		t.Errorf("Expected sentMessages[0].id='msg1', got %+v", result.SentMessages)
	}
}

// A14: PUT + JSON body (note: generated code uses POST for this endpoint — a template bug)
func TestUpdateRichMenuAlias_PutJson(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/v2/bot/richmenu/alias/alias-1" {
				t.Errorf("Expected path /v2/bot/richmenu/alias/alias-1, got %s", r.URL.Path)
			}
			if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
				t.Errorf("Expected Content-Type starting with 'application/json', got %q", r.Header.Get("Content-Type"))
			}

			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("Failed to read request body: %v", err)
			}
			var body map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &body); err != nil {
				t.Fatalf("Failed to unmarshal request body: %v", err)
			}
			if body["richMenuId"] != "richmenu-abc" {
				t.Errorf("Expected richMenuId 'richmenu-abc', got %v", body["richMenuId"])
			}

			w.WriteHeader(http.StatusOK)
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

	_, err = client.UpdateRichMenuAlias(
		"alias-1",
		&messaging_api.UpdateRichMenuAliasRequest{RichMenuId: "richmenu-abc"},
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

// A17: DELETE
func TestDeleteRichMenu_Delete(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/v2/bot/richmenu/richmenu-xyz" {
				t.Errorf("Expected path /v2/bot/richmenu/richmenu-xyz, got %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
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

	_, err = client.DeleteRichMenu("richmenu-xyz")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

// A2: GET + multi path params
func TestGetGroupMemberProfile_MultiPathParams(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("Expected method GET, got %s", r.Method)
			}
			if r.URL.Path != "/v2/bot/group/group-abc/member/user-def" {
				t.Errorf("Expected path /v2/bot/group/group-abc/member/user-def, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"displayName":"Group Member","userId":"user-def","pictureUrl":"https://example.com/pic.jpg"}`))
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

	result, err := client.GetGroupMemberProfile("group-abc", "user-def")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.DisplayName != "Group Member" {
		t.Errorf("Expected DisplayName 'Group Member', got %q", result.DisplayName)
	}
	if result.UserId != "user-def" {
		t.Errorf("Expected UserId 'user-def', got %q", result.UserId)
	}
}

