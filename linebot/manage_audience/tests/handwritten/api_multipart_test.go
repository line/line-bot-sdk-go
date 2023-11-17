package tests

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/line/line-bot-sdk-go/v8/linebot/manage_audience"
)

func generateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40 // Set the version to 4
	b[8] = (b[8] & 0x3f) | 0x80 // Set the variant to 10
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func TestStickerMessage(t *testing.T) {
	// Write some random UUIDs to the temporary file.
	tempFile, err := os.CreateTemp("", "temp-uuids-")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Delete the temporary file after use
	defer tempFile.Close()           // Close the file after use

	// Generate some random UUIDs
	for i := 0; i < 10; i++ {
		randomUUID, err := generateUUID()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// Write the UUIDs to the temporary file
		_, err = tempFile.WriteString(randomUUID + "\n")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	tempFile.Sync()
	tempFile.Seek(0, 0)

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.Header.Get("Content-type"), "multipart/form-data; boundary=") {
				t.Fatalf("Unexpected content-type: %v", r.Header.Get("Content-type"))
			}

			err := r.ParseMultipartForm(10 << 20) // Limit: 10MB
			if err != nil {
				log.Println("Cannot parse multipart form")
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Get the 'file' field from the form
			file, _, err := r.FormFile("file")
			if err != nil {
				log.Println("Error Retrieving the File")
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()

			// Check if the file is empty
			fileContent, err := io.ReadAll(file)
			if err != nil {
				log.Println("Error reading the file")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if len(fileContent) == 0 {
				log.Println("Empty file")
				http.Error(w, "Empty file", http.StatusBadRequest)
				return
			}

			w.Write([]byte(`{
				"audienceGroupId": 1234567890123,
				"createRoute": "MESSAGING_API",
				"type": "UPLOAD",
				"description": "audienceGroupName_01",
				"created": 1613700237,
				"permission": "READ_WRITE",
				"expireTimestamp": 1629252237,
				"isIfaAudience": false
			  }`))
		}),
	)
	client, err := manage_audience.NewManageAudienceBlobAPI(
		"channelToken",
		manage_audience.WithBlobEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.CreateAudienceForUploadingUserIds(tempFile, "hello", true, "foobar")
	if err != nil {
		t.Fatalf("Failed to create audience: %v", err)
	}
	log.Printf("Got response: %v", resp)
}
