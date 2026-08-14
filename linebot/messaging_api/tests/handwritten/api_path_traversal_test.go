package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func TestPathTraversal_GetProfile_DotSegmentsRejected(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("request should not be sent, but got path: %s", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	dotSegments := []string{
		"..",
		".",
	}

	for _, input := range dotSegments {
		t.Run(input, func(t *testing.T) {
			_, err := client.GetProfile(input)
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
			server := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if got := r.URL.EscapedPath(); got != tc.wantEscaped {
						t.Errorf("EscapedPath = %s; want %s", got, tc.wantEscaped)
					}
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"userId":"test","displayName":"Test"}`))
				}),
			)
			defer server.Close()

			client, err := messaging_api.NewMessagingApiAPI(
				"channelToken",
				messaging_api.WithEndpoint(server.URL),
			)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			_, err = client.GetProfile(tc.input)
			if err != nil {
				t.Errorf("percent-encoded dots should be double-encoded and allowed, but got error: %v", err)
			}
		})
	}
}

func TestPathTraversal_GetProfile_ReservedCharsEncoded(t *testing.T) {
	cases := []struct {
		input       string
		wantEscaped string
	}{
		{"../message/quota", "/v2/bot/profile/..%2Fmessage%2Fquota"},
		{"a/b", "/v2/bot/profile/a%2Fb"},
		{"a?b", "/v2/bot/profile/a%3Fb"},
		{"a#b", "/v2/bot/profile/a%23b"},
		{"..%2Fmessage%2Fquota", "/v2/bot/profile/..%252Fmessage%252Fquota"},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			server := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if got := r.URL.EscapedPath(); got != tc.wantEscaped {
						t.Errorf("EscapedPath = %s; want %s", got, tc.wantEscaped)
					}
					if strings.Contains(r.RequestURI, "%257B") || strings.Contains(r.RequestURI, "%257D") {
						t.Errorf("double-encoded path detected: %s", r.RequestURI)
					}
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"userId":"test","displayName":"Test"}`))
				}),
			)
			defer server.Close()

			client, err := messaging_api.NewMessagingApiAPI(
				"channelToken",
				messaging_api.WithEndpoint(server.URL),
			)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			_, err = client.GetProfile(tc.input)
			if err != nil {
				t.Errorf("reserved chars should be encoded, not rejected: %v", err)
			}
		})
	}
}

func TestPathTraversal_GetProfile_NormalInput(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantEscaped := "/v2/bot/profile/U0047556f2e40dba2456887320ba7c76d"
			if got := r.URL.EscapedPath(); got != wantEscaped {
				t.Errorf("EscapedPath = %s; want %s", got, wantEscaped)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"userId":"U0047556f2e40dba2456887320ba7c76d","displayName":"Test"}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	resp, err := client.GetProfile("U0047556f2e40dba2456887320ba7c76d")
	if err != nil {
		t.Errorf("unexpected error for normal input: %v", err)
	}
	if resp != nil && resp.UserId != "U0047556f2e40dba2456887320ba7c76d" {
		t.Errorf("unexpected userId: %s", resp.UserId)
	}
}

func TestPathTraversal_PlaceholderBleed(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantEscaped := "/v2/bot/user/%7BrichMenuId%7D/richmenu/real-menu-id"
			if got := r.URL.EscapedPath(); got != wantEscaped {
				t.Errorf("EscapedPath = %s; want %s", got, wantEscaped)
			}
			if strings.Contains(r.RequestURI, "%257B") {
				t.Errorf("double-encoded placeholder: %s", r.RequestURI)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.LinkRichMenuIdToUser("{richMenuId}", "real-menu-id")
	if err != nil {
		t.Errorf("placeholder in value should be escaped, not expanded: %v", err)
	}
}

func TestPathTraversal_GetProfile_DotsInMiddle(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantEscaped := "/v2/bot/profile/abc..def"
			if got := r.URL.EscapedPath(); got != wantEscaped {
				t.Errorf("EscapedPath = %s; want %s", got, wantEscaped)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"userId":"abc..def","displayName":"Test"}`))
		}),
	)
	defer server.Close()

	client, err := messaging_api.NewMessagingApiAPI(
		"channelToken",
		messaging_api.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.GetProfile("abc..def")
	if err != nil {
		t.Errorf("dots in middle of value should be allowed, but got error: %v", err)
	}
}
