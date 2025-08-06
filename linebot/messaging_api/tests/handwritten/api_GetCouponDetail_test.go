package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	// "reflect"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestGetCouponDetailWithHttpInfo(t *testing.T) {
	expectedCouponId := "COUPON123"
	expectedBody := `{
			"acquisitionCondition": {
				"type": "lottery",
				"lotteryProbability": 50
			},
			"barcodeImageUrl": "https://example.com/barcode.png",
			"couponCode": "UNIQUECODE123",
			"description": "Get 100 Yen off your purchase",
			"endTimestamp": 1700000000,
			"imageUrl": "https://example.com/image.png",
			"maxAcquireCount": 1000,
			"maxUseCountPerTicket": 1,
			"startTimestamp": 1600000000,
			"title": "100 Yen OFF",
			"usageCondition": "Minimum purchase of 500 Yen",
			"reward": {
				"type": "discount",
				"priceInfo": {
					"type": "fixed",
					"fixedAmount": 100,
					"currency": "JPY"
				}
			},
			"visibility": "PUBLIC",
			"timezone": "ASIA_TOKYO",
			"couponId": "COUPON123",
			"createdTimestamp": 1600000000,
			"status": "RUNNING"
		}`

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify the request method
			if r.Method != http.MethodGet {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			// Verify the request path
			expectedPath := fmt.Sprintf("/v2/bot/coupon/%s", expectedCouponId)
			if r.URL.Path != expectedPath {
				t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
			}

			// Simulate a successful response with the expected coupon details
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(expectedBody))
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

	resp, result, err := client.GetCouponDetailWithHttpInfo(expectedCouponId)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if acquisitionCondition, ok := result.AcquisitionCondition.(messaging_api.LotteryAcquisitionConditionResponse); ok {
		if acquisitionCondition.LotteryProbability != 50 {
			t.Errorf("Expected AcquisitionCondition LotteryProbability 50, got %v", acquisitionCondition.LotteryProbability)
		}
	} else {
		t.Errorf("Expected AcquisitionCondition to be of type LotteryAcquisitionCondition, got %T", result.AcquisitionCondition)
	}

	if result.BarcodeImageUrl != "https://example.com/barcode.png" {
		t.Errorf("Expected BarcodeImageUrl 'https://example.com/barcode.png', got %s", result.BarcodeImageUrl)
	}
	if result.CouponCode != "UNIQUECODE123" {
		t.Errorf("Expected CouponCode 'UNIQUECODE123', got %s", result.CouponCode)
	}
	if result.Description != "Get 100 Yen off your purchase" {
		t.Errorf("Expected Description 'Get 100 Yen off your purchase', got %s", result.Description)
	}
	if result.EndTimestamp != 1700000000 {
		t.Errorf("Expected EndTimestamp 1700000000, got %d", result.EndTimestamp)
	}
	if result.ImageUrl != "https://example.com/image.png" {
		t.Errorf("Expected ImageUrl 'https://example.com/image.png', got %s", result.ImageUrl)
	}
	if result.MaxAcquireCount != 1000 {
		t.Errorf("Expected MaxAcquireCount 1000, got %d", result.MaxAcquireCount)
	}
	if result.MaxUseCountPerTicket != 1 {
		t.Errorf("Expected MaxUseCountPerTicket 1, got %d", result.MaxUseCountPerTicket)
	}
	if result.StartTimestamp != 1600000000 {
		t.Errorf("Expected StartTimestamp 1600000000, got %d", result.StartTimestamp)
	}
	if result.Title != "100 Yen OFF" {
		t.Errorf("Expected Title '100 Yen OFF', got %s", result.Title)
	}
	if result.UsageCondition != "Minimum purchase of 500 Yen" {
		t.Errorf("Expected UsageCondition 'Minimum purchase of 500 Yen', got %s", result.UsageCondition)
	}
	if reward, ok := result.Reward.(messaging_api.CouponDiscountRewardResponse); ok {
		if priceInfo, ok := reward.PriceInfo.(messaging_api.DiscountFixedPriceInfoResponse); ok {
			if priceInfo.FixedAmount != 100 {
				t.Errorf("Expected Reward PriceInfo FixedAmount 100, got %d", priceInfo.FixedAmount)
			}
			if priceInfo.Currency != "JPY" {
				t.Errorf("Expected Reward PriceInfo Currency 'JPY', got %s", priceInfo.Currency)
			}
		} else {
			t.Errorf("Expected Reward PriceInfo to be of type DiscountFixedPriceInfoResponse, got %T", reward.PriceInfo)
		}
	} else {
		t.Errorf("Expected Reward to be of type CouponDiscountRewardResponse, got %T", result.Reward)
	}
	if result.Visibility != messaging_api.CouponResponseVISIBILITY_PUBLIC {
		t.Errorf("Expected Visibility 'PUBLIC', got %s", result.Visibility)
	}
	if result.Timezone != messaging_api.CouponResponseTIMEZONE_ASIA_TOKYO {
		t.Errorf("Expected Timezone 'ASIA_TOKYO', got %s", result.Timezone)
	}
	if result.CouponId != expectedCouponId {
		t.Errorf("Expected CouponId '%s', got %s", expectedCouponId, result.CouponId)
	}
	if result.CreatedTimestamp != 1600000000 {
		t.Errorf("Expected CreatedTimestamp 1600000000, got %d", result.CreatedTimestamp)
	}
	if result.Status != "RUNNING" {
		t.Errorf("Expected Status 'RUNNING', got %s", result.Status)
	}
}
