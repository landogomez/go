package main

import (
	"fmt"
	"time"
)

// Channels are a way to communicate between goroutines
// Channels are a way to send and receive values between goroutines
/*
   Like maps and slices, channels must be created before use:

   ch := make(chan int)

   SEND DATA TO A CHANNEL
   ch <- 42

   The "<-" is called the channel operator. Data flows in the direction of the arrow. This operation will block another goroutine until the data is received.

   RECEIVE DATA FROM A CHANNEL

    x := <- ch

*/

func waitForDbs(numDbs int, dbChan chan struct{}) {
	for i := 0; i < numDbs; i++ {
		<-dbChan // This will block until a value is received
	}
}

func testCH(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	fmt.Println("=========")
	time.Sleep(time.Millisecond * 10)
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}
