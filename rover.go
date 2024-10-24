package mars_rover

import "fmt"

type Rover struct {
	X, Y      int
	Direction rune
	PlanetMap *Map
	Output    func(string)
}

func (r *Rover) Interpret(commands string) {
	for _, command := range commands {
	var newX, newY int
	switch command {
	case 'f':
		newX, newY = r.FutureLocation(1)
	case 'b':
		newX, newY = r.FutureLocation(-1)
	case 'l':
		r.Turn('l')
	case 'r':
		r.Turn('r')
	} 
	
	if r.PlanetMap.IsObstacle(newX, newY) {
		r.Output(fmt.Sprintf("You hit rock at coordinates (%d, %d)", newX, newY))
		return
	}

	r.X, r.Y = newX, newY
	r.Output(fmt.Sprintf("New location: (%d, %d)", r.X, r.Y))
	}

}


func (r *Rover) FutureLocation(distance int) (int, int) {
	switch r.Direction {
	case 'N':
		return r.X, r.Y + distance
	case 'S':
		return r.X, r.Y - distance
	case 'E':
		return r.X + distance, r.Y
	case 'W':
		return r.X - distance, r.Y
	}

	return r.X, r.Y
}

func (r *Rover) Turn(direction rune) {
	switch direction {
	case 'l':
		switch r.Direction {
		case 'N':
			r.Direction = 'W'
		case 'E':
			r.Direction = 'N'
		case 'S':
			r.Direction = 'E'
		case 'W':
			r.Direction = 'S'
		}
	case 'r':
		switch r.Direction {
		case 'N':
			r.Direction = 'E'
		case 'E':
			r.Direction = 'S'
		case 'S':
			r.Direction = 'W'
		case 'W':
			r.Direction = 'N'
		}
	}
}

func NewRover(x int, y int, direction rune, m *Map) *Rover {
	return &Rover{
		X:         x,
		Y:         y,
		Direction: direction,
		PlanetMap: m,
		Output: func(message string) {
			fmt.Println(message)
		},
	}
}
