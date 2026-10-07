package main

import "fmt"

type Student struct {
	Name string
	Age  int
	GPA  float64
}

func changeGPA(GPA *float64) {
	*GPA = 4.0
}

func main() {
	student1 := Student{
		Name: "babby",
		Age:  29,
		GPA:  2.2,
	}
	changeGPA(&student1.GPA)
	fmt.Print(student1.GPA)
}
