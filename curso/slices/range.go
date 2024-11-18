package main

// RANGE
// Go provides synctatic sugar for looping over slices and maps using the range keyword

/*
for INDEX, ELEMENT := range SLICE {
}
*/

// FOR EXAMPLE

/*
fruits := []string{"apple", "banana", "cherry"}
for i, fruit := range fruits {
	fmt.Println(i, fruit)
}
	// 0 apple
	// 1 banana
	// 2 cherry
*/

// ASSIGNMENT

/*
Complete the indexOfFirstBadWord function. If it finds any bad words in the message
it should return the index of the first bad word in the msg
slice. If there are no bad words, it should return -1.

*/

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badbadWords := range badWords {
			if word == badbadWords {
				return i
			}
		}
	}
	return -1
}

func test(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` _`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("====================================")
}

func rangeGo() {
	test([]string{"hello", "world"}, []string{"world"})
	test([]string{"Hola", "todos"}, []string{"si"})
}
