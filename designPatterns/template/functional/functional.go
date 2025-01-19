package functional

import "fmt"

func PlayGame(start, takeTurn func(), haveWinner func() bool, winningPlayer func() int) {
	start()

	for !haveWinner() {
		takeTurn()
	}

	fmt.Println("Player wins:", winningPlayer())
}

func main() {
	//turn, maxTurns, currentPlayer := 1, 10, 0

	// define functions
	// then call PlayerGame with given methods

}
