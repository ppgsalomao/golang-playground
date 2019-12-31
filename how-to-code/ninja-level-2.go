package main

import "fmt"

func main() {
	fmt.Println("------------------\nExercise 1\n------------------")
	ninjaLevel2exercise1()
	fmt.Println("------------------\nExercise 2\n------------------")
	ninjaLevel2exercise2()
	fmt.Println("------------------\nExercise 3\n------------------")
	ninjaLevel2exercise3()
	fmt.Println("------------------\nExercise 4\n------------------")
	ninjaLevel2exercise4()
	fmt.Println("------------------\nExercise 5\n------------------")
	ninjaLevel2exercise5()
	fmt.Println("------------------\nExercise 6\n------------------")
	ninjaLevel2exercise6()
}

// ------------------------------------------
// Exercise 1
// ------------------------------------------

func ninjaLevel2exercise1() {
	number := 42
	fmt.Printf("Binary: %b\nDecimal: %d\nHex: %#x\n", number, number, number)
}

// ------------------------------------------
// Exercise 2
// ------------------------------------------

func ninjaLevel2exercise2() {

	test_condition_1 := 7 == 42
	fmt.Println("Test Condition 1 (7 == 42):", test_condition_1)
	test_condition_2 := 7 <= 42
	fmt.Println("Test Condition 2 (7 <= 42):", test_condition_2)
	test_condition_3 := 7 >= 42
	fmt.Println("Test Condition 3 (7 >= 42):", test_condition_3)
	test_condition_4 := 7 != 42
	fmt.Println("Test Condition 4 (7 != 42):", test_condition_4)
	test_condition_5 := 7 < 42
	fmt.Println("Test Condition 5 (7 < 42):", test_condition_5)
	test_condition_6 := 7 > 42
	fmt.Println("Test Condition 6 (7 > 42):", test_condition_6)
}

// ------------------------------------------
// Exercise 3
// ------------------------------------------

const coconut int = 42
const orange = "James Bond"

func ninjaLevel2exercise3() {
	fmt.Println(coconut)
	fmt.Println(orange)
}

// ------------------------------------------
// Exercise 4
// ------------------------------------------

func ninjaLevel2exercise4() {
	number := 42
	fmt.Printf("Binary: %b\nDecimal: %d\nHex: %#x\n", number, number, number)
	new_number := number << 1
	fmt.Printf("Binary: %b\nDecimal: %d\nHex: %#x\n", new_number, new_number, new_number)
}

// ------------------------------------------
// Exercise 5
// ------------------------------------------

func ninjaLevel2exercise5() {
	s := `Bond, James Bond`
	fmt.Println(s)
}

// ------------------------------------------
// Exercise 6
// ------------------------------------------

const (
	current_year = 2019 - iota
	last_year = current_year - iota
	two_years_before = current_year - iota
	three_years_before = current_year - iota
	four_years_before = current_year - iota
)

func ninjaLevel2exercise6() {
	fmt.Printf("%v\t%v\t%v\t%v", last_year, two_years_before, three_years_before, four_years_before)
}