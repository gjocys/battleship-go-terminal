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

var coordRef = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
}

type game struct {
	playerName string
	isRunning  bool
	buffer     *bytes.Buffer
	ocean      ocean
	startedAt  time.Time
	error      string
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
	game.error = ""
	userInput := game.getUserInput()
	ok, x, y := game.getCoordinates(userInput)
	if ok {
		game.ocean.incomingMissile(x-1, y-1)
	}
}

func (game *game) getUserInput() []string {
	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')
	userInput := strings.Fields(rawInput)
	return userInput
}

func (game *game) getCoordinates(str []string) (bool, int, int) {
	if len(str) != 1 {
		game.error = "Invalid coordinates"
		return false, 0, 0
	}

	coord := strings.Join(str, "")

	x := coordRef[strings.ToLower(string(coord[0]))]
	y, _ := strconv.Atoi(string(coord[1:]))

	if x < 1 || y < 1 || x > 10 || y > 10 {
		game.error = "Invalid coordinates"
		return false, 0, 0
	}

	return true, x, y
}

func (game *game) loop() {
	for game.isRunning {
		game.render()
		game.update()
	}
}
