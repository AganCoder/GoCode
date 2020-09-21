package main

import (
	"fmt"
	"math"
)

// 结构体 值传递,
type Vertex struct {
	X, Y  float64
}

func (v Vertex) Abs() float64  {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// 如果不使用指针的话，形参地址与实参地址是不一样的，相当于 copy 了一份
func (v Vertex) Scale(f float64)  {
	fmt.Printf("%p\n", &v)

	v.X = v.X * f
	v.Y = v.Y * f

	//fmt.Printf("%p\n", &v)
}

func main() {
	v := Vertex{3, 4}

	v1 := v

	fmt.Printf("%p ----- %p\n", &v, &v1)

	v.Scale(10)
	fmt.Println(v.Abs())
}

