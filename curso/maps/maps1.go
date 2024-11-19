// MAPS

//	Maps are a data structure that store key-value pairs. The keys are unique and the values are indexed by the keys.
// We can create map by using a literal or by using the make() function.
/*
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	ages["bob"] = 32

	// or

	ages := map[string]int{
		"alice": 31,
		"charlie": 34,
		"bob": 32,
	}

	// The len() function works on a map, it returns the total number of a key/value pairs.
	fmt.Println(len(ages)) -> 3
*/

// ASSIGNMENT

/*
	Complete the getUserMap function. It takes a slice of names and a slice of phone numbers,
	and returns a map of name -> user structs and potentially an error. A user struct just contains
	a user's name and phone number.

	If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".

	If the first name in the names slice mathes the first phone number,and so on.
*/

package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil

}

type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("============")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println("- number:", users[name].phoneNumber)
	}
}

func maps1() {
	test([]string{"alice", "bob", "charlie"}, []int{123, 456, 789})
	test([]string{"alice", "bob", "charlie"}, []int{123, 456})
}

// MUTATIONS

/*
	INSERT AN ELEMENT
	m[key] = value

	GET AN ELEMENT
	value = m[key]

	DELETE AN ELEMENT
	delete(m, key)

	CHECK IF A KEY EXISTS
	value, ok := m[key]

*/

// ASSIGNMENT
/*
	Complete the deleteIfNecessary function.

	- If the user doesn't exist in the map, return the error not found
	- If they exist, but aren't schedule for deletion, return deleted as false with
	  no errors.
	- If they exist and are scheduled for deletion, return deleted as true with no errors
	  and delete the user from the map.

*/

func deleteIfNecessary(m map[string]user2, name string) (deleted bool, err error) {
	existingUser, ok := m[name]
	if !ok {
		return false, errors.New("not found")
	}
	if existingUser.scheduleForDeletion {
		delete(m, name)
		return true, nil
	}
	return false, nil
}

type user2 struct {
	name                string
	number              int
	scheduleForDeletion bool
}

// You can use structs as map keys if the struct has no fields that are maps, slices, or functions.
// Effective go book.
