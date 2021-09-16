package main

import "fmt"

type I interface {
}

func main() {

	var num1 interface{} = 10
	var num2 interface{} = "10"

	if num1 == num2 {
		fmt.Println("equal")
	} else {
		fmt.Println("Not equal")
	}

}
