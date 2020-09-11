package main

import "fmt"

type data struct {
	x int
}

func main() {
	d := data{
		x: 100,
	}

	var t interface{} = d

	println(t.(data).x)

	fmt.Printf("%p \n", &d) // 0xc000016070
	fmt.Printf("%p \n", &t) // 0xc000010200

	// t.(data).x = 200 // cannot assign to t.(data).x 无法进行修改

	// 接口进行复制的时候，值类型复制
	var t2 interface{} = &d
	println(t2.(*data).x)

	t2.(*data).x = 200
	println(t2.(*data).x)
}
