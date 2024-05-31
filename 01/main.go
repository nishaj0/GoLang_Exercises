// ? 1. Create a function called "Add" which accepts two integers
// ?  as input and return the sum. Main function should call this
// ?  function and print the returned value

package main

import "fmt"

func main() {

	sum := Add(3, 5)
	fmt.Println(sum)
}

func Add(valOne int, valTwo int) int {
	return valOne + valTwo
}
