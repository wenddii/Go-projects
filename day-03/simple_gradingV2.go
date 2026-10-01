package main

import "fmt"

func main() {
	name := "wondwosen"
	score := 40
	credits := 125
	grade := getGrade(score)
	status := canGraduate(credits, grade)

	fmt.Println("Hello", name, "your grade is", grade)
	fmt.Println("can graduate : ", status)
}

func getGrade(score int) string {
	if score > 100 || score < 0 {
		return "not valid "
	}

	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func canGraduate(credits int, grade string) bool {
	if credits > 120 && grade != "F" {
		return true
	} else {
		return false
	}
}
