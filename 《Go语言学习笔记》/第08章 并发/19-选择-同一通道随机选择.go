package main

import "sync"

// 会随机选择某个通道

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)

	go func() {
		defer  wg.Done()

		for {
			var i int
			var ok bool

			select {
			case i, ok = <-c:
				println("a1: ", i)
				case i, ok = <-c:
					println("a2:", i)
			}

			if !ok {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 10; i ++ {
			select {
			case c <- i:
				case c <- i * 10:

			}
		}
	}()

	wg.Wait()
}