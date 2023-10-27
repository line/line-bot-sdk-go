package linebot

// MentionedTargetType type
type MentionedTargetType string

// MentionedTargetType constants
const (
	MentionedTargetTypeUser MentionedTargetType = "user"
	MentionedTargetTypeAll  MentionedTargetType = "all"
)

// Mention type
// Deprecated: Use OpenAPI based classes instead.
type Mention struct {
	Mentionees []*Mentionee `json:"mentionees"`
}

// Mentionee type
// Deprecated: Use OpenAPI based classes instead.
type Mentionee struct {
	Index  int                 `json:"index"`
	Length int                 `json:"length"`
	Type   MentionedTargetType `json:"type"`
	UserID string              `json:"userId,omitempty"`
}
