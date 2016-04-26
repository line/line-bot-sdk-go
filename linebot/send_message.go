package linebot

import (
	"strconv"
)

// SendText function
func (client *Client) SendText(to []string, text string) (result *ResponseContent, err error) {
	return client.sendSingleMessage(to, SingleMessageContent{
		ContentType: ContentTypeText,
		ToType:      RecipientTypeUser,
		Text:        text,
	})
}

// SendImage function
func (client *Client) SendImage(to []string, imageURL, previewURL string) (result *ResponseContent, err error) {
	return client.sendSingleMessage(to, SingleMessageContent{
		ContentType:        ContentTypeImage,
		ToType:             RecipientTypeUser,
		OriginalContentURL: imageURL,
		PreviewImageURL:    previewURL,
	})
}

// SendVideo function
func (client *Client) SendVideo(to []string, videoURL, previewURL string) (result *ResponseContent, err error) {
	return client.sendSingleMessage(to, SingleMessageContent{
		ContentType:        ContentTypeVideo,
		ToType:             RecipientTypeUser,
		OriginalContentURL: videoURL,
		PreviewImageURL:    previewURL,
	})
}

// SendAudio function
func (client *Client) SendAudio(to []string, audioURL string, duration int) (result *ResponseContent, err error) {
	return client.sendSingleMessage(to, SingleMessageContent{
		ContentType:        ContentTypeAudio,
		ToType:             RecipientTypeUser,
		OriginalContentURL: audioURL,
		ContentMetaData:    map[string]string{"AUDLEN": strconv.Itoa(duration)},
	})
}

// SendLocation function
func (client *Client) SendLocation(to []string, title, address string, latitude, longitude float64) (result *ResponseContent, err error) {
	return client.sendSingleMessage(to, SingleMessageContent{
		ContentType: ContentTypeLocation,
		ToType:      RecipientTypeUser,
		Text:        title,
		Location: &MessageContentLocation{
			Title:     title,
			Address:   address,
			Latitude:  latitude,
			Longitude: longitude,
		},
	})
}

// SendSticker function
func (client *Client) SendSticker(to []string, stkID, stkPkgID, stkVer int) (result *ResponseContent, err error) {
	return client.sendSingleMessage(to, SingleMessageContent{
		ContentType: ContentTypeSticker,
		ToType:      RecipientTypeUser,
		ContentMetaData: map[string]string{
			"STKID":    strconv.Itoa(stkID),
			"STKPKGID": strconv.Itoa(stkPkgID),
			"STKVER":   strconv.Itoa(stkVer),
		},
	})
}
