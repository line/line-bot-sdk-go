package linebot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSendText(t *testing.T) {
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
		if !reflect.DeepEqual(message.To, []string{"DUMMY_MID"}) {
			t.Errorf("invalid request to: %v", message.To)
			return
		}

		if message.Content.Text != "hello!" {
			t.Errorf("invalid request content text: %v", message.Content.Text)
			return
		}
		if message.Content.ContentType != ContentTypeText {
			t.Errorf("invalid request content type: %v", message.Content.ContentType)
			return
		}
		if message.Content.ToType != RecipientTypeUser {
			t.Errorf("invalid request content to_type: %v", message.Content.ToType)
			return
		}

		w.Write([]byte(`{"failed":[],"messageId":"1347940533207","timestamp":1347940533207,"version":1}`))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := client.SendText([]string{"DUMMY_MID"}, "hello!")
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error("response is nil")
		return
	}
	if res.Version != 1 {
		t.Errorf("invalid version: %v", res.Version)
	}
	if res.MessageID != "1347940533207" {
		t.Errorf("invalid messageID: %v", res.MessageID)
	}
	if len(res.Failed) > 0 {
		t.Errorf("failed: %v", len(res.Failed))
	}
	if res.Timestamp != 1347940533207 {
		t.Errorf("invalid timestamp: %v", res.Timestamp)
	}
}

func TestSendImage(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var message SingleMessage
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			t.Error(err)
			return
		}
		if message.Content.ContentType != ContentTypeImage {
			t.Errorf("invalid request content type: %v", message.Content.ContentType)
			return
		}
		if message.Content.ToType != RecipientTypeUser {
			t.Errorf("invalid request content to_type: %v", message.Content.ToType)
			return
		}
		if message.Content.OriginalContentURL != "http://example.com/image.jpg" {
			t.Errorf("invalid request content original_content_url: %v", message.Content.OriginalContentURL)
			return
		}
		if message.Content.PreviewImageURL != "http://example.com/image_preview.jpg" {
			t.Errorf("invalid request content preview_image_url: %v", message.Content.PreviewImageURL)
			return
		}
		w.Write([]byte(`{"failed":[],"messageId":"1347940533207","timestamp":1347940533207,"version":1}`))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := client.SendImage([]string{"DUMMY_MID"}, "http://example.com/image.jpg", "http://example.com/image_preview.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error("response is nil")
		return
	}
}

func TestSendVideo(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var message SingleMessage
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			t.Error(err)
			return
		}
		if message.Content.ContentType != ContentTypeVideo {
			t.Errorf("invalid request content type: %v", message.Content.ContentType)
			return
		}
		if message.Content.ToType != RecipientTypeUser {
			t.Errorf("invalid request content to_type: %v", message.Content.ToType)
			return
		}
		if message.Content.OriginalContentURL != "http://example.com/video.mp4" {
			t.Errorf("invalid request content original_content_url: %v", message.Content.OriginalContentURL)
			return
		}
		if message.Content.PreviewImageURL != "http://example.com/image_preview.jpg" {
			t.Errorf("invalid request content preview_image_url: %v", message.Content.PreviewImageURL)
			return
		}
		w.Write([]byte(`{"failed":[],"messageId":"1347940533207","timestamp":1347940533207,"version":1}`))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := client.SendVideo([]string{"DUMMY_MID"}, "http://example.com/video.mp4", "http://example.com/image_preview.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error("response is nil")
		return
	}
}

func TestSendAudio(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var message SingleMessage
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			t.Error(err)
			return
		}
		if message.Content.ContentType != ContentTypeAudio {
			t.Errorf("invalid request content type: %v", message.Content.ContentType)
			return
		}
		if message.Content.ToType != RecipientTypeUser {
			t.Errorf("invalid request content to_type: %v", message.Content.ToType)
			return
		}
		if message.Content.OriginalContentURL != "http://example.com/audio.mp3" {
			t.Errorf("invalid request content original_content_url: %v", message.Content.OriginalContentURL)
			return
		}
		if !reflect.DeepEqual(message.Content.ContentMetaData, map[string]string{"AUDLEN": "3601"}) {
			t.Errorf("invalid request content content_meta_data: %v", message.Content.ContentMetaData)
			return
		}
		w.Write([]byte(`{"failed":[],"messageId":"1347940533207","timestamp":1347940533207,"version":1}`))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := client.SendAudio([]string{"DUMMY_MID"}, "http://example.com/audio.mp3", 3601)
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error("response is nil")
		return
	}
}

func TestSendLocation(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var message SingleMessage
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			t.Error(err)
			return
		}
		if message.Content.ContentType != ContentTypeLocation {
			t.Errorf("invalid request content type: %v", message.Content.ContentType)
			return
		}
		if message.Content.ToType != RecipientTypeUser {
			t.Errorf("invalid request content to_type: %v", message.Content.ToType)
			return
		}
		if message.Content.Location.Title != "位置ラベル" || message.Content.Location.Address != "tokyo shibuya-ku" {
			t.Errorf("invalid location: %v", message.Content.Location)
			return
		}
		w.Write([]byte(`{"failed":[],"messageId":"1347940533207","timestamp":1347940533207,"version":1}`))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := client.SendLocation([]string{"DUMMY_MID"}, "位置ラベル", "tokyo shibuya-ku", 35.61823286112982, 139.72824096679688)
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error("response is nil")
		return
	}
}

func TestSendSticker(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var message SingleMessage
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			t.Error(err)
			return
		}
		if message.Content.ContentType != ContentTypeSticker {
			t.Errorf("invalid request content type: %v", message.Content.ContentType)
			return
		}
		if message.Content.ToType != RecipientTypeUser {
			t.Errorf("invalid request content to_type: %v", message.Content.ToType)
			return
		}
		if !reflect.DeepEqual(message.Content.ContentMetaData, map[string]string{"STKID": "1", "STKPKGID": "2", "STKVER": "3"}) {
			t.Errorf("invalid request content content_meta_data: %v", message.Content.ContentMetaData)
			return
		}
		w.Write([]byte(`{"failed":[],"messageId":"1347940533207","timestamp":1347940533207,"version":1}`))
	}))
	defer server.Close()
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := client.SendSticker([]string{"DUMMY_MID"}, 1, 2, 3)
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error("response is nil")
		return
	}
}
