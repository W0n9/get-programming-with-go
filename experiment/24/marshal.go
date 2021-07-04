package main

import (
	"encoding/json"
	"fmt"
)

type jsonout struct {
	Decimal float64 `json:"decimal"`
	DMS     string  `json:"dms"`
	Deg     float64 `json:"degress"`
	Min     float64 `json:"minutes"`
	Sec     float64 `json:"seconds"`
	Hemi    rune    `json:"hemisphere"`
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	var A jsonout
	A.Decimal = c.decimal()
	A.DMS = fmt.Sprintf("%vº%v'%.1f\" %c", c.d, c.m, c.s, c.h)
	A.Deg = c.d
	A.Min = c.m
	A.Sec = c.s
	A.Hemi = c.h
	return json.MarshalIndent(A, "", "  ")
}

// location with a latitude, longitude in decimal degrees.
type location struct {
	lat, long coordinate
}

// String formats a location with latitude, longitude.
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {
	elysium := location{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0, 'E'},
	}

	fmt.Println("Elysium Planitia is at", elysium)
}
