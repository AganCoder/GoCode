package main

import "fmt"

// 指定了零个方法的接口值被称为 *空接口：*
// 空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
// 空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。

// 这就相当于 Swift 里面的 Any 或者 AnyObject
// 			OC 里面的 id
//			C 里的 void * 万能指针

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}