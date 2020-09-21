package main

import (
	"fmt"
	"time"
)

// Go语言里面的 没有条件的 switch 与 Swift 里面 case let 很像

func main() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")
	}
}
