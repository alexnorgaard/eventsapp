package handler

import (
	"fmt"
	"net/http"

	"github.com/alexnorgaard/eventsapp/cmd/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetEvent(c echo.Context, db *gorm.DB) error {
	return nil
}
func CreateEvent(c echo.Context, db *gorm.DB) error {
	fmt.Printf("Creating event: %v\n", c)
	event := model.Event{}
	err := c.Bind(&event)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	db.Create(&event)

	return c.JSON(http.StatusOK, event)
}
