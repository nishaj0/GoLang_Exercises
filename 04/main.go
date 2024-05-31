// ? 4. Create a program  to print "Hello <username>" from an array of usernames

package main

import "fmt"

func main() {
	users := [5]string{"john", "doe", "james", "smith", "ervin"}

	for _, user := range users {
		fmt.Printf("Hello %v\n", user)
	}
}

// * output
// Hello john
// Hello doe
// Hello james
// Hello smith
// Hello ervin
