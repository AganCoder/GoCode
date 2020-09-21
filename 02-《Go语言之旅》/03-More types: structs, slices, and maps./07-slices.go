package main

import "fmt"
import "reflect"

// a[low : high] 表示 [low, high) 半开区间

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	fmt.Println("%T Value: %v", s, reflect.TypeOf(s))
}