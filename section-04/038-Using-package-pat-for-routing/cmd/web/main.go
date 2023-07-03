package main

import (
	"log"
	"net/http"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/037-Sharing-data-with-templates/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/037-Sharing-data-with-templates/pkg/handlers"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/037-Sharing-data-with-templates/pkg/renderers"
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
