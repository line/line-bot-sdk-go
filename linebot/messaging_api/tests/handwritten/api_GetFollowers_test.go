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
	start := "some-start"
	var limit int32 = 1000
	resp, _, _ := client.GetFollowersWithHttpInfo(&start, &limit)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Not getting 200 response back: %s", resp.Header.Get("TEST-ERROR"))
	}
}

func TestGetFollowers_ItShouldNotPassStartWhenItIsEmptyString(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Has("start") {
				w.Header().Set("TEST-ERROR", "incorrect start being sent from client. does not expect to have `start` in query-parameter")
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
	var limit int32 = 1000
	resp, _, _ := client.GetFollowersWithHttpInfo(nil, &limit)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Not getting 200 response back: %s", resp.Header.Get("TEST-ERROR"))
	}
}
