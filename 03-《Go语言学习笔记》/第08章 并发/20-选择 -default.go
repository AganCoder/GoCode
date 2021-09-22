package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)

	go func() {
		defer wg.Done()

		for {
			select {
			case x, ok := <-c:
				if !ok {
					return
				}

				fmt.Println("data:", x)

			default: // 避免 select 阻塞
			}

			fmt.Println(time.Now())
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 3)

	c <- 100
	close(c)

}
