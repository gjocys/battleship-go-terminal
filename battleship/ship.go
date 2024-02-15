package battleship

import (
	"math/rand"
	"time"
)

type ship struct {
	coord [][]int
	size  int
	hits  int
	sunk  bool
}

func newShip(size int) *ship {
	return &ship{
		coord: generateCoordinates(size),
		size:  size,
		hits:  0,
		sunk:  false,
	}
}

func (s *ship) hit() bool {
	s.hits++
	if s.hits == s.size {
		s.sunk = true
	}
	return s.sunk
}

func randomDirection() bool {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(2) == 1
}

func generateCoordinates(size int) [][]int {
	var x, y int

	ship := make([][]int, size)
	for i := range ship {
		ship[i] = make([]int, 2)
	}

	//  true - vertical, false - horizontal
	vertical := randomDirection()

	// gen first coordinate - random x and random y depending on direction and ship size
	if vertical {
		x = rand.Intn(10-0) + 0
		y = rand.Intn(10-size-0) + 0

	} else {
		x = rand.Intn(10-size-0) + 0
		y = rand.Intn(10-0) + 0
	}

	ship[0][0] = x
	ship[0][1] = y

	for i := 0; i < size; i++ {
		if !vertical {
			x++
		} else {
			y++
		}
		ship[i][0] = x
		ship[i][1] = y
	}

	return ship
}
