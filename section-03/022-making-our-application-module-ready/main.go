package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			log.Println(err)
		}
		log.Println("written byte:", n)
	})
	_ = http.ListenAndServe("localhost:8080", nil)
}
