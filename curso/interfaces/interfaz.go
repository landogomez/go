package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage()) // Notice that we don't really know if its a birthday message or a sending report
}

// message is an interface that has a method called send
type message interface {
	getMessage() string // Requirement for the interface
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string { // Messages
	return fmt.Sprintf("Hi: %s, iti is your birthday on %s", bm.recipientName, bm.birthdayTime)
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string { // This is the implementation of the interface so we can think of both as messages
	return fmt.Sprintf("Your %s report is ready. You've sent %v messages", sr.reportName, sr.numberOfSends)
}

func test(m message) { // This function can take any type that implements the message interface
	sendMessage(m)
}

func interaz() {
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		birthdayTime:  time.Now(),
		recipientName: "John",
	})
	test(sendingReport{
		reportName:    "Second Report",
		numberOfSends: 20,
	})
}
