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

func TestLoadingExampleMap(t *testing.T) {

	stubMap := mars_rover.ExampleMap()
	start := mars_rover.Position{X: 0, Y: 1, D: "N"}
	rover := mars_rover.NewRover(&start, stubMap)
	
	if rover.GetPosition() != start {
		t.Fatalf("Expected (0, 0, N), got (%+v)", start)
	}
}
func TestRoverCanMoveForwardNorthFacing(t *testing.T) {
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

func TestRoverCanMoveBackwardNorthFacing(t *testing.T) {
	start := mars_rover.Position{X: 5, Y: 5, D: "N"}
	rover := mars_rover.NewRover(&start, emptyplanet)
	rover.Command("b")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 5, Y: 4, D: "N"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
func TestRoverCanMoveBackwardSouthFacing(t *testing.T) {

	start := mars_rover.Position{X: 5, Y: 5, D: "S"}
	rover := mars_rover.NewRover(&start, emptyplanet)
	rover.Command("b")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 5, Y: 6, D: "S"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}

func TestRoverCanMoveLeft(t *testing.T) {
	start := mars_rover.Position{X: 5, Y: 5, D: "N"}
	rover := mars_rover.NewRover(&start, emptyplanet)
	rover.Command("l")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 5, Y: 5, D: "W"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
func TestRoverCanMoveRight(t *testing.T) {
	start := mars_rover.Position{X: 5, Y: 5, D: "N"}
	rover := mars_rover.NewRover(&start, emptyplanet)
	rover.Command("r")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 5, Y: 5, D: "E"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
func TestRoverShouldNotDriveThroughObstacle(t *testing.T) {

	start := mars_rover.Position{X: 3, Y: 3, D: "N"}
	rover := mars_rover.NewRover(&start, planet)
	rover.Command("f")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 3, Y: 3, D: "N"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
func TestReportMsgWhenObstacleObstructing(t *testing.T) {

	start := mars_rover.Position{X: 3, Y: 3, D: "N"}
	rover := mars_rover.NewRover(&start, planet)
	rover.Command("f")
	actual := rover.Output()
	expected := "Obstacle encountered! Current coordinates : ({X:3 Y:3 D:N})"
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
func TestRoverCanHandleMultipleCommands(t *testing.T) {

	start := mars_rover.Position{X: 0, Y: 0, D: "N"}
	rover := mars_rover.NewRover(&start, planet)
	rover.Command("fffrff")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 2, Y: 3, D: "E"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
func TestRoverCanStopInFrontOfObstacleafterMultipleCommands(t *testing.T) {

	start := mars_rover.Position{X: 0, Y: 0, D: "N"}
	rover := mars_rover.NewRover(&start, planet)
	rover.Command("rffflfffffff")
	actual := rover.GetPosition()
	expected := mars_rover.Position{X: 3, Y: 3, D: "N"}
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}