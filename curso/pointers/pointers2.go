package main

// ASSIGNMENT

/*
Complete the removeProfanity function. It should use the strings, ReplaceAll function to replace all instances of the following words in the input message with asterisks

"dang" -> ****
"shoot" -> *****
"heck" -> ****

It should mutate the value in the pointer and return nothing

*/

import (
	"fmt"
	"strings"
)

func removeProfanity(m *string) {
	messageVal := *m
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*m = messageVal
}

func test(messages []string) {
	for _, m := range messages {
		removeProfanity(&m)
		fmt.Println(m)
	}
}

func assignment2() {
	messages := []string{
		"Hey, dang it!",
		"shoot! I'm late",
		"What the heck?",
	}
	test(messages)
}

func cambiarVal() {
	var x int = 50
	var y *int = &x // y = 50
	*y = 100        // x = 100
}

// ASSIGNMENT

/* Update the removeProfanity function. If messages is nil, return eartly to avoid a panic. */

func removeProfanity2(m *string) error {
	if m == nil {
		return fmt.Errorf("nil pointer")
	}
	messageVal := *m
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*m = messageVal
	return nil
}
