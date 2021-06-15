package main

import "fmt"

func main() {
	distance := 56000000
	const hoursperday = 24
	const days = 28
	fmt.Println(distance / (hoursperday * days))
}
