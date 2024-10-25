package mars_rover

import (
	"mars_rover"
	"testing"
)

func TestReportWhenObstacleObstructing(t *testing.T) {

	start := mars_rover.Position{X: 3, Y: 3, D: "N"}
	rover := mars_rover.NewRover(&start, planet)

	rover.Command("f")

	actual := rover.Report()
	expected := "Obstacle encountered! Current coordinates : ({X:3 Y:3 D:N})"
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
func TestReportLocationAfterMovingUnimpeeded(t *testing.T) {

	start := mars_rover.Position{X: 0, Y: 0, D: "N"}
	rover := mars_rover.NewRover(&start, planet)

	rover.Command("f")

	actual := rover.Report()
	expected := "New location: ({X:0 Y:1 D:N})"
	if actual != expected {
		t.Fatalf("Expected %+v) , got %+v ", expected, actual)
	}
}
