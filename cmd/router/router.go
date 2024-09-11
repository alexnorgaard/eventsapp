package router

import (
	"fmt"
	"net/http"

	"github.com/alexnorgaard/eventsapp/cmd/handler"
	"github.com/alexnorgaard/eventsapp/cmd/model"
	echo "github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, h *handler.Handler) {
	fmt.Println("Registering routes")
	v1 := e.Group("/v1")

	event := v1.Group("/event")
	event.POST("/", func(c echo.Context) error {
		return h.EventStore.Create(c)
	})
	event.GET("/", func(c echo.Context) error {
		return h.EventStore.Get(c)
	})
	event.GET("/:id", func(c echo.Context) error {
		fmt.Println("Getting event by ID")
		return h.EventStore.GetByID(c)
	})
	event.PUT("/:id", func(c echo.Context) error {
		return h.EventStore.Update(c)
	})

	event.DELETE("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})

	user := v1.Group("/user")
	user.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	user.POST("/", func(c echo.Context) error {
		var user model.User
		err := c.Bind(&user)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, user)
	})
}
