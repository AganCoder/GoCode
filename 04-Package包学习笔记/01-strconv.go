package main

import (
	"fmt"
	"strconv"
)

// conver ASCII to integer & integer to ASCII
func atoiFunc() {
	i, err := strconv.Atoi("-42")
	if err != nil {
		fmt.Println("err")
		return
	}
	fmt.Println(i)
}

func itoaFunc() {
	s := strconv.Itoa(-3000)
	fmt.Println(s)
}

func main() {

	// ================
	v32 := "354634382"
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v32, 16, 32); err == nil { // 转换超过了 32 bitSize
		fmt.Printf("%T, %v\n", s, s)
	}

	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v64, 16, 64); err == nil { // 转换超过了 64 bitSize
		fmt.Printf("%T, %v\n", s, s)
	}

	// ================

	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))

	// ================
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32)
	fmt.Println(string(b32))

	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64))

	// ===========
	quote := []byte("quote:")
	quote = strconv.AppendQuote(quote, `"Fran & Freddie's Diner"`)
	fmt.Println(string(quote))

}
