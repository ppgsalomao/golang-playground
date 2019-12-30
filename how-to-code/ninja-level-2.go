package main

import "fmt"

func main() {
	fmt.Println("------------------\nExercise 1\n------------------")
	ninjaLevel2exercise1()
}

// ------------------------------------------
// Exercise 1
// ------------------------------------------

func ninjaLevel2exercise1() {
	number := 42
	fmt.Printf("Binary: %b\nDecimal: %d\nHex: %#x", number, number, number)
}
