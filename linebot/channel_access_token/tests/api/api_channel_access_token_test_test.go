package tests

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func TestIssueStatelessChannelTokenByJWTAssertion(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
				t.Fatalf("Invalid content-type: %s", r.Header.Get("Content-Type"))
				return
			}
			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			params, err := url.ParseQuery(string(body))
			if err != nil {
				t.Fatal(err)
			}
			if params.Get("grant_type") != "client_credentials" {
				t.Fatalf("Expected grant_type=client_credentials, got %s", params.Get("grant_type"))
			}
			if params.Get("client_assertion_type") != "urn:ietf:params:oauth:client-assertion-type:jwt-bearer" {
				t.Fatalf("Expected client_assertion_type=urn:ietf:params:oauth:client-assertion-type:jwt-bearer, got %s", params.Get("client_assertion_type"))
			}
			if params.Get("client_assertion") != "my_assertion" {
				t.Fatalf("Expected client_assertion=my_assertion, got %s", params.Get("client_assertion"))
			}
			if params.Get("client_id") != "" {
				t.Fatalf("Expected client_id to be empty, got %s", params.Get("client_id"))
			}
			if params.Get("client_secret") != "" {
				t.Fatalf("Expected client_secret to be empty, got %s", params.Get("client_secret"))
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"access_token":"token","expires_in":900,"token_type":"Bearer"}`))
		}),
	)

	client, err := channel_access_token.NewChannelAccessTokenAPI(
		channel_access_token.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.IssueStatelessChannelTokenByJWTAssertion(
		"my_assertion",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}

func TestIssueStatelessChannelTokenByClientSecret(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
				t.Fatalf("Invalid content-type: %s", r.Header.Get("Content-Type"))
				return
			}
			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			params, err := url.ParseQuery(string(body))
			if err != nil {
				t.Fatal(err)
			}
			if params.Get("grant_type") != "client_credentials" {
				t.Fatalf("Expected grant_type=client_credentials, got %s", params.Get("grant_type"))
			}
			if params.Get("client_assertion_type") != "" {
				t.Fatalf("Expected client_assertion_type to be empty, got %s", params.Get("client_assertion_type"))
			}
			if params.Get("client_assertion") != "" {
				t.Fatalf("Expected client_assertion to be empty, got %s", params.Get("client_assertion"))
			}
			if params.Get("client_id") != "my_id" {
				t.Fatalf("Expected client_id=my_id, got %s", params.Get("client_id"))
			}
			if params.Get("client_secret") != "my_secret" {
				t.Fatalf("Expected client_secret=my_secret, got %s", params.Get("client_secret"))
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"access_token":"token","expires_in":900,"token_type":"Bearer"}`))
		}),
	)

	client, err := channel_access_token.NewChannelAccessTokenAPI(
		channel_access_token.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.IssueStatelessChannelTokenByClientSecret(
		"my_id",
		"my_secret",
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
	log.Printf("Got response: %v", resp)
}
