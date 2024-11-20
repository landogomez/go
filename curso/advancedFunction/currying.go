package main

import "fmt"

// CURRYING
// Function currying is the practice of writing a function that takes a function (or multiple functions) as an argument
// and returns a new function.

//EXAMPLE

func multiply(x, y int) int {
	return x * y
}

func add2(x, y int) int {
	return x + y
}

func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func currying() {
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(add2)

	fmt.Println(squareFunc(5))
	fmt.Println(doubleFunc(5))
}

// ASSIGNMENT

// Complete the getLogger function. It should retutrn a new function that prints the formatted
// inputs using the given FORMATTER function.

func getLogger(formatter func(string, string) string) func(string, string) {
	return func(a string, b string) {
		fmt.Println(formatter(a, b))
	}
}
