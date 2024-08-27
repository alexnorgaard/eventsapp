package event

import (
	"github.com/labstack/echo/v4"
)

type Store interface {
	GetByID(echo.Context) error
	GetEvent(echo.Context) error //TODO: Figure out how to use query parameters for this
	CreateEvent(echo.Context) error
	// Update(*model.Event) error
	// Delete(uuid.UUID) error
}
