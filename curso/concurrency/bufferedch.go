// BUFFERED CHANNELS

// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

// ch := make(chan int, 100)

// This creates a buffered channel with a capacity of 100. Normally, sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
package main

import (
	"fmt"
)

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email: ", email)
	}
}

func testBuff(emails ...string) {
	fmt.Printf("Adding %v emails to queve...\n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Println("Sending emails...")
	sendEmails(len(emails), ch)
	fmt.Println("=========")
}
