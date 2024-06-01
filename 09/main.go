// ? 9. Continuing the above task, create a slice of 5 students using
// ?	the create_student function. The slice should NOT be created manually
// ? 	like arr=[]string{"Bala","Surya","Chan","Saffek",}. An empty
// ?  slice (of datatype student) called "students" needs to be created.
// ?	Each student needs to be created by calling create_student
// ?	with different values. The structure variable returned by create_student
// ?	has to be appended to the slice, "students". Once all 5 students are added
// ?	to the slice, loop the slice and print the contents of it

package main

import "fmt"

func main() {
	students := make([]Student, 0)

	students = append(students, create_student(1, "Bala", 24))
	students = append(students, create_student(2, "Surya", 25))
	students = append(students, create_student(3, "Chan", 33))
	students = append(students, create_student(4, "Saffek", 22))
	students = append(students, create_student(5, "John", 43))

	fmt.Println(len(students)) // 5

	for _, student := range students {
		fmt.Printf("%v is a student with ID %v and age %v\n", student.Name, student.ID, student.Age)
	}
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
// 5
// Bala is a student with ID 1 and age 24
// Surya is a student with ID 2 and age 25
// Chan is a student with ID 3 and age 33
// Saffek is a student with ID 4 and age 22
// John is a student with ID 5 and age 43
