package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestCallingAPITwice(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/v2/bot/message/push" {
				t.Errorf("Expected path '/v2/bot/message/push', but got '%s'", r.URL.Path)
			}
			w.Header().Add("X-Line-Request-Id", "1234567890")
			w.Write([]byte(`{}`))
		}),
	)
	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	_, err = client.PushMessage(&messaging_api.PushMessageRequest{
		Messages: []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: "Hello, world",
			},
		},
	}, "")
	if err != nil {
		t.Fatalf("Failed to get response: %v", err)
	}

	// call again
	_, err = client.PushMessage(&messaging_api.PushMessageRequest{
		Messages: []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: "Hello, world",
			},
		},
	}, "")
	if err != nil {
		t.Fatalf("Failed to get response: %v", err)
	}
}
