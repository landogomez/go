package main

import (
	"fmt"
)

type fundsError struct {
	funds float64
}

func (f fundsError) Error() string {
	return fmt.Sprintf("Insufficient funds. Your account balance: %.2f", f.funds)
}

func withdraw(balance, amount float64) (float64, error) {
	if balance < amount {
		return 0, fundsError{funds: balance}
	}
	return balance - amount, nil
}

func testpractice(x, y float64) {
	defer fmt.Println("=========")
	result, err := withdraw(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Succesful Withdraw: %.2f\n", y)
	fmt.Printf("Remaining balance: %.2f\n", result)
}

func practice() {
	testpractice(150.0, 50.0)
	testpractice(100.0, 150.0)
}
