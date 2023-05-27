package main

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-02/helpers"
	"log"
)

const numPool = 100

func CalculateValue(intChan chan int) {
	randNumber := helpers.RandomNumber(numPool)
	intChan <- randNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)
	num := <-intChan
	log.Println(num)
}
