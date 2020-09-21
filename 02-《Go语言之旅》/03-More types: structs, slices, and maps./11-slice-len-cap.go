package main

import "fmt"


func main() {

	s := []int {2, 3, 5, 7, 11, 13}
	printSlice(s)
	fmt.Println(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)
	fmt.Println(s)
		
	// 动态扩容了，虽然切片 s 为空，但是底层的数组依旧存在，s[:4] 就从空的位置往后 4 位
	s = s[:4]
	fmt.Println(s)
	printSlice(s)

	// s = s[2: ]
	// printSlice(s)
	// fmt.Println(s)

	// s = s[1:4]
	// fmt.Println(s)

	// s = s[:2]
	// fmt.Println(s)

	// s = s[1:]
	// fmt.Println(s)
	
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}
