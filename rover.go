package mars_rover

import "fmt"

type Rover struct {
	Position  Position
	PlanetMap *Map
}
type Position struct {
	X int
	Y int
	D string
}

func (r *Rover) Command(commands string) {
	for _, command := range commands {
		var newX, newY int
		switch command {
		case 'f':
			newX, newY = r.FutureLocation(1)
		case 'b':
			newX, newY = r.FutureLocation(-1)
		case 'l':
			r.Turn("l")
			newX, newY = r.Position.X, r.Position.Y
		case 'r':
			r.Turn("r")
			newX, newY = r.Position.X, r.Position.Y
		}

		if r.PlanetMap.IsObstacle(newX, newY) {
			r.Output(fmt.Sprintf("You hit rock at coordinates (%d, %d)", newX, newY))
			return
		}

		r.Position.X, r.Position.Y = newX, newY
		r.Output(fmt.Sprintf("New location: (%d, %d)", r.Position.X, r.Position.Y))
	}

}

func (r *Rover) FutureLocation(distance int) (int, int) {
	switch r.Position.D {
	case "N":
		return r.Position.X, r.Position.Y + distance
	case "S":
		return r.Position.X, r.Position.Y - distance
	case "E":
		return r.Position.X + distance, r.Position.Y
	case "W":
		return r.Position.X - distance, r.Position.Y
	}

	return r.Position.X, r.Position.Y
}

func (r *Rover) Turn(direction string) {
	switch direction {
	case "l":
		switch r.Position.D {
		case "N":
			r.Position.D = "W"
		case "E":
			r.Position.D = "N"
		case "S":
			r.Position.D = "E"
		case "W":
			r.Position.D = "S"
		}
	case "r":
		switch r.Position.D {
		case "N":
			r.Position.D = "E"
		case "E":
			r.Position.D = "S"
		case "S":
			r.Position.D = "W"
		case "W":
			r.Position.D = "N"
		}
	}
}

func (r *Rover) GetPosition() Position { return r.Position }

func (r *Rover) Output(message string) { fmt.Println(message) }

func NewRover(p *Position, m *Map) *Rover {
	return &Rover{
		Position:  *p,
		PlanetMap: m,
	}
}
