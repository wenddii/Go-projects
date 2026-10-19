package main

import "fmt"

func subtractExpense(balance int, expense int) int {
	return balance - expense
}

func addIncome(balance int, income int) int {
	return balance + income
}

func main() {
	balance := 1000

	expenses := []int{100, 200, 300, 100}

	for _, expense := range expenses {
		balance = subtractExpense(balance, expense)
	}

	
	balance = addIncome(balance, 500)

	fmt.Println("Remaining balance:", balance)
}
