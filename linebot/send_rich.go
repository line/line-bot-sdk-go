package linebot

import (
	"encoding/json"
)

// RichMessageRequest type
type RichMessageRequest struct {
	client    *Client
	height    int
	actions   map[string]richMessageAction
	listeners []richMessageListener
}

type richMessageMarkup struct {
	Canvas  richMessageCanvas            `json:"canvas"`
	Images  map[string]richMessageImage  `json:"images"`
	Actions map[string]richMessageAction `json:"actions"`
	Scenes  map[string]richMessageScene  `json:"scenes"`
}
type richMessageCanvas struct {
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	InitialScene string `json:"initialScene"`
}
type richMessageImage struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}
type richMessageAction struct {
	Type   string            `json:"type"`
	Text   string            `json:"text"`
	Params map[string]string `json:"params"`
}
type richMessageListener struct {
	Type   string `json:"type"`
	Params [4]int `json:"params"`
	Action string `json:"action"`
}
type richMessageScene struct {
	Draws     []richMessageSceneImage `json:"draws"`
	Listeners []richMessageListener   `json:"listeners"`
}
type richMessageSceneImage struct {
	Image string `json:"image"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
	W     int    `json:"w"`
	H     int    `json:"h"`
}

// NewRichMessage returns rich message request
func (client *Client) NewRichMessage(height int) *RichMessageRequest {
	return &RichMessageRequest{
		client:    client,
		height:    height,
		actions:   map[string]richMessageAction{},
		listeners: []richMessageListener{},
	}
}

// SetAction function
func (rmr *RichMessageRequest) SetAction(name, text, linkURI string) *RichMessageRequest {
	rmr.actions[name] = richMessageAction{
		Type:   "web",
		Text:   text,
		Params: map[string]string{"linkUri": linkURI},
	}
	return rmr
}

// SetListener function
func (rmr *RichMessageRequest) SetListener(actionName string, x, y, width, height int) *RichMessageRequest {
	rmr.listeners = append(rmr.listeners, richMessageListener{
		Type:   "touch",
		Params: [4]int{x, y, width, height},
		Action: actionName,
	})
	return rmr
}

// Send function
func (rmr *RichMessageRequest) Send(to []string, imageURL, altText string) (result *ResponseContent, err error) {
	markup, err := json.Marshal(richMessageMarkup{
		Canvas: richMessageCanvas{
			Width:        1040,
			Height:       rmr.height,
			InitialScene: "scene1",
		},
		Images: map[string]richMessageImage{
			"image1": richMessageImage{X: 0, Y: 0, W: 1040, H: rmr.height},
		},
		Actions: rmr.actions,
		Scenes: map[string]richMessageScene{
			"scene1": richMessageScene{
				Draws: []richMessageSceneImage{
					richMessageSceneImage{Image: "image1", X: 0, Y: 0, W: 1040, H: rmr.height},
				},
				Listeners: rmr.listeners,
			},
		},
	})
	if err != nil {
		return
	}
	return rmr.client.sendSingleMessage(to, SingleMessageContent{
		ContentType: ContentTypeRichMessage,
		ToType:      RecipientTypeUser,
		ContentMetaData: map[string]string{
			"DOWNLOAD_URL": imageURL,
			"SPEC_REV":     RichMessageSpecRev,
			"ALT_TEXT":     altText,
			"MARKUP_JSON":  string(markup),
		},
	})
}
