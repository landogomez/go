// THERE IS NO WHILE LOOP IN GO

/*A while loop is just a for loop that only has a CONDITION

for CONDITION {
	// do some stuff while CONDITION is true
}
*/

package main

import "fmt"

func whileGo() {
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}

	fmt.Println("plant has grown to ", plantHeight, "inches")
}

func for2() {
	whileGo()
}
