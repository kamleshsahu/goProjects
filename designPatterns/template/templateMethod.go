package main

import "fmt"

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

func PlayGame(g Game) {
	g.Start()

	for !g.HaveWinner() {
		g.TakeTurn()
	}

	fmt.Printf("Player %d winns.\n", g.WinningPlayer())
}

type chess struct {
	turn, maxTurns, currentPlayer int
}

func (c *chess) Start() {
	fmt.Println("starting a chess game")
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Println("Turn taken by player ", c.currentPlayer)
	c.currentPlayer = 1 - (c.currentPlayer)
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}

func NewGameOfChess() Game {
	return &chess{1, 10, 0}
}

func main() {
	chess := NewGameOfChess()
	PlayGame(chess)
}
