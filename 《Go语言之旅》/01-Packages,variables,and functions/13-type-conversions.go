package main

import (
	"fmt"
	"math"
)

// 类型之间需要显示转换
// error: cannot use f (type float64) as type uint in assignment

func main() {
	var x, y int = 3, 4

	var f float64 = math.Sqrt(float64(x * x + y * y) )

	var z uint = uint(f)

	fmt.Println(x, y, z)
}