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

type BankTransfer struct{}

func (B BankTransfer) Pay() {
	fmt.Println("paid through bank transfer")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

func main() {
	ProcessPayment(Telebirr{})
	ProcessPayment(Cash{})
	ProcessPayment(BankTransfer{})
}
