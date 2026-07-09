package tests

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestGetRichMenuImage_BinaryResponse(t *testing.T) {
	pngMagic := []byte{0x89, 0x50, 0x4E, 0x47}

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			expectedPath := "/v2/bot/richmenu/richmenu-abc/content"
			if r.URL.Path != expectedPath {
				t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
			}

			if auth := r.Header.Get("Authorization"); auth != "Bearer channelToken" {
				t.Errorf("Expected Authorization 'Bearer channelToken', got '%s'", auth)
			}

			w.Header().Set("Content-Type", "image/png")
			w.Write(pngMagic)
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiBlobAPI(
		"channelToken",
		messaging_api.WithBlobEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	resp, err := client.GetRichMenuImage("richmenu-abc")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("Expected non-nil response")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if !bytes.Equal(body, pngMagic) {
		t.Errorf("Expected body %v, got %v", pngMagic, body)
	}
}

func TestSetRichMenuImage_BinaryBody(t *testing.T) {
	imageData := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			expectedPath := "/v2/bot/richmenu/richmenu-xyz/content"
			if r.URL.Path != expectedPath {
				t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
			}

			if ct := r.Header.Get("Content-Type"); ct != "image/png" {
				t.Errorf("Expected Content-Type 'image/png', got '%s'", ct)
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("Failed to read request body: %v", err)
			}
			if !bytes.Equal(body, imageData) {
				t.Errorf("Expected body %v, got %v", imageData, body)
			}

			w.WriteHeader(http.StatusOK)
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiBlobAPI(
		"channelToken",
		messaging_api.WithBlobEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.SetRichMenuImage("richmenu-xyz", "image/png", bytes.NewReader(imageData))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
