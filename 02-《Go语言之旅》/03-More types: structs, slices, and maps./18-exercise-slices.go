package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]int, dy)

	for i := range a {
		b := make([]int, dx)

		for j := range b {
			b[j] = (i + j) / 2
		}

		a[i] = b
	}
	return a
}

func main() {
	pic.Show(Pic)
}
