// ? 6. Create a program  to print "Hello my name is
// ?  '<username>'. I am a student".  From an array of students.
// ?  Notice the single quote surrounding the <username>.

package main

import "fmt"

func main() {
	students := [5]string{"john", "doe", "james", "smith", "ervin"}

	for _, student := range students {
		fmt.Printf("Hello my name is  '%v'. I am a student\n", student)
	}
}

// * output
// Hello my name is  'john'. I am a student
// Hello my name is  'doe'. I am a student
// Hello my name is  'james'. I am a student
// Hello my name is  'smith'. I am a student
// Hello my name is  'ervin'. I am a student
