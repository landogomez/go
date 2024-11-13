package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return x / y, nil
}

func test3(x, y float64) {
	defer fmt.Println("=========")
	result, err := divide(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %.2f\n", result)

}

func errors3() {
	test3(1.0, 0.0)
	test3(1.0, 2.0)
	test3(1.0, 3.0)

}
