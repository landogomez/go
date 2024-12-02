package main

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
