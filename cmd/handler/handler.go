package handler

import (
	"github.com/alexnorgaard/eventsapp/cmd/event"
)

type Handler struct {
	EventStore event.Store
}

func NewHandler(eventStore event.Store) *Handler {
	return &Handler{
		EventStore: eventStore,
	}
}
