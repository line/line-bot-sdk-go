package linebot

import (
	"errors"
	"net/url"
	"strings"
)

// ErrPathTraversal is returned when a path parameter or endpoint path
// contains a dot segment that could change the resolved endpoint.
var ErrPathTraversal = errors.New("path parameter must not perform path traversal")

// BuildPath substitutes path parameters into a URL path template,
// percent-encodes each value, and validates the final path against
// dot-segment traversal.
func BuildPath(pathTemplate string, params map[string]string) (string, error) {
	for name, value := range params {
		pathTemplate = strings.ReplaceAll(pathTemplate, "{"+name+"}", url.PathEscape(value))
	}
	if err := validateEscapedPath(pathTemplate); err != nil {
		return "", err
	}
	return pathTemplate, nil
}

func validateEndpoint(endpoint string) error {
	decoded, err := url.PathUnescape(endpoint)
	if err != nil {
		return err
	}
	for _, segment := range strings.Split(decoded, "/") {
		if isDotSegment(segment) {
			return ErrPathTraversal
		}
	}
	return nil
}

func validateEscapedPath(p string) error {
	for _, segment := range strings.Split(p, "/") {
		if isDotSegment(segment) {
			return ErrPathTraversal
		}
	}
	return nil
}

func isDotSegment(s string) bool {
	normalized := strings.ReplaceAll(strings.ToLower(s), "%2e", ".")
	return normalized == "." || normalized == ".."
}
