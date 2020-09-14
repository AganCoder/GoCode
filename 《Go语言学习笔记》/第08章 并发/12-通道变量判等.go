package main

import (
	"fmt"
	"unsafe"
)

// 通道变量本身是一个指针，可以用操作符判断是否为同一对象或者 nil

func main() {
	var a, b chan  int = make(chan  int, 3), make(chan  int)

	var c chan  int
	var d chan  int

	println( a == b )
	println( c == nil )

	println(c == d) // true, 都为 nil

	a <- 3
	a <- 2
	a <- 1

	// a <- 0 满了，fatal error: all goroutines are asleep - deadlock!

	println("a: ", len(a), cap(a))

	fmt.Printf("%p, %d \n", a, unsafe.Sizeof(a))
}
