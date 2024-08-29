package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alexnorgaard/eventsapp/cmd/model"
	geolocationclient "github.com/alexnorgaard/eventsapp/internal/geolocation_client"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EventStore struct {
	db *gorm.DB
}

type APIEvent struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Distance float64   `json:"distance"`
	//TODO: distance  float64   `json:"distance"`
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

func (es *EventStore) GetEvent(c echo.Context) error {
	params := c.QueryParams()
	tags := params.Get("tags") //Should probably not be called tags anymore, as its now also a string search
	title := strings.ReplaceAll(tags, ",", " ")
	lat := params.Get("lat")
	long := params.Get("long")
	coordinates := []string{long, lat}
	//Smart select fields - only returns the fields in APIEvent when used in Find
	var events = []APIEvent{}
	//ILIKE makes the LIKE case insensitive
	result := es.db.Model(&model.Event{}).Where("tags @> ? OR title ILIKE ?", "{"+tags+"}", "%"+title+"%").Select("id, title, lng, lat, ST_DistanceSphere(ST_MakePoint(lng,lat),ST_MakePoint(?))/1000 as distance", coordinates).Order("distance asc").Find(&events)
	// result := es.db.Model(&model.Event{}).Find(&events, "tags @> ?", "{"+tags+"}")
	// result := es.db.Raw({"SELECT events.id,events.title FROM "events" WHERE tags @> ?, '{'+party+'}'})
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, events)
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
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		fmt.Printf("Location is: %v\n", location)
		event.Geolocation = location
	}

	if es.db.Create(&event).Error != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, event)
}

func (es *EventStore) UpdateEvent(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Bad Request - Invalid UUID")
	}
	//TODO: Check for existence of event
	//TODO: Check if user is owner of event
	event := model.Event{Model: model.Model{ID: uuid}}
	fmt.Printf("Event is: %v\n", event)
	err = c.Bind(&event)
	fmt.Printf("Event is: %v\n", event)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	//should probably be ignored on update
	err = c.Validate(&event)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	result := es.db.Model(&event).Updates(&event)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, event)
}
