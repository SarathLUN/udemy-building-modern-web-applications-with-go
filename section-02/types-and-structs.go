package main

import (
	"log"
	"time"
)

var s = "Hello World!"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Tony",
		LastName:  "Stark",
	}
	log.Println(user.FirstName, user.LastName, "birth date:", user.BirthDate)
}
