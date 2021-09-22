package main

import (
	"fmt"
	//"hello"
	"os"
	//"hello/say"
)

func main() {
	f, err := os.Open("/dev/randon")

	fmt.Println(f, err)

	//hello.SayHello()
	//say.SayHello()

	//f, err := os.Open("/dev/random") // no new variables on left side of :=
	//fmt.Println(f, err)

}
