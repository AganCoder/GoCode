package main

import "fmt"

func test() *int {
	a := 0x10
	fmt.Printf("%p \n", &a)
	return &a
}

func main()  {
	a := test()

	fmt.Printf("%p\n", a)
	fmt.Printf("a = %d", *a)
}
