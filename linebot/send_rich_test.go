package linebot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendRich(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/v1/events" {
			t.Errorf("invalid request: %s %s", r.Method, r.URL.Path)
			return
		}
		if r.Header.Get("Content-Type") != "application/json; charset=UTF-8" ||
			r.Header.Get("X-Line-ChannelID") != "1000000000" ||
			r.Header.Get("X-Line-ChannelSecret") != "testsecret" ||
			r.Header.Get("X-Line-Trusted-User-With-ACL") != "TEST_MID" {
			t.Errorf("invalid request header: %v", r.Header)
			return
		}

		defer r.Body.Close()
		var message SingleMessage
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			t.Error(err)
			return
		}
		if message.EventType != "138311608800106203" {
			t.Errorf("invalid request event type: %v", message.EventType)
			return
		}
		if message.ToChannel != 1383378250 {
			t.Errorf("invalid request to_channel: %v", message.ToChannel)
			return
		}
		if len(message.To) != 1 || message.To[0] != "DUMMY_MID" {
			t.Errorf("invalid request to: %v", message.To)
			return
		}
		if message.Content.ContentMetaData["DOWNLOAD_URL"] != "https://example.com/rich-image/foo" ||
			message.Content.ContentMetaData["SPEC_REV"] != "1" ||
			message.Content.ContentMetaData["ALT_TEXT"] != "This is a alt text." {
			t.Errorf("invalid content_metadata: %v", message.Content.ContentMetaData)
		}
		// TODO check MARKUP_JSON
		w.Write([]byte(`{"failed":[],"messageId":"1347940533207","timestamp":1347940533207,"version":1}`))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := client.NewRichMessage(1040).
		SetAction("MANGA", "manga", "https://example.com/family/manga/en").
		SetListener("MANGA", 0, 0, 520, 520).
		Send([]string{"DUMMY_MID"}, "https://example.com/rich-image/foo", "This is a alt text.")
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error("response is nil")
	}
}
