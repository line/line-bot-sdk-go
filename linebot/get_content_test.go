package linebot

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetContent(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/v1/bot/message/123456789/content" {
			t.Errorf("invalid request: %s %s", r.Method, r.URL.Path)
			return
		}
		if r.Header.Get("X-Line-ChannelID") != "1000000000" ||
			r.Header.Get("X-Line-ChannelSecret") != "testsecret" ||
			r.Header.Get("X-Line-Trusted-User-With-ACL") != "TEST_MID" {
			t.Errorf("invalid request header: %v", r.Header)
			return
		}
		w.Header().Set("Content-type", "image/jpeg")
		w.Header().Set("Content-Disposition", `attachment; filename="image.jpg"`)
		w.WriteHeader(200)
	}))
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
		return
	}

	// success
	{
		res, err := client.GetMessageContent(&ReceivedContent{
			ID:          "123456789",
			ContentType: ContentTypeImage,
		})
		if err != nil {
			t.Error(err)
			return
		}
		if res == nil {
			t.Error("response is nil")
			return
		}
		defer res.Content.Close()
		if res.FileName != "image.jpg" {
			t.Errorf("content: %v", res)
			return
		}
	}
}

func TestGetContentPreview(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/v1/bot/message/123456789/content/preview" {
			t.Errorf("invalid request: %s %s", r.Method, r.URL.Path)
			return
		}
		if r.Header.Get("X-Line-ChannelID") != "1000000000" ||
			r.Header.Get("X-Line-ChannelSecret") != "testsecret" ||
			r.Header.Get("X-Line-Trusted-User-With-ACL") != "TEST_MID" {
			t.Errorf("invalid request header: %v", r.Header)
			return
		}
		w.Header().Set("Content-type", "image/jpeg")
		w.Header().Set("Content-Disposition", `attachment; filename="image.jpg"`)
		w.WriteHeader(200)
	}))
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
		return
	}

	// success
	{
		res, err := client.GetMessageContentPreview(&ReceivedContent{
			ID:          "123456789",
			ContentType: ContentTypeImage,
		})
		if err != nil {
			t.Error(err)
			return
		}
		if res == nil {
			t.Error("response is nil")
			return
		}
		defer res.Content.Close()
		if res.FileName != "image.jpg" {
			t.Errorf("content: %v", res)
			return
		}
	}
}
