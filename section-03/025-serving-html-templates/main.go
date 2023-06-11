package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

// Home : handler function for "/"
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About : handler function for "/about"
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("starting web server on port", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalln("can't start server;", err.Error())
	}
}
