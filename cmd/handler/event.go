package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alexnorgaard/eventsapp/cmd/model"
	geolocationclient "github.com/alexnorgaard/eventsapp/internal/geolocation_client"
	minioClient "github.com/alexnorgaard/eventsapp/internal/minio_client"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EventStore struct {
	db *gorm.DB
}

type EventSearchItemDTO struct {
	ID               uuid.UUID `json:"id"`
	Title            string    `json:"title"`
	Distance         float64   `json:"distance"`
	FormattedAddress string    `json:"address"`
	Lat              float64   `json:"lat"`
	Lng              float64   `json:"lng"`
}

type EventDTO struct {
	ID               uuid.UUID `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Banner_url       string    `json:"banner_url"`
	Time_start       string    `json:"time_start"`
	Time_end         string    `json:"time_end"`
	FormattedAddress string    `json:"address"`
	Lat              float64   `json:"lat"`
	Lng              float64   `json:"lng"`
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
	// var event = model.Event{Model: model.Model{ID: uuid}}
	// result := es.db.First(&event)
	var eventDTO = EventDTO{}
	result := es.db.Model(&model.Event{}).Where("id = ?", uuid).First(&eventDTO)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, eventDTO)
}

func (es *EventStore) Get(c echo.Context) error {
	//Smart select fields - only returns the fields in APIEvent when used in Find
	var events = []EventSearchItemDTO{}
	var result *gorm.DB

	params := c.QueryParams()
	if len(params) != 0 {
		tags := params.Get("tags") //Should probably not be called tags anymore, as its now also a string search
		title := strings.ReplaceAll(tags, ",", " ")
		lat := params.Get("lat")
		long := params.Get("long")
		coordinates := []string{long, lat}
		//ILIKE makes the LIKE case insensitive
		result = es.db.Model(&model.Event{}).Where("tags @> ? OR title ILIKE ?", "{"+tags+"}", "%"+title+"%").Select("id, title, formatted_address, lng, lat, ST_DistanceSphere(ST_MakePoint(lng,lat),ST_MakePoint(?))/1000 as distance", coordinates).Order("distance asc").Find(&events)
	} else {
		result = es.db.Model(&model.Event{}).Select("id, title").Find(&events)
	}
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}
	// result := es.db.Model(&model.Event{}).Find(&events, "tags @> ?", "{"+tags+"}")
	// result := es.db.Raw({"SELECT events.id,events.title FROM "events" WHERE tags @> ?, '{'+party+'}'})

	return c.JSON(http.StatusOK, events)
}

func (es *EventStore) Create(c echo.Context) error {
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
			return c.String(http.StatusInternalServerError, "Location not found")
		}
		fmt.Printf("Location is: %v\n", location)
		event.Geolocation = location
	}

	if es.db.Create(&event).Error != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, event)
}

func (es *EventStore) Update(c echo.Context) error {
	ct := c.Request().Header.Get("Content-Type")
	fmt.Println("Content-Type is: ", ct)
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
	// err = c.Validate(&event)
	// if err != nil {
	// 	return c.String(http.StatusBadRequest, err.Error())
	// }
	// file_header, err := c.FormFile("image")
	result := es.db.Model(&event).Updates(&event)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, event)
}

func (es *EventStore) UploadImage(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Bad Request - Invalid UUID")
	}
	err = ValidateImage(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	file_header, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Bad Request - No image")
	}
	client, err := minioClient.GetClient()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	location, err := minioClient.UploadFile(client, file_header)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Could not upload image")
	}

	event := model.Event{Model: model.Model{ID: uuid}}
	c.Set("Banner_s3_url", location)
	err = c.Bind(&event)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	result := es.db.Model(&event).Updates(&event)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, event)
}

// func StoreImage(c echo.Context) string, error {
// 	image, err := c.FormFile("image")
// 	return c.String(http.StatusOK, "Image stored")
// }
