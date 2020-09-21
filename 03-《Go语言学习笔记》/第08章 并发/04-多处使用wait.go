package main

import (
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		wg.Wait()

		time.Sleep(time.Second) // 因为延时，这个没走
		println("wait exit")
	}()

	go func() {
		time.Sleep(time.Second)
		println("done")
		wg.Done()
	}()

	wg.Wait()

	println("main exit")

	println("GOMAXPROCS", runtime.GOMAXPROCS(0))

}
