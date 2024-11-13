package main

import "fmt"

func sendsSMStoCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	cost, err := sendsSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	cost2, err := sendsSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return cost + cost2, nil

}

func sendsSMS(message string) (float64, error) {
	const maxTextLen = 45
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("Can´t send text over %v character", maxTextLen) // fmt.Errorf is a function that returns an error
	}
	return costPerChar * float64(len(message)), nil // nil means no error
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("=========")
	fmt.Println("Message for customer: ", msgToCustomer)
	fmt.Println("Message for spouse: ", msgToSpouse)
	totalCost, err := sendsSMStoCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func errors1() {

	test("I'm going to be late for dinner", "Please don´t wait for me")
	test("Hi honey, I´m going to be late for dinner, I´ll be back in 30 minutes", "Please don´t wait for me")
}
