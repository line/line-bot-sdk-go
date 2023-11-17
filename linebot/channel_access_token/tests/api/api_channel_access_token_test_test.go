package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/channel_access_token"
)

func TestGetsAllValidChannelAccessTokenKeyIds(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
		}),
	)

	client, err := channel_access_token.NewChannelAccessTokenAPI(

		channel_access_token.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.GetsAllValidChannelAccessTokenKeyIds(
		"hello",

		"hello",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestIssueChannelToken(t *testing.T) {
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

	client, err := channel_access_token.NewChannelAccessTokenAPI(

		channel_access_token.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.IssueChannelToken(
		"hello",

		"hello",

		"hello",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestIssueChannelTokenByJWT(t *testing.T) {
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

	client, err := channel_access_token.NewChannelAccessTokenAPI(

		channel_access_token.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.IssueChannelTokenByJWT(
		"hello",

		"hello",

		"hello",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestIssueStatelessChannelToken(t *testing.T) {
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

	client, err := channel_access_token.NewChannelAccessTokenAPI(

		channel_access_token.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.IssueStatelessChannelToken(
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

func TestRevokeChannelToken(t *testing.T) {
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

	client, err := channel_access_token.NewChannelAccessTokenAPI(

		channel_access_token.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.RevokeChannelToken(
		"hello",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestRevokeChannelTokenByJWT(t *testing.T) {
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

	client, err := channel_access_token.NewChannelAccessTokenAPI(

		channel_access_token.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.RevokeChannelTokenByJWT(
		"hello",

		"hello",

		"hello",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestVerifyChannelToken(t *testing.T) {
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

	client, err := channel_access_token.NewChannelAccessTokenAPI(

		channel_access_token.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.VerifyChannelToken(
		"hello",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestVerifyChannelTokenByJWT(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
		}),
	)

	client, err := channel_access_token.NewChannelAccessTokenAPI(

		channel_access_token.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.VerifyChannelTokenByJWT(
		"hello",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}
