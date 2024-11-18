package main

import (
	"errors"
	"fmt"
)

// Slices In GO
// 99 times out of 100, you'll use slices instead of arrays when working with ordered lists

// Arrays are fixed in size. Slices are not.
// A slice is a dynamically-sized, flexible view into the elements of an array.

// Slices ALWAYS have an underlying array.

/*
EXAMPLE
primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
mySlice = {3, 5, 7}

*/

// ASSIGNMENT

// Complete the getMessageWothRetriesForPLan function. It takes a plan variable as input, and you´'ve been provided with constants
// for the plan types at the top of the program
/*
	1. If the plan is a pro plan, return all the strings from getMessageWithRetries()
	2. If the plan is a free plan, return the first 2 strings from getMessageWithRetries()
	3. If the plan isn´t either of those, return an error that says unsupported plan.
*/

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries2()
	if plan == planPro {
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
}

func getMessageWithRetries2() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func slicesGo() {
	fmt.Println(getMessageWithRetriesForPlan(planPro))
	fmt.Println(getMessageWithRetriesForPlan(planFree))
	fmt.Println(getMessageWithRetriesForPlan("enterprise"))
}
