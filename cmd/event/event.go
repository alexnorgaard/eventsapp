package event

import (
	"github.com/alexnorgaard/eventsapp/cmd/model"
	"github.com/google/uuid"
)

type Store interface {
	GetByID(uuid.UUID) (*model.Event, error)
	Get() (string, error) //TODO: Figure out how to use query parameters for this
	Create(*model.Event) error
	Update(*model.Event) error
	Delete(uuid.UUID) error
}
