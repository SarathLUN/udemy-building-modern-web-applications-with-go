package main

import (
	"log"
	"reflect"
)

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func (g *Gorilla) Says() string {
	return "Grooo"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}

func main() {
	d := Dog{
		"Samson",
		"German Shepered",
	}
	PrintInfo(&d)

	g := Gorilla{
		"Grill",
		"Gray",
		32,
	}
	PrintInfo(&g)
}

func PrintInfo(a Animal) {
	t := reflect.TypeOf(a)
	log.Println(t, "says", a.Says(), "and has", a.NumberOfLegs(), "legs.")
}
