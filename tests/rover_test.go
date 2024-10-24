package tests

import (
	"mars_rover"
	"testing"
	
)
var planet *mars_rover.Map

func TestMain(m *testing.M) {
	var err error
	planet, err = mars_rover.LoadMap("./../map.json")
	if err != nil {
		panic(err)
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
	
	rover := mars_rover.NewRover(0, 0, 'N', planet)
	if rover.X != 0 || rover.Y != 0 || rover.Direction != 'N' {
		t.Fatalf("Expected (0, 0, N), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveForwardNorth(t *testing.T) {
	rover := mars_rover.NewRover(0, 0, 'N', planet)

	rover.Interpret('f')
	if rover.X != 0 || rover.Y != 1 {
		t.Fatalf("Expected (0, 1, N), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveForwardEast(t *testing.T) {
	rover := mars_rover.NewRover(0, 0, 'E', planet)

	rover.Interpret('f')
	if rover.X != 1 || rover.Y != 0 || rover.Direction != 'E' {
		t.Fatalf("Expected (1, 0, E), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveBackwardNorth(t *testing.T) {
	rover := mars_rover.NewRover(5, 5, 'N', planet)
	rover.Interpret('b')
	if rover.X != 5 || rover.Y != 4 {
		t.Fatalf("Expected (5, 4, N), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMoveBackwardSouth(t *testing.T) {
	rover := mars_rover.NewRover(5, 5, 'S', planet)
	rover.Interpret('b')
	if rover.X != 5 || rover.Y != 6 {
		t.Fatalf("Expected (5, 6, S), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}

func TestRoverMovesLeft(t *testing.T) {
	rover := mars_rover.NewRover(5, 5, 'N', planet)
	rover.Interpret('l')
	if rover.Direction != 'W' {
		t.Fatalf("Expected (5, 5, W), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
func TestRoverMovesRight(t *testing.T) {
	rover := mars_rover.NewRover(2, 3, 'E', planet)
	rover.Interpret('r')
	if rover.Direction != 'S' {
		t.Fatalf("Expected (2, 3, S), got (%d, %d, %c)", rover.X, rover.Y, rover.Direction)
	}
}
//multiple commands
//obstactle