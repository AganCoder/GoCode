package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M()  {
	fmt.Println(t.S)
}

type F float64

func (f F) M()  {
	fmt.Println(f)
}

// 接口可以传递，可以用作函数的参数或者返回值
// 接口值保存了一个具体底层类型的具体值
func describe(i I)  {
	fmt.Printf("%v %T \n", i, i)
}

func main()  {
	var i I
	i = &T{"Hello"}
	describe(i)

	i = F(math.Pi)
	describe(i)
	i.M()

}