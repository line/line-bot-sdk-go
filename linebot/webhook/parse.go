package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrInvalidSignature = errors.New("invalid signature")
)

// ParseRequest func
func ParseRequest(channelSecret string, r *http.Request) (*CallbackRequest, error) {
	defer func() { _ = r.Body.Close() }()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if !ValidateSignature(channelSecret, r.Header.Get("x-line-signature"), body) {
		return nil, ErrInvalidSignature
	}

	var cb CallbackRequest
	if err = json.Unmarshal(body, &cb); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request body: %w, %s", err, body)
	}
	return &cb, nil
}

func ValidateSignature(channelSecret, signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(channelSecret))

	_, err = hash.Write(body)
	if err != nil {
		return false
	}

	return hmac.Equal(decoded, hash.Sum(nil))
}
