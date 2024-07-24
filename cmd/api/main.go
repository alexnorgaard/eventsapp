package main

import (
	"net/http"

	"github.com/alexnorgaard/eventsapp/cmd/models"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/events", func(c echo.Context) error {
		// event := NewEvent()
		// return c.JSON(http.StatusOK, event)
		return nil
	})
	e.GET("/events/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	e.PUT("/events/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	e.POST("/events", func(c echo.Context) error {
		var event models.Event
		err := c.Bind(&event)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, event)
	})
	e.DELETE("/events/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
