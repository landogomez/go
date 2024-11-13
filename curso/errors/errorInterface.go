// Error is actually an interface type. It´s defined in the standard library´s errors package. The error interface is defined as:
// type error interface {
// 	Error() string
// }

package main

import "fmt"

type useError struct {
	name string
}

func (e useError) Error() string {
	return fmt.Sprintf("%v has a priblem with their account", e.name)
}

/*
func sendSMS(msg, userName string) error {
	if !canSendToUser(userName) {
	return useError{name: userName}
	}
}
*/

// ASSIGNMENT

type divideError struct {
	dividend float64
}

func (e divideError) Error() string {
	return fmt.Sprintf("Cannot divide %v by zero", e.dividend)
}

// ?

func dive(dividen, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.0, divideError{dividend: dividen}
	}
	return dividen / divisor, nil
}

func test2() {
	defer fmt.Println("=========")
	result, err := dive(1.0, 0.0)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Result: %.4f\n", result)
}

func errors2() {
	test2()
}
