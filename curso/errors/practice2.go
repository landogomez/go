package main

import "fmt"

type lenPasswordError struct {
	Len int
}

func (l lenPasswordError) Error() string {
	return fmt.Sprintf("Password must be at least %d characters long", l.Len)
}

func checkPassword(password string) error {
	if len(password) < 8 {
		return lenPasswordError{Len: 8}
	}
	return nil
}

func test5(password string) {
	defer fmt.Println("=========")
	err := checkPassword(password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Password is valid")
}

func practice2() {
	test5("1234567")
	test5("12345678")
}
