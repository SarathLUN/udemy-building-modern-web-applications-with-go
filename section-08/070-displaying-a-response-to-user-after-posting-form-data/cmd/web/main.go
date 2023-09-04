package main

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/069-server-side-form-validation-4/internal/handlers"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/069-server-side-form-validation-4/internal/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/069-server-side-form-validation-4/internal/renderers"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

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
