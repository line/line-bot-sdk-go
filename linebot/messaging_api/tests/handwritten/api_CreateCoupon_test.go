package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestCreateCoupon_ItShouldCorrectlySendRequestBody(t *testing.T) {
	expectedBody := `{
		"acquisitionCondition": {
			"type": "lottery",
			"lotteryProbability": 50,
			"maxAcquireCount": 1000
		},
		"barcodeImageUrl": "https://example.com/barcode.png",
		"couponCode": "UNIQUECODE123",
		"description": "Get 100 Yen off your purchase",
		"endTimestamp": 1700000000,
		"imageUrl": "https://example.com/image.png",
		"maxUseCountPerTicket": 1,
		"startTimestamp": 1600000000,
		"title": "100 Yen OFF",
		"usageCondition": "Minimum purchase of 500 Yen",
		"reward": {
			"type": "discount",
			"priceInfo": {
				"type": "fixed",
				"fixedAmount": 100
			}
		},
		"visibility": "PUBLIC",
		"timezone": "ASIA_TOKYO"
	}`


	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Read the request body
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var gotBody map[string]any
			var expectedBodyMap map[string]any

			// Unmarshal the received JSON
			if err := json.Unmarshal(bodyBytes, &gotBody); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Unmarshal the expected JSON
			if err := json.Unmarshal([]byte(expectedBody), &expectedBodyMap); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Compare the unmarshalled JSON objects
			if !reflect.DeepEqual(gotBody, expectedBodyMap) {
				w.Header().Set("TEST-ERROR", fmt.Sprintf("incorrect request body being sent from client. expected %v, got %v", expectedBodyMap, gotBody))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(messaging_api.CouponCreateResponse{
				CouponId: "COUPON_ID_1",
			})
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

	req := &messaging_api.CouponCreateRequest{
		AcquisitionCondition: &messaging_api.LotteryAcquisitionConditionRequest{
			LotteryProbability: 50,
			MaxAcquireCount:      1000,
		},
		BarcodeImageUrl:      "https://example.com/barcode.png",
		CouponCode:           "UNIQUECODE123",
		Description:          "Get 100 Yen off your purchase",
		EndTimestamp:         1700000000,
		ImageUrl:             "https://example.com/image.png",
		MaxUseCountPerTicket: 1,
		StartTimestamp:       1600000000,
		Title:                "100 Yen OFF",
		UsageCondition:       "Minimum purchase of 500 Yen",
		Reward:               &messaging_api.CouponDiscountRewardRequest{
			PriceInfo: &messaging_api.DiscountFixedPriceInfoRequest{
				FixedAmount: 100,
			},
		},
		Visibility:           messaging_api.CouponCreateRequestVISIBILITY_PUBLIC,
		Timezone:             messaging_api.CouponCreateRequestTIMEZONE_ASIA_TOKYO,
	}

	resp, result, err := client.CreateCouponWithHttpInfo(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Not getting 200 response back: %s", resp.Header.Get("TEST-ERROR"))
	}
	if result.CouponId != "COUPON_ID_1" {
		t.Errorf("Expected CouponId: COUPON_ID_1, got: %s", result.CouponId)
	}
}
