package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello world")
	key := rand.Intn(100)
	validate_random(key)
}

func validate_random(Value int) bool {
	matches := false
	val := rand.Intn(100)
	if val == Value {
		fmt.Println("It matches", val, Value)
		matches = true
	} else {
		fmt.Println("Doesnt matches", val, Value)
		matches = false
	}
	return matches
}
