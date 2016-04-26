package linebot

import (
	"strconv"
)

// MultipleMessageRequest type
type MultipleMessageRequest struct {
	client   *Client
	messages []SingleMessageContent
}

// NewMultipleMessage returns multiple request instance
func (client *Client) NewMultipleMessage() *MultipleMessageRequest {
	return &MultipleMessageRequest{
		client:   client,
		messages: []SingleMessageContent{},
	}
}

// AddText function
func (mmr *MultipleMessageRequest) AddText(text string) *MultipleMessageRequest {
	mmr.messages = append(mmr.messages, SingleMessageContent{
		ContentType: ContentTypeText,
		Text:        text,
	})
	return mmr
}

// AddImage function
func (mmr *MultipleMessageRequest) AddImage(imageURL, previewURL string) *MultipleMessageRequest {
	mmr.messages = append(mmr.messages, SingleMessageContent{
		ContentType:        ContentTypeImage,
		OriginalContentURL: imageURL,
		PreviewImageURL:    previewURL,
	})
	return mmr
}

// AddVideo function
func (mmr *MultipleMessageRequest) AddVideo(videoURL, previewURL string) *MultipleMessageRequest {
	mmr.messages = append(mmr.messages, SingleMessageContent{
		ContentType:        ContentTypeVideo,
		OriginalContentURL: videoURL,
		PreviewImageURL:    previewURL,
	})
	return mmr
}

// AddAudio function
func (mmr *MultipleMessageRequest) AddAudio(audioURL string, duration int) *MultipleMessageRequest {
	mmr.messages = append(mmr.messages, SingleMessageContent{
		ContentType:        ContentTypeAudio,
		OriginalContentURL: audioURL,
		ContentMetaData:    map[string]string{"AUDLEN": strconv.Itoa(duration)},
	})
	return mmr
}

// AddLocation function
func (mmr *MultipleMessageRequest) AddLocation(title, address string, latitude, longitude float64) *MultipleMessageRequest {
	mmr.messages = append(mmr.messages, SingleMessageContent{
		ContentType: ContentTypeLocation,
		Text:        title,
		Location: &MessageContentLocation{
			Title:     title,
			Address:   address,
			Latitude:  latitude,
			Longitude: longitude,
		},
	})
	return mmr
}

// AddSticker function
func (mmr *MultipleMessageRequest) AddSticker(stkID, stkPkgID, stkVer int) *MultipleMessageRequest {
	mmr.messages = append(mmr.messages, SingleMessageContent{
		ContentType: ContentTypeSticker,
		ContentMetaData: map[string]string{
			"STKID":    strconv.Itoa(stkID),
			"STKPKGID": strconv.Itoa(stkPkgID),
			"STKVER":   strconv.Itoa(stkVer),
		},
	})
	return mmr
}

// Send function
func (mmr *MultipleMessageRequest) Send(to []string) (result *ResponseContent, err error) {
	return mmr.client.sendMultipleMessage(to, MultipleMessageContent{
		MessageNotified: 0,
		Messages:        mmr.messages,
	})
}
