package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 映射的零值为 nil 。nil 映射既没有键，也不能添加键。
var m map[string]Vertex

func main() {
	fmt.Println(m)

	if m == nil {
		fmt.Println("nil!")
	}

	m = make( map[string]Vertex)

	fmt.Println(m)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"] )
}