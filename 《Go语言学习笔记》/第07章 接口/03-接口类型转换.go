package main

import "fmt"

type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type data struct {
}

func print(s stringer) {
	fmt.Println(s.string())
}

func (d *data) test()  {
	fmt.Printf("test")
}

func (d *data) string() string {
	fmt.Println("string")
	return ""
}

func main() {
	var d = data{}

	var t tester = &d
	t.test()
	t.string()

	var s stringer = t
	print(s)

	// var t2 tester = s //cannot use s (type stringer) as type tester in assignment:
}