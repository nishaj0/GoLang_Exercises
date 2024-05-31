// ? 5. Create a program  to print "Hello <username>. How are you doing?"  from an array of usernames

package main

import "fmt"

func main() {
	users := [5]string{"john", "doe", "james", "smith", "ervin"}

	for _, username := range users {
		fmt.Printf("Hello %v. How are you doing?\n", username)
	}
}

// * output
// Hello john. How are you doing?
// Hello doe. How are you doing?
// Hello james. How are you doing?
// Hello smith. How are you doing?
// Hello ervin. How are you doing?
