package main

import "fmt"

func main() {
	students := map[string]float64{}
	students["wendwosen"] = 3.5
	students["abebe"] = 3.8
	students["emeye"] = 4.0

	_, exists := students["emeye"]

	if !exists {
		fmt.Println("Student not found")
	}
}


