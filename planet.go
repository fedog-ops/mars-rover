package mars_rover

type Coordinates struct {
	X, Y int
}

type Map struct {
	Width, Height int
	Obstacles     []Coordinates
}

func (m *Map) IsObstacle(x, y int) bool {
	for _, obstactle := range m.Obstacles {
		if obstactle.X == x && obstactle.Y == y {
			return true
		}
	}
	return false
}

func ExampleMap() *Map {
	obstacles := []Coordinates{
		{X: 3, Y: 4},
		{X: 7, Y: 1},
	}

	return &Map{
		Width:     10,
		Height:    10,
		Obstacles: obstacles,
	}
}
