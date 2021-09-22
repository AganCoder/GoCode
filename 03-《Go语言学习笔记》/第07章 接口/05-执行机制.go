package main

import "fmt"

type Ner interface {
	a()
	b(int)
	c(string) string
}
type N int

func (n N) a() {
	fmt.Printf("a")
}

func (n N) b(int2 int) {
	fmt.Printf("b")
}

func (n N) c(string2 string) string {
	fmt.Printf("c")
	return ""
}

func main() {
	var n N

	var t Ner = &n

	t.a()
	t.b(10)
	t.c("aaa")

}
