package main

import (
	"fmt"
	"net/http"
)

const portNumber = "localhost:8080"

// Home : handler function for "/"
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the home page")
}

// About : handler function for "/about"
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 3)
	_, _ = fmt.Fprintf(w, "This is the about page, and 2 + 3 = %d", sum)
}

// addValues : required 2 param as int, return the sum as int.
func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Printf("starting web server on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
