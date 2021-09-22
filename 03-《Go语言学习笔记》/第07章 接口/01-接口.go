package main

import "fmt"

type tester interface {
	test()
	string() string
}

type data struct {
}

func (d *data) test() {
	fmt.Printf("test \n")
}

func (d *data) string() string {
	fmt.Printf("string \n")
	return ""
}

func main() {
	d := data{}

	d.test()
	d.string()

	var t tester = &d
	t.string()
	t.test()
}
