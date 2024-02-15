package battleship

type ocean struct {
	ships []*ship
	grid  [][]int
}

func newOcean() ocean {

	ships := addShips()
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

func (o *ocean) incomingMissile(x int, y int) (hit bool, sunk bool) {
	ok, ship := o.findShipByCoord(x, y)

	//missed
	if !ok {
		o.grid[x][y] = 2
		return false, false
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

	return true, isSunk
}

func addShips() []*ship {

	ships := []*ship{}

	shipSizes := []int{5, 4, 3, 3, 2}

	for i := range shipSizes {
		for {
			s := newShip(shipSizes[i])
			if !shipsCollide(ships, s) {
				ships = append(ships, s)
				break
			}
		}
	}

	return ships
}

func shipsCollide(a []*ship, b *ship) bool {

	for i := range b.coord {
		bCoordX := b.coord[i][0]
		bCoordY := b.coord[i][1]
		for n := range a {
			for v := range a[n].coord {
				aCoordX := a[n].coord[v][0]
				aCoordY := a[n].coord[v][1]
				if aCoordX == bCoordX && aCoordY == bCoordY {
					return true
				}
			}
		}
	}

	return false
}
