package main

import (
	"encoding/gob"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/071-finishing-up-our-response-to-user-and-add-alerts/internal/handlers"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/071-finishing-up-our-response-to-user-and-add-alerts/internal/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/071-finishing-up-our-response-to-user-and-add-alerts/internal/models"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/071-finishing-up-our-response-to-user-and-add-alerts/internal/renderers"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// register none-primitive data type into session
	gob.Register(models.Reservation{})

	// change this to true when in production mode
	app.InProduction = false

	// working on session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := renderers.CreateTemplateCache()
	if err != nil {
		log.Fatalln("cannot create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	// create new repo
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// render new template
	renderers.NewTemplate(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: setupRoutes(&app),
	}

	log.Println("starting web server on port", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln("can't start server;", err.Error())
	}
}
