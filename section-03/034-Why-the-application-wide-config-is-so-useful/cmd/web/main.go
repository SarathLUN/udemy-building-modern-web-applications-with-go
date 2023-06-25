package main

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/034-Why-the-application-wide-config-is-so-useful/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/034-Why-the-application-wide-config-is-so-useful/pkg/handlers"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/034-Why-the-application-wide-config-is-so-useful/pkg/renderers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := renderers.CreateTemplateCache()
	if err != nil {
		log.Fatalln("cannot create template cache", err)
	}
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("starting web server on port", portNumber)
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalln("can't start server;", err.Error())
	}
}
