package linebot

type Mention struct {
	Mentionees []*Mentionee `json:"mentionees"`
}

type Mentionee struct {
	Index  int    `json:"index"`
	Length int    `json:"length"`
	UserId string `json:"userId,omitempty"`
}
