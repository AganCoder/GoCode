package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0

	tmp := 0.00

	for {
		z = z - (z*z-x)/(2*z)

		fmt.Println(z)

		if math.Abs(z-tmp) < 0.0000000001 {
			break
		} else {
			tmp = z
		}
	}
	return z

}

func main() {
	fmt.Println(Sqrt(2))
}
