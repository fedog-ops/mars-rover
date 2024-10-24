package mars_rover

type Rover struct {
	X, Y      int
	Direction rune
	PlanetMap *Map
}

func (r *Rover) Interpret(command rune) {
	switch command {
	case 'f':
		//distance := 1
		r.Move(1)
	case 'b':
		r.Move(-1)
	case 'l':
		r.Turn('l')
	case 'r':
		r.Turn('r')
	}
}

func (r *Rover) Move(distance int) {
	switch r.Direction {
	case 'N':
		r.Y += distance
	case 'S':
		r.Y -= distance
	case 'E':
		r.X += distance
	case 'W':
		r.X -= distance
	}

	//if new X,Y == obstacle
	//return	'can't move there' + last coord

	//goes off map
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
	}
}
