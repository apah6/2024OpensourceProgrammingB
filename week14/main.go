package main

import "fmt"

type student struct {
	id   int
	name string
	gpa  float32
}

func main() {
	var student1 student
	student1.id = 202444049
	student1.name = "Maeng Eunjae"
	student1.gpa = 3.5
	fmt.Println(student1.gpa)
	var student2 student
	student2.id = 202444040
	student2.name = "Unknown"
	student2.gpa = 3.6
	fmt.Println(student2.gpa)
}
