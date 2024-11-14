package main

import "fmt"

// Arrays In GO
// Arrays are fixed size, meaning you can't change the size of an array once it's created.

// ASSIGNMENT

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func send(name string, doneAt int) {
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf("sending: %v", msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("complete failure")
		}
	}
}
func arrays1() {

	// Declare an array of 10 integers:
	//var myInts [10]int

	// Or declared an initialized literal:
	// primes := [6]int{2, 3, 5, 7, 11, 13}

	send("Bob", 0)
	send("Alice", 1)
	send("Mangalam", 2)
	send("George", 3)

}
