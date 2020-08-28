package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	 X, Y float64
}

// 方法调用 p.Abs() 会被解释为 (*p).Abs()。
// 此处还是值传递
func (v Vertex) Abs() float64 {
	fmt.Printf("%p \n", &v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("%p \n", &v)
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{3, 4}
	fmt.Printf("%p \n", p)

	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
