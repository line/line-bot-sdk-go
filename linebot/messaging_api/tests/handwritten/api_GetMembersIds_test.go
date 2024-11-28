package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestGetGroupMembersIdsWithHttpInfo(t *testing.T) {
	tests := []struct {
		name                string
		groupId             string
		start               string
		expectedQueryParams url.Values
	}{
		{
			name:    "With start token",
			groupId: "testGroup",
			start:   "start",
			expectedQueryParams: url.Values{
				"start": {"start"},
			},
		},
		{
			name:                "Without start token",
			groupId:             "testGroup",
			start:               "",
			expectedQueryParams: url.Values{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if !reflect.DeepEqual(r.URL.Query(), tt.expectedQueryParams) {
					t.Fatalf("Expected query: %v, got: %v", tt.expectedQueryParams, r.URL.Query())
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(messaging_api.MembersIdsResponse{MemberIds: []string{"member1", "member2"}, Next: "abcdef"})
			}))
			defer mockServer.Close()

			client, err := messaging_api.NewMessagingApiAPI(
				"channelToken",
				messaging_api.WithEndpoint(mockServer.URL),
			)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			_, result, err := client.GetGroupMembersIdsWithHttpInfo(tt.groupId, tt.start)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(result.MemberIds, []string{"member1", "member2"}) {
				t.Errorf("Expected MemberIds: [\"member1\", \"member2\"], got: %s", result.MemberIds)
			}
			if result.Next != "abcdef" {
				t.Errorf("Expected Next: abcdef, got: %s", result.Next)
			}
		})
	}
}
