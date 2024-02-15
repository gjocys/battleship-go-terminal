package battleship

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var hHeader = "     " + strings.Join(strings.Split("1,2,3,4,5,6,7,8,9,10", ","), "  ") + "\n"
var vHeader = map[int]string{
	0: "A",
	1: "B",
	2: "C",
	3: "D",
	4: "E",
	5: "F",
	6: "G",
	7: "H",
	8: "I",
	9: "J",
}

var hLine = "  +" + strings.Repeat("-", 32) + "+\n"

func (game *game) renderInfo() {
	game.buffer.WriteString("Player: " + game.playerName)
}

func (game *game) renderOcean() {
	game.buffer.WriteString(hHeader + hLine)
	for row := range game.ocean.grid {

		for col := range game.ocean.grid[row] {
			if col == 0 {
				game.buffer.WriteString(vHeader[row] + " | ")
			}
			str := ""
			cx := game.ocean.grid[row][col]
			resetColor := "\u001b[0m"
			switch cx {
			case 2:
				// missed
				color := "\u001b[30;1m"
				str = color + " * " + resetColor
			case 3:
				// hit
				color := "\u001b[33;1m"
				str = color + " X " + resetColor
			case 4:
				//sunk
				color := "\u001b[47;1m\u001b[31;1m"
				str = color + " X " + resetColor
			default:
				color := "\u001b[36;1m"
				str = color + " ~ " + resetColor
			}
			game.buffer.WriteString(str)
			if col == 9 {
				game.buffer.WriteString(" |")
			}
		}
		game.buffer.WriteString("\n")
		if row == 9 {
			game.buffer.WriteString(hLine + "\n")
		}
	}
}

func (game *game) renderFooter() {
	game.buffer.WriteString("Error: " + game.error + "\n")
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
