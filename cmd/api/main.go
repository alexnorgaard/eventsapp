package main

import (
	"net/http"
	"time"

	"github.com/codingsince1985/geo-golang"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Event struct {
	ID            uuid.UUID    `json:"id"`
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	Owner         [3]User      `json:"owner"`
	Private_event bool         `json:"private_event"`
	Time_start    time.Time    `json:"time_start"`
	Time_end      time.Time    `json:"time_end"`
	Address       geo.Address  `json:"address"`
	Geolocation   geo.Location `json:"geolocation"`
}

func NewEvent() *Event {
	return &Event{
		ID:            uuid.New(),
		Title:         "Event Title",
		Description:   "Event Description",
		Owner:         [3]User{},
		Private_event: false,
		Time_start:    time.Now(),
		Time_end:      time.Now(),
		Address:       geo.Address{},
		Geolocation:   geo.Location{},
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/events", func(c echo.Context) error {
		event := NewEvent()
		return c.JSON(http.StatusOK, event)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
