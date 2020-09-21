package main

import "fmt"

// := 结构不能在函数外使用。

func main() {
	var i, j int = 1, 2

	k := 3

	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}