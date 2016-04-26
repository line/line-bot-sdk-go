package linebot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// ReceivedResults type
type ReceivedResults struct {
	Results []ReceivedResult `json:"result"`
}

// ReceivedResult type
type ReceivedResult struct {
	ID          string    `json:"id"`
	From        string    `json:"from"`
	FromChannel int64     `json:"fromChannel"`
	To          []string  `json:"to"`
	ToChannel   int64     `json:"toChannel"`
	EventType   EventType `json:"eventType"`
	RawContent  struct {
		ID              string                 `json:"id"`
		ContentType     ContentType            `json:"contentType"`
		From            string                 `json:"from"`
		CreatedTime     int64                  `json:"createdTime"`
		To              []string               `json:"to"`
		ToType          RecipientType          `json:"toType"`
		ContentMetaData map[string]string      `json:"contentMetadata"`
		Text            string                 `json:"text"`
		Location        MessageContentLocation `json:"location"`
		Revision        int                    `json:"revision"`
		OpType          OpType                 `json:"opType"`
		Params          []string               `json:"params"`
	} `json:"content"`
}

// Content function
func (rr *ReceivedResult) Content() *ReceivedContent {
	return &ReceivedContent{
		parent:      rr,
		ID:          rr.RawContent.ID,
		From:        rr.RawContent.From,
		CreatedTime: rr.RawContent.CreatedTime,
		To:          rr.RawContent.To,
		ToType:      rr.RawContent.ToType,
		IsOperation: rr.EventType == EventTypeReceivingOperation,
		IsMessage:   rr.EventType == EventTypeReceivingMessage,
		OpType:      rr.RawContent.OpType,
		ContentType: rr.RawContent.ContentType,
	}
}

// ReceivedTextContent type
type ReceivedTextContent struct {
	*ReceivedContent
	Text string
}

// ReceivedImageContent type
type ReceivedImageContent struct {
	*ReceivedContent
}

// ReceivedVideoContent type
type ReceivedVideoContent struct {
	*ReceivedContent
}

// ReceivedAudioContent type
type ReceivedAudioContent struct {
	*ReceivedContent
	Duration int
}

// ReceivedLocationContent type
type ReceivedLocationContent struct {
	*ReceivedContent
	Text      string
	Title     string
	Address   string
	Latitude  float64
	Longitude float64
}

// ReceivedStickerContent type
type ReceivedStickerContent struct {
	*ReceivedContent
	PackageID int
	ID        int
	Version   int
}

// ReceivedContactContent type
type ReceivedContactContent struct {
	*ReceivedContent
	Mid         string
	DisplayName string
}

// ReceivedOperation type
type ReceivedOperation struct {
	*ReceivedContent
	Revision int
	Params   []string
}

// ReceivedContent type
type ReceivedContent struct {
	parent      *ReceivedResult
	ID          string
	From        string
	CreatedTime int64
	To          []string
	ToType      RecipientType
	IsOperation bool
	IsMessage   bool
	OpType      OpType
	ContentType ContentType
}

// TextContent function
func (rc *ReceivedContent) TextContent() (*ReceivedTextContent, error) {
	if !rc.IsMessage {
		return nil, ErrInvalidEventType
	}
	if rc.ContentType != ContentTypeText {
		return nil, ErrInvalidContentType
	}
	return &ReceivedTextContent{
		ReceivedContent: rc,
		Text:            rc.parent.RawContent.Text,
	}, nil
}

// ImageContent function
func (rc *ReceivedContent) ImageContent() (*ReceivedImageContent, error) {
	if !rc.IsMessage {
		return nil, ErrInvalidEventType
	}
	if rc.ContentType != ContentTypeImage {
		return nil, ErrInvalidContentType
	}
	return &ReceivedImageContent{
		ReceivedContent: rc,
	}, nil
}

