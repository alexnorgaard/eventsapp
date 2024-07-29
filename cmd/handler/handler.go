package handler

import (
	"github.com/alexnorgaard/eventsapp/cmd/event"
)

type Handler struct {
	eventStore event.Store
}

func NewHandler(eventStore event.Store) *Handler {
	return &Handler{
		eventStore: eventStore,
	}
}
