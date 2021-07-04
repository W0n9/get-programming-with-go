package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	Lats := []float64{-4.5895, -14.5684, -1.9462}
	Longs := []float64{137.4417, 175.472636, 354.4734}

	_ = Lats
	_ = Longs
	type location struct {
		Name string
		Lat  float64
		Long float64
	}

	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}

	_ = locations

	bytes, _ := json.MarshalIndent(locations, "", "  ")
	fmt.Println(string(bytes))
}
