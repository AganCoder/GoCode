package main

import (
	"fmt"
	"runtime"
)

func main()  {
	var goods interface{} = runtime.GOOS
	switch goods {
	case 1:
		fmt.Printf("Int value")
	case "windos":
		fmt.Printf("windows")
	default:
		fmt.Printf("default")
	}
}
