package main

import "fmt"

type data int

func (d data) String() string {
	return fmt.Sprintf("data: %d", d)
}

func main() {
	var d data = 15

	var x interface{} = d

	if n, ok := x.(fmt.Stringer); ok {
		fmt.Println(n)
	}

	if n, ok := x.(data); ok {
		fmt.Println(n)
	}

	//e := x.(error)
	//fmt.Println(e)

	// ==========
	var xiface interface{} = func(x int) string {
		return fmt.Sprintf("d: %d", x)
	}

	switch v := xiface.(type) {
	case nil:
		println("nil")
	case *int:
		println("*int")
	case func(int2 int) string:
		println(v(123))
	case fmt.Stringer:
		fmt.Println(v)
	default:
		println("unknown")
	}
}
