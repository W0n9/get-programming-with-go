package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(86)
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}
