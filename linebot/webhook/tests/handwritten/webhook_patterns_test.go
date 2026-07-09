package tests

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

// C4: Unknown event type falls back to UnknownEvent
func TestUnknownEventTypeFallback(t *testing.T) {
	var cb webhook.CallbackRequest
	err := json.Unmarshal([]byte(`{
		"destination": "Uaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"events": [
			{
				"type": "futureEventType",
				"timestamp": 1234567890,
				"source": {"type": "user", "userId": "U123"}
			}
		]
	}`), &cb)
	if err != nil {
		t.Fatalf("Should not error on unknown event type: %v", err)
	}
	if len(cb.Events) != 1 {
		t.Fatalf("Expected 1 event, got %d", len(cb.Events))
	}
	unknown, ok := cb.Events[0].(webhook.UnknownEvent)
	if !ok {
		t.Fatalf("Expected UnknownEvent, got %T", cb.Events[0])
	}
	if unknown.Type != "futureEventType" {
		t.Errorf("Expected type 'futureEventType', got %q", unknown.Type)
	}
}

// C7: PnpDeliveryCompletionEvent has type="delivery" (type/classname mismatch)
func TestPnpDeliveryCompletionEvent_TypeMismatch(t *testing.T) {
	var cb webhook.CallbackRequest
	err := json.Unmarshal([]byte(`{
		"destination": "Uaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"events": [
			{
				"type": "delivery",
				"timestamp": 1234567890,
				"mode": "active",
				"webhookEventId": "01H810E2VKWTRCJER1MMDP2BA5",
				"deliveryContext": {"isRedelivery": false},
				"delivery": {
					"data": "abc123"
				}
			}
		]
	}`), &cb)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}
	if len(cb.Events) != 1 {
		t.Fatalf("Expected 1 event, got %d", len(cb.Events))
	}
	event, ok := cb.Events[0].(webhook.PnpDeliveryCompletionEvent)
	if !ok {
		t.Fatalf("Expected PnpDeliveryCompletionEvent, got %T", cb.Events[0])
	}
	if event.Type != "delivery" {
		t.Errorf("Expected type 'delivery', got %q", event.Type)
	}
	if event.Delivery == nil || event.Delivery.Data != "abc123" {
		t.Errorf("Expected delivery.data='abc123', got %+v", event.Delivery)
	}
}

// D6: Webhook parsing of PnpDeliveryCompletionEvent (type/classname mismatch)
func TestWebhookParsePnpDeliveryCompletionEvent(t *testing.T) {
	const channelSecret = "testsecret"
	body := []byte(`{
		"destination": "U0123456789abcdef",
		"events": [
			{
				"type": "delivery",
				"timestamp": 1234567890,
				"mode": "active",
				"webhookEventId": "01H810E2VKWTRCJER1MMDP2BA5",
				"deliveryContext": {"isRedelivery": false},
				"delivery": {"data": "hashed_phone"}
			}
		]
	}`)

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cb, err := webhook.ParseRequest(channelSecret, req)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(cb.Events) != 1 {
			t.Errorf("expected 1 event, got %d", len(cb.Events))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		event, ok := cb.Events[0].(webhook.PnpDeliveryCompletionEvent)
		if !ok {
			t.Errorf("Expected PnpDeliveryCompletionEvent, got %T", cb.Events[0])
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if event.Delivery == nil || event.Delivery.Data != "hashed_phone" {
			t.Errorf("Expected delivery.data='hashed_phone', got %+v", event.Delivery)
		}
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewTLSServer(handler)
	defer server.Close()

	signature := generateSignature(channelSecret, body)
	req := makeRequest(t, server.URL, body, signature)

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	res, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d; want %d", res.StatusCode, http.StatusOK)
	}
}

// D8: Webhook parsing with unknown message type falls back to UnknownMessageContent
func TestWebhookParseUnknownMessageType(t *testing.T) {
	const channelSecret = "testsecret"
	body := []byte(`{
		"destination": "U0123456789abcdef",
		"events": [
			{
				"type": "message",
				"timestamp": 1234567890,
				"mode": "active",
				"webhookEventId": "01H810E2VKWTRCJER1MMDP2BA5",
				"deliveryContext": {"isRedelivery": false},
				"replyToken": "nHuyWiB7yP5Zw52FIkcQobQuGDXCTA",
				"source": {"type": "user", "userId": "U206d25c2ea6bd87c17655609a1c37cb8"},
				"message": {
					"id": "325708",
					"type": "futureMessageType",
					"customField": "customValue"
				}
			}
		]
	}`)

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cb, err := webhook.ParseRequest(channelSecret, req)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(cb.Events) != 1 {
			t.Errorf("expected 1 event, got %d", len(cb.Events))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		msgEvent, ok := cb.Events[0].(webhook.MessageEvent)
		if !ok {
			t.Errorf("Expected MessageEvent, got %T", cb.Events[0])
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		unknownMsg, ok := msgEvent.Message.(webhook.UnknownMessageContent)
		if !ok {
			t.Errorf("Expected UnknownMessageContent, got %T", msgEvent.Message)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if unknownMsg.Type != "futureMessageType" {
			t.Errorf("Expected message type 'futureMessageType', got %q", unknownMsg.Type)
		}
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewTLSServer(handler)
	defer server.Close()

	signature := generateSignature(channelSecret, body)
	req := makeRequest(t, server.URL, body, signature)

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	res, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d; want %d", res.StatusCode, http.StatusOK)
	}
}
