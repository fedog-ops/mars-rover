package mars_rover

type Rover struct {
    X, Y      int
    Direction rune
}

func (r *Rover) Move(command rune) {
	switch command {
        case 'f':
            switch r.Direction {
            case 'N':
                r.Y += 1
            case 'E':
                r.X += 1
            case 'S':
                r.Y -= 1
            case 'W':
                r.X -= 1
            }
		case 'b':
            switch r.Direction {
            case 'N':
                r.Y -= 1
            case 'E':
                r.X -= 1
            case 'S':
                r.Y += 1
            case 'W':
                r.X += 1
            }
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
		
    if r.X < 0 {r.X += 20}
	if r.X > 19 {r.X -= 20}
	if r.Y < 0 {r.Y += 20}
	if r.Y > 19 {r.Y -= 20}
}

func NewRover(x int, y int, direction rune) *Rover {
	return &Rover{
		X : x ,
		Y:	y,
		Direction: direction,
	}
}   