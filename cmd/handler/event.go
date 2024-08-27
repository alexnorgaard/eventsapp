package handler

import (
	"fmt"
	"net/http"

	"github.com/alexnorgaard/eventsapp/cmd/model"
	geolocationclient "github.com/alexnorgaard/eventsapp/internal/geolocation_client"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EventStore struct {
	db *gorm.DB
}

func NewEventStore(db *gorm.DB) *EventStore {
	return &EventStore{db: db}
}

func (es *EventStore) GetByID(c echo.Context) error {
	fmt.Println("UUID is: ", c.Param("id"))
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Bad Request - Invalid UUID")
	}
	var event = model.Event{Model: model.Model{ID: uuid}}
	result := es.db.First(&event)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, event)
}

func (es *EventStore) GetEvent(c echo.Context) (string, error) {
	return "hai", nil
}
func (es *EventStore) CreateEvent(c echo.Context) error {
	fmt.Printf("Creating event: %v\n", c)
	event := model.Event{}
	err := c.Bind(&event)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	err = c.Validate(&event)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
    
	if event.Address != nil {
		location, err := geolocationclient.GetGeolocation(event.Address.FormattedAddress)
		if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		fmt.Printf("Location is: %v\n", location)
		event.Geolocation = location
	}
	
	result := es.db.Create(&event)
	if es.db.Create(&event).Error != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, event)
}
