package main

import "fmt"

func main() {
	var student struct {
		id   int
		name string
		gpa  float32
	}

	student.id = 202444049
	student.name = "Maeng Eunjae"
	student.gpa = 3.5
	fmt.Println(student.gpa)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}

	student2.id = 202444040
	student2.name = "Unkwon"
	student2.gpa = 3.6
	fmt.Println(student2.gpa)
}
