package main

// APPEND
// We can append to a slice using the append function
// The append function returns a new slice with the new value added to the end
// If the underlying array is too small, a new array will be created and the values copied over
// The original slice will not be modified
// Example:
/*
slice = append(slice, firstThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)

*/

// ASSIGNMENT
/*
	We've been asked to "bucket" costs for an entire month into the cost ocurring on each day of the month

	Complete the getCostsByDay function. It should return a slice of float64, where each element is the total
	cost for that day. The length of the slice should be equal to the number of days represented in the
	costs slice, including any days that have no costs.

*/

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}

	return costsByDay
}
