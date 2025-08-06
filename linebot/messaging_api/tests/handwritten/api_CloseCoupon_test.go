package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestCloseCouponWithHttpInfo(t *testing.T) {
	expectedCouponId := "COUPON123"

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify the request method
			if r.Method != http.MethodPut {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			// Verify the request path
			expectedPath := fmt.Sprintf("/v2/bot/coupon/%s/close", expectedCouponId)
			if r.URL.Path != expectedPath {
				t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
			}

			// Simulate a successful response
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

	resp, _, err := client.CloseCouponWithHttpInfo(expectedCouponId)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}
