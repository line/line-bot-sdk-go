package linebot

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPathTraversal_GetProfile_DotSegmentsRejected(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("request should not be sent, but got path: %s", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	dotSegments := []string{
		"..",
		".",
		"%2e%2e",
		"%2e",
		"%2E%2E",
		".%2e",
		"%2e.",
	}

	for _, input := range dotSegments {
		t.Run(input, func(t *testing.T) {
			_, err := client.GetProfile(input).Do()
			if err == nil {
				t.Errorf("expected error for dot segment %q, but got nil", input)
			}
		})
	}
}

func TestPathTraversal_GetProfile_SlashInValue(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("request should not be sent, but got path: %s", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	// Legacy API uses fmt.Sprintf (no encoding), so "../message/quota"
	// creates endpoint "/v2/bot/profile/../message/quota" which contains
	// ".." as a path segment and is correctly rejected.
	_, err = client.GetProfile("../message/quota").Do()
	if err == nil {
		t.Error("expected error for ../message/quota in legacy API")
	}
}

func TestPathTraversal_GetProfile_NormalInput(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expected := "/v2/bot/profile/U0047556f2e40dba2456887320ba7c76d"
		if r.URL.Path != expected {
			t.Errorf("URLPath %s; want %s", r.URL.Path, expected)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"userId":"U0047556f2e40dba2456887320ba7c76d","displayName":"Test","pictureUrl":"","statusMessage":""}`))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.GetProfile("U0047556f2e40dba2456887320ba7c76d").Do()
	if err != nil {
		t.Errorf("unexpected error for normal input: %v", err)
	}
	if resp != nil && resp.UserID != "U0047556f2e40dba2456887320ba7c76d" {
		t.Errorf("unexpected userId: %s", resp.UserID)
	}
}

func TestPathTraversal_GetProfile_DotsInMiddle(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expected := "/v2/bot/profile/abc..def"
		if r.URL.Path != expected {
			t.Errorf("URLPath %s; want %s", r.URL.Path, expected)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"userId":"abc..def","displayName":"Test","pictureUrl":"","statusMessage":""}`))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Unexpected Data API call")
		w.WriteHeader(404)
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.GetProfile("abc..def").Do()
	if err != nil {
		t.Errorf("dots in middle of value should be allowed, but got error: %v", err)
	}
}
