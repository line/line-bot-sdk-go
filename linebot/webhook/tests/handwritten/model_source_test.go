package tests

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func TestStickerMessage(t *testing.T) {
	var cb webhook.CallbackRequest
	if err := json.Unmarshal([]byte(`{
		"destination": "Uaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"events": [
			{
				"type": "UNKNOWN",
				"great": "new-field"
			}
		]
	}`), &cb); err != nil {
		t.Fatalf("Failed to unmarshal callback request: %v", err)
	}

	_, ok := cb.Events[0].(webhook.UnknownEvent)
	if !ok {
		t.Fatalf("Failed to cast to UnknownEvent: %v", cb.Events[0])
	}
}

func generateSignature(secret string, body []byte) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func makeRequest(t *testing.T, url string, body []byte, signature string) *http.Request {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("X-Line-Signature", signature)
	return req
}

func TestWebhookParseRequest(t *testing.T) {
	const channelSecret = "testsecret"
	body := []byte(`{"destination":"U0123456789abcdef","events":[]}`)

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cb, err := webhook.ParseRequest(channelSecret, req)
		if err != nil {
			if err == webhook.ErrInvalidSignature {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			t.Errorf("unexpected error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if cb.Destination != "U0123456789abcdef" {
			t.Errorf("destination = %s; want %s", cb.Destination, "U0123456789abcdef")
		}
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewTLSServer(handler)
	defer server.Close()

	signature := generateSignature(channelSecret, body)
	req := makeRequest(t, server.URL, body, signature)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d; want %d", res.StatusCode, http.StatusOK)
	}
}

func TestWebhookParseRequestWithOption(t *testing.T) {
	const channelSecret = "testsecret"
	body := []byte(`{"destination":"U0123456789abcdef","events":[]}`)

	tests := []struct {
		name           string
		skipValidation bool
		useValidSig    bool
		expectedCode   int
	}{
		{
			name:           "valid signature, no skip",
			skipValidation: false,
			useValidSig:    true,
			expectedCode:   http.StatusOK,
		},
		{
			name:           "invalid signature, no skip",
			skipValidation: false,
			useValidSig:    false,
			expectedCode:   http.StatusBadRequest,
		},
		{
			name:           "invalid signature, but skip = true",
			skipValidation: true,
			useValidSig:    false,
			expectedCode:   http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				opt := &webhook.ParseOption{
					SkipSignatureValidation: func() bool { return tt.skipValidation },
				}
				cb, err := webhook.ParseRequestWithOption(channelSecret, req, opt)
				if err != nil {
					if err == webhook.ErrInvalidSignature {
						w.WriteHeader(http.StatusBadRequest)
						return
					}
					t.Errorf("unexpected error: %v", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				if cb.Destination != "U0123456789abcdef" {
					t.Errorf("destination = %s; want %s", cb.Destination, "U0123456789abcdef")
				}
				w.WriteHeader(http.StatusOK)
			})

			server := httptest.NewTLSServer(handler)
			defer server.Close()

			signature := "invalid"
			if tt.useValidSig {
				signature = generateSignature(channelSecret, body)
			}
			req := makeRequest(t, server.URL, body, signature)

			client := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			}
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			if res.StatusCode != tt.expectedCode {
				t.Errorf("StatusCode = %d; want %d", res.StatusCode, tt.expectedCode)
			}
		})
	}
}
