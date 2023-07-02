package main

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/035-Optimizing-our-template-cache-by-using-an-application-config/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/035-Optimizing-our-template-cache-by-using-an-application-config/pkg/handlers"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/035-Optimizing-our-template-cache-by-using-an-application-config/pkg/renderers"
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
	app.UseCache = false

	renderers.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Println("starting web server on port", portNumber)
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalln("can't start server;", err.Error())
	}
}
