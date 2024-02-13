package battleship

type ship struct {
	coord [][]int
	size  int
	hits  int
	sunk  bool
}

func newShip(coord [][]int, size int) *ship {
	return &ship{
		coord: coord,
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
