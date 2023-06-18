package main

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/028-Enabling-Go-Modules-and-refactoring-our-code-to-use-packages/pkg/handlers"
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
