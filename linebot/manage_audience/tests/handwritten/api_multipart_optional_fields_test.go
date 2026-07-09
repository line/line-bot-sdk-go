package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/manage_audience"
)

func createTempFileWithContent(t *testing.T) *os.File {
	t.Helper()
	f, err := os.CreateTemp("", "test-upload-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	if _, err := f.WriteString("user-id-001\nuser-id-002\n"); err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}
	f.Sync()
	f.Seek(0, 0)
	return f
}

// TestAddUserIdsToAudience_PutMultipart (A16)
// Verifies PUT + multipart for AddUserIdsToAudience:
// correct HTTP method, path, content type, and form fields.
func TestAddUserIdsToAudience_PutMultipart(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			expectedPath := "/v2/bot/audienceGroup/upload/byFile"
			if r.URL.Path != expectedPath {
				t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
			}

			if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data; boundary=") {
				t.Errorf("Expected multipart/form-data content type, got %s", r.Header.Get("Content-Type"))
			}

			if err := r.ParseMultipartForm(10 << 20); err != nil {
				t.Fatalf("Failed to parse multipart form: %v", err)
			}

			// audienceGroupId should be present as a string
			if values, ok := r.MultipartForm.Value["audienceGroupId"]; !ok {
				t.Error("Expected audienceGroupId field to be present")
			} else if len(values) != 1 || values[0] != "12345" {
				t.Errorf("Expected audienceGroupId \"12345\", got %q", values)
			}

			// file field must be present and non-empty
			file, _, err := r.FormFile("file")
			if err != nil {
				t.Errorf("Expected file field to be present: %v", err)
			} else {
				content, _ := io.ReadAll(file)
				if len(content) == 0 {
					t.Error("Expected non-empty file content")
				}
				file.Close()
			}

			w.WriteHeader(http.StatusOK)
		}),
	)
	defer server.Close()

	tempFile := createTempFileWithContent(t)
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	client, err := manage_audience.NewManageAudienceBlobAPI(
		"channelToken",
		manage_audience.WithBlobEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.AddUserIdsToAudience(
		tempFile,
		12345,
		"test upload description",
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
