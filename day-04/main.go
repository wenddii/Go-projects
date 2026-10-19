package main

import "fmt"

type GameSession struct {
	customerName string
	gameName     string
	gamesPlayed  int
	pricePerGame float64
}

func main() {

	game(x) := GameSession{
		customerName: "Wendwosen",
		gameName:     "FIFA",
		gamesPlayed:  3,
		pricePerGame: 50,
	}
	x int 
	x := x + 1

	fmt.Println(game1)
	total := game1.pricePerGame * float64(game1.gamesPlayed)
	

	fmt.Print(total)

	
	playagain bool 
	fmt.Print("want to play agian? " playagain)
	if playagain == true {
		game(x)
	}else { 
		return
	}
}



