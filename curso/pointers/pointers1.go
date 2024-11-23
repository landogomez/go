package main

// Pointers are used to store the memory address of another variable.

/*
	The * syntax defines a pointer:
	var p *int


	The & operator generates a pointer to its operand:
	myString := "hello"
	myStringPtr ;= &myString

*/

// ASSIGNMENT

/*
	Fix the bug in the sendMessage function. It's supposed to to print a nicely
	formatted message to the console containing an SMS recipient message and messages body

*/

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

func sendMessage(m Message) {
	fmt.Printf("To %v\n", m.Recipient) // Pointer to Recipient
	fmt.Printf("Message %v\n", m.Text) // Pointer to Text
}

func assignment1() {
	fmt.Println("Hola")
}
