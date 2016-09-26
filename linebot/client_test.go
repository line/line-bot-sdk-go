package linebot

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestNewClient(t *testing.T) {
	secret := "testsecret"
	token := "testtoken"
	client, err := New(secret, token)
	if err != nil {
		t.Error(err)
		return
	}
	if client.channelSecret != secret {
		t.Errorf("channelSecret %s; want %s", client.channelSecret, secret)
	}
	if client.channelToken != token {
		t.Errorf("channelToken %s; want %s", client.channelSecret, secret)
	}
	if client.endpointBase != APIEndpointBase {
		t.Errorf("endpointBase %s; want %s", client.endpointBase, APIEndpointBase)
	}
	if client.httpClient != http.DefaultClient {
		t.Errorf("httpClient %p; want %p", client.httpClient, http.DefaultClient)
	}
}

func TestNewClientWithOptions(t *testing.T) {
	secret := "testsecret"
	token := "testtoken"
	endpoint := "https://example.test/"
	httpClient := http.Client{}
	client, err := New(
		secret,
		token,
		WithHTTPClient(&httpClient),
		WithEndpointBase(endpoint),
	)
	if err != nil {
		t.Error(err)
		return
	}
	if client.endpointBase != endpoint {
		t.Errorf("endpointBase %s; want %s", client.endpointBase, APIEndpointBase)
	}
	if client.httpClient != &httpClient {
		t.Errorf("httpClient %p; want %p", client.httpClient, &httpClient)
	}
}
