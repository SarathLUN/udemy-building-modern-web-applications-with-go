package main

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/030-Building-a-simple-template-cache/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("starting web server on port", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalln("can't start server;", err.Error())
	}
}
