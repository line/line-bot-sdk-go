package linebot

// ResponseContent type
type ResponseContent struct {
	Version   int      `json:"version"`
	MessageID string   `json:"messageId"`
	Failed    []string `json:"failed"`
	Timestamp int64    `json:"timestamp"`
}

// ErrorResponseContent type
type ErrorResponseContent struct {
	Code    string `json:"statusCode"`
	Message string `json:"statusMessage"`
}

// SingleMessage type
type SingleMessage struct {
	To        []string             `json:"to"`
	ToChannel int64                `json:"toChannel"`
	EventType EventType            `json:"eventType"`
	Content   SingleMessageContent `json:"content"`
}

// MultipleMessage type
type MultipleMessage struct {
	To        []string               `json:"to"`
	ToChannel int64                  `json:"toChannel"`
	EventType EventType              `json:"eventType"`
	Content   MultipleMessageContent `json:"content"`
}

// SingleMessageContent type
type SingleMessageContent struct {
	ContentType        ContentType             `json:"contentType"`
	ToType             RecipientType           `json:"toType,omitempty"`
	Text               string                  `json:"text,omitempty"`
	OriginalContentURL string                  `json:"originalContentUrl,omitempty"`
	PreviewImageURL    string                  `json:"previewImageUrl,omitempty"`
	ContentMetaData    map[string]string       `json:"contentMetadata,omitempty"`
	Location           *MessageContentLocation `json:"location,omitempty"`
}

// MultipleMessageContent type
type MultipleMessageContent struct {
	MessageNotified int                    `json:"messageNotified,omitempty"`
	Messages        []SingleMessageContent `json:"messages"`
}

// MessageContentLocation type
type MessageContentLocation struct {
	Title     string  `json:"title"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
