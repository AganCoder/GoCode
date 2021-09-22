package main

import "sync"

type receiver struct {
	sync.WaitGroup
	data chan int
}

// 使用工厂方法将 goroutine 和通道进行绑定
func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}

	r.Add(1)

	go func() {
		defer r.Done()
		for x := range r.data {
			println("recv:", x)
		}
	}()

	return r
}

func main() {
	r := newReceiver()
	r.data <- 1
	r.data <- 2

	close(r.data)
	r.Wait()
}
