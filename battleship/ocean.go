package battleship

type ocean struct {
	ships []*ship
	grid  [][]int
}

func newOcean() ocean {

	// TODO generate ships
	ships := []*ship{
		newShip([][]int{{0, 1}, {0, 2}}, 2),
	}

	grid := make([][]int, 10)
	for i := range grid {
		grid[i] = make([]int, 10)
	}

	return ocean{
		ships: ships,
		grid:  grid,
	}
}

func (o *ocean) findShipByCoord(x int, y int) (found bool, ship *ship) {
	ships := o.ships
	for _, ship := range ships {
		for _, coord := range ship.coord {
			if x == coord[0] && y == coord[1] {
				return true, ship
			}
		}
	}
	return false, nil
}

func (o *ocean) incomingMissile(x int, y int) {
	ok, ship := o.findShipByCoord(x, y)

	//missed
	if !ok {
		o.grid[x][y] = 2
		return
	}

	//hit
	isSunk := ship.hit()

	if isSunk {
		for _, coord := range ship.coord {
			o.grid[coord[0]][coord[1]] = 4
		}
	} else {
		o.grid[x][y] = 3
	}
}
