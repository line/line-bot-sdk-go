package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/channel_access_token"
)

func TestIssueStatelessChannelTokenWithHttpInfo(t *testing.T) {
	tests := []struct {
		name                string
		grantType           string
		clientAssertionType string
		clientAssertion     string
		clientId            string
		clientSecret        string
		expectedFormData    url.Values
	}{
		{
			name:                "Using client assertion",
			grantType:           "client_credentials",
			clientAssertionType: "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
			clientAssertion:     "assertion",
			clientId:            "",
			clientSecret:        "",
			expectedFormData: url.Values{
				"grant_type":            {"client_credentials"},
				"client_assertion_type": {"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
				"client_assertion":      {"assertion"},
			},
		},
		{
			name:                "Using clientId and clientSecret",
			grantType:           "client_credentials",
			clientAssertionType: "",
			clientAssertion:     "",
			clientId:            "client_id",
			clientSecret:        "client_secret",
			expectedFormData: url.Values{
				"grant_type":    {"client_credentials"},
				"client_id":     {"client_id"},
				"client_secret": {"client_secret"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if err := r.ParseForm(); err != nil {
					t.Fatalf("Failed to parse form: %v", err)
				}
				if !reflect.DeepEqual(r.PostForm, tt.expectedFormData) {
					t.Errorf("Expected form data: %v, got: %v", tt.expectedFormData, r.PostForm)
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(channel_access_token.IssueStatelessChannelAccessTokenResponse{
					AccessToken: "test_token",
					TokenType:   "Bearer",
					ExpiresIn:   3600,
				})
			}))
			defer mockServer.Close()

			client, err := channel_access_token.NewChannelAccessTokenAPI(
				channel_access_token.WithEndpoint(mockServer.URL),
			)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			_, result, err := client.IssueStatelessChannelTokenWithHttpInfo(
				tt.grantType,
				tt.clientAssertionType,
				tt.clientAssertion,
				tt.clientId,
				tt.clientSecret,
			)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result.AccessToken != "test_token" {
				t.Errorf("Expected AccessToken: test_token, got: %s", result.AccessToken)
			}
		})
	}
}
