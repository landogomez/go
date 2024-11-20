// CLOSURES

// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense,
// the function is "bound" to the variables.

package main

import "fmt"

// Example

func concatter() func(string) string {
	doc := "" // Variable that will be shared with the inner function, but not with the outside world
	return func(word string) string {
		doc += word + " " // We concatenate the word with the shared variable
		return doc
	}
}

func exClosure() {
	harryPotterAggregator := concatter() // We create a new function that will use the shared variable
	harryPotterAggregator("Harry")       // We call the function with a word
	harryPotterAggregator("Potter")
	harryPotterAggregator("and")
	harryPotterAggregator("the")
	harryPotterAggregator("Philosopher's")
	harryPotterAggregator("Stone")
	harryPotterAggregator("is")
	harryPotterAggregator("a")
	harryPotterAggregator("book")
	harryPotterAggregator("by")
	harryPotterAggregator("J.K.")

	fmt.Println(harryPotterAggregator("Rowling"))
	// Output: Harry Potter and the Philosopher's Stone is a book by J.K. Rowling
}

// ASSIGNMENT

/*
	Complete the adder function.

	It should return a function that adds its input (an int) to an enclosed sum value, then return the new sum.
	In other words, it keeps a running total of the sum variable within a closure

*/

func adder() func(int) int {
	total := 0
	return func(x int) int {
		total += x
		return total
	}
}

type emailBill struct {
	costInPennies int
}

func test2(bills []emailBill) {
	defer fmt.Println("===============")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d ", countAdder(1), costAdder(bill.costInPennies))
	}
}

func assignmentClosure() {
	test2([]emailBill{{costInPennies: 10}, {costInPennies: 20}, {costInPennies: 30}})
}
