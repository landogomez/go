package main

// VARIADIC FUNCTIONS
// A variadic function is a function that can accept a variable number of arguments
// The three dots (...) are what makes a function
// Example:

func sum(nums ...int) int {
	// nums is a slice of ints
	num := 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	return num
}
