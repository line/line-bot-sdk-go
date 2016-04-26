package linebot

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetProfiles(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/v1/profiles" ||
			reflect.DeepEqual(r.URL.Query(), map[string]string{"mids": "TEST_MID"}) {
			t.Errorf("invalid request: %s %s", r.Method, r.URL.Path)
			return
		}
		if r.Header.Get("X-Line-ChannelID") != "1000000000" ||
			r.Header.Get("X-Line-ChannelSecret") != "testsecret" ||
			r.Header.Get("X-Line-Trusted-User-With-ACL") != "TEST_MID" {
			t.Errorf("invalid request header: %v", r.Header)
			return
		}
		w.Write([]byte(`{"contacts":[{"displayName":"BOT API","mid":"u0047556f2e40dba2456887320ba7c76d","pictureUrl":"http://example.com/abcdefghijklmn","statusMessage":"Hello, LINE!"}],"count":1,"display":1,"start":1,"total":1}`))
	}))
	client, err := mockClient(server)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := client.GetUserProfile([]string{"DUMMY_MID"})
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error("response is nil")
		return
	}
	if res.Count != 1 {
		t.Errorf("count: %v", res.Count)
		return
	}
	if len(res.Contacts) != 1 {
		t.Errorf("contacts: %v", res.Contacts)
		return
	}
	if res.Contacts[0].DisplayName != "BOT API" {
		t.Errorf("contact 0: %v", res.Contacts[0])
		return
	}
}
