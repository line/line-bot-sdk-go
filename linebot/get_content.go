package linebot

import (
	"io"
	"mime"
	"net/http"
)

// MessageContentResponse type
type MessageContentResponse struct {
	Content  io.ReadCloser
	FileName string
}

func newMessageContentResponse(res *http.Response) (mc *MessageContentResponse) {
	mc = &MessageContentResponse{
		Content: res.Body,
	}
	_, params, err := mime.ParseMediaType(res.Header.Get("Content-Disposition"))
	if err != nil {
		return
	}
	mc.FileName = params["filename"]
	return
}

// GetMessageContent function
func (client *Client) GetMessageContent(content *ReceivedContent) (mc *MessageContentResponse, err error) {
	res, err := client.get(APIEndpointMessage+"/"+content.ID+"/content", "")
	if err != nil {
		return
	}
	return newMessageContentResponse(res), nil
}

// GetMessageContentPreview function
func (client *Client) GetMessageContentPreview(content *ReceivedContent) (mc *MessageContentResponse, err error) {
	res, err := client.get(APIEndpointMessage+"/"+content.ID+"/content/preview", "")
	if err != nil {
		return
	}
	return newMessageContentResponse(res), nil
}
