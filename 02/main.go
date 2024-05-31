//	? 2. Create an array of 10 integers. Pass the array as input
// ?	to a function called "AddArray" which accepts as input an
// ?	array and it's size and returns the sum of all elements in the array

package main

import "fmt"

func main() {
	nums := [10]int{1, 4, 2, 5, 2, 6, 2, 5, 7, 8}

	fmt.Println(AddArray(nums[:], len(nums)))
}

func AddArray(arr []int, size int) int {
	arrSum := 0

	for i := 0; i < size; i++ {
		arrSum += arr[i]
	}

	return arrSum
}
