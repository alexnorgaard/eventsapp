package main

import (
	"github.com/alexnorgaard/eventsapp/cmd/handler"
	router "github.com/alexnorgaard/eventsapp/cmd/router"
	dbmodule "github.com/alexnorgaard/eventsapp/db"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := dbmodule.Connect()
	if err != nil {
		panic(err)
	}
	dbmodule.Migrate(db)
	e := echo.New()
	e.Validator = handler.NewValidator()
	es := handler.NewEventStore(db)
	h := handler.NewHandler(es)
	router.RegisterRoutes(e, h)
	e.Logger.Fatal(e.Start(":1323"))
}
