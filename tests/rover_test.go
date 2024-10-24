package tests

import (
	"mars_rover"
	"testing"
)

func TestRoverInitialization(t *testing.T) {
	rover := mars_rover.NewRover(0,0,'N')
	if rover.X != 0 || rover.Y != 0 || rover.Direction != 'N' {
		t.Fatalf("Expected (0, 0, N), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveForwardNorth(t *testing.T) {
	rover := mars_rover.NewRover(0, 0, 'N')
	
	rover.Move('f')
	if rover.X != 0 || rover.Y != 1 {
		t.Fatalf("Expected (0, 1, N), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveForwardEast(t *testing.T) {
	rover := mars_rover.NewRover(0, 0, 'E')
	
	rover.Move('f')
	if rover.X != 1 || rover.Y != 0 || rover.Direction != 'E' {
		t.Fatalf("Expected (1, 0, E), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveBackwardNorth(t *testing.T) {
	rover := mars_rover.NewRover(5, 5, 'N')
	rover.Move('b')
	if rover.X != 5 || rover.Y != 4 {
		t.Fatalf("Expected (5, 4, N), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveBackwardSouth(t *testing.T) {
	rover := mars_rover.NewRover(5, 5, 'S')
	rover.Move('b')
	if rover.X != 5 || rover.Y != 6 {
		t.Fatalf("Expected (5, 6, S), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}

func TestRoverMovesLeft(t *testing.T) {
	rover := mars_rover.NewRover(5,5,'N')
	rover.Move('l')
	if rover.Direction != 'W'{
		t.Fatalf("Expected (5, 5, W), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMovesRight(t *testing.T) {
	rover := mars_rover.NewRover(2,3,'E')
	rover.Move('r')
	if rover.Direction != 'S'{
		t.Fatalf("Expected (2, 3, S), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestEdgeOfMap(t *testing.T) {
	rover := mars_rover.NewRover(0, 4, 'W')
	rover.Move('f')
	if rover.X != 19 || rover.Y != 4 {
		t.Fatalf("Expected (19, 4, W), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}

	roverB := mars_rover.NewRover(19, 4, 'E')
	roverB.Move('f')
	if roverB.X != 0 || roverB.Y != 4 {
		t.Fatalf("Expected (0, 4, E), got (%d, %d, %c)", roverB.X, roverB.Y, roverB.Direction)
	}
	roverC := mars_rover.NewRover(19, 19, 'N')
	roverC.Move('f')
	if roverC.X != 19 || roverC.Y != 0 {
		t.Fatalf("Expected (19, 0, N), got (%d, %d, %c)", roverC.X, roverC.Y, roverC.Direction)
	}
}

//multiple commands
//obstactle