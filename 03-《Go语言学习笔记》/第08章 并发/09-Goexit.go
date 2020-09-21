package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 2; i ++ {
		go func(x int) {
			for n := 0; n < 2; n++ {
				fmt.Printf("%c: %d \n", 'a'+x, n)
				time.Sleep(time.Second)
			}
		}(i)

	}
	runtime.Goexit() // 等待所有的任务结束，然后让进程直接崩溃
	println("main exit") // 不会执行
}
