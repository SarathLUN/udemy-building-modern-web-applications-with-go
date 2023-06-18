package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("starting web server on port", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalln("can't start server;", err.Error())
	}
}
