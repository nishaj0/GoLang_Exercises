// ? 3. Create a function called "Search" which accepts an array
// ? 	  and an integer as input and returns the position within the
// ?    array if the integer is found within the array. If not found return -1.
// ?    Print the returned value in main function

package main

import "fmt"

func main() {
	nums := [5]int{1, 4, 2, 5, 2}

	fmt.Println(Search(nums[:], 1))
}

func Search(arr []int, searchItem int) int {

	for index, val := range arr {
		// if the item found return the position
		if searchItem == val {
			return index
		}
	}

	// if the item not found return -1
	return -1
}
