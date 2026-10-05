package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

func (c Car) Introduce() {
	fmt.Println("This is a", c.Brand, c.Model)
}

func (c Car) Describe() {
	fmt.Println(c.Brand, c.Model, "was made in", c.Year)
}

func main() {

	car1 := Car{
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2020,
	}
	car2 := Car{
		Brand: "Suzuki",
		Model: "Dzire",
		Year:  2022,
	}
	car1.Year = 2022
	fmt.Println(car1.Brand, car1.Model)
	fmt.Println(car2.Brand, car2.Model)
	car1.Introduce()
	car2.Introduce()
	car1.Describe()
	car2.Describe()
}
