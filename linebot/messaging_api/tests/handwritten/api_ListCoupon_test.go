package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestListCoupon_ItShouldCorrectlyPassQueryParameter(t *testing.T) {
	expectedLimit := "10"
	expectedStart := "some-start"
	expectedStatus := []string{"CLOSED", "RUNNING"}
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotLimit := r.URL.Query().Get("limit")
			gotStart := r.URL.Query().Get("start")
			r.ParseForm()
			gotStatus := r.Form["status"]
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
			if !reflect.DeepEqual(gotStatus, expectedStatus) {
				w.Header().Set("TEST-ERROR", fmt.Sprintf("incorrect status being sent from client. expected %s, got %s", expectedStatus, gotStatus))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(messaging_api.MessagingApiPagerCouponListResponse{
				Items: []messaging_api.CouponListResponse{
					{
						CouponId: "COUPON_ID_1",
						Title: "100Yen OFF",
					},
				},
				Next: "abcdef",
			})
		}),
	)
	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	status := []string{"CLOSED", "RUNNING"}

	resp, result, err := client.ListCouponWithHttpInfo(&status, "some-start", 10)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Not getting 200 response back: %s", resp.Header.Get("TEST-ERROR"))
	}
	if !reflect.DeepEqual(result.Items, []messaging_api.CouponListResponse{
		{
			CouponId: "COUPON_ID_1",
			Title: "100Yen OFF",
		},
	}) {
		t.Errorf("Expected Items: [{CouponId: COUPON_ID_1, Title: 100Yen OFF}], got: %v", result.Items)
	}
	if result.Next != "abcdef" {
		t.Errorf("Expected Next: abcdef, got: %s", result.Next)
	}
}
