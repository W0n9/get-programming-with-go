package main

import "fmt"

func main() {
	type celsius float64

	const degrees = 20
	var temperature celsius = degrees

	temperature += 10
	var warm float64 = 10
	temperature += celsius(warm)
	fmt.Println(temperature)
}
