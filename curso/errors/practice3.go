package main

import (
	"errors"
	"fmt"
)

func checkPassword2(password string) (string, error) {
	if len(password) < 8 {
		return "", errors.New("ERROR! Password must be at least 8 characters long")
	}
	return fmt.Sprint("Password is correct"), nil
}

func test6(password string) {
	defer fmt.Println("=========")
	result, err := checkPassword2(password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func practice3() {
	test6("1234567")
	test6("12345678")
}
