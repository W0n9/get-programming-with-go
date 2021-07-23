package main

import "fmt"

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello", "hello", "world", "hello"} {
		downstream <- v
	}
	close(downstream)
}

func remove(upstream, downstream chan string) {
	poor := []string{""}

	for v := range upstream {
		flag := false
		for _, j := range poor {
			if v == j {
				flag = true
			}
		}
		if !flag {
			poor = append(poor, v)
			downstream <- v
		}
	}
	close(downstream)
}

func print(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go sourceGopher(c1)
	go remove(c1, c2)
	print(c2)
}
