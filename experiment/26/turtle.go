package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y++
}
func (t *turtle) down() {
	t.y--
}
func (t *turtle) left() {
	t.x--
}
func (t *turtle) right() {
	t.x++
}

func main() {
	turtle := turtle{}
	fmt.Printf("%+v\n", turtle)
	turtle.up()
	fmt.Printf("%+v\n", turtle)
	turtle.right()
	fmt.Printf("%+v\n", turtle)
	turtle.down()
	fmt.Printf("%+v\n", turtle)
	turtle.left()

}
