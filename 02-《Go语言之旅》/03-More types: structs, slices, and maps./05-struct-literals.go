package main

import "fmt"

type Vertex struct {
	X, Y int
}

// 结构体文法通过直接列出字段的值来新分配一个结构体。
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
