package main

import "fmt"

// 1. const 修饰变量值不能改变
// 2. 常量不能用 := 语法
// 3. 常量首字母大写

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	// Pi = 1 cannot assign to Pi

	const Truth = true
	fmt.Println("Go rules?", Truth)
}