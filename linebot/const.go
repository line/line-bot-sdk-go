package linebot

// constants
const (
	APIEndpointBaseTrial = "https://trialbot-api.line.me"
	APIEndpointEvents    = "/v1/events"
	APIEndpointMessage   = "/v1/bot/message"
	APIEndpointProfiles  = "/v1/profiles"

	SendingMessageChannelID = 1383378250

	RichMessageSpecRev = "1"
)

// EventType type
type EventType string

// EventType constants
const (
	EventTypeSendingMessage         EventType = "138311608800106203"
	EventTypeSendingMultipleMessage EventType = "140177271400161403"
	EventTypeReceivingMessage       EventType = "138311609000106303"
	EventTypeReceivingOperation     EventType = "138311609100106403"
)

// ContentType type
type ContentType int

// ContentType constants
const (
	ContentTypeText        ContentType = 1
	ContentTypeImage       ContentType = 2
	ContentTypeVideo       ContentType = 3
	ContentTypeAudio       ContentType = 4
	ContentTypeLocation    ContentType = 7
	ContentTypeSticker     ContentType = 8
	ContentTypeContact     ContentType = 10
	ContentTypeRichMessage ContentType = 12
)

// RecipientType type
type RecipientType int

// RecipientType constants
const (
	RecipientTypeUser RecipientType = 1
)

// OpType type
type OpType int

// OpType constants
const (
	OpTypeAddedAsFriend OpType = 4
	OpTypeBlocked       OpType = 8
)
