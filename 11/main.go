// ? 11. Continuing the above task, convert the json back to structure
//	? 	  and print the contents of the structure. Hint: json.Unmarshall

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

	// convert slice to json
	jsonByteData := sliceToJson(students)

	// convert json to struct
	studentStruct := jsonToStruct(jsonByteData)
	for _, student := range studentStruct {
		fmt.Printf("id: %v, name: %v, age:%v\n", student.ID, student.Name, student.Age)
	}
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

func jsonToStruct(jsonByteData []byte) []Student {

	// ? used to store struct data
	var studentStruct []Student

	// ? check json is valid
	checkValid := json.Valid(jsonByteData)

	if checkValid {
		err := json.Unmarshal(jsonByteData, &studentStruct)
		checkNilError(err)
	} else {
		fmt.Println("json not valid")
	}

	return studentStruct
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

// * output
// id: 1, name: Bala, age:24
// id: 2, name: Surya, age:25
// id: 3, name: Chan, age:33
// id: 4, name: Saffek, age:22
// id: 5, name: John, age:43
