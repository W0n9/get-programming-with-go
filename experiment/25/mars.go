package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cat struct {
	name string
}

func (c cat) move() string {
	return "miao"
}
func (c cat) eat() string {
	return "eating fish"
}
func (c cat) String() string {
	return c.name
}

type dog struct {
	name string
}

func (d dog) move() string {
	return "bark"
}
func (d dog) eat() string {
	return "eating bone"
}
func (d dog) String() string {
	return d.name
}

type animal interface {
	move() string
	eat() string
}

func step(a animal) string {
	if rand.Intn(2) == 0 {
		return fmt.Sprintf("%v %v", a, a.move())
	} else {
		return fmt.Sprintf("%v %v", a, a.eat())
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	animals := []animal{
		cat{name: "ccc"},
		dog{name: "ddd"},
	}

	var day, hour int
	for {
		i := rand.Intn(2)
		fmt.Printf("%v:00 %v\n", hour, step(animals[i]))
		time.Sleep(time.Second)
		hour++
		if hour == 24 {
			hour = 0
			day++
		}
		if day == 3 {
			break
		}
	}
}
