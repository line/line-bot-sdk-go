package linebot

import (
	"errors"
	"testing"
)

func TestBuildPath_DotSegmentsRejected(t *testing.T) {
	cases := []struct {
		name     string
		template string
		params   map[string]string
	}{
		{
			name:     "literal dot-dot",
			template: "/v2/bot/profile/{userId}",
			params:   map[string]string{"userId": ".."},
		},
		{
			name:     "literal dot",
			template: "/v2/bot/profile/{userId}",
			params:   map[string]string{"userId": "."},
		},
		{
			name:     "template contains dot segment",
			template: "/v2/bot/../profile/{userId}",
			params:   map[string]string{"userId": "Uxxx"},
		},
		{
			name:     "template dot segment no params",
			template: "/v2/bot/../secret",
			params:   nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := BuildPath(tc.template, tc.params)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !errors.Is(err, ErrPathTraversal) {
				t.Errorf("expected ErrPathTraversal, got %v", err)
			}
		})
	}
}

func TestBuildPath_EncodedDotsAllowed(t *testing.T) {
	cases := []struct {
		name string
		val  string
		want string
	}{
		{"%2e%2e", "%2e%2e", "/v2/bot/profile/%252e%252e"},
		{"%2e", "%2e", "/v2/bot/profile/%252e"},
		{"%2E%2E", "%2E%2E", "/v2/bot/profile/%252E%252E"},
		{".%2e", ".%2e", "/v2/bot/profile/.%252e"},
		{"%2e.", "%2e.", "/v2/bot/profile/%252e."},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := BuildPath("/v2/bot/profile/{userId}", map[string]string{"userId": tc.val})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func TestBuildPath_ReservedCharsEncoded(t *testing.T) {
	cases := []struct {
		name string
		val  string
		want string
	}{
		{"slash", "a/b", "/v2/bot/profile/a%2Fb"},
		{"question", "a?b", "/v2/bot/profile/a%3Fb"},
		{"hash", "a#b", "/v2/bot/profile/a%23b"},
		{"space", "a b", "/v2/bot/profile/a%20b"},
		{"backslash", `a\b`, "/v2/bot/profile/a%5Cb"},
		{"at", "a@b", "/v2/bot/profile/a@b"},
		{"traversal-with-slash", "../message/quota", "/v2/bot/profile/..%2Fmessage%2Fquota"},
		{"double-encoded-slash", "..%2Fmsg", "/v2/bot/profile/..%252Fmsg"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := BuildPath("/v2/bot/profile/{userId}", map[string]string{"userId": tc.val})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func TestBuildPath_PlaceholderBleed(t *testing.T) {
	got, err := BuildPath("/v2/bot/group/{groupId}/member/{userId}", map[string]string{
		"groupId": "{userId}",
		"userId":  "Ub",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "/v2/bot/group/%7BuserId%7D/member/Ub"
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

func TestBuildPath_NormalInput(t *testing.T) {
	got, err := BuildPath("/v2/bot/profile/{userId}", map[string]string{
		"userId": "U0047556f2e40dba2456887320ba7c76d",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "/v2/bot/profile/U0047556f2e40dba2456887320ba7c76d"
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

func TestBuildPath_DotsInMiddle(t *testing.T) {
	got, err := BuildPath("/v2/bot/profile/{userId}", map[string]string{
		"userId": "abc..def",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "/v2/bot/profile/abc..def"
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

func TestBuildPath_NilParams(t *testing.T) {
	got, err := BuildPath("/v2/bot/message/push", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "/v2/bot/message/push"
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

func TestBuildPath_MultipleParams(t *testing.T) {
	got, err := BuildPath("/v2/bot/user/{userId}/richmenu/{richMenuId}", map[string]string{
		"userId":     "Uabc",
		"richMenuId": "rm-123",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "/v2/bot/user/Uabc/richmenu/rm-123"
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

func TestBuildEndpoint(t *testing.T) {
	cases := []struct {
		name   string
		format string
		values []string
		want   string
	}{
		{
			name:   "single param",
			format: "/v2/bot/profile/%s",
			values: []string{"Uabc"},
			want:   "/v2/bot/profile/Uabc",
		},
		{
			name:   "slash in param",
			format: "/v2/bot/richmenu/%s",
			values: []string{"alias/list"},
			want:   "/v2/bot/richmenu/alias%2Flist",
		},
		{
			name:   "multiple params",
			format: "/v2/bot/group/%s/member/%s",
			values: []string{"Gabc", "Udef"},
			want:   "/v2/bot/group/Gabc/member/Udef",
		},
		{
			name:   "dot-dot encoded",
			format: "/v2/bot/profile/%s",
			values: []string{".."},
			want:   "/v2/bot/profile/..",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := buildEndpoint(tc.format, tc.values...)
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}
