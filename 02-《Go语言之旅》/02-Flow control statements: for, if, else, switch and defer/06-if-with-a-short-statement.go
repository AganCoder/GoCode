package main

import (
	"fmt"
	"math"
)

// if 语句可以在条件表达式前执行一个简单的语句。
// 只能有一条语句

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 20),
	)
}
