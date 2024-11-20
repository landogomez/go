// ANONYMOUS FUNCTION
// An anonymous function is a function that does not have a name.

// EXAMPLE

// doMath accepts a function that converts one int into another
// and a slice of ints that have been
// converted by the passed in function.

package main

import "fmt"

func doMath(f func(int) int, nums []int) []int {
	var results []int
	for _, n := range nums {
		results = append(results, f(n))
	}
	return results
}

func anonymousEx() {
	nums := []int{1, 2, 3, 4, 5}

	// Here we define an anonymous function that doubles an int
	// and pass it to doMath
	allNumsDoubled := doMath(func(x int) int { // Anonymous function
		return x + x
	}, nums)

	fmt.Println(allNumsDoubled)
}

// ASSIGNMENT

// Complete the printCostReport once for each message. Pass in an anonymous function as the costCalculator that return an int
// equal twice the length of the message.

func printReport(messages []string) {
	for _, message := range messages {
		printCostReport(func(string) int {
			return len(message) * 2
		}, message)
	}
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message) // We call the anonymous function
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
