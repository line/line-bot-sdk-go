package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestPathTraversal_GetProfile(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("request should not be sent, but got path: %s", r.URL.Path)
			w.WriteHeader(http.StatusOK)
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

	traversalInputs := []string{
		"../message/quota",
		"..%2Fmessage%2Fquota",
		"..",
		".",
	}

	for _, input := range traversalInputs {
		t.Run(input, func(t *testing.T) {
			_, err := client.GetProfile(input)
			if err == nil {
				t.Errorf("expected error for path traversal input %q, but got nil", input)
			}
		})
	}
}

func TestPathTraversal_GetProfile_NormalInput(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expected := "/v2/bot/profile/U0047556f2e40dba2456887320ba7c76d"
			if r.URL.Path != expected {
				t.Errorf("URLPath %s; want %s", r.URL.Path, expected)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"userId":"U0047556f2e40dba2456887320ba7c76d","displayName":"Test"}`))
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

	resp, err := client.GetProfile("U0047556f2e40dba2456887320ba7c76d")
	if err != nil {
		t.Errorf("unexpected error for normal input: %v", err)
	}
	if resp != nil && resp.UserId != "U0047556f2e40dba2456887320ba7c76d" {
		t.Errorf("unexpected userId: %s", resp.UserId)
	}
}

func TestPathTraversal_GetProfile_DotsInMiddle(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expected := "/v2/bot/profile/abc..def"
			if r.URL.Path != expected {
				t.Errorf("URLPath %s; want %s", r.URL.Path, expected)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"userId":"abc..def","displayName":"Test"}`))
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

	_, err = client.GetProfile("abc..def")
	if err != nil {
		t.Errorf("dots in middle of value should be allowed, but got error: %v", err)
	}
}
