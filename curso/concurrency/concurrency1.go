package main

/*
	Concurrent programming is a form of programming in which we can execute multiple tasks simultaneously.
	Concurrency is not parallelism. Concurrency is the ability to run multiple tasks at the same time, but not necessarily simultaneously.
*/

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func sendEmail2(message string) {
	go func() { // This is a goroutine, it is a lightweight thread managed by the Go runtime (Concurrency)
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func test() {
	//sendEmail("Hello World")
	sendEmail2("Hello World")
	time.Sleep(time.Second)
}
