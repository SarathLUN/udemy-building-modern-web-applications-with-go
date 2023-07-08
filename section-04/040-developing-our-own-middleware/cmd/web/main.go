package main

import (
	"log"
	"net/http"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-04/040-developing-our-own-middleware/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-04/040-developing-our-own-middleware/pkg/renderers"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := renderers.CreateTemplateCache()
	if err != nil {
		log.Fatalln("cannot create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	renderers.NewTemplate(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: setup_routes(&app),
	}

	log.Println("starting web server on port", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln("can't start server;", err.Error())
	}
}
