package main

import "fmt"

// 切片的零值是 nil。
// nil 切片的长度和容量为 0 且没有底层数组。

// 1. 切片如果底层没有数组，那么切片为零值 nil
// 2.            有数组， 那么切片就为空，而不是 nil
// 3. 数组申明不进行实例化，会有一个默认值

func main() {
	// 相当于声明了一个变量为s 整形数组切片,但未实现
	var s []int

	fmt.Println(s, len(s), cap(s)) // [] 0 0

	if s == nil {
		fmt.Println("nil!") // nil!
	}

	s1 := []int{ }

	fmt.Println(s1, len(s1), cap(s1))


	if s1 == nil {
		fmt.Println("s1 nil!")
	}

	var s2 [6]int
	fmt.Println(s2, len(s2), cap(s2)) // [0 0 0 0 0 0] 6 6


}