package main

import "fmt"

// 当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。

func main() {
	var s []int 
	printSlice(s)
	fmt.Printf("%p \n", &s)

	if s == nil {
		fmt.Println("nil! ")
	}

	s = append(s, 0)
	printSlice(s)
	fmt.Printf("%p \n", &s)

	s = append(s, 2, 3, 4)
	printSlice(s)
	fmt.Printf("%p \n", &s)
}

func printSlice(s []int) {
	fmt.Printf(" len = %d, cap = %d, %v \n ", len(s), cap(s), s)
}