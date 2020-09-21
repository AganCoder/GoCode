package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// 对于方法而言，方法接受者如果是一个指针，在调用的时候，系统会默认把接受者转化为指针接受者
// 如 下面的 Scale 方法接受者是 Vertex 指针，你依然可以使用 v.Scale(10) 来调用，Go会默认转为(&v).Scale(10)
func (v *Vertex) Scale(f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 对于函数而言，传递的参数类型必须与实际类型一致，Go不会进行默认的转换
func ScaleFunc(v *Vertex, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := Vertex{3, 4}
	p.Scale(3)
	ScaleFunc(&p, 8)

	fmt.Println(v, p)
}


