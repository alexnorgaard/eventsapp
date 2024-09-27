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
	// autoTLSManager := autocert.Manager{
	// 	Prompt: autocert.AcceptTOS,
	// 	// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)
	// 	Cache:      autocert.DirCache("/usr/local/nginx/conf"),
	// 	HostPolicy: autocert.HostWhitelist("app.alexnorgaard.dk"),
	// }

	// s := http.Server{
	// 	Addr:    ":https",
	// 	Handler: e, // set Echo as handler
	// 	TLSConfig: &tls.Config{
	// 		//Certificates: nil, // <-- s.ListenAndServeTLS will populate this field
	// 		GetCertificate: autoTLSManager.GetCertificate,
	// 		NextProtos:     []string{acme.ALPNProto},
	// 	},
	// 	//ReadTimeout: 30 * time.Second, // use custom timeouts
	// }
	// fmt.Println("starting listen and serve")
	// if err := s.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
	// 	e.Logger.Fatal(err)
	// }
	e.Logger.Fatal(e.Start(":8080"))
}
