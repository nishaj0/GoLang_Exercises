// ? 7. Create a structure called "student" with elements
// ? 	ID(int), Name(string), Age(int). Initialize the structure
// ? 	with values (1, Bala, 24). Print the structure elements
// ?  in this format "Bala is a student with ID 1 and age 24"

package main

import "fmt"

func main() {
	studentBala := Student{1, "Bala", 24}

	fmt.Printf("%v is a student with ID %v and age %v", studentBala.Name, studentBala.ID, studentBala.Age)
}

type Student struct {
	ID   int
	Name string
	Age  int
}

// * output
// Bala is a student with ID 1 and age 24
