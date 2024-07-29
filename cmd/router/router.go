package router

import (
	"fmt"
	"net/http"

	"github.com/alexnorgaard/eventsapp/cmd/handler"
	"github.com/alexnorgaard/eventsapp/cmd/model"
	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	fmt.Println("Registering routes")
	v1 := e.Group("/v1")

	events := v1.Group("/events")
	events.GET("/", func(c echo.Context) error {
		return nil
	})
	events.POST("/", func(c echo.Context) error {
		return handler.CreateEvent(c, db)
	})
	events.GET("/", func(c echo.Context) error {
		// event := NewEvent()
		// return c.JSON(http.StatusOK, event)
		return nil
	})
	events.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	events.PUT("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})

	events.DELETE("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})

	users := v1.Group("/users")
	users.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	users.POST("/", func(c echo.Context) error {
		var user model.User
		err := c.Bind(&user)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, user)
	})
}
