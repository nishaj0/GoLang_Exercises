// ? 10. Continuing the above task create a json from the student
// ? 		structure and print it. Hint: Use json.Marshall

package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID   int
	Name string
	Age  int
}

func main() {
	students := make([]Student, 0)

	students = append(students, create_student(1, "Bala", 24))
	students = append(students, create_student(2, "Surya", 25))
	students = append(students, create_student(3, "Chan", 33))
	students = append(students, create_student(4, "Saffek", 22))
	students = append(students, create_student(5, "John", 43))

	// convert to json
	fmt.Println(string(sliceToJson(students)))
}

func create_student(id int, name string, age int) Student {
	student := Student{id, name, age}
	return student
}

func sliceToJson(sliceData []Student) []byte {
	createdJson, err := json.MarshalIndent(sliceData, "", "\t")
	checkNilError(err)

	return createdJson
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

// * output

// [
//         {
//                 "ID": 1,
//                 "Name": "Bala",
//                 "Age": 24
//         },
//         {
//                 "ID": 2,
//                 "Name": "Surya",
//                 "Age": 25
//         },
//         {
//                 "ID": 3,
//                 "Name": "Chan",
//                 "Age": 33
//         },
//         {
//                 "ID": 4,
//                 "Name": "Saffek",
//                 "Age": 22
//         },
//         {
//                 "ID": 5,
//                 "Name": "John",
//                 "Age": 43
//         }
// ]
