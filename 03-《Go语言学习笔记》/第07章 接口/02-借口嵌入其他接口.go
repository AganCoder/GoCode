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

func (d *data) test() {
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
}
