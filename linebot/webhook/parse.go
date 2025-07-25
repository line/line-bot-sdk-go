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

type ParseOption struct {
	// SkipSignatureValidation is a function that determines whether to skip
	// webhook signature verification.
	//
	// If the function returns true, the signature verification step is skipped.
	// This can be useful in scenarios such as when you're in the process of updating
	// the channel secret and need to temporarily bypass verification to avoid disruptions.
	SkipSignatureValidation func() bool
}

// ParseRequestWithOption parses a LINE webhook request with optional behavior.
//
// Use this when you need to customize parsing, such as skipping signature validation
// via ParseOption. This is useful during channel secret rotation or local development.
//
// For standard use, prefer ParseRequest.
func ParseRequestWithOption(channelSecret string, r *http.Request, opt *ParseOption) (*CallbackRequest, error) {
	defer func() { _ = r.Body.Close() }()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	skip := opt != nil && opt.SkipSignatureValidation != nil && opt.SkipSignatureValidation()
	if !skip && !ValidateSignature(channelSecret, r.Header.Get("x-line-signature"), body) {
		return nil, ErrInvalidSignature
	}

	var cb CallbackRequest
	if err = json.Unmarshal(body, &cb); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request body: %w, %s", err, body)
	}
	return &cb, nil
}

// ParseRequest parses a LINE webhook request with signature verification.
//
// If you need to customize behavior (e.g. skip signature verification),
// use ParseRequestWithOption instead.
func ParseRequest(channelSecret string, r *http.Request) (*CallbackRequest, error) {
	return ParseRequestWithOption(channelSecret, r, nil)
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
