package helpers

import "math/rand"

type SomeType struct {
	TypeName   string
	TypeNumber int
}

func RandomNumber(poolSize int) int {
	value := rand.Intn(poolSize)
	return value
}
