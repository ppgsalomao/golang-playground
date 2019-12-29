package main

import "fmt"

func main() {
	fmt.Println("------------------\nExercise 1\n------------------")
	exercise1()
	fmt.Println("------------------\nExercise 2\n------------------")
	exercise2()
	fmt.Println("------------------\nExercise 3\n------------------")
	exercise3()
	fmt.Println("------------------\nExercise 4\n------------------")
	exercise4()
	fmt.Println("------------------\nExercise 5\n------------------")
	exercise5()
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

// ------------------------------------------
// Exercise 4
// ------------------------------------------

type hotdog int
var g hotdog

func exercise4() {
	fmt.Println(g)
	fmt.Printf("%T\n", g)
	g = 42
	fmt.Println(g)
}

// ------------------------------------------
// Exercise 5
// ------------------------------------------

type banana int
var h banana
var i int

func exercise5() {
	fmt.Println(h)
	fmt.Printf("%T\n", h)
	h = 42
	fmt.Println(h)

	i = int(h)
	fmt.Println(i)
	fmt.Printf("%T", i)
}
