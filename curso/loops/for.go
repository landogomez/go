package main

import "fmt"

// ASSIGNEMT 1 bulkSend() should return the total cost (as float65) to send a batch of numMessages messages. Each message costs 1.0, plus an additional
// fee. The fee structure is as follows:
// 1st message: 1.0 + 0.0
// 2nd message: 1.0 + 0.01
// 3rd message: 1.0 + 0.02
// 4th message: 1.0 + 0.03

func bulkSend(numMessages int) float64 {
	var cost float64
	for i := 1; i <= numMessages; i++ {
		cost += 1.0 + float64(i-1)*0.01
	}

	return cost
}

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("====================================")
}

// ASSIGNMENT 2
// Complete the maxMessages function. Given a cost threshold. it should calculate the maximum number of messages that can be sent

// Each message costs 1.0, plus an additional fee. The fee structure is

// 1st message: 1.0 + 0.00
// 2nd message: 1.0 + 0.01
// 3rd message: 1.0 + 0.02
// 4th message: 1.0 + 0.03

func maxMessages(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ { // In go you con ommit the condition in a for loop
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}

func test2(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: %v\n", max)
	fmt.Println("===============")
}

func forloop() {
	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	// Ommiting conditions from a for loop

	test(10)
	test2(10.00)
}
