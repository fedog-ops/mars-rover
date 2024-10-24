package mars_rover

import (
	"mars_rover"
	"testing"
)
var emptyplanet *mars_rover.Map
var planet *mars_rover.Map

func TestMain(m *testing.M) {
	var err error
	emptyplanet, err = mars_rover.LoadMap("./../emptymap.json")
	if err != nil {
		panic(err)
	}
	var err2 error
	planet, err2 = mars_rover.LoadMap("./../map.json")
	if err2 != nil {
		panic(err2)
	}
	m.Run()
}
func TestEaxmapleMap(t *testing.T) {
	
	stubMap := mars_rover.ExampleMap()
	
	rover := mars_rover.NewRover(0, 0, 'N', stubMap)
	if rover.X != 0 || rover.Y != 0 || rover.Direction != 'N' {
		t.Fatalf("Expected (0, 0, N), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestLoadingMap(t *testing.T) {
	
	rover := mars_rover.NewRover(0, 0, 'N', emptyplanet)
	if rover.X != 0 || rover.Y != 0 || rover.Direction != 'N' {
		t.Fatalf("Expected (0, 0, N), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveForwardNorth(t *testing.T) {
	rover := mars_rover.NewRover(0, 0, 'N', emptyplanet)

	rover.Interpret('f')
	if rover.X != 0 || rover.Y != 1 {
		t.Fatalf("Expected (0, 1, N), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveForwardEast(t *testing.T) {
	rover := mars_rover.NewRover(0, 0, 'E', emptyplanet)

	rover.Interpret('f')
	if rover.X != 1 || rover.Y != 0 || rover.Direction != 'E' {
		t.Fatalf("Expected (1, 0, E), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveBackwardNorth(t *testing.T) {
	rover := mars_rover.NewRover(5, 5, 'N', emptyplanet)
	rover.Interpret('b')
	if rover.X != 5 || rover.Y != 4 {
		t.Fatalf("Expected (5, 4, N), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveBackwardSouth(t *testing.T) {
	rover := mars_rover.NewRover(5, 5, 'S', emptyplanet)
	rover.Interpret('b')
	if rover.X != 5 || rover.Y != 6 {
		t.Fatalf("Expected (5, 6, S), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}

func TestRoverMovesLeft(t *testing.T) {
	rover := mars_rover.NewRover(5, 5, 'N', emptyplanet)
	rover.Interpret('l')
	if rover.Direction != 'W' {
		t.Fatalf("Expected (5, 5, W), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMovesRight(t *testing.T) {
	rover := mars_rover.NewRover(2, 3, 'E', emptyplanet)
	rover.Interpret('r')
	if rover.Direction != 'S' {
		t.Fatalf("Expected (2, 3, S), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestHittingObject(t *testing.T){
	rover := mars_rover.NewRover(3, 3, 'N', planet)
	rover.Interpret('f')
	if  rover.X != 3 || rover.Y != 3 || rover.Direction != 'N' {
		t.Fatalf("Expected rock message and (3, 3, N), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}}

//multiple commands
//obstactle