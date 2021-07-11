package main

import "fmt"

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v pickup a %v\n", c.name, i.name)
	c.leftHand = i
}
func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if to.leftHand != nil {
		fmt.Printf("%v no hand\n", to.name)
	}
	if c.leftHand == nil {
		fmt.Printf("%v has nothing\n", c.name)
	}
	fmt.Printf("%v give %v to %v\n", c.name, c.leftHand.name, to.name)
	to.leftHand = c.leftHand
	c.leftHand = nil
}

func (c character) String() string {
	if c.leftHand != nil {
		return fmt.Sprintf("%v has %v", c.name, c.leftHand.name)
	} else {
		return fmt.Sprintf("%v has nothing", c.name)
	}

}

func main() {
	arthur := character{name: "arthur"}
	knight := character{name: "knight"}
	knife := item{name: "knife"}
	fmt.Println(arthur)
	fmt.Println(knight)
	arthur.pickup(&knife)
	fmt.Println(arthur)
	fmt.Println(knight)
	arthur.give(&knight)
	fmt.Println(arthur)
	fmt.Println(knight)

}
