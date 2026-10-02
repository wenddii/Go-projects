package main

import "fmt"

type Payment interface {
	Pay()
}

type Cash struct{}

func (c Cash) Pay() {
	fmt.Println("paid with cash")
}

type Telebirr struct{}

func (c Telebirr) Pay() {
	fmt.Println("paid with Telebirr")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

func main() {
	ProcessPayment(Telebirr{})
	ProcessPayment(Cash{})
}
