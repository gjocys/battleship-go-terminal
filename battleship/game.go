package battleship

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type game struct {
	playerName string
	isRunning  bool
	buffer     *bytes.Buffer
	ocean      ocean
	startedAt  time.Time
}

func NewGame() *game {
	fmt.Printf("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	return &game{
		playerName: name,
		buffer:     new(bytes.Buffer),
		startedAt:  time.Now(),
	}
}

func (game *game) Start() {
	game.ocean = newOcean()
	game.isRunning = true
	game.loop()
}

func (game *game) update() {
	//TODO proper way to read user input
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	split := strings.Split(userInput, "")

	//check if hit
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	game.ocean.incomingMissile(x, y)
}

func (game *game) loop() {
	for game.isRunning {
		game.render()
		game.update()
	}
}
