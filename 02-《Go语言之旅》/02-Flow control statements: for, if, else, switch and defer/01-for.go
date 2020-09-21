package main

import "fmt"

// 1. for 后面不需要括号
// 2. for 最后的大括号必须的

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}
