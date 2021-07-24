package main

import (
	"fmt"
	"strings"
)

func source(downstream chan string) {
	for _, v := range []string{"foo bar", "bar baz", "foo baz"} {
		downstream <- v
	}
	close(downstream)
}

func split(upstream, downstream chan string) {
	for v := range upstream {
		for _, j := range strings.Fields(v) {
			downstream <- j
		}
	}
	close(downstream)
}

func printout(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go source(c1)
	go split(c1, c2)
	printout(c2)
}
