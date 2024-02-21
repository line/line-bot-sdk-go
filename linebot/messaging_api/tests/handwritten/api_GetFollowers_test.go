package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestGetFollowers_ItShouldCorrectlyPassLimitAndStartQueryParameter(t *testing.T) {
	expectedLimit := "1000"
	expectedStart := "some-start"
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotLimit := r.URL.Query().Get("limit")
			gotStart := r.URL.Query().Get("start")
			if gotLimit != expectedLimit {
				w.Header().Set("TEST-ERROR", fmt.Sprintf("incorrect limit being sent from client. expected %s, got %s", expectedLimit, gotLimit))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if gotStart != expectedStart {
				w.Header().Set("TEST-ERROR", fmt.Sprintf("incorrect start being sent from client. expected %s, got %s", expectedStart, gotStart))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(messaging_api.GetFollowersResponse{UserIds: []string{}, Next: "abcdef"})
		}),
	)
	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, _, _ := client.GetFollowersWithHttpInfo("some-start", 1000)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Not getting 200 response back: %s", resp.Header.Get("TEST-ERROR"))
	}
}
