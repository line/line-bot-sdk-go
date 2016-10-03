package httphandler

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// EventHandler type
type EventHandler struct {
	botClient  *linebot.Client
	handleFunc func([]*linebot.Event)
}

// New returns a new EventHander instance.
func New(botClient *linebot.Client, handler func([]*linebot.Event)) *EventHandler {
	return &EventHandler{
		botClient:  botClient,
		handleFunc: handler,
	}
}

// ServeHTTP method of EventHandler
func (h *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := h.botClient.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	h.handleFunc(events)
}
