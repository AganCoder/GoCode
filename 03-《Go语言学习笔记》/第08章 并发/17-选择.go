package main

import (
	"sync"
)

// 选择的作用在于处理多个通道问题

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	go func() { // 接收端
		defer wg.Done()

		for {
			var (
				name string
				x    int
				ok   bool
			)
			select { // 随机选择可用 channel 接受数据
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}

			if !ok {
				return
			}

			println(name, x)
		}
	}()

	go func() { // 发送端
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {
			select { // 随机选择发送 channel
			case a <- i:
			case b <- i * 10:

			}
		}
	}()

	wg.Wait()
}
