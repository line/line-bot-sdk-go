package webhook

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// ParseRequest func
func ParseRequest(channelSecret string, r *http.Request) (*CallbackRequest, error) {
	defer func() { _ = r.Body.Close() }()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if !linebot.ValidateSignature(channelSecret, r.Header.Get("x-line-signature"), body) {
		return nil, linebot.ErrInvalidSignature
	}

	var cb CallbackRequest
	if err = json.Unmarshal(body, &cb); err != nil {
		return nil, err
	}
	return &cb, nil
}
