package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 当接口内部具体值为 nil， 仍然会调用 nil 接收者方法
// 在 OC 中， nil 接收者不会调用任何方法 --- GO 语言会调用
// Python  Java 或者 C里面会触发空指针异常，Go 语言不会
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	var t *T

	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}
