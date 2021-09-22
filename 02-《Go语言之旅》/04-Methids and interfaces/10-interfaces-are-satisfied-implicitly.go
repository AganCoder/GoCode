package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。
// 这样带来的一个问题是： 无法知道 类型实现了哪些接口，不够明确
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i T = T{"hello"}
	i.M()
}
