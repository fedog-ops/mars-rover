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
	start := mars_rover.Position{X: 0, Y: 1, D: "N"}
	rover := mars_rover.NewRover(&start, stubMap)
	
	if rover.GetPosition() != start {
		t.Fatalf("Expected (0, 0, N), got (%+v)", start)
	}
}
func TestRoverMoveForwardNorth(t *testing.T) {
	start := mars_rover.Position{X: 0, Y: 1, D: "N"}
	rover := mars_rover.NewRover(&start, emptyplanet)

	rover.Command("f")

	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 0, Y: 2, D: "N"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}

}

func TestRoverMoveForwardEastFacing(t *testing.T) {
	start := mars_rover.Position{X: 0, Y: 0, D: "E"}
	rover := mars_rover.NewRover(&start, emptyplanet)

	rover.Command("f")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 1, Y: 0, D: "E"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}

func TestRoverMoveBackwardNorth(t *testing.T) {
	start := mars_rover.Position{X: 5, Y: 5, D: "N"}
	rover := mars_rover.NewRover(&start, emptyplanet)
	rover.Command("b")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 5, Y: 4, D: "N"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
func TestRoverMoveBackwardSouth(t *testing.T) {

	start := mars_rover.Position{X: 5, Y: 5, D: "S"}
	rover := mars_rover.NewRover(&start, emptyplanet)
	rover.Command("b")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 5, Y: 6, D: "S"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}

func TestRoverMovesLeft(t *testing.T) {
	start := mars_rover.Position{X: 5, Y: 5, D: "N"}
	rover := mars_rover.NewRover(&start, emptyplanet)
	rover.Command("l")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 5, Y: 5, D: "W"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
func TestRoverMovesRight(t *testing.T) {
	start := mars_rover.Position{X: 5, Y: 5, D: "N"}
	rover := mars_rover.NewRover(&start, emptyplanet)
	rover.Command("r")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 5, Y: 5, D: "E"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
func TestDoesntDriveThroughObstacle(t *testing.T) {

	start := mars_rover.Position{X: 3, Y: 3, D: "N"}
	rover := mars_rover.NewRover(&start, emptyplanet)
	rover.Command("f")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 3, Y: 3, D: "N"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
