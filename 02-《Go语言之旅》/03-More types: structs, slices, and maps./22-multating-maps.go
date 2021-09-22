package main

import "fmt"

// 删除元素
// 更改元素
// 添加元素

// 若 key 在 m 中，ok 为 true ；否则，ok 为 false。
// 若 key 不在映射中，那么 elem 是该映射元素类型的零值。

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The values:", m["Answer"])

	m["Answer"] = 42
	fmt.Println("The values:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The values:", m["Answer"])

	v, ok := m["Answer"]

	fmt.Println("The value:", v, "Present?", ok)

}
