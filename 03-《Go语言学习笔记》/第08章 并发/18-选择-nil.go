package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	a := make(chan int)
	b := make(chan int)

	go func() { // 准备好接受端
		defer wg.Done()

		for {
			select { // 随机选择一个可用的通道
			case x, ok := <-a:
				if !ok { // 如果通道被关闭了，设置为 nil， 阻塞
					a = nil
					break
				}
				println(x, "a")
			case x, ok := <-b:
				if !ok { // 如果通道被关闭了，设置为 nil， 阻塞
					b = nil
					break
				}
				println(x, "b")
			}

			if a == nil && b == nil { // 全部 close 直接return
				return
			}
		}
	}()

	go func() { // 发送端
		defer wg.Done()
		defer close(a)

		for i := 0; i < 3; i++ {
			a <- i
		}
	}()

	go func() { // 发送端
		defer wg.Done()
		defer close(b)

		for i := 0; i < 5; i++ {
			b <- i
		}
	}()

	wg.Wait()
}
