package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestTextMessageWithHttpInfo(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	resp, _, err := client.PushMessageWithHttpInfo(&messaging_api.PushMessageRequest{
		Messages: []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: "Hello, world",
			},
		},
	}, "")
	if err != nil {
		t.Fatalf("Failed to create audience: %v", err)
	}
	log.Printf("Got response: %v", resp)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status: %d", resp.StatusCode)
	}
	if resp.Header.Get("X-Line-Request-Id") != "1234567890" {
		t.Errorf("X-Line-Request-Id: %s", resp.Header.Get("X-Line-Request-Id"))
	}
}
