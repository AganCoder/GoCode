package main

import (
	"fmt"
	"time"
)

type T struct {
	name string
	age  int
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")

	say("Hello")
}
