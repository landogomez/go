package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n ", m.message, m.phoneNumber)
}

func assignment1() {
	test(messageToSend{1234567890, "Hello World"})
	test(messageToSend{9876543210, "Bye World"})
}
