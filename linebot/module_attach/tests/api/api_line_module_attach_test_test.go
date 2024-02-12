package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/module_attach"
)

func TestAttachModule(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
				t.Fatalf("Invalid content-type: %s", r.Header.Get("Content-Type"))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
		}),
	)

	client, err := module_attach.NewLineModuleAttachAPI(
		"MY_CHANNEL_TOKEN",
		module_attach.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.AttachModule(
		"hello",

		"hello",

		"hello",

		"hello",

		"hello",

		"hello",

		"hello",

		"hello",

		"hello",

		"hello",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}
