package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func nestedFunc() {
	var funcvar = func(x int) int {
		return 10
	}

	fmt.Println(funcvar(10))
}

func hello() {
	fmt.Println("hello world")
}

func exec(f func()) {
	f()
}

func main() {
	fmt.Println(add(42, 13))
	nestedFunc()

	f := hello
	exec(f)
}
