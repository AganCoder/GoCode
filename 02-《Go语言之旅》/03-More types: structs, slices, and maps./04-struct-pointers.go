package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v

	(*p).X = 1e9 // 应该为这种写法
	p.X = 1e8   // 语言层面提供了隐式间接引用

	fmt.Println(v)
	fmt.Println(p)
}