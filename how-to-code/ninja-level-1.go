package main

import "fmt"

func main() {
	fmt.Println("------------------\nExercise 1\n------------------")
	exercise1()
	fmt.Println("------------------\nExercise 2\n------------------")
	exercise2()
	fmt.Println("------------------\nExercise 3\n------------------")
	exercise3()
}

// ------------------------------------------
// Exercise 1
// ------------------------------------------

func exercise1() {
	// int
	x := 42
	fmt.Println(x)

	// string
	y := "James Bond"
	fmt.Println(y)

	// boolean
	z := true
	fmt.Println(z)

	fmt.Println("Value for x: ", x, "\nValue for y: ", y, "\nValue for z: ", z)
}

// ------------------------------------------
// Exercise 2
// ------------------------------------------

var a int
var b string
var c bool

func exercise2() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// ------------------------------------------
// Exercise 3
// ------------------------------------------

var d int = 42
var e string = "James Bond"
var f bool = true

func exercise3() {
	s := fmt.Sprintf("%v\t%v\t%v", d, e, f)
	fmt.Println(s)
}
