package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/module"
)

func TestAcquireChatControl(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Header.Get("Content-Type") != "application/json; charset=UTF-8" {
				t.Fatalf("Invalid content-type: %s", r.Header.Get("Content-Type"))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
		}),
	)

	client, err := module.NewLineModuleAPI(
		"MY_CHANNEL_TOKEN",
		module.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.AcquireChatControl(
		"hello",

		nil,
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestDetachModule(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Header.Get("Content-Type") != "application/json; charset=UTF-8" {
				t.Fatalf("Invalid content-type: %s", r.Header.Get("Content-Type"))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
		}),
	)

	client, err := module.NewLineModuleAPI(
		"MY_CHANNEL_TOKEN",
		module.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.DetachModule(
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestGetModules(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
		}),
	)

	client, err := module.NewLineModuleAPI(
		"MY_CHANNEL_TOKEN",
		module.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.GetModules(
		"hello",

		0,
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestReleaseChatControl(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
		}),
	)

	client, err := module.NewLineModuleAPI(
		"MY_CHANNEL_TOKEN",
		module.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.ReleaseChatControl(
		"hello",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}
