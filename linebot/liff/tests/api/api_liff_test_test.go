package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/liff"
)

func TestAddLIFFApp(t *testing.T) {
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

	client, err := liff.NewLiffAPI(
		"MY_CHANNEL_TOKEN",
		liff.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.AddLIFFApp(
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestDeleteLIFFApp(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
		}),
	)

	client, err := liff.NewLiffAPI(
		"MY_CHANNEL_TOKEN",
		liff.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.DeleteLIFFApp(
		"hello",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestGetAllLIFFApps(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
		}),
	)

	client, err := liff.NewLiffAPI(
		"MY_CHANNEL_TOKEN",
		liff.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.GetAllLIFFApps()
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestUpdateLIFFApp(t *testing.T) {
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

	client, err := liff.NewLiffAPI(
		"MY_CHANNEL_TOKEN",
		liff.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.UpdateLIFFApp(
		"hello",

		nil,
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}
