/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py

package messaging_api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type MessagingApiBlobAPI struct {
	httpClient   *http.Client
	endpoint     *url.URL
	channelToken string
	ctx          context.Context
}

// MessagingApiBlobAPIOption type
type MessagingApiBlobAPIOption func(*MessagingApiBlobAPI) error

// New returns a new bot client instance.
func NewMessagingApiBlobAPI(channelToken string, options ...MessagingApiBlobAPIOption) (*MessagingApiBlobAPI, error) {
	if channelToken == "" {
		return nil, errors.New("missing channel access token")
	}

	c := &MessagingApiBlobAPI{
		channelToken: channelToken,
		httpClient:   http.DefaultClient,
	}

	u, err := url.ParseRequestURI("https://api-data.line.me")
	if err != nil {
		return nil, err
	}
	c.endpoint = u

	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// WithContext method
func (call *MessagingApiBlobAPI) WithContext(ctx context.Context) *MessagingApiBlobAPI {
	call.ctx = ctx
	return call
}

func (client *MessagingApiBlobAPI) Do(req *http.Request) (*http.Response, error) {
	if client.channelToken != "" {
		req.Header.Set("Authorization", "Bearer "+client.channelToken)
	}
	req.Header.Set("User-Agent", "LINE-BotSDK-Go/"+linebot.GetVersion())
	if client.ctx != nil {
		req = req.WithContext(client.ctx)
	}
	return client.httpClient.Do(req)
}

func (client *MessagingApiBlobAPI) Url(endpointPath string) string {
	u := client.endpoint
	u.Path = path.Join(u.Path, endpointPath)
	return u.String()
}

// WithHTTPClient function
func WithBlobHTTPClient(c *http.Client) MessagingApiBlobAPIOption {
	return func(client *MessagingApiBlobAPI) error {
		client.httpClient = c
		return nil
	}
}

// WithEndpointClient function
func WithBlobEndpoint(endpoint string) MessagingApiBlobAPIOption {
	return func(client *MessagingApiBlobAPI) error {
		u, err := url.ParseRequestURI(endpoint)
		if err != nil {
			return err
		}
		client.endpoint = u
		return nil
	}
}

// GetMessageContent
//
// Download image, video, and audio data sent from users.
// Parameters:
//        messageId             Message ID of video or audio

// You must close the response body when finished with it.
// https://developers.line.biz/en/reference/messaging-api/#get-content
func (client *MessagingApiBlobAPI) GetMessageContent(

	messageId string,

) (*http.Response, error) {
	_, body, error := client.GetMessageContentWithHttpInfo(

		messageId,
	)
	return body, error
}

// GetMessageContent
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Download image, video, and audio data sent from users.
// Parameters:
//        messageId             Message ID of video or audio

// You must close the response body when finished with it.
// https://developers.line.biz/en/reference/messaging-api/#get-content
func (client *MessagingApiBlobAPI) GetMessageContentWithHttpInfo(

	messageId string,

) (*http.Response, *http.Response, error) {
	path := "/v2/bot/message/{messageId}/content"

	path = strings.Replace(path, "{messageId}", messageId, -1)

	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return res, nil, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, nil, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	return res, res, nil

}

// GetMessageContentPreview
//
// Get a preview image of the image or video
// Parameters:
//        messageId             Message ID of image or video

// You must close the response body when finished with it.
// https://developers.line.biz/en/reference/messaging-api/#get-image-or-video-preview
func (client *MessagingApiBlobAPI) GetMessageContentPreview(

	messageId string,

) (*http.Response, error) {
	_, body, error := client.GetMessageContentPreviewWithHttpInfo(

		messageId,
	)
	return body, error
}

// GetMessageContentPreview
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Get a preview image of the image or video
// Parameters:
//        messageId             Message ID of image or video

// You must close the response body when finished with it.
// https://developers.line.biz/en/reference/messaging-api/#get-image-or-video-preview
func (client *MessagingApiBlobAPI) GetMessageContentPreviewWithHttpInfo(

	messageId string,

) (*http.Response, *http.Response, error) {
	path := "/v2/bot/message/{messageId}/content/preview"

	path = strings.Replace(path, "{messageId}", messageId, -1)

	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return res, nil, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, nil, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	return res, res, nil

}

// GetMessageContentTranscodingByMessageId
//
// Verify the preparation status of a video or audio for getting
// Parameters:
//        messageId             Message ID of video or audio

// https://developers.line.biz/en/reference/messaging-api/#verify-video-or-audio-preparation-status
func (client *MessagingApiBlobAPI) GetMessageContentTranscodingByMessageId(

	messageId string,

) (*GetMessageContentTranscodingResponse, error) {
	_, body, error := client.GetMessageContentTranscodingByMessageIdWithHttpInfo(

		messageId,
	)
	return body, error
}

// GetMessageContentTranscodingByMessageId
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Verify the preparation status of a video or audio for getting
// Parameters:
//        messageId             Message ID of video or audio

// https://developers.line.biz/en/reference/messaging-api/#verify-video-or-audio-preparation-status
func (client *MessagingApiBlobAPI) GetMessageContentTranscodingByMessageIdWithHttpInfo(

	messageId string,

) (*http.Response, *GetMessageContentTranscodingResponse, error) {
	path := "/v2/bot/message/{messageId}/content/transcoding"

	path = strings.Replace(path, "{messageId}", messageId, -1)

	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return res, nil, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, nil, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	result := GetMessageContentTranscodingResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return res, &result, nil

}

// GetRichMenuImage
//
// Download rich menu image.
// Parameters:
//        richMenuId             ID of the rich menu with the image to be downloaded

// You must close the response body when finished with it.
// https://developers.line.biz/en/reference/messaging-api/#download-rich-menu-image
func (client *MessagingApiBlobAPI) GetRichMenuImage(

	richMenuId string,

) (*http.Response, error) {
	_, body, error := client.GetRichMenuImageWithHttpInfo(

		richMenuId,
	)
	return body, error
}

// GetRichMenuImage
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Download rich menu image.
// Parameters:
//        richMenuId             ID of the rich menu with the image to be downloaded

// You must close the response body when finished with it.
// https://developers.line.biz/en/reference/messaging-api/#download-rich-menu-image
func (client *MessagingApiBlobAPI) GetRichMenuImageWithHttpInfo(

	richMenuId string,

) (*http.Response, *http.Response, error) {
	path := "/v2/bot/richmenu/{richMenuId}/content"

	path = strings.Replace(path, "{richMenuId}", richMenuId, -1)

	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return res, nil, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, nil, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	return res, res, nil

}

// SetRichMenuImage
//
// Upload rich menu image
// Parameters:
//        richMenuId             The ID of the rich menu to attach the image to
//        bodyContentType   content-type
//        bodyReader        file content

// https://developers.line.biz/en/reference/messaging-api/#upload-rich-menu-image
func (client *MessagingApiBlobAPI) SetRichMenuImage(

	richMenuId string,

	bodyContentType string,
	bodyReader io.Reader,

) (struct{}, error) {
	_, body, error := client.SetRichMenuImageWithHttpInfo(

		richMenuId,

		bodyContentType,
		bodyReader,
	)
	return body, error
}

// SetRichMenuImage
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Upload rich menu image
// Parameters:
//        richMenuId             The ID of the rich menu to attach the image to
//        bodyContentType   content-type
//        bodyReader        file content

// https://developers.line.biz/en/reference/messaging-api/#upload-rich-menu-image
func (client *MessagingApiBlobAPI) SetRichMenuImageWithHttpInfo(

	richMenuId string,

	bodyContentType string,
	bodyReader io.Reader,

) (*http.Response, struct{}, error) {
	path := "/v2/bot/richmenu/{richMenuId}/content"

	path = strings.Replace(path, "{richMenuId}", richMenuId, -1)

	req, err := http.NewRequest(http.MethodPost, client.Url(path), bodyReader)
	if err != nil {
		return nil, struct{}{}, err
	}
	req.Header.Set("Content-Type", bodyContentType)

	res, err := client.Do(req)

	if err != nil {
		return res, struct{}{}, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, struct{}{}, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, struct{}{}, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	return res, struct{}{}, nil

}
