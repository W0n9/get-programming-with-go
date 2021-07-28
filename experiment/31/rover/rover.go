package main

import (
	"image"
	"log"
	"time"
)

func main() {
	r := NewRoverDriver()
	time.Sleep(time.Second)
	r.Start()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(time.Second)
}

// RoverDriver drives a rover around the surface of Mars.
type RoverDriver struct {
	commandc chan command
}

// NewRoverDriver starts a new RoverDriver and returns it.
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop  = command(3)
)

// drive is responsible for driving the rover. It
// is expected to be started in a goroutine.
func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	status := false
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case start:
				status = true
				log.Println("Rover Started")
			case stop:
				status = false
				log.Println("Rover Stopped")
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			if status {
				pos = pos.Add(direction)
			}
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

// Left turns the rover left (90° counterclockwise).
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right turns the rover right (90° clockwise).
func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Start() {
	r.commandc <- start
}
func (r *RoverDriver) Stop() {
	r.commandc <- stop
}
