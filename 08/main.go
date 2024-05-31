// ? 8. Create a structure called student with ID, Name and Age.
// ?  Create a function called "create_student" which should take Id,
// ?  Name and Age as input and return a variable of structure
// ?  student initialized with the input values. Call the function
// ?  create_student(1,Bala,24) from main and print the returned
// ?  structure data of the format "Bala is a student with ID 1 and Age 24"

package main

import "fmt"

func main() {
	student := create_student(1, "Bala", 24)

	fmt.Printf("%v is a student with ID %v and age %v", student.Name, student.ID, student.Age)
}

func create_student(id int, name string, age int) Student {
	student := Student{id, name, age}
	return student
}

type Student struct {
	ID   int
	Name string
	Age  int
}

// * output
// Bala is a student with ID 1 and age 24