// VideoContent function
func (rc *ReceivedContent) VideoContent() (*ReceivedVideoContent, error) {
	if !rc.IsMessage {
		return nil, ErrInvalidEventType
	}
	if rc.ContentType != ContentTypeVideo {
		return nil, ErrInvalidContentType
	}
	return &ReceivedVideoContent{
		ReceivedContent: rc,
	}, nil
}

// AudioContent function
func (rc *ReceivedContent) AudioContent() (*ReceivedAudioContent, error) {
	if !rc.IsMessage {
		return nil, ErrInvalidEventType
	}
	if rc.ContentType != ContentTypeAudio {
		return nil, ErrInvalidContentType
	}
	audlen, err := strconv.Atoi(rc.parent.RawContent.ContentMetaData["AUDLEN"])
	if err != nil {
		return nil, err
	}
	return &ReceivedAudioContent{
		ReceivedContent: rc,
		Duration:        audlen,
	}, nil
}

// LocationContent function
func (rc *ReceivedContent) LocationContent() (*ReceivedLocationContent, error) {
	if !rc.IsMessage {
		return nil, ErrInvalidEventType
	}
	if rc.ContentType != ContentTypeLocation {
		return nil, ErrInvalidContentType
	}
	return &ReceivedLocationContent{
		ReceivedContent: rc,
		Text:            rc.parent.RawContent.Text,
		Title:           rc.parent.RawContent.Location.Title,
		Address:         rc.parent.RawContent.Location.Address,
		Latitude:        rc.parent.RawContent.Location.Latitude,
		Longitude:       rc.parent.RawContent.Location.Longitude,
	}, nil
}

// StickerContent function
func (rc *ReceivedContent) StickerContent() (*ReceivedStickerContent, error) {
	if !rc.IsMessage {
		return nil, ErrInvalidEventType
	}
	if rc.ContentType != ContentTypeSticker {
		return nil, ErrInvalidContentType
	}
	stkPkgID, err := strconv.Atoi(rc.parent.RawContent.ContentMetaData["STKPKGID"])
	if err != nil {
		return nil, err
	}
	stkID, err := strconv.Atoi(rc.parent.RawContent.ContentMetaData["STKID"])
	if err != nil {
		return nil, err
	}
	stkVer, err := strconv.Atoi(rc.parent.RawContent.ContentMetaData["STKVER"])
	if err != nil {
		return nil, err
	}
	return &ReceivedStickerContent{
		ReceivedContent: rc,
		PackageID:       stkPkgID,
		ID:              stkID,
		Version:         stkVer,
	}, nil
}

// ContactContent function
func (rc *ReceivedContent) ContactContent() (*ReceivedContactContent, error) {
	if !rc.IsMessage {
		return nil, ErrInvalidEventType
	}
	if rc.ContentType != ContentTypeContact {
		return nil, ErrInvalidContentType
	}
	return &ReceivedContactContent{
		ReceivedContent: rc,
		Mid:             rc.parent.RawContent.ContentMetaData["mid"],
		DisplayName:     rc.parent.RawContent.ContentMetaData["displayName"],
	}, nil
}

// OperationContent function
func (rc *ReceivedContent) OperationContent() (*ReceivedOperation, error) {
	if !rc.IsOperation {
		return nil, ErrInvalidEventType
	}
	return &ReceivedOperation{
		ReceivedContent: rc,
		Revision:        rc.parent.RawContent.Revision,
		Params:          rc.parent.RawContent.Params,
	}, nil
}

// ReceivedOperationContent type
type ReceivedOperationContent struct {
	Revision int
	OpType   OpType
	Params   []string
}

// ParseRequest function
func (client *Client) ParseRequest(r *http.Request) (results *ReceivedResults, err error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	if !client.validateSignature(r.Header.Get("X-LINE-ChannelSignature"), body) {
		return nil, ErrInvalidSignature
	}

	results = &ReceivedResults{}
	if err = json.Unmarshal(body, results); err != nil {
		return nil, err
	}
	return
}

func (client *Client) validateSignature(signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(client.channelSecret))
	hash.Write(body)
	return hmac.Equal(decoded, hash.Sum(nil))
}
