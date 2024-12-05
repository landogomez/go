package main

// Generics OR Type Parameters
// Generics allow us to use variables to refer to specfic types. This allows us to write functions and data structures that can work with any type.

// EXAMPLE

/*
	func splitAnySlice[T any](s []T) ([]T, []T){
		mid := len(s) / 2
		return s[:mid], s[mid:]
	}

*/

// ASSIGNMENT

/*
	Complete the getLast() function. It shoulde be a generic function that returns the last element from a slice, no matter the types stored in the slice.
*/
import (
	"fmt"
)

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	return s[len(s)-1]
}

// CONSTRAINTS
// Constraints allow us to restrict the types that can be used with a generic function. We can use constraints to ensure that a generic
// function only works with types that have certain methods.

// EXAMPLE

type stringer interface {
	String() string
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String() // We know that T has a String() method
		// because of the constraint
	}
	return result
}

// ASSIGNMENT

/*
	Complete the chargeForLineItem function. First, it should check if the user has a balance with enough funds to be able to pay for the cost
	of the newItem. If they don't then return an "insufficient funds" error.
	If they do have enough funs:
		1. Add the line item to the user's history by appending the newItem to the slice of oldItems. This new slice is your first value to return
		2. Calculate the user's new balance by subtracting the cost of the newItem from the user's balance. This new balance is your second value to return.
*/

func chargeForLineItem[T lineItem](
	newItem T, oldItems []T, balance float64,
) ([]T, float64, error) {
	newBalance := balance - newItem.getCost()
	if newBalance < 0 {
		return nil, 0.0, fmt.Errorf("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	return oldItems, newBalance, nil
}

type lineItem interface {
	getCost() float64
	getName() string
}

// INTERFACE TYPE LISTS

// EXAMPLE

type stringer2 interface {
	~int | ~string | ~float64 | ~bool | ~rune | ~byte
}
