package mars_rover

import (
	"encoding/json"
	"fmt"
	"os"
)

type Coordinates2 struct {
	X, Y int `json:"x",json:"y"`
}

type MapJSON struct {
	Width     int         `json:"width"`
	Height    int         `json:"height"`
	Obstacles []Coordinates `json:"obstacles"`
}

func LoadMap(filename string) (*Map, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	var m Map
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&m); err != nil {
		return nil, fmt.Errorf("could not decode JSON: %w", err)
	}

	return &m, nil
}
