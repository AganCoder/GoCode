package main

import "fmt"

// 发送者可通过 close 关闭一个信道来表示没有需要发送的值了
// 循环 for i := range c 会不断从信道接收值，直到它被关闭。
// 向一个已经关闭的信道发送数据会引发程序恐慌（panic）。

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
