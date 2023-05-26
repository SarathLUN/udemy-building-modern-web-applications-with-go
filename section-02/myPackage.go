package main

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-02/helpers"
	"log"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	log.Println(myVar)
}
