package linebot

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
)

func mockClient(server *httptest.Server) (*Client, error) {
	client, err := New(
		"testsecret",
		"testtoken",
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
