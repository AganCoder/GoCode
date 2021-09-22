package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{}
	v.X = 100
	v.Y = 200

	fmt.Println(v)

	fmt.Println(Vertex{1, 2})
}
