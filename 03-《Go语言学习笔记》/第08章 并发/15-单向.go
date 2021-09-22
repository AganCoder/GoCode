package main

import (
	"fmt"
	"sync"
)

// 通道默认是双向的，并不区分发送端和接受端

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	c := make(chan int)

	var send chan<- int = c

	fmt.Printf("send: %p, c: %p \n", send, c)

	var recv <-chan int = c

	fmt.Printf("recv: %p, c: %p \n", send, c)

	go func() {
		defer wg.Done()

		for x := range recv {
			println(x)
		}
	}()

	go func() {
		defer wg.Done()
		//defer close(recv) // invalid operation: close(recv) (cannot close receive-only channel)
		defer close(c)

		for i := 0; i < 3; i++ {
			send <- i
		}

		c <- 100
	}()
	wg.Wait()
}
