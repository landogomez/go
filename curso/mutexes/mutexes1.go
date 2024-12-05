package main

// Mutexes
// A mutex is used to synchronize access to shared state across multiple goroutines.
// A mutex has two methods: Lock and Unlock.
// We can use a mutex to safely access data across multiple goroutines.

// MAPS ARE NOT THREAD-SAFE
// Maps are not safe for concurrent use: it's not safe to read and write to them simultaneously.
// To avoid this, we can use a mutex to synchronize access to the map.

// ASSIGNMENT
/*
	We send emails across many different goroutines at Mailio. To keep track of how many we've sent to given email address, we use an in-memory
	map.

	Our safeCounter struct is unsafe! Updtate the inc() and val() methods so that they utilize the safeCounter's mutex yo ensure that the map is
	not accessed by multiple goroutines at the same time.

*/

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mux    *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	return sc.counts[key]
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter

}

type emailTest struct {
	email string
	count int
}

func test(sc safeCounter, emailTests []emailTest) {
	emails := make(map[string]struct{})

	var wg sync.WaitGroup
	for _, emailT := range emailTests {
		emails[emailT.email] = struct{}{}
		for i := 0; i < emailT.count; i++ {
			wg.Add(1)
			go func(emailT emailTest) {
				sc.inc(emailT.email)
				wg.Done()
			}(emailT)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailsSorted = append(emailsSorted, email)
	}
	sort.Strings(emailsSorted)

	for _, email := range emailsSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("=========")
}
