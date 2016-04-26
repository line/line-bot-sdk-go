package linebot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSendMultiple(t *testing.T) {
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
		var message MultipleMessage
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			t.Error(err)
			return
		}
		if message.EventType != "140177271400161403" {
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
		if len(message.Content.Messages) != 6 ||
			message.Content.Messages[0].ContentType != ContentTypeText ||
			message.Content.Messages[1].ContentType != ContentTypeImage ||
			message.Content.Messages[2].ContentType != ContentTypeVideo ||
			message.Content.Messages[3].ContentType != ContentTypeAudio ||
			message.Content.Messages[4].ContentType != ContentTypeLocation ||
			message.Content.Messages[5].ContentType != ContentTypeSticker {
			t.Errorf("invalid messages: %v", message.Content.Messages)
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

	res, err := client.NewMultipleMessage().
		AddText("hello!").
		AddImage("http://example.com/image.jpg", "http://example.com/image_preview.jpg").
		AddVideo("http://example.com/video.mp4", "http://example.com/image_preview.jpg").
		AddAudio("http://example.com/audio.mp3", 3601).
		AddLocation("位置ラベル", "tokyo shibuya-ku", 35.61823286112982, 139.72824096679688).
		AddSticker(1, 2, 3).
		Send([]string{"DUMMY_MID"})
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error("response is nil")
		return
	}
}
