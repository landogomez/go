package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

// aggregate applies the the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func advFunc1() {
	fmt.Println(aggregate(2, 3, 4, add))
	fmt.Println(aggregate(2, 3, 4, mul))
}

// DEFINITION: FIRST-CLASS FUNCTION
// A first-class function is a function that can be treated like any other value.

// DEFINITION: HIGHER-ORDER FUNCTION
// A higher-order function is a function that takes one or more functions as arguments or returns a function as its result.

// ASSIGNEMT
/*
	Textio is launching a new email messaging product, "Malio"!

	Fix the compile-time bug in the getFormattedMessages function. The function body is correct, but the function
	signature is not

*/

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}

	return formattedMessages
}

func addSignature(message string) string {
	return message + "Kind regards."
}

func addGreeting(message string) string {
	return "Hello!" + message
}

func test(messages []string, formatter func(string) string) {
	formattedMessages := getFormattedMessages(messages, formatter)
	for _, message := range formattedMessages {
		fmt.Println(message)
	}
}

func assignment1() {
	messages := []string{"Hi i wosh you are havin a great day", "Please let me know if you need something else", "Hey"}
	test(messages, addSignature)
	test(messages, addGreeting)
}
