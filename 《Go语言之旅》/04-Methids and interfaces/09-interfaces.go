package main

import (
	"fmt"
	"math"
)
// 接口类型 是由一组方法签名定义的集合。
//
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func main()  {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	// 因为 Vertex 未实现 Abser， 而 *Vertex 实现了该方法
	//a = v
	fmt.Println(a.Abs())
}
