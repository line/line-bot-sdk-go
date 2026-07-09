package linebot

import (
	"errors"
	"net/url"
	"strings"
)

var ErrPathTraversal = errors.New("path parameter must not perform path traversal")

func ValidatePathParam(value string) error {
	if strings.ContainsAny(value, "/\\") {
		return ErrPathTraversal
	}
	lower := strings.ToLower(value)
	if strings.Contains(lower, "%2f") {
		return ErrPathTraversal
	}
	normalized := strings.NewReplacer("%2e", ".", "%2E", ".").Replace(value)
	if normalized == "." || normalized == ".." {
		return ErrPathTraversal
	}
	return nil
}

func validateEndpoint(endpoint string) error {
	decoded, _ := url.PathUnescape(endpoint)
	for _, segment := range strings.Split(decoded, "/") {
		clean := strings.NewReplacer("%2e", ".", "%2E", ".").Replace(segment)
		if clean == "." || clean == ".." {
			return ErrPathTraversal
		}
	}
	return nil
}

func BuildPath(pathTemplate string, params map[string]string) (string, error) {
	for name, value := range params {
		if err := ValidatePathParam(value); err != nil {
			return "", err
		}
		pathTemplate = strings.ReplaceAll(pathTemplate, "{"+name+"}", url.PathEscape(value))
	}
	return pathTemplate, nil
}
