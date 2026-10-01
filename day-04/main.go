package main

import "fmt"

type GameSession struct {
	customerName string
	gameName     string
	gamesPlayed  int
	pricePerGame float64
}

func main() {

	game1 := GameSession{
		customerName: "Wendwosen",
		gameName:     "FIFA",
		gamesPlayed:  3,
		pricePerGame: 50,
	}

	fmt.Println(game1)
	total := game1.pricePerGame * game1.pricePerGame

	fmt.Print(total)

}
