package main

import "fmt"

var a string 

// G O G

func main() {
	a = "G"

	fmt.Printf("%p  - %v \n", &a, a)

	print(a)
	f1()
}

func f1() {
	a := "O"

	fmt.Printf("%p  - %v \n", &a, a)

	print(a)

	f2()
}

func f2() {

	fmt.Printf("%p  - %v \n", &a, a)
}

