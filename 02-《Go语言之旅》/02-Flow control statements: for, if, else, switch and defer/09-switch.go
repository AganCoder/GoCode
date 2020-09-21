package main

import (
	"fmt"
	"runtime"
)

// 1. Switch 和 if 一样，可以插入一条语句，然后接着判断
// 2. switch 的 case 无需为常量，取值不必为整数
// 3. 默认在最后面加了一个 break
// 4. 使用 fallthrough 后，分支不会自动终止（相当于后面没有加 break）

func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
