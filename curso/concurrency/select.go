package main

// SELECT
// The select statement lets a goroutine wait on multiple communication operations.
// It is similar to a switch statement, but with the cases all being channel operations.

/*
	select {
		case i, ok := <- chInts:
		// Do something
		case s, ok := <- chStrings:
		// Do something
	}

	The ok variable will be false if the channel is closed.

*/

// ASSIGNMENT

/*
	Complete the logMessages function

*/

import (
	"fmt"
)

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case email, ok := <-chEmails:
			if !ok {
				return
			}
			logEmail(email)
		case sms, ok := <-chSms:
			if !ok {
				return
			}
			logSms(sms)
		}
	}

}

func logSms(sms string) {
	fmt.Println("SMS: ", sms)
}

func logEmail(email string) {
	fmt.Println("Email: ", email)
}

func testSel(sms []string, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmail := sendToLogger(sms, emails)

	logMessages(chEmail, chSms)
	fmt.Println("=========")

}

func sendToLogger(sms []string, emails []string) (chan string, chan string) {
	chSms := make(chan string)
	chEmail := make(chan string)

	go func() {
		for _, sms := range sms {
			chSms <- sms
		}
		close(chSms)
	}()

	go func() {
		for _, email := range emails {
			chEmail <- email
		}
		close(chEmail)
	}()

	return chSms, chEmail
}
