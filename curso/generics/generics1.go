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

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	return s[len(s)-1]
}
