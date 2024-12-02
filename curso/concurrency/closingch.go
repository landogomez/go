package main

import (
	"fmt"
	"time"
)

// Channels can be explicty closed by the sender to indicate that no more values will be sent.

/*
ch := make(chan int)

close(ch)

CHECKING IF A CHANNEL IS CLOSED
v, ok := <- ch
// ok is false if there are no more values to receive and the channel is closed
*/
func countReports(numSentCh chan int) int {
	total := 0

	for {
		numSent, ok := <-numSentCh // Check if the channel is closed
		if !ok {                   // If the channel is closed
			break // Exit the loop
		}
		total += numSent
	}
	return total
}

// RANGE
// Similar to slices, maps, and strings, channels can be iterated over with a range loop.

/*
	for item := range ch{
		fmt.Println(item)
	}
*/

// ASSIGNMENT

/*
	Complete the concurrentFib it should:
	- Create a new channel of int s.
	- Call a fibonacci in a goroutine, passing it the channel and the number of fibonacci numbers to generate.
	- Use a range loop to read from the channel and print out the numbers one by one, each on a new line.

*/

func concurrentFib(n int) {
	chInts := make(chan int)
	go func() {
		fibonacci(n, chInts)
	}()

	for num := range chInts {
		fmt.Println(num)
	}

}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
