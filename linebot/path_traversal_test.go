package linebot

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestPathTraversal_GetProfile_EncodedDotsAllowed(t *testing.T) {
	cases := []struct {
		input       string
		wantEscaped string
	}{
		{"%2e%2e", "/v2/bot/profile/%252e%252e"},
		{"%2e", "/v2/bot/profile/%252e"},
		{"%2E%2E", "/v2/bot/profile/%252E%252E"},
		{".%2e", "/v2/bot/profile/.%252e"},
		{"%2e.", "/v2/bot/profile/%252e."},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if got := r.URL.EscapedPath(); got != tc.wantEscaped {
					t.Errorf("EscapedPath = %s; want %s", got, tc.wantEscaped)
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"userId":"test","displayName":"Test","pictureUrl":"","statusMessage":""}`))
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

			_, err = client.GetProfile(tc.input).Do()
			if err != nil {
				t.Errorf("percent-encoded dots should be double-encoded and allowed, but got error: %v", err)
			}
		})
	}
}

func TestPathTraversal_GetProfile_SlashEncoded(t *testing.T) {
	cases := []struct {
		input       string
		wantEscaped string
	}{
		{"../message/quota", "/v2/bot/profile/..%2Fmessage%2Fquota"},
		{"a/b", "/v2/bot/profile/a%2Fb"},
		{"..%2Fmessage%2Fquota", "/v2/bot/profile/..%252Fmessage%252Fquota"},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if got := r.URL.EscapedPath(); got != tc.wantEscaped {
					t.Errorf("EscapedPath = %s; want %s", got, tc.wantEscaped)
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"userId":"test","displayName":"Test","pictureUrl":"","statusMessage":""}`))
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

			_, err = client.GetProfile(tc.input).Do()
			if err != nil {
				t.Errorf("slash in value should be encoded, not rejected: %v", err)
			}
		})
	}
}

func TestPathTraversal_GetRichMenu_EndpointCollision(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantEscaped := "/v2/bot/richmenu/alias%2Flist"
		if got := r.URL.EscapedPath(); got != wantEscaped {
			t.Errorf("EscapedPath = %s; want %s", got, wantEscaped)
		}
		if !strings.Contains(r.URL.EscapedPath(), "%2F") {
			t.Errorf("slash in richMenuID was not encoded, endpoint collision possible: %s", r.URL.EscapedPath())
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"richMenuId":"alias/list","size":{"width":2500,"height":1686},"selected":false,"areas":[]}`))
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

	_, err = client.GetRichMenu("alias/list").Do()
	if err != nil {
		t.Errorf("slash in richMenuID should be encoded: %v", err)
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

func TestPathTraversal_NewRawCall_DotSegmentsRejected(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("request should not be sent, but got path: %s", r.URL.Path)
		w.WriteHeader(http.StatusOK)
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

	traversalEndpoints := []string{
		"/../secret",
		"/v2/bot/../../secret",
		"/%2e%2e/secret",
		"/%2e./secret",
		"/.%2e/secret",
	}

	for _, ep := range traversalEndpoints {
		t.Run(ep, func(t *testing.T) {
			_, err := client.NewRawCall("GET", ep)
			if err == nil {
				t.Errorf("expected error for traversal endpoint %q, but got nil", ep)
			}
		})
	}
}
