package main

import (
	"github.com/gjocys/battleship-go-terminal/battleship"
)

func main() {
	game := battleship.NewGame()
	game.Start()
}
