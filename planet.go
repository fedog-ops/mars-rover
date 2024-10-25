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
func (m *Map) WrapCoordinates(x, y int) (int, int) {
	if x < 0 {
		x = m.Width - 1
	} else if x >= m.Width {
		x = 0
	}
	if y < 0 {
		y = m.Height - 1
	} else if y >= m.Height {
		y = 0
	}
	return x, y
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
