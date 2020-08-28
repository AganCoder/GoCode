package main

import "fmt"

// Go 语言里面没有提供 while 和 do-while

func main() {
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
