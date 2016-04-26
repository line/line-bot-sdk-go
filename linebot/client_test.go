package linebot

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
)

func mockClient(server *httptest.Server) (*Client, error) {
	client, err := NewClient(
		1000000000,
		"testsecret",
		"TEST_MID",
		WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}),
		WithEndpointBase(server.URL),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}
