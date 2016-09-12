package linebot

import (
	"bytes"
	"errors"
	"fmt"
)

// errors
var (
	ErrInvalidSignature   = errors.New("Invalid Signature")
	ErrInvalidContentType = errors.New("Invalid ContentType")
	ErrInvalidEventType   = errors.New("Invalid EventType")
	ErrUnknown            = errors.New("Unknown Error")
)

// APIError type
type APIError struct {
	Code     int
	Response *ErrorResponse
}

// Error method
func (e *APIError) Error() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "linebot: APIError %d ", e.Code)
	if e.Response != nil {
		fmt.Fprintf(&buf, "%s", e.Response.Message)
		for _, d := range e.Response.Details {
			fmt.Fprintf(&buf, "\n[%s] %s", d.Property, d.Message)
		}
	}
	return buf.String()
}
