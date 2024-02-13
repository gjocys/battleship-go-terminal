package battleship

import (
	"fmt"
	"os"
	"time"
)

func (game *game) renderInfo() {
	game.buffer.WriteString("Player: " + game.playerName)
}

func (game *game) renderOcean() {
	for row := range game.ocean.grid {
		for col := range game.ocean.grid[row] {
			str := ""
			cx := game.ocean.grid[row][col]
			switch cx {
			case 2:
				// missed
				str = " O "
			case 3:
				// hit
				str = " X "
			case 4:
				//sunk
				str = " Z "
			default:
				str = " ~ "
			}
			game.buffer.WriteString(str)
		}
		game.buffer.WriteString("\n")
	}
}

func (game *game) renderFooter() {
	game.buffer.WriteString("Enter coordinates: ")
}

func (game *game) render() {
	fmt.Fprintf(os.Stdout, "\033[2J\033[1;1H") // clear terminal

	game.renderInfo()
	game.renderOcean()
	game.renderFooter()

	fmt.Fprintf(os.Stdout, game.buffer.String())

	time.Sleep(time.Millisecond * 100) // cap FPS

	game.buffer.Reset()
}
