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

package manage_audience

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

type ManageAudienceBlobAPI struct {
	httpClient   *http.Client
	endpoint     *url.URL
	channelToken string
	ctx          context.Context
}

// ManageAudienceBlobAPIOption type
type ManageAudienceBlobAPIOption func(*ManageAudienceBlobAPI) error

// New returns a new bot client instance.
func NewManageAudienceBlobAPI(channelToken string, options ...ManageAudienceBlobAPIOption) (*ManageAudienceBlobAPI, error) {
	if channelToken == "" {
		return nil, errors.New("missing channel access token")
	}

	c := &ManageAudienceBlobAPI{
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
func (call *ManageAudienceBlobAPI) WithContext(ctx context.Context) *ManageAudienceBlobAPI {
	call.ctx = ctx
	return call
}

func (client *ManageAudienceBlobAPI) Do(req *http.Request) (*http.Response, error) {
	if client.channelToken != "" {
		req.Header.Set("Authorization", "Bearer "+client.channelToken)
	}
	req.Header.Set("User-Agent", "LINE-BotSDK-Go/"+linebot.GetVersion())
	if client.ctx != nil {
		req = req.WithContext(client.ctx)
	}
	return client.httpClient.Do(req)
}

func (client *ManageAudienceBlobAPI) Url(endpointPath string) string {
	newPath := path.Join(client.endpoint.Path, endpointPath)
	u := *client.endpoint
	u.Path = newPath
	return u.String()
}

// WithHTTPClient function
func WithBlobHTTPClient(c *http.Client) ManageAudienceBlobAPIOption {
	return func(client *ManageAudienceBlobAPI) error {
		client.httpClient = c
		return nil
	}
}

// WithEndpointClient function
func WithBlobEndpoint(endpoint string) ManageAudienceBlobAPIOption {
	return func(client *ManageAudienceBlobAPI) error {
		u, err := url.ParseRequestURI(endpoint)
		if err != nil {
			return err
		}
		client.endpoint = u
		return nil
	}
}

// AddUserIdsToAudience
//
// Add user IDs or Identifiers for Advertisers (IFAs) to an audience for uploading user IDs (by file).
// Parameters:
//        file             A text file with one user ID or IFA entered per line. Specify text/plain as Content-Type. Max file number: 1 Max number: 1,500,000
//        audienceGroupId             The audience ID.
//        uploadDescription             The description to register with the job

// https://developers.line.biz/en/reference/messaging-api/#update-upload-audience-group-by-file
func (client *ManageAudienceBlobAPI) AddUserIdsToAudience(

	file *os.File,

	audienceGroupId int64,

	uploadDescription string,

) (struct{}, error) {
	_, body, error := client.AddUserIdsToAudienceWithHttpInfo(

		file,

		audienceGroupId,

		uploadDescription,
	)
	return body, error
}

// AddUserIdsToAudience
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Add user IDs or Identifiers for Advertisers (IFAs) to an audience for uploading user IDs (by file).
// Parameters:
//        file             A text file with one user ID or IFA entered per line. Specify text/plain as Content-Type. Max file number: 1 Max number: 1,500,000
//        audienceGroupId             The audience ID.
//        uploadDescription             The description to register with the job

// https://developers.line.biz/en/reference/messaging-api/#update-upload-audience-group-by-file
func (client *ManageAudienceBlobAPI) AddUserIdsToAudienceWithHttpInfo(

	file *os.File,

	audienceGroupId int64,

	uploadDescription string,

) (*http.Response, struct{}, error) {
	path := "/v2/bot/audienceGroup/upload/byFile"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("audienceGroupId", strconv.FormatInt(audienceGroupId, 10))

	writer.WriteField("uploadDescription", string(uploadDescription))

	fileWriter, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return nil, struct{}{}, err
	}
	io.Copy(fileWriter, file)

	err = writer.Close()
	if err != nil {
		return nil, struct{}{}, err
	}

	req, err := http.NewRequest(http.MethodPut, client.Url(path), body)
	if err != nil {
		return nil, struct{}{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

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

// CreateAudienceForUploadingUserIds
//
// Create audience for uploading user IDs (by file).
// Parameters:
//        file             A text file with one user ID or IFA entered per line. Specify text/plain as Content-Type. Max file number: 1 Max number: 1,500,000
//        description             The audience's name. This is case-insensitive, meaning AUDIENCE and audience are considered identical. Max character limit: 120
//        isIfaAudience             To specify recipients by IFAs: set `true`. To specify recipients by user IDs: set `false` or omit isIfaAudience property.
//        uploadDescription             The description to register for the job (in `jobs[].description`).

// https://developers.line.biz/en/reference/messaging-api/#create-upload-audience-group-by-file
func (client *ManageAudienceBlobAPI) CreateAudienceForUploadingUserIds(

	file *os.File,

	description string,

	isIfaAudience bool,

	uploadDescription string,

) (*CreateAudienceGroupResponse, error) {
	_, body, error := client.CreateAudienceForUploadingUserIdsWithHttpInfo(

		file,

		description,

		isIfaAudience,

		uploadDescription,
	)
	return body, error
}

// CreateAudienceForUploadingUserIds
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Create audience for uploading user IDs (by file).
// Parameters:
//        file             A text file with one user ID or IFA entered per line. Specify text/plain as Content-Type. Max file number: 1 Max number: 1,500,000
//        description             The audience's name. This is case-insensitive, meaning AUDIENCE and audience are considered identical. Max character limit: 120
//        isIfaAudience             To specify recipients by IFAs: set `true`. To specify recipients by user IDs: set `false` or omit isIfaAudience property.
//        uploadDescription             The description to register for the job (in `jobs[].description`).

// https://developers.line.biz/en/reference/messaging-api/#create-upload-audience-group-by-file
func (client *ManageAudienceBlobAPI) CreateAudienceForUploadingUserIdsWithHttpInfo(

	file *os.File,

	description string,

	isIfaAudience bool,

	uploadDescription string,

) (*http.Response, *CreateAudienceGroupResponse, error) {
	path := "/v2/bot/audienceGroup/upload/byFile"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("description", string(description))

	writer.WriteField("isIfaAudience", strconv.FormatBool(isIfaAudience))

	writer.WriteField("uploadDescription", string(uploadDescription))

	fileWriter, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return nil, nil, err
	}
	io.Copy(fileWriter, file)

	err = writer.Close()
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(http.MethodPost, client.Url(path), body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

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
	result := CreateAudienceGroupResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return res, &result, nil

}
