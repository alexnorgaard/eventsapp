package event

import (
	"github.com/labstack/echo/v4"
)

type Store interface {
	GetByID(echo.Context) error
	Get(echo.Context) error
	Create(echo.Context) error
	Update(echo.Context) error
	UpdateImage(echo.Context) error
}
