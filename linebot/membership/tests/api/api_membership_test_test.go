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

	"github.com/line/line-bot-sdk-go/v8/linebot/membership"
)


func TestGetMembershipList(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
        }),
    )

	client, err := membership.NewMembershipAPI(
		"MY_CHANNEL_TOKEN",
        membership.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.GetMembershipList(
	
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
}

func TestGetMembershipSubscription(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
        }),
    )

	client, err := membership.NewMembershipAPI(
		"MY_CHANNEL_TOKEN",
        membership.WithEndpoint(server.URL),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	resp, err := client.GetMembershipSubscription(
	"hello",
        
	
	)
	if err != nil {
		t.Fatalf("Failed to call API: %v", err)
	}
}

